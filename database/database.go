package main

// What i want to implement:
// Students will enter their details in the console and i will save them to the database
// Details will be
/*
1. Name
2. username
3. did you complete level 1?(yes/no)
4. did you complete level 2?(yes/no)
5. did you complete level 3?(yes/no)
// 6. did you complete level 4?(yes/no)
// 7. did you complete level 5?(yes/no)

*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"level_6/database/handlers"
	"level_6/database/models"
	"level_6/database/repository"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	studentHandler *handlers.StudentsHandler
}

// This is a factory function ? That creates and returns a new instance of an App struct
/*
What is a factory function?
 - It follows dependency injection and separation of concerns principles
*/

/*
NOTE:
func NewApp(db *gorm.DB) *App
- `NewApp` -> Its a constructor like function
- it takse `db` which is a type of *grom.DB (a GORM database connection)
- It returns a pointer to an App (*App)
*/

func NewApp(db *gorm.DB) *App {
	studentRepo := repository.NewStudentRepository(db)
	/*
		this calls a constructor function `NewStudentRepository` from the repository package
		- It passes the db to it so the repository can access the database
		- This repository handles all data access for students (like queries, inserts, etc.)
	*/

	studentHandler := handlers.NewStudentsHandler(studentRepo)
	/*
			Calls another constructor function from the handlers package
		- Passes `studentRepo` to it

	*/
	return &App{
		studentHandler: studentHandler,
	}
}

func (a *App) getMenuChoice() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nChoose an operation:")
	fmt.Println("1. Read (show all students)")
	fmt.Println("2. Create (add new student)")
	fmt.Println("3. Update (modify existing student)")
	fmt.Println("4. Delete (remove student)")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice (1-5): ")

	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func (a *App) Run() {
	for {
		choice := a.getMenuChoice()

		switch choice {
		case "1":
			a.studentHandler.GetAllStudents()
		case "2":
			a.studentHandler.CreateStudent()
		case "3":
			a.studentHandler.UpdateStudent()
		case "4":
			a.studentHandler.DeleteStudent()
		case "5":
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Err")
		}
	}
}

func Database(chApp chan *App) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Students{})

	app := NewApp(db)
	fmt.Println("Sending Back the database")
	chApp <- app
}

// func CreateNewStudentFromLocal() {
// }

//	 // In repository/studentRepository.go
//
//	2 func (r *StudentRepository) GetByUsername(username string) (*models.Students, error) {
//	4     var student models.Students
//	5     // Note we are using r.db to build the query
//	6     err := r.db.Where("username = ?", username).First(&student).Error
//	7     if err != nil {
//	8         return nil, err
//	9     }
//
// 10     return &student, nil
// 11 }
//
//	func CheckIfUserAlreadyExist() {
//		u := models.Students
//
//		// Get first matched record
//		// user, err := u.WithContext(ctx).Where(u.Name.Eq("modi")).First()
//		return true
//	}

func CheckIfUserAlreadyExist(h *handlers.StudentsHandler) {
	// students, err := h.Repo.GetAll()
	// fmt.Println(students)
	// fmt.Println(err)
	// var u models.Students
	// Check issue issue@var_def_init obsidian
	u := &models.Students{} 
	defer fmt.Println("Checking if the user already exists")
	user , _ := u.

}

// HACK:

func CheckIfUserAlreadyExist(h *handlers.StudentsHandler) bool {
    defer func() {
        fmt.Println("Checking if the user already exists")
    }()

    // Assuming you have a method to get all users from the repository
    students, err := h.Repo.GetAll()
    if err != nil {
        fmt.Printf("Error fetching students: %v\n", err)
        return false
    }

    // Assuming you want to check if any student with the same ID exists
    for _, s := range students {
        if s.ID == u.ID { // Assuming 'u' is the user object you want to check
            return true
        }
    }

    return false
}

func main() {
	defer fmt.Println("Database Initialized")
	chDatabase := make(chan *App)
	go Database(chDatabase)
	time.Sleep(5 * time.Second)
	thisApp := <-chDatabase
	CheckIfUserAlreadyExist(&handlers.StudentsHandler{})
	thisApp.Run()
}
