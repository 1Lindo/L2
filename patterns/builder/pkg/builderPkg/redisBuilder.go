package builderPkg

type RedisBuilder struct {
	Config     string
	Connection string
	Ping       string
}

func NewIglooBuilder() *RedisBuilder {
	return &RedisBuilder{}
}

func (i *RedisBuilder) SetConfig() {
	i.Config = "Redis config"
}

func (i *RedisBuilder) SetConn() {
	i.Connection = "Redis connection"
}

func (i *RedisBuilder) PingDb() {
	i.Ping = "Redis ping"
}

func (i *RedisBuilder) GetDb() Db {
	return Db{
		Config:     i.Config,
		Connection: i.Connection,
		Ping:       i.Ping,
	}
}
