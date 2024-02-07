package main

import (
	"fmt"
)

type StorageFacade struct {
	account      *Account
	securityCode *SecurityCode
	storage      *Storage
	notification *Notification
}

// Инициализация фасада склада для нового юзера
func newStorageFacade(newUserName string, newCode int) *StorageFacade {
	fmt.Println("Creating new account")
	storageFacade := &StorageFacade{
		account:      newAccount(newUserName),
		securityCode: newSecurityCode(newCode),
		storage:      newStorage(),
		notification: &Notification{},
	}
	fmt.Println("Account created!")
	return storageFacade
}

// Верификация юзера и добавление определенного количества товара юзера на склад
func (s *StorageFacade) addItemToStorage(userName string, code int, items int) error {
	fmt.Println("Starting adding items to users storage account")
	err := s.account.accountCheck(userName)
	if err != nil {
		return err
	}
	err = s.securityCode.securityCodeCheck(code)
	if err != nil {
		return err
	}
	s.storage.addItem(items)
	s.notification.sendStorageAddNotification()
	return nil
}

// Верификация пользователя и вывоз товара со склада
func (s *StorageFacade) takeItemFromStorage(userName string, code int, items int) error {
	fmt.Println("Starting taking items from users storage account")
	err := s.account.accountCheck(userName)
	if err != nil {
		return err
	}
	err = s.securityCode.securityCodeCheck(code)
	if err != nil {
		return err
	}
	s.storage.takeItem(items)
	s.notification.sendStorageTakeItemNotification()
	return nil
}
