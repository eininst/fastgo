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
	buffer, err := os.ReadFile(conf_path)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(buffer, &data)
	if err != nil {
		log.Fatal(err)
	}

	newValue, er := json.Marshal(&data)
	if er != nil {
		log.Println(er)
	}
	ret = gjson.Parse(string(newValue))
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
