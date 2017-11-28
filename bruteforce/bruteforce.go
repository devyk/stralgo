package bruteforce

func StrBruteForce(haystack string, needle string) (int) {
	var i, j int
	hLength := len(haystack)
	nLength := len(needle)

	for i = 0; i < hLength - nLength; i++ {
		for j = 0; j < nLength && haystack[i+j] == needle[j]; j++ {}
		if j >= nLength-1 {
			return i
		}
	}

	return -1
}
