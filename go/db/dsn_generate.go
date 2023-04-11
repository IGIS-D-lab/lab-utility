package db

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type DataSourceName interface {
	generate() string
}

type Postgres struct {
	// Database network location
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	// Login user information
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	// Database data location
	DatabaseName string `yaml:"dbname"`
	Schema       string `yaml:"schema"`
	// Additional information
	SslMode string `yaml:"sslmode,omitempty"`
}

const (
	POSTGRES = "postgres"
)

func (p Postgres) generate() string {
	dsnElement := []string{
		fmt.Sprintf("host=%s", p.Host),
		fmt.Sprintf("port=%v", p.Port),
		fmt.Sprintf("user=%s", p.UserName),
		fmt.Sprintf("dbname=%s", p.DatabaseName),
		fmt.Sprintf("password=%s", p.Password),
		fmt.Sprintf("search_path=%s", p.Schema),
		fmt.Sprintf("sslmode=%s", p.SslMode),
	}
	return strings.Join(dsnElement, " ")
}

func New(dbType, config string) string {
	var d DataSourceName
	switch dbType {
	case POSTGRES:
		d = populate[Postgres](dbType, config)
	}
	return d.generate()
}

func populate[T DataSourceName](iam, f string) DataSourceName {
	parsed := map[string]T{}
	d, err := os.ReadFile(f)
	if err != nil {
		log.Panicf("File %s not found\n", f)
	}

	err = yaml.Unmarshal(d, &parsed)
	if err != nil {
		log.Panicf("Yaml file to map[string]T failed\n")
	}
	return parsed[iam]
}
