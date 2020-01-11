package utils

import "fmt"

type Alert struct {
	Message string
	Type    string
}

func NewAlert(message, alert string) Alert {
	fmt.Println("alert incorrect")
	return Alert{message, alert}
}
