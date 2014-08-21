package romankana

import (
	//	"fmt"
	"strings"
)

func KanaRoman(input string) string {
	temp := strings.Split(input, "")
	var array []string
	for i, s := range temp {
		if i+1 < len(temp) {
			next_str := temp[i+1]
			if strings.Contains("ァィゥェォャュョ", next_str) {
				s = s + next_str
				temp[i+1] = ""
			}
		}
		if len(s) > 0 {
			array = append(array, s)
		}
	}

	var ret []string
	for _, s := range array {
		e, ok := k2r[s]
		if !ok {
			e = s
		}
		ret = append(ret, e)
	}

	for i, s := range array {
		if s == "ッ" {
			if i+1 < len(ret) {
				a := strings.Split(ret[i+1], "")
				c := a[0]
				if !strings.Contains("aiueo", c) {
					ret[i] = c
				}
			}
		} else if s == "ー" {
			if i-1 >= 0 {
				a := strings.Split(ret[i-1], "")
				c := a[len(a)-1]
				if strings.Contains("aiueo", c) {
					ret[i] = c
				}
			}
		}
	}

	return strings.Join(ret, "")
}
