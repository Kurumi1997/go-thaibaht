# go-thaibaht
go-thaibaht is a simple library for convert number to thai baht as text

> Hit the project with a star if you find it useful ⭐

## Installation
```bash
go get github.com/Kurumichan666/go-thaibaht
```

## Example
```go
package main

import (
	"log"
	gothaibaht "github.com/Kurumichan666/go-thaibaht"
)

func main() {
	floatToText := gothaibaht.FloatToText(123456789.99)
	fmt.Println(floatToText)
	// output: หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเก้าสิบเก้าสตางค์

	stringToText := gothaibaht.StringToText("123456789.99")
	// output: หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเก้าสิบเก้าสตางค์
	fmt.Println(stringToText)
}
```



