package config

import (
	"fmt"
	"log"
	"os"

	gcfg "gopkg.in/gcfg.v1"
)

// Config application configuration struct relecting `.ini` file structure
type (
	Config struct {
		Environment string
		Server      struct {
			Port          string
			TracerPort    int
			LocalTimeZone string
		}
		Database struct {
			CoreSlave   string
			CoreMaster  string
			MaxOpenConn int
			MaxIdleConn int
		}
	}
)

// FilePathList list of possible config file relative path to binary location
func FilePathList() []string {
	return []string{
		"/etc/web-personal/",
		"./files/etc/web-personal/",
		"../../files/etc/web-personal/",
		"../../../files/etc/web-personal/",
	}
}

// ReadConfig read `*.ini` configuration file and save it to variable of `*Config` type
func ReadConfig() *Config {
	var (
		cfg     Config
		err     error
		environ string
	)

	environ = os.Getenv("ENVSYS")
	if environ == "" {
		environ = "development"
	}

	path := FilePathList()

	for _, val := range path {
		file := fmt.Sprintf("%sweb-personal.%s.ini", val, environ)
		log.Printf("%s\n", file)
		err := gcfg.ReadFileInto(&cfg, file)
		if err == nil {
			break
		}
	}

	if err != nil {
		log.Fatalf("[web-personal] Cannot load config env:%s :%+v\n ", environ, err)
	}

	log.Printf("[web-personal] Config load success, using \"%s\".\n", environ)
	cfg.Environment = environ

	return &cfg
}
