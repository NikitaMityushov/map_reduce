package config

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"local"`
	GRPC GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() (*Config, []string, int) {
	path, chunks, nReduce := fetchArgs()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exists; " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg, chunks, nReduce
}

// PRIORITY: flag > env > default
func fetchArgs() (string, []string, int) {
	var res string
	var lChunks string
	var nReduce int

	flag.StringVar(&res, "config", "", "path to config file")
	flag.StringVar(&lChunks, "chunks", "", "chunks addresses")
	flag.IntVar(&nReduce, "nReduce", 0, "number of reduce tasks")

	flag.Parse()

	var chunks []string
	if lChunks == "" {
		panic("no chunks loaded")
	} else {
		chunks = strings.Fields(lChunks)
	}

	for i, el := range chunks {
		fmt.Printf("%d: %s \n", i, el)
	}

	if res == "" {
		res = os.Getenv("COORDINATOR_CONFIG_PATH")
	}

	return res, chunks, nReduce
}
