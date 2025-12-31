package day05

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func Sql() {
	_ = initDB()
	//id, _ := InsertUser(User{Name: "xiaoming", Age: 18})
	//_ = UpdateUser(User{Id: 1, Name: "小明", Age: 12})
	users, _ := QueryUsers()
	for _, u := range users {
		fmt.Println(u.Id, u.Name, u.Age)
	}

	fmt.Println("------")
	//u1, _ := QueryUserByID(1)
	//fmt.Println(u1.Id, u1.Name, u1.Age)
	transation, err := db.Begin()
	if err != nil {
		fmt.Println("transation err")

	}
	_ = DeleteUser(1)

	fmt.Println("------")
	transation.Rollback()
	//transation.Commit()
	users1, _ := QueryUsers()
	for _, u := range users1 {
		fmt.Println(u.Id, u.Name, u.Age)
	}
}

var db *sql.DB

func initDB() error {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	return db.Ping()
}

func QueryUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, name, age from user ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []User
	for rows.Next() {

		var u User
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		list = append(list, u)
	}
	return list, rows.Err()
}

func QueryUserByID(id int) (User, error) {
	var u User
	err := db.QueryRow("SELECT id, name, age from user where id = ?", id).Scan(&u.Id, &u.Name, &u.Age)
	return u, err
}
func InsertUser(u User) (int64, error) {
	result, err := db.Exec("INSERT INTO user(name,age) values (? ,?)", u.Name, u.Age)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateUser(u User) error {
	_, err := db.Exec(
		"UPDATE  user SET name =?,age=? where  id =?",
		u.Name, u.Age, u.Id)
	return err
}

func DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM user WHERE id = ?", id)
	return err
}
