package ratelimit

//default cgi limit

const (
	DefaultLimit int64 = 15
)

func init() {

	//old  (filename . key)
	limitConfig["resource_tc_instance"] = 50
	limitConfig["resource_tc_instance.create"] = 10
	limitConfig["resource_tc_instance.update"] = 10
	limitConfig["resource_tc_instance.delete"] = 10

	//new(filename . action)
	limitConfig["service_tencentcloudenterprise_mysql"] = 50
	limitConfig["service_tencentcloudenterprise_mysql.CreateDBInstanceHour"] = 20
	limitConfig["service_tencentcloudenterprise_mysql.OfflineIsolatedInstances"] = 20
	limitConfig["service_tencentcloudenterprise_mysql.CreateBackup"] = 5
	limitConfig["service_tencentcloudenterprise_mysql.ModifyInstanceParam"] = 20

	//new(filename)
	limitConfig["service_tencentcloudenterprise_cos"] = DefaultLimit
	limitConfig["service_tencentcloudenterprise_vpc"] = DefaultLimit
	limitConfig["service_tencentcloudenterprise_redis"] = DefaultLimit
	limitConfig["service_tencentcloudenterprise_mongodb"] = DefaultLimit
	limitConfig["service_tencentcloudenterprise_dcg"] = DefaultLimit
	limitConfig["service_tencentcloudenterprise_dc"] = 5
	limitConfig["service_tencentcloudenterprise_ccn"] = DefaultLimit
	limitConfig["service_tencentcloudenterprise_cbs"] = DefaultLimit
	limitConfig["service_tencentcloudenterprise_as"] = DefaultLimit
}
