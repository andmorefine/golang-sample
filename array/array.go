package array

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, num := range numbersToSum {
		sums = append(sums, Sum(num))
	}
	return sums
}
