package db

import "log"

func recoverDB() {
	err := recover()
	if err != nil {
		log.Fatalln("Cannot connect to database:", err)
	}
}

func recoverTable() {
	if err := recover(); err != nil {
		log.Fatalln("CANNOT CREATE TABLE:", err)
	}
}

// RecoverEnv defer error when .env file is not found
func RecoverEnv() {
	err := recover()
	if err != nil {
		log.Fatalln("Unable to load .env file:", err)
	}
}
