package main

import (
	"log"
	"sdm_demo_todolist/raw_sql/dbal"
	"sdm_demo_todolist/raw_sql/router"
)

func main() {
	err := dbal.OpenDB()
	if err != nil {
		println(err.Error())
		return
	}
	defer func() {
		_ = dbal.CloseDB()
	}()

	myRouter := router.New()
	// log.Fatal
	// https://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
	// log.Fatal(http.ListenAndServe(":8080", myRouter))
	// https://stackoverflow.com/questions/57354389/how-to-render-static-files-within-gin-router
	log.Fatal(myRouter.Run(":8080"))
}
