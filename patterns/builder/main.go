package main

import (
	"fmt"
	"main/pkg/builderPkg"
)

func main() {
	postgresBuilder := builderPkg.GetBuilder("postgres")
	redisBuilder := builderPkg.GetBuilder("redis")

	director := builderPkg.NewDirector(postgresBuilder)
	postgresDb := director.BuildDb()

	fmt.Printf("PostgresDB config[%+v], conn[%+v], ping[%+v] \n", postgresDb.Config, postgresDb.Connection, postgresDb.Ping)

	director.SetDirector(redisBuilder)
	redisDB := director.BuildDb()
	fmt.Printf("RedisDB config[%+v], conn[%+v], ping[%+v] \n", redisDB.Config, redisDB.Connection, redisDB.Ping)
}
