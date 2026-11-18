type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	buf := strings.Builder{}
	for _, str := range strs {
		sl := fmt.Sprintf("%03d", len(str))
		buf.WriteString(sl)
		buf.WriteString(str)
	}
	return buf.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := make([]string, 0)
	bs := []byte(encoded)
	for len(bs) > 0 {
		sl := bs[:3]
		bs = bs[3:]
		l, _ := strconv.Atoi(string(sl))
		str := string(bs[:l])
		bs = bs[l:]
		res = append(res, str)
	}
	return res
}
