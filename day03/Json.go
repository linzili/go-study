package day03

import (
	"encoding/json"
	"fmt"
	"os"
)

func Json() {
	user := User{
		ID:   1,
		Name: "张三",
		Age:  18,
	}
	_ = saveToFile(user, "user.json")
	u2, _ := loadFromFile("user.json")
	fmt.Println(u2)

}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func saveToFile(u User, filename string) error {

	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func loadFromFile(filename string) (User, error) {
	var u User
	data, err := os.ReadFile(filename)
	if err != nil {
		return u, err
	}
	err = json.Unmarshal(data, &u)
	return u, err
}
