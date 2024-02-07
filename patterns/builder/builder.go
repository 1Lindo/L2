package main

type IBuilder interface {
	setConfig()
	setConn()
	pingDb()
	getDb() Db
}

func getBuilder(builderType string) IBuilder {
	if builderType == "postgres" {
		return newPostgresBuilder()
	}

	if builderType == "redis" {
		return newIglooBuilder()
	}

	return nil
}
