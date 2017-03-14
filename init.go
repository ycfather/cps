package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	DEFAULT_PAGE_SIZE = 10
	DEFAULT_PAGE      = 1
)

var (
	rsaPublicKey  []byte
	rsaPrivateKey []byte
	cpsConfig     *CpsConfig = &CpsConfig{}
	db            *sql.DB
)

type CpsConfig struct {
	Server struct {
		port uint16
	}

	Encrypt struct {
		Rsa struct {
			PrivatePem string `yaml:"private_pem"`
			PublicPem  string `yaml:"public_pem"`
		}
	}

	Sso struct {
		Url  string
		Site string
	}

	Datasource struct {
		Host     string
		Port     uint16
		Db       string
		Username string
		Password string
		Params   string
	}
}

type JsonResult struct {
	Status uint16 `json:"status"`
	Data   struct {
		List struct {
			Total       uint32        `json:"total"`
			PerPage     uint8         `json:"per_page"`
			CurrentPage uint32        `json:"current_page"`
			Data        []interface{} `json:"data"`
		} `json:"list"`
	} `json:"data"`
}

func ParseCpsConfig() {
	buf, _ := ioutil.ReadFile("config/cps.yaml")
	yaml.Unmarshal(buf, cpsConfig)
	fmt.Printf("config : %+v\n", cpsConfig)
}

func init() {
	ParseCpsConfig()
	rsaPrivateKey, _ = ioutil.ReadFile(cpsConfig.Encrypt.Rsa.PrivatePem)
	rsaPublicKey, _ = ioutil.ReadFile(cpsConfig.Encrypt.Rsa.PublicPem)

	datasourceName := []string{
		cpsConfig.Datasource.Username,
		":",
		cpsConfig.Datasource.Password,
		"@tcp(",
		cpsConfig.Datasource.Host,
		":",
		strconv.FormatUint(uint64(cpsConfig.Datasource.Port), 10),
		")/",
		cpsConfig.Datasource.Db,
		"?",
		cpsConfig.Datasource.Params,
	}
	db, _ = sql.Open("mysql", strings.Join(datasourceName, ""))
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(20)
	db.Ping()
}
