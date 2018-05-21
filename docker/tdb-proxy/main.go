package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const TIMEOUT = 10 * time.Second

var configs Configs

type Config struct {
	Pattern string
	Value   string
}

type Configs struct {
	Cfgs []Config `routes`
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("recv proxy req: %+v\n", r.URL.String())
	fmt.Println("counting routes: ", len(configs.Cfgs))

	for _, route := range configs.Cfgs {
		if strings.Contains(r.URL.String(), route.Pattern) {
			fmt.Println("Finding this: ", route.Value)
		}
	}
}

func main() {
	source, err := ioutil.ReadFile("routing.yaml")
	if err != nil {
		fmt.Println("Panic: ", err)
		panic(err)
	}

	yaml.Unmarshal(source, &configs)
	fmt.Printf("--- config:\n%v\n\n", configs)
	fmt.Println("counting routes: ", len(configs.Cfgs))

	s := &http.Server{Addr: ":7000", Handler: http.HandlerFunc(handleReq), ReadTimeout: TIMEOUT, WriteTimeout: TIMEOUT}
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
