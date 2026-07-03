type Solution struct{}

func (s *Solution) Encode(strs []string) string {

	res := ""

	for _, v := range strs {
		res += strconv.Itoa(len(v)) + "#" + v
	}

	return res

}

func (s *Solution) Decode(str string) []string {
	res := make([]string, 0)
	i := 0
	j := 0

	for i < len(str) {
		v := str[i]
		length := 0
		if v == '#' {
			substr := str[j:i]
			length, _ = strconv.Atoi(substr)

			j = i + length + 1 
			res = append(res, str[i+1:j])
			i = j
			continue
		}
		i++
	}

	return res
}
