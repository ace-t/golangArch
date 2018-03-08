package main

import "fmt"

type AccessBody struct {
	AllAccess []AccessStateItem
}
type AccessStateItem struct {
	Token string
}
func main() {
	aBody := AccessBody{
		AllAccess:[]AccessStateItem{
			{
				Token:"abc-token-123",
			},
		},
	}
	fmt.Println(aBody)
}
