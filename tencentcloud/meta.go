package tencentcloud

const (
	TypeInvalid = "invalid"
	TypeBool    = "bool"
	TypeInt     = "int"
	TypeFloat   = "float"
	TypeString  = "string"
	TypeList    = "list"
	TypeMap     = "map"
	TypeSet     = "set"
)

type CNDescription struct {
	TerraformTypeCN string `json:"TerraformTypeCN"`
	DescriptionCN   string `json:"DescriptionCN"`
	AttributesCN    map[string]string
}

// resourceDescriptionProviders resource中文描述提供全局变量
var resourceDescriptionProviders = map[string]CNDescription{}

// dataDescriptionProviders data中文描述提供全局变量
var dataDescriptionProviders = map[string]CNDescription{}

func registerResourceDescriptionProvider(resourceName string, description CNDescription) {
	resourceDescriptionProviders[resourceName] = description
}

func registerDataDescriptionProvider(resourceName string, description CNDescription) {
	dataDescriptionProviders[resourceName] = description
}
