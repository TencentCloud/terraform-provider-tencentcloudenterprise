package tencentcloud

const (
	CFS              = "CFS"
	DISK             = "DISK"
	INSTANCE         = "INSTANCE"
	COS              = "COS"
	CSP              = "CSP"
	MySQL_MariaDB    = "MySQL_MariaDB"
	//PostgreSQL       = "PostgreSQL"
	TDSQL_MySQL      = "TDSQL_MySQL"
	//TDSQL_PostgreSQL = "TDSQL_PostgreSQL"
	// filter
	BACKUP_GROUP_ID = "backup-group-id"
	BACKUP_ID       = "backup-id"
)

var BackupResouceTypes = []string{
	CFS,
	COS,
	CSP,
	DISK,
	INSTANCE,
	MySQL_MariaDB,
	TDSQL_MySQL,
}

var CreateResourceTypes = []string{
	COS,
	CSP,
	MySQL_MariaDB,
	TDSQL_MySQL,
}

var ObejectStorage = []string {
	COS,
	CSP,
}

var DataBaseStorage = []string {
	MySQL_MariaDB,
	//PostgreSQL,
	TDSQL_MySQL,
	//TDSQL_PostgreSQL,
}