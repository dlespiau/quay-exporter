package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {

}

func main() {
	quayToken := flag.String("quay-token", "", "Quay.io OAuth 2 Bearer Token")
	flag.Parse()

	repository := NewRegistry()
	if *quayToken != "" {
		repository.SetBearerToken(*quayToken)
	}

	repositories, err := repository.ListRepositories("weaveworks")
	if err != nil {
		log.Fatal(err)
	}

	for _, repository := range repositories {
		fmt.Printf("%s/%s\n", repository.Namespace, repository.Name)
	}
}
