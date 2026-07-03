func groupAnagrams(strs []string) [][]string {

	mp := make(map[string][]string)

	res := make([][]string,0)

	for _,v := range strs {
		key := buildKey(v)

		mp[key] = append (mp[key], v)
	}

	for _,v := range mp {
		res = append (res, v)
	}

	return res
}

func buildKey(s string) string {
	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	var res string
	for i, v := range freq {
		res += strconv.Itoa(i) + "#" + strconv.Itoa(v) + "|"
	}

	return res
}