package primes

import "errors"

func First(count int) []int {
	return List(count, 0)
}

func List(count int, start int) []int {
	primes := make([]int, 0)
	num := start
	for len(primes) < count {
		if IsPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	return primes
}

func Between(start int, end int) []int {
	if start > end {
		start, end = end, start
	}

	primes := make([]int, 0)

	for n := start; n <= end; n++ {
		if IsPrime(n) {
			primes = append(primes, n)
		}
	}

	return primes
}

func UpTo(num int) []int {
	return Between(0, num)
}

func Next(num int) int {
	return List(1, num)[0]
}

func Prev(num int) (int, error) {
	for n := (num - 1); n >= 2; n-- {
		if IsPrime(n) {
			return n, nil
		}
	}

	return 0, errors.New("there are no prime numbers less than 2")
}

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
