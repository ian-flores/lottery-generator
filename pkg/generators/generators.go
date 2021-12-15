package generators

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// Powerball generates the 5 numbers and the 1 special number for the powerball draw
func Powerball() {
	draw_numbers, special_number := contest_generator(5, 69, 26)
	color.Cyan("Powerball Numbers")
	printResults(draw_numbers, special_number)
}

// Traditional generates the 5 numbers and the 1 special number for the
// traditional draw
func Traditional() {
	draw_numbers, special_number := contest_generator(5, 35, 10)
	color.Cyan("Traditional Numbers")
	printResults(draw_numbers, special_number)
}

// Pega generates the different sets of numbers for Pega 2, Pega 3 and Pega 4
func Pega(pega_num int) {
	draw_numbers, special_number := contest_generator(pega_num, 9, 0)
	fmt.Println("Pega", pega_num)
	printResults(draw_numbers, special_number)
}

// contest_generator generates the numbers for any random contest. It is the
// main function of the package.
func contest_generator(draw_numbers_size int, draw_number_max int, special_number_max int) ([]int, int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	numbers := make([]int, draw_numbers_size)

	for i := 0; i < draw_numbers_size; i++ {
		numbers[i] = r1.Intn(draw_number_max) + 1
	}

	if special_number_max > 0 {
		return numbers, r1.Intn(special_number_max) + 1
	}

	return numbers, 0
}

// printResults prints the results of the contest
func printResults(draw_numbers []int, special_number int) {
	fmt.Println("Draw Numbers:", draw_numbers)

	if special_number != 0 {
		fmt.Println("Special Number:", special_number)
	}
	fmt.Println("")
}
