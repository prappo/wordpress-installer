package libs

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	DBHost , DBUser , DBPass string
}

func ConfigStruct() Config {
	return Config{
		DBHost: "localhost",
		DBUser: "root",
		DBPass: "",
	}
}

func MakeConfigFIle(){

	config := ConfigStruct()

	file , _ := json.MarshalIndent(config,""," ")
	_ = ioutil.WriteFile("config.json",file, 0644)

}
