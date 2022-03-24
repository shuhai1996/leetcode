package main
//NO.344
func main() {
	reverseString([]byte{'H','a','n','n','a','h'})
}

func reverseString(s []byte)  {
	var m byte
	l := len(s)
	for i:=0;i<(len(s)+1)/2;i++ {
		m = s[i]
		s[i] = s[l-1-i]
		s[l-1-i] = m
	}
}