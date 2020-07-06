package acronym

import (
	
	"strings"
)

// Abbreviate returns the acronym of a given string.
func Abbreviate(s string) string {

	s=strings.ReplaceAll(s, "-"," ")

	s=strings.ReplaceAll(s, "_"," ")

	var result string

	temp := strings.Split(s, " ")
	for i:=0;i<len(temp);i++ {

		if len(temp[i])!=0 {

		
		result+=string(temp[i][0])
		}
	}
	return strings.ToUpper(result)
}
