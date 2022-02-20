package main

import (
	"Ant-Man-Url/api/middleware"
	"fmt"
)

func main() {

	check, _, _, err := middleware.AuthenticateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJSb2xlIjoidXNlciIsIklkIjozMSwiZXhwIjoxNjQ0ODI2MzIyfQ.PDWnmT2UsWeWZofEE79NN68bUUbZa0sK62FF9P6pTdY")

	if err != nil {
		fmt.Printf("ERR : %s\n", err)
	}
	fmt.Printf("Check : %v", check)
}
