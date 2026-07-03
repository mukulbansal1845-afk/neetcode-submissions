func longestConsecutive(nums []int) int {

	set := make(map[int]struct{})

	for _, v := range nums {
		set[v] = struct{}{}
	}

	ans := 0

	for v := range set {
		
		if _, ok := set[v-1]; ok {
			continue
		}

		length := 1
		for {
			if _, ok := set[v+length]; !ok {
				break
			}
			length++
		}

		if length > ans {
			ans = length
		}
	}

	return ans

}
