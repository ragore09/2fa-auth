package main

import "fmt"

func main() {
	token := GetTOTP("16DIGITSECRET")
	fmt.Println(token)
}
