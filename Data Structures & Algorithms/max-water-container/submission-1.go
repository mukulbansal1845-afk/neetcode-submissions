func maxArea(arr []int) int {

	n := len(arr)

	i := 0
	j := n-1

	ans := 0

	for i < j {
		area := min(arr[i], arr[j]) * (j - i)

		if area > ans {
			ans = area
		}

		if arr[i] > arr[j] {
			j--
		} else {
			i++
		}
	}

	return ans

}
