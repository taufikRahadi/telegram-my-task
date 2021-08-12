package main

import "MyTodoList/configs"

func main() {
	configs.InitDB() // create connection to database
	db := configs.DbManager()
	configs.InitApi(db) // init api configurations
}
