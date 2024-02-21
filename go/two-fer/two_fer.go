// Package twofer was created to determine what to say as I give away the extra cookie
package twofer


import "fmt"

// ShareWith returns a string that shows who gets an extra cookie - my friend (name) or next persan in the line if "name" argument has zero value.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
