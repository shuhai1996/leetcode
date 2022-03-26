package main

import "strings"

func main() {

}

func letterCasePermutation(s string) []string {
	s =  strings.ToLower(s)
	res := []string{s}
	for i := range s {
		if s[i] >= 'a' {
			end := len(res)
			for  j:=0; j < end; j++{
				newTemp := []byte(res[j])
				newTemp[i] = s[i]-'a'+'A'
				res = append(res, string(newTemp))
			}
		}
	}

	return res
}