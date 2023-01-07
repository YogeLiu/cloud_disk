package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitPathNotExsit(t *testing.T) {
	asserts := assert.New(t)

	asserts.NotPanics(func() {
		Init("not/exist/path/conf.ini")
	})
}

func TestInitContentError(t *testing.T) {
	asserts := assert.New(t)
	testCase := `[Database]
Type = mysql
User = root
Password233root
Host = 127.0.0.1:3306
Name = v3
TablePrefix = v3_`
	err := ioutil.WriteFile("testConf.toml", []byte(testCase), 0644)
	defer func() { err = os.Remove("testConf.ini") }()
	if err != nil {
		panic(err)
	}
	asserts.Panics(func() {
		Init("testConf.toml")
	})
}

func TestInit(t *testing.T) {
	asserts := assert.New(t)
	testCase := `
		[System]
	Listen = ":3000"
	Debug = false

	[Database]
	Type = "mysql"
	User = "root"
	Password = "root"
	Host = "127.0.0.1:3306"
	Name = "v3"
	`
	err := ioutil.WriteFile("testConf.toml", []byte(testCase), 0644)
	defer func() { err = os.Remove("testConf.toml") }()
	if err != nil {
		panic(err)
	}
	Init("testConf.toml")
	dbConfig := &database{Type: "mysql", User: "root", Password: "root", Host: "127.0.0.1:3306", Name: "v3"}
	std, _ := json.Marshal(dbConfig)
	candidate, _ := json.Marshal(DatabaseConfig)
	asserts.Equal(candidate, std)
}
