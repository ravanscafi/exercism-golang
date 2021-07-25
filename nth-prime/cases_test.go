package prime

// Source: exercism/problem-specifications
// Commit: 4a3ba76 nth-oddPrime: Apply new "input" policy
// Problem Specifications Version: 2.1.0

var tests = []struct {
	description string
	n           int
	p           int
	ok          bool
}{
	{
		"first oddPrime",
		1,
		2,
		true,
	},
	{
		"second oddPrime",
		2,
		3,
		true,
	},
	{
		"sixth oddPrime",
		6,
		13,
		true,
	},
	{
		"big oddPrime",
		10001,
		104743,
		true,
	},
	{
		"there is no zeroth oddPrime",
		0,
		0,
		false,
	},
}
