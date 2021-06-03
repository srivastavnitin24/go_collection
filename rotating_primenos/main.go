package main

import (
	"math"
	"strconv"
)

func rotatenum(i float64) float64 {
	return math.Mod(i, 10)*(math.Pow(10, math.Floor(math.Log10(i)))) + math.Floor(i/10)
}

func checkIt(num int64, primes []bool) {
	for i := num * 2; i < int64(len(primes)); i += num {
		primes[i] = false
	}
}

// given a number, "rotates" the number by one by moving the last digit into
// the position of the first digit: e.g. 337 -> 733
func main() {
	var num float64 = 25
	primes := make([]bool, int64(num))
	for i := range primes {
		primes[i] = true
	}

	for i := 2; i <= int(math.Sqrt(num)); i++ {
		checkIt(int64(i), primes)
	}

	count := 0
	for i := range primes {
		if i > 1 && primes[i] {
			// this prime might be circular
			candidate := float64(i)
			ok := true
			desc := strconv.FormatInt(int64(i), 10)

			var rotations = int(math.Floor(math.Log10(float64(i))))
			var primenorotate = 0
			for j := 0; j < rotations; j++ {
				candidate = rotatenum(candidate)
				if (candidate < num) && !primes[int(candidate)] {
					ok = false
					break
				} else if candidate < num && candidate != float64(i) {
					primes[int(candidate)] = false
					desc += ", " + strconv.FormatInt(int64(candidate), 10)
					primenorotate++
				}
			}
			if ok {
				println(desc)
				count += (1 + primenorotate)
			}
		}
	}
	println("found", count, "primes")
}
