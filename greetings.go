package greetings;

import "fmt";

func Hello(name string) string { // exported functions are capitalized
	return fmt.Sprintf("Hi, %v. Welcome!", name);
}