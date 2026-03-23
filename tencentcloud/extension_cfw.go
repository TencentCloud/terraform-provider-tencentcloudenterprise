package tencentcloud

const (
	BAND_WIDTH = 20

	MODE_0 = 0
	MODE_1 = 1

	CROSS_A_ZONE_0 = 0
	CROSS_A_ZONE_1 = 1

	SWITCH_MODE_1 = 1
	SWITCH_MODE_2 = 2
	SWITCH_MODE_4 = 4

	POLICY_RULE_ACTION_ACCEPT = "accept"
	POLICY_RULE_ACTION_DROP   = "drop"
	POLICY_RULE_ACTION_LOG    = "log"

	POLICY_ENABLE_TRUE  = "true"
	POLICY_ENABLE_FALSE = "false"

	DIRECTION_0 = "0"
	DIRECTION_1 = "1"
	DIRECTION_3 = "3"

	RULE_TYPE_1 = 1
	RULE_TYPE_2 = 2
)

var MODE = []int{
	MODE_0,
	MODE_1,
}

var CROSS_A_ZONE = []int{
	CROSS_A_ZONE_0,
	CROSS_A_ZONE_1,
}

var SWITCH_MODE = []int{
	SWITCH_MODE_1,
	SWITCH_MODE_2,
	SWITCH_MODE_4,
}

var POLICY_RULE_ACTION = []string{
	POLICY_RULE_ACTION_ACCEPT,
	POLICY_RULE_ACTION_DROP,
	POLICY_RULE_ACTION_LOG,
}

var POLICY_ENABLE = []string{
	POLICY_ENABLE_TRUE,
	POLICY_ENABLE_FALSE,
}

var DIRECTION = []string{
	DIRECTION_0,
	DIRECTION_1,
	DIRECTION_3,
}

var RULE_TYPE = []int{
	RULE_TYPE_1,
	RULE_TYPE_2,
}

const (
	ADDRESS_TEMPLATE_TYPE_1 = 1
	ADDRESS_TEMPLATE_TYPE_5 = 5

	POLICY_SCOPE_SERIAL = "serial"
	POLICY_SCOPE_SIDE   = "side"
	POLICY_SCOPE_ALL    = "all"

	FW_TYPE_NAT = "nat"
	FW_TYPE_EW  = "ew"
)

var ADDRESS_TEMPLATE_TYPE = []int{
	ADDRESS_TEMPLATE_TYPE_1,
	ADDRESS_TEMPLATE_TYPE_5,
}

var POLICY_SCOPE = []string{
	POLICY_SCOPE_SERIAL,
	POLICY_SCOPE_SIDE,
	POLICY_SCOPE_ALL,
}

var FW_TYPE = []string{
	FW_TYPE_NAT,
	FW_TYPE_EW,
}

type SourceContentJson struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type TargetContentJson struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
