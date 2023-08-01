package piscine

func Split(s, sep string) []string {
	lenOfSep := len(sep)
	lenOfS := len(s)

	for i := 0; i < lenOfS-lenOfSep; i++ {
		if s[i:i+lenOfSep] == sep {
			s = s[:i] + " " + s[i+lenOfSep:]
			lenOfS -= lenOfSep
		}
	}

	return SplitWhiteSpaces(s)
}
