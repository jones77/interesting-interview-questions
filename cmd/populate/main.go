package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Question struct {
	ShortName    string
	Title        string
	QuestionText string
	Solution     string
	Hint         string
	Difficulty   string
	Category     string
}

func main() {
	db, err := sql.Open("sqlite", "interview-questions.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec(`DROP TABLE IF EXISTS questions`); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(`
		CREATE TABLE questions (
			id            INTEGER PRIMARY KEY,
			short_name    TEXT NOT NULL UNIQUE,
			title         TEXT NOT NULL,
			question_text TEXT NOT NULL,
			solution      TEXT NOT NULL,
			hint          TEXT,
			difficulty    TEXT NOT NULL DEFAULT 'medium',
			category      TEXT NOT NULL DEFAULT 'general'
		)
	`); err != nil {
		log.Fatal(err)
	}

	questions := seedData()

	stmt, err := db.Prepare(`INSERT INTO questions (short_name, title, question_text, solution, hint, difficulty, category) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}

	inserted := 0
	for _, q := range questions {
		if _, err := stmt.Exec(q.ShortName, q.Title, q.QuestionText, q.Solution, q.Hint, q.Difficulty, q.Category); err != nil {
			log.Printf("SKIP %q: %v", q.ShortName, err)
			continue
		}
		inserted++
	}

	log.Printf("inserted %d questions", inserted)
}

func seedData() []Question {
	return []Question{
		{
			ShortName: "two-sum",
			Title:     "Two Sum",
			QuestionText: `Given an array of integers ` + "`nums`" + ` and an integer ` + "`target`" + `, return indices of the two numbers that add up to ` + "`target`" + `.

You may assume that each input has exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

**Example 1:**

` + "```" + `
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
` + "```" + `

**Example 2:**

` + "```" + `
Input: nums = [3,2,4], target = 6
Output: [1,2]
` + "```" + `

**Example 3:**

` + "```" + `
Input: nums = [3,3], target = 6
Output: [0,1]
` + "```" + `

**Constraints:**

- ` + "`2 <= nums.length <= 10^4`" + `
- ` + "`-10^9 <= nums[i] <= 10^9`" + `
- ` + "`-10^9 <= target <= 10^9`" + `
- Only one valid answer exists.

**Follow-up:** Can you come up with an algorithm that is less than O(n²) time complexity?`,
			Solution: `func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, v := range nums {
        if j, ok := m[target-v]; ok {
            return []int{j, i}
        }
        m[v] = i
    }
    return nil
}`,
			Difficulty: "easy",
			Category:   "arrays",
		},
		{
			ShortName: "fizzbuzz",
			Title:     "Fizz Buzz",
			QuestionText: `Given an integer ` + "`n`" + `, return a string array ` + "`answer`" + ` (1-indexed) where:

- ` + "`answer[i] == \"FizzBuzz\"`" + ` if ` + "`i`" + ` is divisible by 3 and 5.
- ` + "`answer[i] == \"Fizz\"`" + ` if ` + "`i`" + ` is divisible by 3.
- ` + "`answer[i] == \"Buzz\"`" + ` if ` + "`i`" + ` is divisible by 5.
- ` + "`answer[i] == i`" + ` (as a string) otherwise.

**Example 1:**

` + "```" + `
Input: n = 3
Output: ["1","2","Fizz"]
` + "```" + `

**Example 2:**

` + "```" + `
Input: n = 5
Output: ["1","2","Fizz","4","Buzz"]
` + "```" + `

**Example 3:**

` + "```" + `
Input: n = 15
Output: ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
` + "```" + `

**Constraints:**

- ` + "`1 <= n <= 10^4`" + ``,
			Solution: `func fizzBuzz(n int) []string {
    ans := make([]string, n)
    for i := 1; i <= n; i++ {
        switch {
        case i%15 == 0:
            ans[i-1] = "FizzBuzz"
        case i%3 == 0:
            ans[i-1] = "Fizz"
        case i%5 == 0:
            ans[i-1] = "Buzz"
        default:
            ans[i-1] = strconv.Itoa(i)
        }
    }
    return ans
}`,
			Difficulty: "easy",
			Category:   "general",
		},
	}
}
