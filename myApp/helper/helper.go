package helper

import (
	"example.com/myApp/helper/internal"
)

func Greet() string {

	return internal.AuthHelper() + "Greetings from helper package"
}
