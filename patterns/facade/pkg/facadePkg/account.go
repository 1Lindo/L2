package facadePkg

import "fmt"

type Account struct {
	userName string
}

func newAccount(newUserName string) *Account {
	return &Account{userName: newUserName}
}

func (a *Account) accountCheck(newUserName string) error {
	if newUserName != a.userName {
		fmt.Errorf("Wrong user name!")
	}
	fmt.Println("Account verified!")
	return nil
}
