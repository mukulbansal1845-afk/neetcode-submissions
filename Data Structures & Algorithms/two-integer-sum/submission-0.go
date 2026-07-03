func twoSum(nums []int, target int) []int {

	mp := make(map[int]int)

	for i,v := range nums {

		if j, ok := mp[v]; ok {
			if i > j{
				return []int{j,i}
			}
			return []int{i,j}
		}
		mp[target - v] = i
	}

	return nil
    
}
