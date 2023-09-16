package array

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersTosum ...[]int) []int {
	var sums []int
	for _, v := range numbersTosum {
		sums = append(sums, Sum(v))
	}

	return sums

}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, v := range numbersToSum {
		if len(v) == 0 {
			sums = append(sums, 0)
		} else {
			tail := v[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
