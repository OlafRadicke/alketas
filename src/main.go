package main

import (
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Tokens []tokenItem `yaml:"tokens"`
}

type tokenItem struct {
	Name  string `yaml:"name"`
	Token string `yaml:"token"`
}

func (c *conf) getConf() *conf {

	renewTokens := os.Getenv("VAULT_RENEW_TOKENS")
	log.Printf("Try to open file %v", renewTokens)
	yamlFile, err := os.ReadFile(renewTokens)
	if err != nil {
		log.Printf("File err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var config conf
	config.getConf()

	//  bao token lookup
	log.Print("==== LOOKUP ====\n\n")
	for _, v := range config.Tokens {
		log.Printf("Lookup token: %v", v.Name)
		os.Setenv("VAULT_TOKEN", v.Token)
		cmd := exec.Command("bao", "token", "lookup", v.Token)
		stdout, err := cmd.Output()
		if err != nil {
			log.Printf("Lookup error: %v", err.Error())
		}
		log.Printf(string(stdout))
		log.Printf("----")
	}

	//  bao token renew
	log.Print("==== RENEW ====\n\n")
	for _, v := range config.Tokens {
		log.Printf("Renew token: %v", v.Name)
		os.Setenv("VAULT_TOKEN", v.Token)
		cmd := exec.Command("bao", "token", "renew", "-increment=31d")
		stdout, err := cmd.Output()
		if err != nil {
			log.Fatalf("Renew error: %v", err.Error())
			os.Exit(1)
		}
		log.Printf(string(stdout))
		log.Printf("----")
	}

}
