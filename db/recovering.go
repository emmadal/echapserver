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

