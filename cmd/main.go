package main

import (
	"fmt"
	"jsondb"
	"os"
	"strconv"
)

func main() {

	args := os.Args
	action := args[1]

	switch action {
	case "list":
		jsondb.PrintList()
	case "add":
		{
			name := args[2]
			age := args[3]
			_arg, _ := strconv.Atoi(age)

			jsondb.AddUser(name, uint8(_arg))
		}
	default:
		fmt.Println("Unknown method", action)
	}

}
