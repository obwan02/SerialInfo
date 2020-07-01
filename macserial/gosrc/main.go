package main

import "fmt"

func main() {
	err := StartServer("80")
	if err != nil {
		fmt.Println(err.Error())
	}
}
