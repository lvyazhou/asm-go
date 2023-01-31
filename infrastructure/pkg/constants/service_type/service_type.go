package constservicetype

// ServiceType 服务类型
type ServiceType int8

const (
	// ServiceFTP 文件共享服务端口
	ServiceFTP   ServiceType = 1  // 21/22/69 FTP/TFTP 文件传输协议
	ServiceNFS   ServiceType = 11 // 2049 NFS 服务
	ServiceSamba ServiceType = 12 // 139 Samba 服务 实现SMB协议的一个免费软件
	ServiceLDAP  ServiceType = 13 // 389 LDAP 目录访问协

	// ServiceSsh 远程连接服务端口
	ServiceSsh    ServiceType = 2
	ServiceRdp    ServiceType = 21
	ServiceTelnet ServiceType = 22
	ServiceVNC    ServiceType = 23

	// ServiceMongodb 数据库端口服务
	ServiceMongodb  ServiceType = 3  // 27017/27018
	ServiceMssql    ServiceType = 31 // 1433
	ServiceMysql    ServiceType = 32 // 3306
	ServiceOracle   ServiceType = 33 // 1521
	ServicePostgres ServiceType = 34 // 5432
	ServiceRedis    ServiceType = 35 // 6379
	ServiceEs       ServiceType = 36 // 9200/9300
	ServiceMemCache ServiceType = 37 // 11211
	ServiceDB2      ServiceType = 38 // 5000

	// ServiceWeb web应用服务端口
	ServiceWeb           ServiceType = 4  // 80/443/8080
	ServiceWebLogic      ServiceType = 41 // 7001/7002
	ServiceJboss         ServiceType = 42 // 8080/8089
	ServiceWebSphere     ServiceType = 43 // 9090
	ServiceGlassFish     ServiceType = 44 // 4848
	ServiceLotusDominion ServiceType = 45 // Lotus dominion邮件服务 1352
	ServiceWebMin        ServiceType = 46 // Webmin-Web控制面板	 10000

	// ServiceEmail 邮件服务
	ServiceEmail ServiceType = 5 // POP3-110/IMAP-143

	// ServiceSNMP 网络常见协议端口
	ServiceSNMP ServiceType = 6
)

func (e ServiceType) String() string {
	switch e {
	case ServiceFTP:
		return "ftp"
	case ServiceNFS:
		return "nfs"
	case ServiceSamba:
		return "samba"
	case ServiceLDAP:
		return "ldap"

	case ServiceSsh:
		return "ssh"
	case ServiceRdp:
		return "rdp"
	case ServiceVNC:
		return "vnc"
	case ServiceTelnet:
		return "telnet"

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
	case ServiceRedis:
		return "redis"
	case ServiceEs:
		return "es"
	case ServiceMemCache:
		return "memcache"
	case ServiceDB2:
		return "SysBase/DB2"

	case ServiceWeb:
		return "web"
	case ServiceWebLogic:
		return "weblogic"
	case ServiceJboss:
		return "jboss"
	case ServiceWebSphere:
		return "websphere"
	case ServiceGlassFish:
		return "glassfish"
	case ServiceLotusDominion:
		return "lotus"
	case ServiceWebMin:
		return "webmin"

	case ServiceEmail:
		return "email"

	case ServiceSNMP:
		return "snmp"

	default:
		return "未知"
	}
}

func (e ServiceType) Val() (val int8) {
	switch e {
	case ServiceFTP:
		return 1
	case ServiceNFS:
		return 11
	case ServiceSamba:
		return 12
	case ServiceLDAP:
		return 13

	case ServiceSsh:
		return 2
	case ServiceRdp:
		return 21
	case ServiceTelnet:
		return 22
	case ServiceVNC:
		return 23

	case ServiceMongodb:
		return 3
	case ServiceMssql:
		return 32
	case ServiceMysql:
		return 32
	case ServiceOracle:
		return 33
	case ServicePostgres:
		return 34
	case ServiceRedis:
		return 35
	case ServiceEs:
		return 36
	case ServiceMemCache:
		return 37
	case ServiceDB2:
		return 38

	case ServiceWeb:
		return 4
	case ServiceWebLogic:
		return 41
	case ServiceJboss:
		return 42
	case ServiceWebSphere:
		return 43
	case ServiceGlassFish:
		return 44
	case ServiceLotusDominion:
		return 45
	case ServiceWebMin:
		return 46

	case ServiceEmail:
		return 5

	case ServiceSNMP:
		return 6
	default:
		return 0
	}
}
