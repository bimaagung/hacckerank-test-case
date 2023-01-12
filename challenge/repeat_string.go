package challenge

import "strings"

func RepeadString(s string, n int64) int64 {
	size := len(s)
	qtFirst := int64(0)
	res := int64(0)

	// check the qt first letter in the string
	c := strings.Split(s, "")
	pos := -1

	for i := 0; i < len(c); i++ {
		if c[i] == "a" {
			if pos == -1 {
				pos = i
			}
			qtFirst++
		}
	}

	if pos == -1 {
		return res
	}

	// calc how many times it will be repeated
	mod := n % int64(size)
	total := n / int64(size)

	if mod == 0 {
		res = total * qtFirst
	} else {
		res = total * qtFirst
		for i := int64(0); i < mod; i++ {
			if c[i] == "a" {
				res++
			}
		}
	}

	return int64(res)
}