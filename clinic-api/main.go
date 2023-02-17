package main

import (
	"clinic-api/api/telegram"
	"clinic-api/repository/sqlite"
	"clinic-api/service"
)

func main() {

	db, err := sqlite.NewDB("sqlite.db")
	if err != nil {
		panic(err)
	}

	repo := sqlite.NewRepository(db)

	service := service.NewService(repo)

	tgBot, err := telegram.New(service)
	if err != nil {
		panic(err)
	}

	err = tgBot.Init()
	if err != nil {
		panic(err)
	}

}
