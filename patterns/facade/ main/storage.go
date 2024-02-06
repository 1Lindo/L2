package main

import "fmt"

type Storage struct {
	items int
}

func newStorage() *Storage {
	return &Storage{
		items: 0,
	}
}

func (s *Storage) storageCheck() {
	fmt.Printf("%v", s.items)
}

func (s *Storage) addItem(items int) {
	s.items += items
	fmt.Printf("%v items have been added!", s.items)
}

func (s *Storage) takeItem(items int) error {
	if s.items < items {
		return fmt.Errorf("Storage is not sufficient")
	}
	fmt.Println("Storage is sufficient")
	s.items = s.items - items
	fmt.Printf("%v items have been deleted!", s.items)
	return nil
}
