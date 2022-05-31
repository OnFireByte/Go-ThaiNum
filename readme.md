# Go-ThaiNum

> **Are you convervative? Wanna use Thai numerals everywhere even coding? Good!**

This package will allow you to convert Thai numerals string to integer and float numbers type and vice versa

> Don't let that gen-Z people judge you 😈😈😈

PS. JK pls don't use it. It's suppose to be a joke package even if it's technically usable but please don't

## Functions

-   ToThaiNum
-   Atoi
-   ParseFloat
-   FormatInt
-   FormatFloat

All functions is behave the same way as the strandard strconv functions with the same name except that they also work with Thai numerals (or return Thai numerals string instead of Arabic numerals string)

There are some weird cases where mixed numerals also work (Atoi("1๑๒42๓") or ParseFloat("1๑๒42๓".๑2) are legit) but I dont care

For ToThaiNum, like the name suggested, it take a string and replace all Arabic numerals to Thai one
