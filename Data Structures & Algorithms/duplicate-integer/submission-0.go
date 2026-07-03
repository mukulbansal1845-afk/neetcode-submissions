func hasDuplicate(nums []int) bool {

	mp := make(map[int]int)

	for _,v := range nums {
		mp[v]++
		if mp[v] > 1 {
			return true
		}
	}

	return false
    
}
