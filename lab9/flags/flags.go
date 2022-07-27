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
		rFlags.Log = rFlags.Log + err.Error() + "\n"
	}
	if err = json.NewDecoder(data).Decode(&rFlags); err != nil {
		rFlags.Log = rFlags.Log + err.Error() + "\n"
	}
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Db_url, "1111", "Db_url") + "\n"
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Jaeger_url, "0011", "jaeger_url") + "\n"
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Sentry_url, "0011", "Sentry_url") + "\n"
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Kafka_broker, "0010", "Kafka_broker") + "\n"
	return
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
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Db_url, "1111", "Db_url") + "\n"
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Jaeger_url, "0011", "jaeger_url") + "\n"
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Sentry_url, "0011", "Sentry_url") + "\n"
	rFlags.Log = rFlags.Log + parseFlag(rFlags.Kafka_broker, "0010", "Kafka_broker") + "\n"
	////////////////////////////////////
	return
}
func parseFlag(url_s string, check, name string) (res string) {
	buf, err := url.Parse(url_s)
	if err != nil {
		res = res + name + " " + err.Error() + "/n"
		return
	}
	m := make(map[string]string)
	m["name"] = name
	if check[0] == '1' {
		m["user"] = buf.User.String()
	}
	if check[1] == '1' {
		m["host"] = buf.Host
	}
	if check[2] == '1' {
		m["port"] = buf.Port()
	}
	if check[3] == '1' {
		m["path"] = buf.Path
	}
	res = testMap(m)
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
		return "in " + r["name"] + " " + s + "Not found"
	} else {
		return ""
	}

}
