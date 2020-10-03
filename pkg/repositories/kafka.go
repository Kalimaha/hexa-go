package repositories

import "fmt"

func Ciao() string {
	return "Ciao!"
}

func Ciaone(yo map[string]interface{}) string {
	fmt.Printf("About to parse %v\n", yo)
	return "Ciaone!"
}
