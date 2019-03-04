package examples

import (
	"fmt"

	"github.com/prabdeb/go-utils/contains"
)

func stringContains() {
	myArray := []string{"My", "Awesome", "Search", "Util"}
	if contains.Contains(myArray, "Hello") {
		fmt.Println("String array found the match")
	}
}
