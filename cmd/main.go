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

			_age, err := strconv.Atoi(age)

			var emsg string
			if err != nil {
				emsg = "Invalid input type [string] [number]"
			}
			if _age > 200 {
				emsg = "Age must lower than 200"
			}

			if len(emsg) > 0 {
				fmt.Println(emsg)
				os.Exit(0)
			}

			jsondb.AddUser(name, uint8(_age))
		}
	default:
		fmt.Println("Unknown method", action)
	}

}
