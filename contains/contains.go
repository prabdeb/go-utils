package contains

// Contains can search in string or int Array and does not work for maps
// for maps the below can be used -
//
//if val, ok := dict["foo"]; ok {
//do something here
//}
func Contains(list interface{}, search interface{}) bool {
	switch list := list.(type) {
	case []string:
		switch search := search.(type) {
		case string:
			for _, a := range list {
				if a == search {
					return true
				}
			}
		}
	case []int:
		switch search := search.(type) {
		case int:
			for _, a := range list {
				if a == search {
					return true
				}
			}
		}
	}
	return false
}
