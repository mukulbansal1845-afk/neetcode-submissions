func twoSum(arr []int, target int) []int {

	n := len(arr)

	i := 0
	j := n-1

	for i<j {

		sum := arr[i] + arr[j]

		if sum > target {
			j--
		} else if sum == target {
			return []int{i+1,j+1}
		} else {
			i++
		}
	}

	return nil

}
