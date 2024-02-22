package main

import (
	"fmt"
	"log"
	"main/pkg/facadePkg"
)

func main() {

	storageFacade := facadePkg.NewStorageFacade("Sam", 123)
	fmt.Printf("%v", storageFacade)

	err := storageFacade.AddItemToStorage("Sam", 123, 25)
	if err != nil {
		log.Fatalf("Error: %+v \n", err)
	}

	err = storageFacade.TakeItemFromStorage("Sam", 123, 10)
	if err != nil {
		log.Fatalf("Error: %+v \n", err)
	}
}
