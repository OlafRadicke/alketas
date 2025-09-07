package main

import (
	"log"
	"os"
    "os/exec"
    "gopkg.in/yaml.v2"
)



type conf struct {
    Tokens  []tokenItem `yaml:"tokens"`
}

type tokenItem struct {
    Name string `yaml:"name"`
    Token string `yaml:"token"`
}

func (c *conf) getConf() *conf {

    yamlFile, err := os.ReadFile("tokens.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
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
	log.Print(config)


	for _, v := range config.Tokens {
		log.Printf(v.Name)
		log.Printf("----")
	}


	oldToken:= os.Getenv("VAULT_TEST_TOKEN")
	// log.Print(string(oldToken))

	//  get status / connection check
	cmd := exec.Command("bao", "status")
    stdout, err := cmd.Output()
    if err != nil {
        log.Print(err.Error())
        return
    }
    log.Print("==== STAUS ====\n\n")
    log.Print(string(stdout))

	//  bao token lookup
	cmd = exec.Command("bao", "token", "lookup", oldToken)
    stdout, err = cmd.Output()
    if err != nil {
        log.Print(err.Error())
        return
    }
    log.Print("==== LOOKUP %v ====\n\n", oldToken)
    log.Print(string(stdout))


	//  bao token renew
	cmd = exec.Command("bao", "token", "renew", "-increment=31d", oldToken)
    stdout, err = cmd.Output()
    if err != nil {
        log.Print(err.Error())
        return
    }
    log.Print("==== RENEW %s ====\n\n", oldToken)
    log.Print(string(stdout))

}