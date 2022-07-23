package flags

import (
	"encoding/json"
	"flag"
	"net/url"
	"os"
)

type Flags struct {
	Port         string
	Db_url       string
	Jaeger_url   string
	Sentry_url   string
	Kafka_broker string
	Some_app_id  string
	Some_app_key string
	Log          string
}

func GetFileflag() (rFlags Flags) {
	data, err := os.Open("conf.json")
	if err != nil {
		rFlags.Log = rFlags.Log + err.Error() + "/n"
	}
	if err = json.NewDecoder(data).Decode(&rFlags); err != nil {
		rFlags.Log = rFlags.Log + err.Error() + "/n"
	}
	return
}
func testMap(r map[string]string) string {
	s := ""
	for k, v := range r {
		if v == "" {
			s += k + ", "
		}
	}
	if s != "" {
		return s + "Not found"
	} else {
		return ""
	}

}
func GetFlag() (rFlags Flags) {
	flag.StringVar(&rFlags.Port, "port", "8080", "Port for database")
	flag.StringVar(&rFlags.Db_url, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "Url for database")
	flag.StringVar(&rFlags.Jaeger_url, "jaeger_url", "http://jaeger:16686", "Какой то егерь и его url..")
	flag.StringVar(&rFlags.Sentry_url, "sentry_url", "http://sentry:9000", "Еще како то url. Лучше не трогать..")
	flag.StringVar(&rFlags.Kafka_broker, "kafka_broker", "kafka:9092", "Какой-то брокер и походу философ")
	flag.StringVar(&rFlags.Some_app_id, "some_app_id", "testid", "Вроде id, но лучше не трогать")
	flag.StringVar(&rFlags.Some_app_key, "some_app_key", "testkey", "Не положите проду пусть дефолт!")
	flag.Parse()
	buf, err := url.Parse(rFlags.Db_url)

	if err != nil {
		rFlags.Log = rFlags.Log + "Db_url: " + err.Error() + "/n"
		rFlags.Db_url = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
	}
	m := make(map[string]string)
	m["path"] = buf.Path
	m["user"] = buf.User.String()
	m["host"] = buf.Host
	m["port"] = buf.Port()
	// ну и здесь можно дописывать параметров для проверки, потом обнулить мапу и проверить остальные флаги
	rFlags.Log = rFlags.Log + testMap(m)
	return
}
