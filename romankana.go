package romankana

import (
	"regexp"
	"strings"
)

func HiraganaToKatakana(input string) string {
	runes := []rune(input)
	var array []rune
	for _, r := range runes {
		if 0x3041 <= r && r <= 0x3096 {
			r += 0x60
		}
		array = append(array, r)
	}
	result := string(array)
	return result
}

func KatakanaToHiragana(input string) string {
	runes := []rune(input)
	var array []rune
	for _, r := range runes {
		if 0x30A1 <= r && r <= 0x30F6 {
			r -= 0x60
		}
		array = append(array, r)
	}
	result := string(array)
	return result
}

func FindKanaFromStr(input string) string {
	if len(input) == 0 {
		return ""
	}

	found, ok := r2k[input]
	if ok {
		return string(found)
	} else if len(input) >= 2 && input[0] == 'n' && input[1] == 'n' {
		return "ン" + FindKanaFromStr(input[2:len(input)])
	} else if len(input) > 2 && input[0] == 'n' && !strings.Contains("aiueoy", string(input[1])) {
		return "ン" + FindKanaFromStr(input[1:len(input)])
	} else if len(input) > 2 && input[0] == 'm' && strings.Contains("bmp", string(input[1])) {
		return "ン" + FindKanaFromStr(input[1:len(input)])
	} else if len(input) >= 2 && input[0] == input[1] && strings.Contains("bcdfghjklmnpqrstvwxyz", string(input[0])) {
		return "ッ" + FindKanaFromStr(input[1:len(input)])
	} else if len(input) >= 2 {
		return FindKanaFromStr(input[0:1]) + FindKanaFromStr(input[1:len(input)])
	} else {
		return string(input)
	}
}

func KanaRoman(input string) string {
	input = HiraganaToKatakana(input)
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

func RomanKana(input string) string {
	temp := strings.Split(input, " ")
	var array []string
	for _, s := range temp {
		s2 := regexp.MustCompile("([^aiueo]*[aiueo])").FindAllString(s, -1)
		var tmp string
		if len(s2) == 0 {
			tmp = s
		} else {
			var array2 []string
			for _, s3 := range s2 {
				array2 = append(array2, FindKanaFromStr(s3))
			}
			tmp = strings.Join(array2, "")
		}
		array = append(array, tmp)
	}
	return strings.Join(array, " ")
}
