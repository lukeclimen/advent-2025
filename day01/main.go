package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part01()
}

func SpinDialRight(starting_number int, num_clicks int) int {
	truncated_num_clicks := num_clicks % 100
	ending_number := starting_number + truncated_num_clicks
	if ending_number > 100 {
		return ending_number % 100
	}
	return ending_number
}

func SpinDialLeft(starting_number int, num_clicks int) int {
	truncated_num_clicks := num_clicks % 100
	ending_number := starting_number - truncated_num_clicks
	if ending_number < 0 {
		return ending_number + 100
	}
	return ending_number
}

func part01() {
	number_of_zeroes, current_position := 0, 50
	file, err := os.Open("day_1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check rotation direction and distance
		direction := string(line[0])
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if strings.ToLower(direction) == "l" {
			// Reduce current_position, rolling over from 0 to 99
			current_position = SpinDialLeft(current_position, distance)
			
		} else if strings.ToLower(direction) == "r" {
			// Increase current_position, rolling over from 99 to 0
			current_position = SpinDialRight(current_position, distance)
		}

		current_position := current_position % 100
		
		// Increment number_of_zeroes if the value is 0
		if current_position == 0 {
			number_of_zeroes += 1
		}
	}
	fmt.Println(number_of_zeroes)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}