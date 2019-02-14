package examples

import (
	"fmt"

	"github.com/prabdeb/go-utils/common"
)

func intContains() {
	myArray := []int{1, 2, 3, 4, 15, 16, 18}
	if common.Contains(myArray, 15) {
		fmt.Println("Int array found the match")
	}
}
