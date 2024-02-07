package main

type PostgresBuilder struct {
	config     string
	connection string
	ping       string
}

func newPostgresBuilder() *PostgresBuilder {
	return &PostgresBuilder{}
}

func (n *PostgresBuilder) setConfig() {
	n.config = "Postgres config"
}

func (n *PostgresBuilder) setConn() {
	n.connection = "Postgres connection"
}

func (n *PostgresBuilder) pingDb() {
	n.ping = "Postgres pinging"
}

func (n *PostgresBuilder) getDb() Db {
	return Db{
		config:     n.config,
		connection: n.connection,
		ping:       n.ping,
	}
}
