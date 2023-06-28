package main

import "fmt"

func main() {
	error := DBConnection.Connect()
	if error != nil {
		fmt.Println(error)
	}
	StartGrpcServer()
}
