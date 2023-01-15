// Greetings package
package greetings

import "fmt"

// hello return a hello to a person
func hello(name string) string {
	message := fmt.Sprintf("hi, %v.", name)
	return message
}
