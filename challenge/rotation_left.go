package challenge

func RotLeft(a []int32, d int32) []int32 {

	for i := int32(1); i < d+1; i++ {
		firstVal := a[0]

		sliceTemp := a[1:]
		a = append(sliceTemp, firstVal)

	}

	return a
}