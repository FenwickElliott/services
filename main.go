package main

import (
	"fmt"

	"github.com/fenwickelliott/sync"
)

func main() {
	fmt.Println("Welcome to the cookie sync service generator!")
	fmt.Println("We need to ask a few questions about the service to be generated")
	var name, address, port, redirect string
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println("A service name must be provided")
		return
	}
	fmt.Print("Address: ")
	fmt.Scanln(&address)
	if address == "" {
		fmt.Println("A service address must be provided")
		return
	}
	fmt.Print("Port: ")
	fmt.Scanln(&port)
	if port == "" {
		port = "80"
	}
	fmt.Print("Redirect: ")
	fmt.Scanln(&redirect)

	service := sync.Service{
		Name:        name,
		Address:     address,
		Port:        port,
		Redirect:    redirect,
		MongoServer: "cookies.fenwickelliott.io",
	}

	fmt.Println(service)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
