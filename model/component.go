package model

import (
	"fmt"
	"github.com/yakhyadabo/go-template/utils"
	"gopkg.in/yaml.v2"
	"os"
)

type Component struct {
	Name string `yaml:"name"`
	Region string `yaml:"region"`
	Environment string `yaml:"environment"`
	HostnamePrefix string `yaml:"hostname-prefix"`
	HostnameSuffix string `yaml:"hostname-suffix"`
	Certificate struct {
		CN string `yaml:"cn"`
		Hosts []string `yaml:"hosts"`
	}
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func GetData(comp *Component) {
	f, err := os.Open("resources/parameters/" + comp.Region + "-" + comp.Environment + ".yml")
	if err != nil {
		processError(err)
	}

	defer utils.CloseFile(f)

	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(comp)
	if err != nil {
		processError(err)
	}
}