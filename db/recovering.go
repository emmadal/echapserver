package db

import "fmt"


func recoverDB() {
	err := recover()
	if err != nil {
		fmt.Println("Cannot connect to database:", err)
	}
}

func recoverTable() {
	if err := recover(); err != nil {
		fmt.Println("CANNOT CREATE TABLE:", err)
	}
}
// RecoverEnv defer error when .env file is not found
func RecoverEnv() {
	err := recover()
	if err != nil {
		fmt.Println("Unable to load .env file:", err)
	}
}
