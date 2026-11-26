package tencentcloud

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"regexp"
	"strings"
	"unicode/utf8"
)

func validateNameRegex(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if _, err := regexp.Compile(value); err != nil {
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid regular expression: %s",
			k, err))
	}
	return
}

func validateNotEmpty(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if value == "" {
		errors = append(errors, fmt.Errorf("%s must not use empty string: %s", k, value))
	}
	return
}

func validateInstanceType(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !strings.Contains(value, ".") {
		errors = append(errors, fmt.Errorf("the format of %s is invalid: %s, it should be like S1.SMALL1", k, value))
		return
	}
	return
}

// NOTE not exactly strict, but ok for now
func validateIntegerInRange(min, max int64) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := int64(v.(int))
		if value < min {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than %d: %d", k, min, value))
		}
		if value > max {
			errors = append(errors, fmt.Errorf(
				"%q cannot be higher than %d: %d", k, max, value))
		}
		return
	}
}

func validateIntegerMin(min int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if value < min {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than %d: %d", k, min, value))
		}
		return
	}
}

func validateStringLengthInRange(min, max int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		length := utf8.RuneCountInString(v.(string))
		if length < min {
			errors = append(errors, fmt.Errorf(
				"length of %q cannot be lower than %d: %d", k, min, length))
		}
		if length > max {
			errors = append(errors, fmt.Errorf(
				"length of %q cannot be higher than %d: %d", k, max, length))
		}
		return
	}
}

func validateAllowedStringValue(ss []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		if !IsContains(ss, value) {
			errors = append(errors, fmt.Errorf("%q must contain a valid string value must in array %#v, got %q", k, ss, value))
		}
		return
	}
}

func validateMysqlPassword(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 64 || len(value) < 8 {
		// errors = append(errors, fmt.Errorf("%s invalid password, len(password) must between 8 and 64,%s", k, value))
		errors = append(errors, fmt.Errorf("the length of %s must be 8-64: %s", k, value))
	}
	var match = make(map[string]bool)
	if strings.ContainsAny(value, "_+-&=!@#$%^*()") {
		match["alien"] = true
	}
	for i := 0; i < len(value); i++ {
		if len(match) >= 2 {
			break
		}
		if value[i] >= '0' && value[i] <= '9' {
			match["number"] = true
			continue
		}
		if (value[i] >= 'a' && value[i] <= 'z') || (value[i] >= 'A' && value[i] <= 'Z') {
			match["letter"] = true
			continue
		}
	}
	if len(match) < 2 {
		errors = append(errors, fmt.Errorf("the format of %s is invalid: %s, it must contain at least letters, Numbers, and characters(_+-&=!@#$%%^*())", k, value))
	}
	return
}

func validateAllowedIntValue(ints []int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if !IsContains(ints, value) {
			errors = append(errors, fmt.Errorf("%q must contain a valid int value in array %#v, got %d", k, ints, value))
		}
		return
	}
}

func validateAsConfigPassword(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 16 || len(value) < 8 {
		errors = append(errors, fmt.Errorf("the length of %s must be 8-16: %s", k, value))
	}
	var match = make(map[string]bool)
	if strings.ContainsAny(value, "()~!@#$%^&*-+={}[]:;',.?/") {
		match["alien"] = true
	}
	for i := 0; i < len(value); i++ {
		if len(match) >= 2 {
			break
		}
		if value[i] >= '0' && value[i] <= '9' {
			match["number"] = true
			continue
		}
		if (value[i] >= 'a' && value[i] <= 'z') || (value[i] >= 'A' && value[i] <= 'Z') {
			match["letter"] = true
			continue
		}
	}
	if len(match) < 2 {
		errors = append(errors, fmt.Errorf("%s is invalid, it must contains at least letters, numbers, and characters(_+-&=!@#$%%^*()): %s", k, value))
	}
	return
}

// check if string has given prefix, if no one prefix matches, errors will have error
func validateStringPrefix(prefix ...string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(string)
		for _, p := range prefix {
			if strings.HasPrefix(value, p) {
				return
			}
		}
		errors = append(errors, fmt.Errorf("%s doesn't have preifx %v", k, prefix))
		return
	}
}

func validateImageID(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !strings.HasPrefix(value, "img-") {
		errors = append(errors, fmt.Errorf("the format of %q is invalid: %s, it should begin with `img-`", k, value))
	}
	return
}

func validateTkeGpuDriverVersion(v interface{}, k string) (ws []string, errors []error) {
	value := v.(map[string]interface{})
	if len(value) > 0 {
		keySet := []string{"name", "version"}
		for _, paraKey := range keySet {
			if value[paraKey] == nil || strings.TrimSpace(value[paraKey].(string)) == "" {
				errors = append(errors, fmt.Errorf("%s in %s cannot be empty", paraKey, k))
			}
		}
	}
	return
}
