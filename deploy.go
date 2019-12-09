package main

import (
	"fmt"
	"github.com/yakhyadabo/go-template/model"
	"github.com/yakhyadabo/go-template/utils"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Print("Region : ")
	var region, env string
	_, err := fmt.Scanln(&region)

	fmt.Print("Env : ")
	_, err = fmt.Scanln(&env)
	fmt.Println(env)

	tpl, err := template.ParseGlob("resources/templates/cert.json")
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("resources/accounts/" + region + "/"+ env + "/cert.json")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	defer utils.CloseFile(output)

	var comp model.Component

	comp.Region = region
	comp.Environment = env

	model.GetData(&comp)
	fmt.Println("%+v", comp)

	err = tpl.Execute(output, comp)
	if err != nil {
		panic(err)
	}

}


// https://dev.to/ilyakaznacheev/a-clean-way-to-pass-configs-in-a-go-application-1g64

// Yaml: https://github.com/go-yaml/yaml