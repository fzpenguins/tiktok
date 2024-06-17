package config

func ConfigForTest() {
	SqlUserName = "root"
	SqlPassword = "123456"
	DataBase = "testing"
	MysqlIP = "127.0.0.1:3306"

	RabbitmqUserName = "127.0.0.1:5672"
	RabbitmqPassword = "guest"
	RabbitmqIP = "guest"

	ExchangeName = "direct"
	RedisAddr = "127.0.0.1:6379"
	RedisPassword = ""
	RedisDB = 0
	EsAddr = ""

	EndPoint = "127.0.0.1:9000"
	AccessKeyID = "minioadmin"
	SecretAccessKey = "minioadmin"
	SSL = false
	BucketName = "testing"
	EtcdAddr = "127.0.0.1:2379"
	JaegerAddr = "127.0.0.1:6831"
}
