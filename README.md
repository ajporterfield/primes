# primes
A Go package for prime numbers

```go
// Returns the first X prime numbers
primes.First(10) // [2 3 5 7 11 13 17 19 23 29]

// Returns the first X prime numbers starting at an offset
primes.List(10, 100) // [101 103 107 109 113 127 131 137 139 149]

// Returns prime numbers between 0 and X
primes.UpTo(10) // [2 3 5 7]

// Returns prime numbers between two number
primes.Between(10, 20) // [11 13 17 19]

// Returns the next prime number starting at an offset
primes.Next(10) // 11

// Returns the previous prime number starting at an offset
primes.Prev(10) // 7

// Checks if a number is prime or not
primes.IsPrime(11) // true
```