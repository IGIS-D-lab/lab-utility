package db_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

type DSN interface {
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

func TestPopulatePostgres(t *testing.T) {
	// Testing env
	fileLoc := "postgres_example.yaml"
	iam := POSTGRES

	f, err := os.ReadFile(fileLoc)
	if err != nil {
		t.Fatalf("File %s not found\n", fileLoc)
	}

	d := map[string]Postgres{}
	err = yaml.Unmarshal(f, &d)
	if err != nil {
		t.Fatalf("Yaml file to map[string]string failed\n")
	}
	fmt.Println(d[iam].generate())
}
