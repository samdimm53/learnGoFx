package main

import (
	"fmt"

	"example.com/myApp/helper"
	"example.com/myApp/internal/advanced"
	"example.com/myApp/internal/utils"
)

func main() {

	fmt.Println(utils.Help())
	fmt.Println(advanced.InvAdvanced())
	fmt.Print(helper.Greet())

}
