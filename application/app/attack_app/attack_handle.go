package attack_app

import (
	"asm_platform/application/dto"
	constservicetype "asm_platform/infrastructure/pkg/constants/service_type"
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

type BruteThreader struct {
	wg        sync.WaitGroup
	signal    bool
	bruteType string
	bruteFunc BruteFunc
	info      *dto.HostInfo
	timeout   int64
	num       int64
	total     int64
	flag      bool
	success   dto.BruteList
}

// 获取服务链接驱动
func getBruteType(bruteType constservicetype.ServiceType) BruteFunc {
	var conn BruteFunc
	switch bruteType {
	case constservicetype.ServiceFTP:
		conn = &FtpConn{}
	case constservicetype.ServiceNFS:

	case constservicetype.ServiceSamba:

	case constservicetype.ServiceLDAP:

	case constservicetype.ServiceSsh:
		conn = &SshConn{}
	case constservicetype.ServiceRdp:
		conn = &RdpConn{}
	case constservicetype.ServiceTelnet:

	case constservicetype.ServiceVNC:

	case constservicetype.ServiceMongodb:
		conn = &MongoConn{}
	case constservicetype.ServiceMssql:
		conn = &MSSqlConn{}
	case constservicetype.ServiceMysql:
		conn = &MysqlConn{}
	case constservicetype.ServiceOracle:
		conn = &OracleConn{}
	case constservicetype.ServicePostgres:
		conn = &PostgresConn{}
	case constservicetype.ServiceRedis:
		conn = &RedisConn{}

	case constservicetype.ServiceEs:

	case constservicetype.ServiceMemCache:

	case constservicetype.ServiceDB2:

	case constservicetype.ServiceWeb:

	case constservicetype.ServiceWebLogic:

	case constservicetype.ServiceJboss:

	case constservicetype.ServiceWebSphere:

	case constservicetype.ServiceGlassFish:

	case constservicetype.ServiceLotusDominion:

	case constservicetype.ServiceWebMin:

	case constservicetype.ServiceEmail:

	case constservicetype.ServiceSNMP:

	default:
		panic("unknown service type")
	}
	return conn
}

// GetBruteConfig 配置
func GetBruteConfig(bruteType constservicetype.ServiceType, info *dto.HostInfo, timeout int64) *BruteThreader {
	conn := getBruteType(bruteType)
	bt := &BruteThreader{
		bruteType: bruteType.String(),
		bruteFunc: conn,
		info:      info,
		timeout:   timeout,
		num:       0,
		total:     0,
		flag:      false,
	}
	return bt
}

func (t *BruteThreader) Run(thread int, brList <-chan *dto.BruteList) bool {
	Context, cancel := context.WithCancel(context.Background())
	for i := 0; i < thread; i++ {
		t.wg.Add(1)
		go t.worker(brList, Context, cancel)
	}
	t.wg.Wait()
	cancel()
	return t.flag
}

func (t *BruteThreader) GetInfo() string {
	var ret string
	if t.flag {
		ret = fmt.Sprintf("[%s] %s:%s", t.bruteType, t.success.User, t.success.Pass)
	}
	return ret
}
func (t *BruteThreader) GetNum() string {
	ret := fmt.Sprintf("%d", t.num)
	return ret
}

func (t *BruteThreader) GenerateData(users []string, password []string) <-chan *dto.BruteList {
	brList := make(chan *dto.BruteList)
	go func() {
		for _, user := range users {
			for _, pass := range password {
				pass = strings.Replace(pass, "{user}", user, -1)
				brList <- &dto.BruteList{User: user, Pass: pass}
			}
		}
		close(brList)
	}()
	return brList
}

func (t *BruteThreader) worker(brList <-chan *dto.BruteList, Context context.Context, cancel context.CancelFunc) {
	defer t.wg.Done()
	for {
		select {
		case <-Context.Done():
			return
		case one, ok := <-brList:
			if !ok {
				return
			}
			user, pass := one.User, one.Pass
			flag, _ := t.bruteFunc.Attack(t.info, user, pass, t.timeout)
			atomic.AddInt64(&t.num, 1)
			if flag == true {
				t.flag = true
				t.success = *one
				cancel()
			}
		}
	}
}
