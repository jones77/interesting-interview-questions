# Fizz Buzz

Source: https://leetcode.com/problems/fizz-buzz/

Given an integer `n`, return a string array `answer` (1-indexed) where:

- `answer[i] == "FizzBuzz"` if `i` is divisible by 3 and 5.
- `answer[i] == "Fizz"` if `i` is divisible by 3.
- `answer[i] == "Buzz"` if `i` is divisible by 5.
- `answer[i] == i` (as a string) otherwise.

**Example 1:**

```
Input: n = 3
Output: ["1","2","Fizz"]
```

**Example 2:**

```
Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
```

**Example 3:**

```
Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
```

**Constraints:**

- `1 <= n <= 10^4`
<<<<<<< HEAD:questions/easy/fizzbuzz/README.md
=======
*/
func FizzBuzz(n int) []string {
	var list []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			list = append(list, "FizzBuzz")
		} else if i%3 == 0 {
			list = append(list, "Fizz")
		} else if i%5 == 0 {
			list = append(list, "Buzz")
		} else {
			list = append(list, fmt.Sprintf("%d", i))
		}
	}

	return list
}
>>>>>>> fecfad5 (Temp solutions):questions/easy/fizzbuzz/fizzbuzz.go
