package main

import (
	"fmt"
	"log"
)

func main() {

	storageFacade := newStorageFacade("Sam", 123)
	fmt.Printf("%v", storageFacade)

	err := storageFacade.addItemToStorage("Sam", 123, 25)
	if err != nil {
		log.Fatalf("Error: %+v \n", err)
	}

	err = storageFacade.takeItemFromStorage("Sam", 123, 10)
	if err != nil {
		log.Fatalf("Error: %+v \n", err)
	}
}
