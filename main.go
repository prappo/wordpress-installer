package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"wordpress-installer/libs"
)

/*
ToDo
- Make download system
- Make Json file for saving configuration
-
*/
const configFileName = "config.json"

func main() {

	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		libs.MakeConfigFIle()
	}

	configFile, err := ioutil.ReadFile(configFileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	var config libs.Config
	err = json.Unmarshal(configFile, &config)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(libs.WelcomeMessage())
	fmt.Println(config.DBHost)

	//wordPtr := flag.String("name","prappo","Your name")
	//
	//
	//flag.Parse()
	//fmt.Println(*wordPtr)
}
