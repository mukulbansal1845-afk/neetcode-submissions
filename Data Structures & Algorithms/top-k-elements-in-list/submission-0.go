func topKFrequent(nums []int, k int) []int {

	var res []int

	mp := make(map[int]int)

	for _, v := range nums {
		mp[v]++
	}

	n := len(nums)

	bucket := make([][]int, n+1)

	for i, v := range mp {
		bucket[v] = append(bucket[v], i)
	}

	for i := n; i >= 0; i-- {
		for _, v := range bucket[i] {
			res = append(res, v)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
