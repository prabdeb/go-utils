package cases

import (
	"testing"

	"github.com/prabdeb/go-utils/common"
)

func TestStringContains(t *testing.T) {
	if !common.Contains([]string{"My", "Awesome", "Search", "Util"}, "My") {
		t.Error("String array search did not worked for true case")
	}
	if common.Contains([]string{"My", "Awesome", "Search", "Util"}, "Hello") {
		t.Error("String array search did not worked for false case")
	}
}

func TestIntContains(t *testing.T) {
	if !common.Contains([]int{1, 2, 3, 4, 15, 16, 18}, 16) {
		t.Error("Int array search did not worked for true case")
	}
	if common.Contains([]int{1, 2, 3, 4, 15, 16, 18}, 100) {
		t.Error("Int array search did not worked for false case")
	}
}
