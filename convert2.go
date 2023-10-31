package convert

import (
	"fmt"
	"strings"
	"unicode"
)

func ToWords2(n int64) string {
	_ = map[int]string{
		1: "mốt",
		5: "lăm",
	}

	_ = map[int64]string{
		// tens exception
		0: "lẻ",
		1: "mười",
		// ones exception
		4: "tư",
		5: "năm",
	}

	_ = map[int64]string{
		0: "linh",
		5: "nhăm",
	}

	_ = map[int64]string{
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
	digits := fmt.Sprintf("%d", n)
	totalDigits := len(digits)
	result := ""
	currentLength := len(digits)
	if totalDigits == 1 {
		result = num[digits]
	} else if totalDigits >= 2 {
		multipleZero := 0
		multipleOne := 0
		print(multipleZero)
		for _, char := range digits {
			formatChar := fmt.Sprintf("%c", char)
			rest := currentLength % 3
			if rest == 0 {
				// abc check
				// todo multiple zero checking
				// định dạng hàng trăm - a
				result += num[formatChar] + "trăm"
			} else if rest == 2 {
				// bc check
				// định dạng hàng chục - b: linh/lẻ, mười, mươi
				if formatChar == "0" {
					multipleZero += 1
				} else if formatChar == "1" {
					multipleOne += 1
				} else {
					result += num[formatChar] + " mươi"
				}
			} else {
				// c check
				// định dạng hàng đơn vị
				if multipleOne == 1 {
					result += num[formatChar] + "mười"
					multipleOne = 0 // reset
				} else {
					result += num[formatChar]
				}
			}
			currentLength -= 1
		}
	}

	charResult := []byte(result)
	charResult[0] = byte(unicode.ToUpper(rune(charResult[0])))
	result = string(charResult)
	result = fmt.Sprintf(strings.TrimRight(result, " "))
	result += " đồng"
	return result
}
