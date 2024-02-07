package main

import "fmt"

type Notification struct {
}

func (n *Notification) sendStorageAddNotification() {
	fmt.Println("Sending storage \"add\" notification")
}

func (n *Notification) sendStorageTakeItemNotification() {
	fmt.Println("Sending storage \"take\" notification")
}
