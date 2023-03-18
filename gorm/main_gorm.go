package main

import (
	"log"
	"sdm_demo_todolist/gorm/dbal"
	"sdm_demo_todolist/gorm/router"
)

func main() {
	err := dbal.OpenDB()
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		_ = dbal.CloseDB()
	}()

	r := router.New()

	log.Fatal(r.Run(":8080"))
}
