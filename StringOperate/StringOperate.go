package StringOperate

func StringReverse(data string) string {
	var mid []byte = []byte(data)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			mid[i], mid[j] = mid[j], mid[i]
		}
	}
	return string(mid)
}
