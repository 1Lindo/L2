package builderPkg

type IBuilder interface {
	SetConfig()
	SetConn()
	PingDb()
	GetDb() Db
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "postgres" {
		return NewPostgresBuilder()
	}

	if builderType == "redis" {
		return NewIglooBuilder()
	}

	return nil
}
