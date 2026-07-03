import "slices"

func threeSum(nums []int) [][]int {

	slices.Sort(nums)

	n := len(nums)

	res := make([][]int, 0)

	for k := 0; k < n; k++ {

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i := k + 1
		j := n - 1

		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if sum < 0 {
				i++
			} else if sum == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				j--
			}
		}
	}

	return res

}
