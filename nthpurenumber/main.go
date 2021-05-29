/*It should use the only digit 4 or 5.
It must have even number of digits.
It should be a palindrome.
Let’s take a look at some of the examples:

N=1 result would be “44”.
N=2 result would be “55”.
N=2 result would be “4444”.
N=2 result would be “4554”.
N=2 result would be “5445”.
N=2 result would be “5555”.
Looking at the pattern we can see New Pure number is a combination of what we get after appending “4”
at the beginning and the last or appending “5” at the starting and the end. “4444” can be derived using
“4” + “44” + “4” , “4554” is derived using “4”+”55“+”4”, “5445” is derived using “5”+”44“+”5” and “5555”
using “5”+”55“+”5” where bolded numbers are previous pure numbers.

So, in each iteration i, we get new power(2, i) pure number using previously power(2, i-1) pure numbers
by appending either “4” or “5” at both ends. We maintain a counter named as count that stops once we reach
the Nth pure number.*/
package nthpurenumber

import "math"

func nthPureNumber(n int) string {
	arr := make([]string, n+1)
	arr[0] = "44"
	arr[1] = "55"

	count, val, i := 2, -1, 0

	for count < n {
		i++
		x := val + 1
		val = int(float64(x) + (math.Pow(2, float64(i)) - 1))

		for k := x; k <= val; k++ {
			arr[count] = "4" + arr[k] + "4"
			count++
			if count >= n {
				break
			}
		}
		for k := x; k <= val; k++ {
			arr[count] = "5" + arr[k] + "5"
			count++
			if count >= n {
				break
			}
		}
	}
	return arr[n-1]
}
