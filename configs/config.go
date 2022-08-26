package configs

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var data map[string]any
var ret gjson.Result

func Setup(conf_path string) {
	profile := os.Getenv("ENV")
	if profile == "" {
		profile = "dev"
	}
	log.Println("profile is:", profile)

	file, err := os.Open(conf_path)
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}
	dec := yaml.NewDecoder(file)
	err = dec.Decode(&data)

	for {
		var t map[string]interface{}
		err = dec.Decode(&t)
		if err != nil {
			break
		}
		if p, ok := t["profile"]; ok {
			if p == profile {
				data = merge(data, t)
				break
			}
		}
	}
	v, er := json.Marshal(&data)
	if er != nil {
		log.Println(er)
	}
	log.Println(string(v))
	ret = gjson.Parse(string(v))
}

func Get(path ...string) gjson.Result {
	var r gjson.Result
	for _, p := range path {
		if r.Value() == nil {
			r = ret.Get(p)
		} else {
			r = r.Get(p)
		}
	}
	return r
}

func merge(m1 map[string]any, m2 map[string]any) map[string]any {
	n := make(map[string]any)
	for k, v := range m1 {
		n[k] = v
	}
	for k, v := range m2 {
		n[k] = v
	}
	return n
}
