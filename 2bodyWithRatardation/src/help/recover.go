package help

import "fmt"

func RecoverSendWheterOr(ch chan bool) {
	if r := recover(); r != nil {
		ch <- false
	} else {
		ch <- true
	}
}

func RecoverSendErrorIfExist(ch chan error) {
	if err := recover(); err != nil {
		ch <- fmt.Errorf(fmt.Sprint(err))
	}
}
