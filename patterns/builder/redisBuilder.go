package main

type RedisBuilder struct {
	config     string
	connection string
	ping       string
}

func newIglooBuilder() *RedisBuilder {
	return &RedisBuilder{}
}

func (i *RedisBuilder) setConfig() {
	i.config = "Redis config"
}

func (i *RedisBuilder) setConn() {
	i.connection = "Redis connection"
}

func (i *RedisBuilder) pingDb() {
	i.ping = "Redis ping"
}

func (i *RedisBuilder) getDb() Db {
	return Db{
		config:     i.config,
		connection: i.connection,
		ping:       i.ping,
	}
}
