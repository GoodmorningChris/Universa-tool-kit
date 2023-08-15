package util

import (
	"errors"
)

func CheckUserName(name string) error {

	if len(name) < 2 {
		return errors.New("The name strength is too low")
	} else if len(name) > 16 {
		return errors.New("The name strength is too long")
	} else if name[0] == ' ' || name[len(name)-1] == ' ' {
		return errors.New("The name cannot start or end with Spaces")
	}
	return nil

}
