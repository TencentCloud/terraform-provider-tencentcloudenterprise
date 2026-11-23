package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cloud "terraform-provider-tencentcloudenterprise/tencentcloud"
)

const (
	cloudMark      = "tencentcloudenterprise"
	cloudTitle     = "TencentCloudEnterprise"
	cloudPrefix    = cloudMark + "_"
	cloudMarkShort = "tc"
	docRoot        = "../docs"
)

var (
	hclMatch   = regexp.MustCompile("(?si)([^`]+)?```(hcl)?(.*?)```")
	usageMatch = regexp.MustCompile(`(?s)(?m)^([^ \n].*?)(?:\n{2}|$)(.*)`)
	bigSymbol  = regexp.MustCompile("([\u007F-\uffff])")

	// Statistics
	statsSuccess int
	statsSkipped int
	statsFailed  int
)

// productNameMap maps resource prefix to human-readable product name for subcategory.
// Resources not matching any key here will have their prefix auto-capitalized.
var productNameMap = map[string]string{
	"apm":          "Application Performance Management(APM)",
	"as":           "Auto Scaling(AS)",
	"availability": "Provider Data Sources",
	"bms":          "Bare Metal Server(BMS)",
	"brc":          "Backup and Recovery(BRC)",
	"cam":          "Cloud Access Management(CAM)",
	"cbs":          "Cloud Block Storage(CBS)",
	"ccn":          "Cloud Connect Network(CCN)",
	"cfs":          "Cloud File Storage(CFS)",
	"cfw":          "Cloud Firewall(CFW)",
	"cic":          "Corporate Identity Center(CIC)",
	"ckafka":       "Cloud Kafka(ckafka)",
	"clb":          "Cloud Load Balancer(CLB)",
	"cls":          "Cloud Log Service(CLS)",
	"cos":          "Cloud Object Storage(COS)",
	"csp":          "Cloud Storage Platform(CSP)",
	"cvm":          "Cloud Virtual Machine(CVM)",
	"cwp":          "Cloud Workload Protection Platform(CWP)",
	"dc":           "Direct Connect(DC)",
	"dcdb":         "TDSQL for MySQL(DCDB)",
	"dcx":          "Direct Connect Gateway(DCX)",
	"eip":          "Cloud Elastic IP(EIP)",
	"eips":         "Cloud Elastic IP(EIP)",
	"kms":          "Key Management Service(KMS)",
	"ngwaf":        "Web Application Firewall(NGWAF)",
	"organization": "Tencent Cloud Organization",
	"redis":        "TencentDB for Redis(crs)",
	"ssm":          "Secrets Manager(SSM)",
	"tag":          "Tag",
	"tbase":        "TDSQL PostgreSQL(Tbase)",
	"tcr":          "Tencent Container Registry(TCR)",
	"tdmq":         "TDMQ",
	"tke":          "Tencent Kubernetes Engine(TKE)",
	"tsf":          "Tencent Service Framework(TSF)",
	"turbofs":      "TurboFS",
	"vpc":          "Virtual Private Cloud(VPC)",
	"vpcdns":       "Virtual Private Cloud DNS(VPCDNS)",
}

func main() {
	provider := cloud.Provider()
	vProvider := runtime.FuncForPC(reflect.ValueOf(cloud.Provider).Pointer())

	filename, _ := vProvider.FileLine(0)
	filePath := filepath.Dir(filename)
	message("generating doc from: %s\n", filePath)

	// Build product list from provider schema (not from comments)
	products := buildProductsFromProvider(provider)
	message("discovered %d products from provider schema\n", len(products))

	// Generate index page
	genIdx(products)

	for _, product := range products {
		// document for DataSources
		for _, dataSource := range product.DataSources {
			if _, ok := provider.DataSourcesMap[dataSource]; ok {
				genDoc(product.Name, "data-sources", filePath, dataSource, provider.DataSourcesMap[dataSource])
			} else {
				message("[WARN]data source not found in provider: %s\n", dataSource)
				statsSkipped++
			}
		}

		// document for Resources
		for _, resource := range product.Resources {
			if _, ok := provider.ResourcesMap[resource]; ok {
				genDoc(product.Name, "resources", filePath, resource, provider.ResourcesMap[resource])
			} else {
				message("[WARN]resource not found in provider: %s\n", resource)
				statsSkipped++
			}
		}
	}

	// Print summary
	message("\n=== Generation Summary ===\n")
	message("✅ Success: %d\n", statsSuccess)
	message("⚠️  Skipped: %d\n", statsSkipped)
	message("❌ Failed:  %d\n", statsFailed)
	message("========================\n")
}

// buildProductsFromProvider scans provider.ResourcesMap and DataSourcesMap
// to discover all products and their resources, instead of parsing comments.
func buildProductsFromProvider(provider *schema.Provider) []Product {
	productMap := make(map[string]*Product)

	for name := range provider.DataSourcesMap {
		prefix := getProductPrefix(name)
		productName := getProductName(prefix)
		if _, ok := productMap[productName]; !ok {
			productMap[productName] = &Product{Name: productName}
		}
		productMap[productName].DataSources = append(productMap[productName].DataSources, name)
	}

	for name := range provider.ResourcesMap {
		prefix := getProductPrefix(name)
		productName := getProductName(prefix)
		if _, ok := productMap[productName]; !ok {
			productMap[productName] = &Product{Name: productName}
		}
		productMap[productName].Resources = append(productMap[productName].Resources, name)
	}

	// Convert map to sorted slice
	var products []Product
	for _, prod := range productMap {
		sort.Strings(prod.DataSources)
		sort.Strings(prod.Resources)
		products = append(products, *prod)
	}

	sort.Slice(products, func(i, j int) bool {
		// Provider Data Sources first
		if products[i].Name == "Provider Data Sources" {
			return true
		}
		if products[j].Name == "Provider Data Sources" {
			return false
		}
		return products[i].Name < products[j].Name
	})

	return products
}

// getProductPrefix extracts the product prefix from a full resource name.
// e.g. "tencentcloudenterprise_cvm_instance" -> "cvm"
// e.g. "tencentcloudenterprise_availability_zones" -> "availability"
func getProductPrefix(name string) string {
	short := strings.TrimPrefix(name, cloudPrefix)
	if strings.HasPrefix(short, "availability_") {
		return "availability"
	}
	parts := strings.SplitN(short, "_", 2)
	if len(parts) > 0 {
		return parts[0]
	}
	return "other"
}

// getProductName maps a product prefix to the human-readable product name.
func getProductName(prefix string) string {
	if name, ok := productNameMap[prefix]; ok {
		return name
	}
	return strings.ToUpper(prefix)
}

// genIdx generates the index page (docs/index.md)
func genIdx(products []Product) {
	data := map[string]interface{}{
		"cloud_mark":  cloudMark,
		"cloud_title": cloudTitle,
		"cloudPrefix": cloudPrefix,
		"Products":    products,
	}

	filename := filepath.Join(docRoot, "index.md")

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		message("[FAIL!]create directory %s failed: %s", filepath.Dir(filename), err)
		os.Exit(1)
	}

	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		message("[FAIL!]open file %s failed: %s", filename, err)
		os.Exit(1)
	}

	defer fd.Close()

	tmpl := template.Must(template.New("t").Funcs(template.FuncMap{"replace": replace}).Parse(idxTPL))

	if err := tmpl.Execute(fd, data); err != nil {
		message("[FAIL!]write file %s failed: %s", filename, err)
		os.Exit(1)
	}

	message("[SUCC.]write doc to file success: %s", filename)
}

// genDoc generating doc for data source and resource
func genDoc(product, dtype, fpath, name string, resource *schema.Resource) {
	// Convert dtype to singular form for sidebar_current
	dtypeSingular := dtype
	if dtype == "data-sources" {
		dtypeSingular = "datasource"
	} else if dtype == "resources" {
		dtypeSingular = "resource"
	}

	data := map[string]string{
		"product":           product,
		"name":              name,
		"dtype":             strings.Replace(dtype, "-", "", -1),
		"dtype_singular":    dtypeSingular,
		"resource":          strings.TrimPrefix(name, cloudPrefix),
		"cloud_mark":        cloudMark,
		"cloud_title":       cloudTitle,
		"example":           "",
		"description":       "",
		"description_short": "",
		"import":            "",
	}

	filename := fmt.Sprintf("resource_%s_%s.go", cloudMarkShort, data["resource"])
	if dtype == "data-sources" {
		filename = fmt.Sprintf("data_source_%s_%s.go", cloudMarkShort, data["resource"])
	}
	message("[START]get description from file: %s\n", filename)

	description, err := getFileDescription(filepath.Join(fpath, filename))
	if err != nil {
		message("[SKIP]get description failed: %s - %s\n", filename, err)
		statsSkipped++
		return
	}

	description = strings.TrimSpace(description)
	if description == "" {
		message("[SKIP]description empty: %s\n", filename)
		statsSkipped++
		return
	}

	// Use regex-based heading parser (supports both "# Import" and "Import")
	if before, after, ok := splitDocSection(description, "Import"); ok {
		data["import"] = strings.TrimSpace(after)
		description = strings.TrimSpace(before)
	}

	if before, after, ok := splitDocSection(description, "Example Usage"); ok {
		data["example"] = formatHCL(after)
		description = strings.TrimSpace(before)
	} else {
		message("[SKIP]example usage missing: %s\n", filename)
		statsSkipped++
		return
	}

	data["description"] = description
	pos := strings.Index(description, "\n\n")
	if pos != -1 {
		data["description_short"] = strings.TrimSpace(description[:pos])
	} else {
		data["description_short"] = description
	}

	var (
		requiredArgs []string
		optionalArgs []string
		attributes   []string
		subStruct    []string
	)

	if _, ok := resource.Schema["result_output_file"]; dtype == "data-sources" && !ok {
		if resource.DeprecationMessage != "" {
			message("[SKIP]argument 'result_output_file' is missing (deprecated): %s\n", filename)
			statsSkipped++
			return
		} else {
			message("[SKIP]argument 'result_output_file' is missing: %s\n", filename)
			statsSkipped++
			return
		}
	}

	for k, v := range resource.Schema {
		if v.Description == "" {
			message("[SKIP]description for '%s' is missing: %s\n", k, filename)
			statsSkipped++
			return
		} else {
			v.Description = checkDescription(k, v.Description)
		}
		if dtype == "data-sources" && v.ForceNew {
			message("[SKIP]Don't set ForceNew on data source: '%s' in %s\n", k, filename)
			statsSkipped++
			return
		}
		if v.Required && v.Optional {
			message("[SKIP]Don't set Required and Optional at the same time: '%s' in %s\n", k, filename)
			statsSkipped++
			return
		}
		if v.Required {
			opt := "Required"
			sub := getSubStruct(0, k, v)
			subStruct = append(subStruct, sub...)
			// get type
			res := parseSubtract(v, sub)
			valueType := parseType(v)
			if res == "" {
				opt += fmt.Sprintf(", %s", valueType)
			} else {
				opt += fmt.Sprintf(", %s: [`%s`]", valueType, res)
			}
			if v.ForceNew {
				opt += ", ForceNew"
			}
			if v.Deprecated != "" {
				opt += ", **Deprecated**"
				v.Description = fmt.Sprintf("%s %s", v.Deprecated, v.Description)
			}
			requiredArgs = append(requiredArgs, fmt.Sprintf("* `%s` - (%s) %s", k, opt, v.Description))
		} else if v.Optional {
			opt := "Optional"
			sub := getSubStruct(0, k, v)
			subStruct = append(subStruct, sub...)
			// get type
			res := parseSubtract(v, sub)
			valueType := parseType(v)
			if res == "" {
				opt += fmt.Sprintf(", %s", valueType)
			} else {
				opt += fmt.Sprintf(", %s: [`%s`]", valueType, res)
			}
			if v.ForceNew {
				opt += ", ForceNew"
			}
			if v.Deprecated != "" {
				opt += ", **Deprecated**"
				v.Description = fmt.Sprintf("%s %s", v.Deprecated, v.Description)
			}
			optionalArgs = append(optionalArgs, fmt.Sprintf("* `%s` - (%s) %s", k, opt, v.Description))
		} else {
			attrs := getAttributes(0, k, v)
			if len(attrs) > 0 {
				attributes = append(attributes, attrs...)
			}
		}
	}

	sort.Strings(requiredArgs)
	sort.Strings(optionalArgs)
	sort.Strings(attributes)
	sort.Strings(subStruct)

	// remove duplicates
	if len(subStruct) > 0 {
		uniqSubStruct := make([]string, 0, len(subStruct))
		var i int
		for i = 0; i < len(subStruct)-1; i++ {
			if subStruct[i] != subStruct[i+1] {
				uniqSubStruct = append(uniqSubStruct, subStruct[i])
			}
		}
		uniqSubStruct = append(uniqSubStruct, subStruct[i])
		subStruct = uniqSubStruct
	}

	requiredArgs = append(requiredArgs, optionalArgs...)
	data["arguments"] = strings.Join(requiredArgs, "\n")
	if len(subStruct) > 0 {
		data["arguments"] += "\n" + strings.Join(subStruct, "\n")
	}
	data["attributes"] = strings.Join(attributes, "\n")
	if dtype == "resources" {
		idAttribute := "* `id` - ID of the resource.\n"
		data["attributes"] = idAttribute + data["attributes"]
	}

	filename = filepath.Join(docRoot, dtype, fmt.Sprintf("%s.md", cloudPrefix+data["resource"]))

	// Ensure output directory exists
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		message("[FAIL]create directory %s failed: %s\n", filepath.Dir(filename), err)
		statsFailed++
		return
	}

	fd, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		message("[FAIL]open file %s failed: %s\n", filename, err)
		statsFailed++
		return
	}

	defer fd.Close()
	t := template.Must(template.New("t").Parse(docTPL))
	err = t.Execute(fd, data)
	if err != nil {
		message("[FAIL]write file %s failed: %s\n", filename, err)
		statsFailed++
		return
	}

	message("[SUCC]write doc to file success: %s\n", filename)
	statsSuccess++
}

// getAttributes get attributes from schema
func getAttributes(step int, k string, v *schema.Schema) []string {
	var attributes []string
	ident := strings.Repeat(" ", step*2)

	if v.Description == "" {
		return attributes
	} else {
		v.Description = checkDescription(k, v.Description)
	}

	if v.Computed {
		if v.Deprecated != "" {
			v.Description = fmt.Sprintf("(**Deprecated**) %s %s", v.Deprecated, v.Description)
		}
		if _, ok := v.Elem.(*schema.Resource); ok {
			var listAttributes []string
			for kk, vv := range v.Elem.(*schema.Resource).Schema {
				attrs := getAttributes(step+1, kk, vv)
				if len(attrs) > 0 {
					listAttributes = append(listAttributes, attrs...)
				}
			}
			var slistAttributes string
			sort.Strings(listAttributes)
			if len(listAttributes) > 0 {
				slistAttributes = "\n" + strings.Join(listAttributes, "\n")
			}
			attributes = append(attributes, fmt.Sprintf("%s* `%s` - %s%s", ident, k, v.Description, slistAttributes))
		} else {
			attributes = append(attributes, fmt.Sprintf("%s* `%s` - %s", ident, k, v.Description))
		}
	}

	return attributes
}

func getMarkdownFileDescription(fname string) (string, error) {
	bytes, err := os.ReadFile(fname)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// getFileDescription get description from go file
func getFileDescription(fname string) (string, error) {
	if strings.HasSuffix(fname, ".go") {
		mdFname := strings.TrimSuffix(fname, ".go") + ".md"
		if _, err := os.Stat(mdFname); err == nil {
			return getMarkdownFileDescription(mdFname)
		}
	}

	fset := token.NewFileSet()

	parsedAst, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	return parsedAst.Doc.Text(), nil
}

// splitDocSection splits doc text by a heading line. Supports headings with optional '#' prefix.
// Returns text before heading, text after heading, and whether the heading was found.
func splitDocSection(text, heading string) (string, string, bool) {
	if text == "" {
		return "", "", false
	}

	re := regexp.MustCompile("(?m)^\\s*#*\\s*" + regexp.QuoteMeta(heading) + "\\s*$")
	loc := re.FindStringIndex(text)
	if loc == nil {
		return text, "", false
	}

	before := strings.TrimSpace(text[:loc[0]])
	afterStart := loc[1]
	if afterStart < len(text) {
		if text[afterStart] == '\r' {
			afterStart++
			if afterStart < len(text) && text[afterStart] == '\n' {
				afterStart++
			}
		} else if text[afterStart] == '\n' {
			afterStart++
		}
	}
	after := strings.TrimLeft(text[afterStart:], "\r\n")

	return before, after, true
}

// getSubStruct get sub structure from go file
func getSubStruct(step int, k string, v *schema.Schema) []string {
	var subStructs []string

	if v.Description == "" {
		return subStructs
	} else {
		v.Description = checkDescription(k, v.Description)
	}

	var subStruct []string
	if v.Type == schema.TypeMap || v.Type == schema.TypeList || v.Type == schema.TypeSet {
		if _, ok := v.Elem.(*schema.Resource); ok {
			subStruct = append(subStruct, fmt.Sprintf("\nThe `%s` object supports the following:\n", k))
			var (
				requiredArgs []string
				optionalArgs []string
			)
			for kk, vv := range v.Elem.(*schema.Resource).Schema {
				desc := checkDescription(kk, vv.Description)
				if vv.Required {
					opt := "Required"
					valueType := parseType(vv)
					opt += fmt.Sprintf(", %s", valueType)
					if vv.ForceNew {
						opt += ", ForceNew"
					}
					requiredArgs = append(requiredArgs, fmt.Sprintf("* `%s` - (%s) %s", kk, opt, desc))
				} else if vv.Optional {
					opt := "Optional"
					valueType := parseType(vv)
					opt += fmt.Sprintf(", %s", valueType)
					if vv.ForceNew {
						opt += ", ForceNew"
					}
					optionalArgs = append(optionalArgs, fmt.Sprintf("* `%s` - (%s) %s", kk, opt, desc))
				}
			}
			sort.Strings(requiredArgs)
			subStruct = append(subStruct, requiredArgs...)
			sort.Strings(optionalArgs)
			subStruct = append(subStruct, optionalArgs...)
			subStructs = append(subStructs, strings.Join(subStruct, "\n"))

			for kk, vv := range v.Elem.(*schema.Resource).Schema {
				subStructs = append(subStructs, getSubStruct(step+1, kk, vv)...)
			}
		}
	}

	return subStructs
}

// formatHCL format HCL code
func formatHCL(s string) string {
	var rr []string

	s = strings.TrimSpace(s)
	m := hclMatch.FindAllStringSubmatch(s, -1)
	if len(m) > 0 {
		for _, v := range m {
			p := strings.TrimSpace(v[1])
			if p != "" {
				p = formatUsageDesc(p)
			}
			b := hclwrite.Format([]byte(strings.TrimSpace(v[3])))
			rr = append(rr, fmt.Sprintf("\n%s\n\n```hcl\n%s\n```", p, string(b)))
		}
	}

	return strings.TrimSpace(strings.Join(rr, "\n"))
}

// checkDescription check description format and auto-fix issues when possible.
// Returns the (possibly fixed) description string.
func checkDescription(k, s string) string {
	if s == "" {
		return s
	}

	// Auto-fix: trim leading spaces
	if strings.TrimLeft(s, " ") != s {
		message("[WARN]Auto-fixed: trimmed leading space from description: '%s'\n", k)
		s = strings.TrimLeft(s, " ")
	}

	// Auto-fix: trim trailing spaces
	if strings.TrimRight(s, " ") != s {
		message("[WARN]Auto-fixed: trimmed trailing space from description: '%s'\n", k)
		s = strings.TrimRight(s, " ")
	}

	// Auto-fix: add ending punctuation if missing
	if len(s) > 0 && s[len(s)-1] != '.' && s[len(s)-1] != ':' {
		message("[WARN]Auto-fixed: added ending '.' to description: '%s'\n", k)
		s = s + "."
	}

	// Warning only: unexpected symbols (non-ASCII)
	if c := containsBigSymbol(s); c != "" {
		message("[WARN]There is unexpected symbol '%s' on the description: '%s': '%s'\n", c, k, s)
	}

	// Warning only: space before punctuation
	for _, v := range []string{",", ".", ";", ":", "?", "!"} {
		if strings.Contains(s, " "+v) {
			message("[WARN]There is space before '%s' on the description: '%s': '%s'\n", v, k, s)
			break
		}
	}

	return s
}

// containsBigSymbol returns the Big symbol if found
func containsBigSymbol(s string) string {
	m := bigSymbol.FindStringSubmatch(s)
	if len(m) > 0 {
		return m[0]
	}

	return ""
}

// message print color message
func message(msg string, v ...interface{}) {
	if strings.Contains(msg, "FAIL") {
		color.Red(fmt.Sprintf(msg, v...))
	} else if strings.Contains(msg, "SUCC") {
		color.Green(fmt.Sprintf(msg, v...))
	} else if strings.Contains(msg, "SKIP") {
		color.Yellow(fmt.Sprintf(msg, v...))
	} else {
		color.White(fmt.Sprintf(msg, v...))
	}
}

func parseType(v *schema.Schema) string {
	res := ""
	switch v.Type {
	case schema.TypeBool:
		res = "Bool"
	case schema.TypeInt:
		res = "Int"
	case schema.TypeFloat:
		res = "Float64"
	case schema.TypeString:
		res = "String"
	case schema.TypeList:
		res = "List"
	case schema.TypeMap:
		res = "Map"
	case schema.TypeSet:
		res = "Set"
	}
	return res
}

func parseSubtract(v *schema.Schema, subStruct []string) string {
	res := ""
	if v.Type == schema.TypeSet || v.Type == schema.TypeList {
		if len(subStruct) == 0 {
			vv := v.Elem.(*schema.Schema)
			res = parseType(vv)
		}
	}
	return res
}

func formatUsageDesc(s string) string {
	var rr []string
	s = strings.TrimSpace(s)
	m := usageMatch.FindAllStringSubmatch(s, -1)

	for _, v := range m {
		title := strings.TrimSpace(v[1])
		descp := strings.TrimSpace(v[2])

		rr = append(rr, fmt.Sprintf("### %s\n\n%s", title, descp))
	}

	ret := strings.TrimSpace(strings.Join(rr, "\n\n"))
	return ret
}
