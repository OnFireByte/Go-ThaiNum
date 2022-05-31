# Go-ThaiNum

**Are you convervative? Wanna use Thai numerals everywhere even coding? Good!**

This package will allow you to convert Thai numerals string to integer and float numbers type and vice versa.

> Don't let that Gen-Z people judge you 😈😈😈

PS. JK pls don't use it. It's suppose to be a joke package even if it's technically usable but please don't.

## Functions

### Stand-alone functions

-   ToThaiNum

### strconv's based functions

-   Atoi
-   ParseFloat
-   FormatInt
-   FormatFloat

### fmt print's based functions

-   Print
-   Println
-   Printf
-   SPrintf

All functions is behave the same way as the functions with the same name from standard library except that they also work with Thai numerals (or return Thai numerals string instead of Arabic numerals string).

There are some weird cases where mixed numerals also work (Atoi("1๑๒42๓") or ParseFloat("1๑๒42๓".๑2) are valid) but I don't care.

For ToThaiNum, as the name suggested, is take a string and replace all Arabic numerals to Thai ones.

Again, THIS IS A JOKE PROJECT (but feel free to add new issue or PR doe)
