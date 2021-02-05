package main

// Repeat repeats the character
func Repeat(character string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += character
	}

	return repeated
}

func main() {
	Repeat("a")
}
