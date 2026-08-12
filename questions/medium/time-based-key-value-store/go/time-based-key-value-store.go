package main

// TimeMap solves Time Based Key-Value Store. See README.md for the problem statement.
type TimeMap struct{}

func Constructor() *TimeMap {
	return &TimeMap{}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {}

func (t *TimeMap) Get(key string, timestamp int) string {
	return ""
}

func main() {}
