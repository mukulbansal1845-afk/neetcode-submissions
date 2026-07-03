func twoSum(nums []int, target int) []int {

	i := 0
	n := len(nums)
	j := n-1

	for i < j {
		sum := nums[i] + nums[j]

		if (sum>target) {
			j--
		} else if (sum == target) {
			return []int{i+1,j+1}
		} else {
			i++
		}
	}

	return nil

}
