package piscine

func SplitWhiteSpaces(s string) []string {
	var strArray []string

	strNew := ""

	for i := 0; i <= len(s)-1; i++ {
		if i == len(s)-1 {
			strNew += string(s[i])
			strArray = append(strArray, strNew)
		} else if string(s[i]) != " " && string(s[i]) != "\t" && string(s[i]) != "\n" {
			strNew += string(s[i])
		} else {
			strArray = append(strArray, strNew)
			strNew = ""
		}
	}
	return strArray
}
