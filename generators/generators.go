package generators

import (
	"fmt"
	"math/rand"
	"time"
)

func powerball() {
	draw_numbers, special_number := contest_generator(5, 69, 26)
	fmt.Println("Powerball Numbers")
	printResults(draw_numbers, special_number)
}

func traditional() {
	draw_numbers, special_number := contest_generator(5, 35, 10)
	fmt.Println("Traditional Numbers")
	printResults(draw_numbers, special_number)
}

func pega(pega_num int) {
	draw_numbers, special_number := contest_generator(pega_num, 9, 0)
	fmt.Println("Pega", pega_num)
	printResults(draw_numbers, special_number)
}

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

func printResults(draw_numbers []int, special_number int) {
	fmt.Println("Draw Numbers:", draw_numbers)

	if special_number != 0 {
		fmt.Println("Special Number:", special_number)
	}
	fmt.Println("")
}
