package twofer

import "fmt"

// ShareWith receives a name and return a sentence with the name
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
