package tencentcloud

const (
	NoneTopicType       = -1
	NonePulsarTopicType = -1
)

// 指定死信和重试主题名称规范，LEGACY表示历史命名规则，COMMUNITY表示Pulsar社区命名规范
const (
	PostFixPatternLegacy    = "LEGACY"
	PostFixPatternCommunity = "COMMUNITY"
)

var PostFixPattern = []string{
	PostFixPatternLegacy,
	PostFixPatternCommunity,
}

const (
	VPCRouteAlreadyExist = "ResourceInUse.VpcRoute"
	RabbitMQVipInstanceRunning = 0
	RabbitMQVipInstanceSuccess = 1
)

const (
	AutoRenewFlagTrue = 1
)

// RabbitMQ User Tag
const (
	RabbitMQUserTagAdministrator = "administrator"
	RabbitMQUserTagMonitoring    = "monitoring"
	RabbitMQUserTagPolicymaker   = "policymaker"
	RabbitMQUserTagManagement    = "management"
	RabbitMQUserTagNone          = "none"
)

var RabbitMQUserTags = []string{
	RabbitMQUserTagAdministrator,
	RabbitMQUserTagMonitoring,
	RabbitMQUserTagPolicymaker,
	RabbitMQUserTagManagement,
	RabbitMQUserTagNone,
}