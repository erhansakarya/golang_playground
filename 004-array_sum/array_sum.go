package main

// Sum sum 5 integers
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {

}
