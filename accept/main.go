package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/redis.v2"
)

func redisClient() *redis.Client {
	return redis.NewTCPClient(&redis.Options{
		Addr:     getRedisAddr(),
		Password: getRedisPassword(),
		DB:       0, // use default DB
	})
}

func getRedisAddr() string {
	redis, err := getServices().WithName("redis")
	if err != nil {
		panic(err)
	}
	return redis.Credentials["hostname"] + ":" + redis.Credentials["port"]
}

func getRedisPassword() string {
	redis, err := getServices().WithName("redis")
	if err != nil {
		panic(err)
	}
	return redis.Credentials["password"]
}

func getServices() *cfenv.Services {
	var rawServices map[string]interface{}
	servicesVar := os.Getenv("VCAP_SERVICES")
	if err := json.Unmarshal([]byte(servicesVar), &rawServices); err != nil {
		return nil
	}

	services := make(cfenv.Services)
	for k, v := range rawServices {
		var serviceInstances []cfenv.Service
		if err := mapstructure.Decode(v, &serviceInstances); err != nil {
			return nil
		}
		services[k] = serviceInstances
	}
	return &services
}

type Word struct {
	Word string
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("add.html")
	t.Execute(w, nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	word := &Word{Word: r.FormValue("word")}
	if err := redisClient().RPush("words", word.Word).Err(); err != nil {
		panic(err)
	}
	t, _ := template.ParseFiles("thanks.html")
	t.Execute(w, word)
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)

}
