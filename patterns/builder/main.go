package main

import "fmt"

func main() {
	postgresBuilder := getBuilder("postgres")
	redisBuilder := getBuilder("redis")

	director := newDirector(postgresBuilder)
	postgresDb := director.BuildDb()

	fmt.Printf("PostgresDB config[%+v], conn[%+v], ping[%+v] \n", postgresDb.config, postgresDb.connection, postgresDb.ping)

	director.setDirector(redisBuilder)
	redisDB := director.BuildDb()
	fmt.Printf("RedisDB config[%+v], conn[%+v], ping[%+v] \n", redisDB.config, redisDB.connection, redisDB.ping)
}
