package challenge

import "fmt"

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func MakeAnagram(a string, b string) int32 {
	var sameValue = []string{}
	var notMatchValue = []string{}
	var appendValue = a + b

	for i := 0; i < len(a); i++ {
		for x := 0; x < len(b); x++ {
			if a[i] == b[x] {
				sameValue = append(sameValue, string(a[i]))
			}
		}
	}

	rmDuplicates := removeDuplicates(sameValue)

	for q := 0; q < len(rmDuplicates); q++ {

		for p := 0; p < len(appendValue); p++ {
			if rmDuplicates[q] != string(appendValue[p]) {

				for c := 0; c < len(rmDuplicates); c++ {
					if rmDuplicates[c] == string(appendValue[p]) {
						notMatchValue = append(notMatchValue, string(appendValue[p]))
					}
				}
			}

		}
	}

	fmt.Println(rmDuplicates)
	return int32(len(notMatchValue))
}
