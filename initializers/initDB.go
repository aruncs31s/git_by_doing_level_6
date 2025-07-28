package initializers

import (
	"fmt"
	"time"

	"level_6/database"
)

var App *database.App

func InitDatabase( /*chDBInitialized chan bool*/ ) {
	defer fmt.Println("Database Initialized")
	chDatabase := make(chan *database.App)
	go database.Database(chDatabase)
	time.Sleep(5 * time.Second)
	App = <-chDatabase
	// chDBInitialized <- true
}
