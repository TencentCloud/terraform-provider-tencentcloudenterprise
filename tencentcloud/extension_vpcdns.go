package tencentcloud

const (
	VPCDNS_SERVICE_TYPE  = "vpcdns"
	VPCDNS_RESOURCE_TYPE = "privatezone"

	DNS_FORWARD_STATUS_ENABLE   = "ENABLED"
	DNS_FORWARD_STATUS_DISABLED = "DISABLED"
)

var PRIVATE_DNS_FORWARD_STATUS = []string{
	DNS_FORWARD_STATUS_ENABLE,
	DNS_FORWARD_STATUS_DISABLED,
}

// VPCDNS create/update can contend on a shared TLD suffix lock (e.g. a.com and
// b.com both lock "com" while assigning dnsId). Retry these temporary conflicts
// so concurrent Terraform creates do not fail without user-side sequencing.
var vpcdnsSuffixLockRetryableErrors = []string{
	"ResourceInUse",
	"FailedOperation",
	"RequestLimitExceeded",
}
