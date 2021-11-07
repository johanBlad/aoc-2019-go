package day4

import (
	"fmt"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	// 6 digit number XXXYYY
	// Two adjecent digits are the same
	// monotonically non-decreasing
	// range input: 108457-562041
	num := 108457
	validPasswords := 0
	for num < 562041 {
		num++
		if isValid2(strconv.Itoa(num)) {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
	// fmt.Println(isValid2("112233"))
	// fmt.Println(isValid2("123444"))
	// fmt.Println(isValid2("111122"))
}

func isValid(num string) bool {
	hasAjacent := false
	for i := 1; i < len(num); i++ {
		prev, err := strconv.Atoi(string(num[i-1]))
		check(err)
		next, err := strconv.Atoi(string(num[i]))
		check(err)
		if prev == next {
			hasAjacent = true
		} else if prev > next {
			return false
		}
	}
	return hasAjacent
}

func isEqual(base int, num string, i int) bool {
	comparator, err := strconv.Atoi(string(num[i]))
	check(err)
	return comparator == base
}

func isValid2(num string) bool {
	valid := false
	for i := 1; i < len(num); i++ {
		prev, err := strconv.Atoi(string(num[i-1]))
		check(err)
		next, err := strconv.Atoi(string(num[i]))
		check(err)
		if prev == next {
			if i == 1 && !isEqual(next, num, i+1) {
				valid = true
			} else if i == len(num)-1 && !isEqual(prev, num, i-2) {
				valid = true
			} else if 1 < i && i < len(num)-1 && !isEqual(next, num, i+1) && !isEqual(prev, num, i-2) {
				valid = true
			}
		} else if prev > next {
			return false
		}
	}
	return valid
}
