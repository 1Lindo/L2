package facadePkg

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(newCode int) *SecurityCode {
	return &SecurityCode{
		code: newCode,
	}
}

func (s *SecurityCode) securityCodeCheck(code int) error {
	if s.code != code {
		fmt.Errorf("Wrong security code!")
	}
	fmt.Println("Security code verified!")
	return nil
}
