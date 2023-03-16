package main

import "log"

func isIsomorphic(s string, t string) bool {
	// s字串對應t字串表
	var tCharMap = make(map[string]string, len(t))
	// t字串對應s字串表
	var sCharMap = make(map[string]string, len(s))
	// 將 s 字串 對應到t並組成新字串
	var sToT string
	// 將 t 字串 對應到s並組成新字串

	var tToS string
	if len(t) != len(s) {
		return false
	}
	for i := range t {
		sChar, sOk := sCharMap[string(t[i])]
		if sOk {
			tToS += sChar
		} else {
			sCharMap[string(t[i])] = string(s[i])
			tToS += string(s[i])
		}
	}
	for i := range s {
		tChar, tOk := tCharMap[string(s[i])]
		if tOk {
			sToT += tChar
		} else {
			tCharMap[string(s[i])] = string(t[i])
			sToT += string(t[i])
		}
	}

	if sToT == t && tToS == s {
		return true
	}

	return false
}

func main() {
	log.Println(isIsomorphic("egg", "add"))
	log.Println(isIsomorphic("foo", "bar"))
	log.Println(isIsomorphic("paper", "title"))
	log.Println(isIsomorphic("he he", "ya ya"))
	log.Println(isIsomorphic("badc", "baba"))
}
