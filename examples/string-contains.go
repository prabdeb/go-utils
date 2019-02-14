package examples

import (
	"fmt"

	"github.com/prabdeb/go-utils/common"
)

func stringContains() {
	myArray := []string{"My", "Awesome", "Search", "Util"}
	if common.Contains(myArray, "Hello") {
		fmt.Println("String array found the match")
	}
}
