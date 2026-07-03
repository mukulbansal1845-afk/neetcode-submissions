func isAnagram(s string, t string) bool {

	freq := make([]int,26)

	if len(s) != len (t){
		return false
	}

	for _,v := range s{
		freq[v - 'a']++
	}

	for _,v := range t{
		freq[v - 'a']--
		if freq[v - 'a'] < 0{
			return false
		}
	}

	return true

}
