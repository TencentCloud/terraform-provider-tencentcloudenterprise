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
	limitConfig["service_cloud_mysql"] = 50
	limitConfig["service_cloud_mysql.CreateDBInstanceHour"] = 20
	limitConfig["service_cloud_mysql.OfflineIsolatedInstances"] = 20
	limitConfig["service_cloud_mysql.CreateBackup"] = 5
	limitConfig["service_cloud_mysql.ModifyInstanceParam"] = 20

	//new(filename)
	limitConfig["service_cloud_cos"] = DefaultLimit
	limitConfig["service_cloud_vpc"] = DefaultLimit
	limitConfig["service_cloud_redis"] = DefaultLimit
	limitConfig["service_cloud_mongodb"] = DefaultLimit
	limitConfig["service_cloud_dcg"] = DefaultLimit
	limitConfig["service_cloud_dc"] = 5
	limitConfig["service_cloud_ccn"] = DefaultLimit
	limitConfig["service_cloud_cbs"] = DefaultLimit
	limitConfig["service_cloud_as"] = DefaultLimit
}
