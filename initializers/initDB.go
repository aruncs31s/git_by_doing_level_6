package initializers

import (
	"level_6/database"
)

var App *database.App

func InitDatabase( /*chDBInitialized chan bool*/ ) {
	chDatabase := make(chan *database.App)
	go database.Database(chDatabase)
	App = <-chDatabase
	// chDBInitialized <- true
}

// func init() {
// 	fmt.Println("init function is called")
// 	// defer fmt.Println("Database Initialized")
// 	chDatabase := make(chan *database.App)
// 	go database.Database(chDatabase)
// 	// time.Sleep(5 * time.Second)
// 	thisApp := <-chDatabase
// 	thisApp.StudentHandler.CreateLocalStudent()
// }
