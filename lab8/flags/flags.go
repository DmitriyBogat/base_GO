package flags

import "flag"

type Flags struct {
	port         string
	db_url       string
	jaeger_url   string
	sentry_url   string
	kafka_broker string
	some_app_id  string
	some_app_key string
	log          string
}

func getFlag() (R_Flags Flags) {
	flag.StringVar(&R_Flags.db_url, "port", "8080", "Port for database")
	flag.StringVar(&R_Flags.db_url, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Url for database")
	flag.StringVar(&R_Flags.db_url, "jaeger_url", "http://jaeger:16686", "Какой то егерь и его url..")
	flag.StringVar(&R_Flags.db_url, "sentry_url", "http://sentry:9000", "Еще како то url. Лучше не трогать..")
	flag.StringVar(&R_Flags.db_url, "kafka_broker", "kafka:9092", "Какой-то брокер и походу философ")
	flag.StringVar(&R_Flags.db_url, "some_app_id", "testid", "Вроде id, но лучше не трогать")
	flag.StringVar(&R_Flags.db_url, "some_app_key", "testkey", "Не положите проду пусть дефолт!")

	return
}

//port: 8080
//db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
//jaeger_url: http://jaeger:16686
//sentry_url: http://sentry:9000
//kafka_broker: kafka:9092
//some_app_id: testid
//some_app_key: testkey
