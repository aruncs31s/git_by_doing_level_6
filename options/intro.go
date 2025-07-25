package options

import "fmt"

func Intro(theirUserName string) {
	/*

	  func DisplayStudents(students []Student) {
	    fmt.Println"📚 Student Details:")

	   for _, student := range students {
	            fmt.Print"   Name: %s\n", student.Name)
	            fmt.Print"   username: %s\n", student.Username)
	    }
	   }
	*/
	fmt.Println("   📖 Intro")
	fmt.Printf("   Hello %v\n", theirUserName)
	fmt.Printf("Possible Commands are ")
	fmt.Printf("   play: attend the level\n")
	fmt.Printf("   check: check if everything is working ")
	fmt.Printf("   show errors")
}
