package raindrops

import "fmt"

func Convert(n int) string {
	strFinal := ""
	if n%3 == 0 {
		strFinal += "Pling"
	}
	if n%5 == 0 {
		strFinal += "Plang"
	}
	if n%7 == 0 {
		strFinal += "Plong"
	}
	if len(strFinal) == 0 {
		strFinal = fmt.Sprint(n)
	}
	return strFinal
}
