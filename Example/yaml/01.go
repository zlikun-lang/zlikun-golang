package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

// Config is the base data structure.
type Config struct {
	Defaults    DefaultsData          `yaml:"defaults"`
	DataSources map[string]DataSource `yaml:"data-sources"`
}

// DefaultsData defines the possible default values to define.
type DefaultsData struct {
	DataSourceRef     string        `yaml:"data-source"`
	QueryInterval     time.Duration `yaml:"query-interval"`
	QueryTimeout      time.Duration `yaml:"query-timeout"`
	QueryValueOnError string        `yaml:"query-value-on-error"`
}

// DataSource is configuration a data source which must be supported by sql-agent.
type DataSource struct {
	Driver     string                 `yaml:"driver"`
	Properties map[string]interface{} `yaml:"properties"`
}

func main() {
	file := "Example/yaml/01/data.yml"

	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	} else {
		//fmt.Println(string(b))
	}

	var c Config

	// 解析YAML填充Config
	yaml.Unmarshal(b, &c)

	fmt.Println(c)
}
