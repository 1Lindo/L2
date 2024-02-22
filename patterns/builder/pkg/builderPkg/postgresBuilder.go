package builderPkg

type PostgresBuilder struct {
	Config     string
	Connection string
	Ping       string
}

func NewPostgresBuilder() *PostgresBuilder {
	return &PostgresBuilder{}
}

func (n *PostgresBuilder) SetConfig() {
	n.Config = "Postgres config"
}

func (n *PostgresBuilder) SetConn() {
	n.Connection = "Postgres connection"
}

func (n *PostgresBuilder) PingDb() {
	n.Ping = "Postgres pinging"
}

func (n *PostgresBuilder) GetDb() Db {
	return Db{
		Config:     n.Config,
		Connection: n.Connection,
		Ping:       n.Ping,
	}
}
