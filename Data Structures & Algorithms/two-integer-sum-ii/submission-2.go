func twoSum(nums []int, target int) []int {

	n := len(nums)

	i := 0
	j := n-1

	for i<j {
		sum := nums[i] + nums[j]

		if sum>target{
			j--
		} else if sum == target{
			return []int{i+1,j+1}
		} else {
			i++
		}
	}

	return nil
}
