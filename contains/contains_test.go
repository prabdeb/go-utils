package contains

import "testing"

func TestStringContains(t *testing.T) {
	if !Contains([]string{"My", "Awesome", "Search", "Util"}, "My") {
		t.Error("String array search did not worked for true case")
	}
	if Contains([]string{"My", "Awesome", "Search", "Util"}, "Hello") {
		t.Error("String array search did not worked for false case")
	}
}

func TestIntContains(t *testing.T) {
	if !Contains([]int{1, 2, 3, 4, 15, 16, 18}, 16) {
		t.Error("Int array search did not worked for true case")
	}
	if Contains([]int{1, 2, 3, 4, 15, 16, 18}, 100) {
		t.Error("Int array search did not worked for false case")
	}
}
