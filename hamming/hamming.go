package hamming

import "errors"

func Distance(a, b string) (int, error) {

	c:=0
	if len(a)==len(b) {
		for i:=0;i<len(a);i++ {

			if a[i]!=b[i] {

				c++
			}
		} 
		return c,nil
	}
	return 0,errors.New("unequal lengths")	
}
