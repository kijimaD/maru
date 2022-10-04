package utils

import (
	"fmt"
)

func CheckErr(err error) (hasErr bool) {
	hasErr = err == nil
	if err != nil {
		fmt.Println(err)
	}
	return hasErr
}
