package main

import (
	"os"
	"gopkg.in/yaml.v2"
	"fmt"
)

type Config struct {
        Function []struct {
                Path  string `yaml:"path"`
                R     int    `yaml:"r"`
                G     int    `yaml:"g"`
                B     int    `yaml:"b"`
                Pixel int    `yaml:"pixel"`
        } `yaml:"function"`
}

func processError(err error) {
        fmt.Println(err)
        os.Exit(2)
}

func readFile(cfg *Config) {
        f, err := os.Open("config.yml")
        if err != nil {
                processError(err)
        }
        defer f.Close()

        decoder := yaml.NewDecoder(f)
        err = decoder.Decode(cfg)
        if err != nil {
                processError(err)
        }
}
