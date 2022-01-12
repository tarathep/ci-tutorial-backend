package main

import (
	"fmt"
	"os"
)

func main() {

	// Env Config
	dbconn := os.Getenv("MONGODB_CONNECTION_STRING") // mongodb://root:password@192.168.1.102:27017
	port := os.Getenv("PORT")                        //8089

	// Default config
	if dbconn == "" {
		dbconn = "mongodb://127.0.0.1:27017"
	}
	if port == "" {
		port = "8089"
	}

	fmt.Println(dbconn)
	// db, err := db.Init(dbconn)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// route := router.Router{apis.TutorialHandler{db}}

	// route.Route().Run(":" + port)
}
