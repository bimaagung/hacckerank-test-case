package challenge

import (
	"fmt"
)

func AppleOrange(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// distance apple trees fall apple
	distApple := []int32{}

	for _, v := range apples {
		result := a + v
		distApple = append(distApple, result)
	}

	// distance orange trees fall orange
	distOrange := []int32{}

	for _, v := range oranges {
		result := b + v
		distOrange = append(distOrange, result)
	}

    // match apple or orange fall in land on Sam's house
    var countApple int32 
    var countOrange int32

    for _, v := range distApple {
        if v >= s && v <= t {
            countApple += 1
        }
    }

    for _, v := range distOrange {
        if v >= s && v <= t {
            countOrange += 1
        }
    }

	fmt.Println(countApple)
	fmt.Println(countOrange)
}