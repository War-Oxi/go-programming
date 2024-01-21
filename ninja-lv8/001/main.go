package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserComment  string `json:"user_comment"`
}

func main() {
	user1 := User{
		"kim",
		"taeji",
		"hello",
	}

	user2 := User{
		"ryu",
		"suhyun",
		"hello",
	}

	users := []User{
		user1,
		user2,
	}

	fmt.Println(users)

	usersMarshal, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error ocurred")
		return
	}
	fmt.Println(string(usersMarshal))
}
