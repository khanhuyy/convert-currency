package convert

import (
	"fmt"
	"strings"
	"unicode"
)

func ToWords(n int64) string {
	exceptionOne := map[int]string{
		1: "mốt",
	}

	exception := map[int64]string{
		// tens exception
		0: "lẻ",
		1: "mười",
		// ones exception
		4: "tư",
		5: "năm",
	}

	group := map[int64]string{
		1: "mươi",
		2: "trăm",
		3: "nghìn",
		6: "triệu",
		9: "tỷ",
	}

	num := map[string]string{
		"0": "không",
		"1": "một",
		"2": "hai",
		"3": "ba",
		"4": "bốn",
		"5": "năm",
		"6": "sáu",
		"7": "bảy",
		"8": "tám",
		"9": "chín",
	}
	s := fmt.Sprintf("%d", n)
	var lenString = int64(len(s))
	checkBillions := lenString - (lenString % 9) // detect nine consecutive similar numbers
	checkTriple := lenString - (lenString % 3)   // detect multi triple position
	offset := lenString - checkTriple + 1        // triple convert
	oneFlag := false                             // flag to check exceptionOne
	// check empty string
	if lenString == 0 {
		return ""
		// one digit
	} else if lenString < 2 {
		result := num[s]
		// todo check used for
		// *
		charResult := []byte(result)
		charResult[0] = byte(unicode.ToUpper(rune(charResult[0])))
		result = string(charResult)
		result = fmt.Sprintf(strings.TrimRight(result, " "))
		// *
		result += " đồng"
		return result
	} else {
		result := ""
		var zeroFlag int64 // zero check
		for pos, char := range s {
			if offset == 1 {
				offset = 3
				checkTriple -= 3
			} else {
				offset -= 1
			}
			unit := len(s) - pos - 1
			residualTriple := unit % 3
			formatChar := fmt.Sprintf("%c", char)
			switch residualTriple {
			case 1:
				if formatChar == "0" {
					zeroFlag += 1
				} else if formatChar == "1" {
					if zeroFlag%3 == 1 {
						result += num["0"] + " " + group[2] + " " + exception[1] + " "
					} else {
						result += exception[1] + " "
					}
					zeroFlag = 0
				} else {
					if zeroFlag%3 == 1 {
						result += num["0"] + " " + group[2] + " " + num[formatChar] + " " + group[1] + " "
					} else {
						result += num[formatChar] + " " + group[1] + " "
					}
					oneFlag = true
					zeroFlag = 0
				}
			case 2:
				if fmt.Sprintf("%c", char) == "0" {
					zeroFlag += 1
				} else {
					zeroFlag = 0
					result += num[fmt.Sprintf("%c", char)] + " " + group[2] + " "
				}
			default:
				if formatChar == "0" {
					zeroFlag += 1
				} else if formatChar == "1" {
					if zeroFlag > 1 {
						result += exception[0] + " " + num["1"] + " "
						oneFlag = false
					} else {
						if oneFlag {
							result += exceptionOne[1] + " "
							oneFlag = true
						} else {
							result += num["1"] + " "
						}
					}
					zeroFlag = 0
				} else if formatChar == "4" {
					if zeroFlag > 1 {
						result += exception[0] + " " + exception[4] + " "
					} else {
						if pos != 0 {
							result += exception[4] + " "
						} else {
							result += num["4"] + " "
						}
					}
					zeroFlag = 0
				} else if formatChar == "5" {
					if zeroFlag > 1 {
						result += exception[0] + " " + num["5"] + " "
					} else {
						result += exception[5] + " "
					}
					zeroFlag = 0
				} else {
					if zeroFlag > 1 {
						result += exception[0] + " " + num[formatChar] + " "
					} else {
						result += num[formatChar] + " "
					}
					zeroFlag = 0
				}
				if checkTriple%9 == 0 && checkTriple > 0 {
					checkBillions -= 9
					result += group[9] + " "
				} else {
					if zeroFlag < 1 {
						result += group[checkTriple%9] + " "
						zeroFlag = 0
					} else if zeroFlag/3 < 1 {
						result += group[checkTriple%9] + " "
						zeroFlag = 0
					}
				}
			}
		}
		charResult := []byte(result)
		charResult[0] = byte(unicode.ToUpper(rune(charResult[0])))
		result = string(charResult)
		result = fmt.Sprintf(strings.TrimRight(result, " "))
		result += " đồng"
		return result
	}
}
