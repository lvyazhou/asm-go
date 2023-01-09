package constservicetype

// ServiceType 服务类型
type ServiceType int8

const (
	ServiceFTP      ServiceType = 1
	ServiceMongodb  ServiceType = 2
	ServiceMssql    ServiceType = 3
	ServiceMysql    ServiceType = 4
	ServiceOracle   ServiceType = 5
	ServicePostgres ServiceType = 6
	ServiceRdp      ServiceType = 7
	ServiceRedis    ServiceType = 8
	ServiceSsh      ServiceType = 9
)

func (e ServiceType) String() string {
	switch e {
	case ServiceFTP:
		return "ftp"
	case ServiceMongodb:
		return "mongodb"
	case ServiceMssql:
		return "mssql"
	case ServiceMysql:
		return "mysql"
	case ServiceOracle:
		return "oracle"
	case ServicePostgres:
		return "postgres"
	case ServiceRdp:
		return "rdp"
	case ServiceRedis:
		return "redis"
	case ServiceSsh:
		return "ssh"
	default:
		return "未知"
	}
}

func (e ServiceType) Val() (val int8) {
	switch e {
	case ServiceFTP:
		return 1
	case ServiceMongodb:
		return 2
	case ServiceMssql:
		return 3
	case ServiceMysql:
		return 4
	case ServiceOracle:
		return 5
	case ServicePostgres:
		return 6
	case ServiceRdp:
		return 7
	case ServiceRedis:
		return 8
	case ServiceSsh:
		return 9
	default:
		return 0
	}
}
