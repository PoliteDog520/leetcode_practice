package main

import "log"

func longestPalindrome(s string) int {
	var charCal = make(map[string]int)
	for i := range s {
		num, ok := charCal[string(s[i])]
		if ok {
			charCal[string(s[i])] = num + 1
		} else {
			charCal[string(s[i])] = 1
		}
	}

	var oddCharNum int
	var palindromeNum int
	for _, v := range charCal {
		if v%2 == 1 {
			if oddCharNum == 0 { // 第一個奇數字母
				oddCharNum = v
			} else if v > oddCharNum { // 不是第一個奇數字母，且字母數>前一個
				// 使用(前個奇數字母-1)=偶數 當作回文內容
				palindromeNum = palindromeNum + oddCharNum - 1
				// 指派新數量最多的奇數字母
				oddCharNum = v
			} else {
				// 使用(此奇數字母-1)=偶數 當作回文內容
				palindromeNum = palindromeNum + v - 1
			}
		} else {
			palindromeNum += v
		}
	}
	return palindromeNum + oddCharNum
}

func main() {
	log.Println(longestPalindrome("abccccdd"))
	log.Println(longestPalindrome("a"))
	log.Println(longestPalindrome("aaa"))
	log.Println(longestPalindrome("aaabbbvv"))
	log.Println(longestPalindrome("aaabbbvvccccc"))
	log.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
