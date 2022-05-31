package thaiNum

import (
	"fmt"
	"strconv"
)

func toArabic(str string) string {
	// It's liberately private because guy who wanna use this 
	// won't convert anything to arabic anyway lmao
	thaiMap := map[string]string{
		"๐": "0",
		"๑": "1",
		"๒": "2",
		"๓": "3",
		"๔": "4",
		"๕": "5",
		"๖": "6",
		"๗": "7",
		"๘": "8",
		"๙": "9",
	}
	res := ""
	for _, t := range str {
		char, ok := thaiMap[string(t)]
		if !ok {
			char = string(t)
		}
		res += char

	}
	return res
}

func ToThaiNum(str string) string {
	thaiMap := map[string]string{
		"0": "๐",
		"1": "๑",
		"2": "๒",
		"3": "๓",
		"4": "๔",
		"5": "๕",
		"6": "๖",
		"7": "๗",
		"8": "๘",
		"9": "๙",
	}
	res := ""
	for _, t := range str {
		char, ok := thaiMap[string(t)]
		if !ok {
			char = string(t)
		}
		res += char

	}
	return res
}


func FormatInt(i int64, base int) string {
	return ToThaiNum(strconv.FormatInt(i, base))
}

func FormatFloat(f float64, fmt byte, prec, bitSize int) string {
	return ToThaiNum(strconv.FormatFloat(f, fmt, prec, bitSize))
}

func Atoi(s string) (int, error) {
	return strconv.Atoi(toArabic(s))
}

func ParseFloat(s string) (float64, error) {
	return strconv.ParseFloat(toArabic(s), 64)
}

func Print(a ...any) (n int, err error){
	var res []interface{}
	for _, v := range a {
		res = append(res, ToThaiNum(fmt.Sprint(v)))
	}
	return fmt.Print(res...)
}

func Println(a ...any) (n int, err error) {
	var res []interface{}
	for _, v := range a {
		res = append(res, ToThaiNum(fmt.Sprint(v)))
	}
	return fmt.Println(res...)
}

func Printf(format string, a ...any) (n int, err error){
	return fmt.Print(Sprintf(format, a...))
	
}

func Sprintf(format string, a ...any)string{
	return ToThaiNum(fmt.Sprintf(format, a...))
}