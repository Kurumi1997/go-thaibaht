package gothaibaht

import "fmt"

func ExampleFloatToText() {
	text := FloatToText(123456789.99)

	fmt.Println(text)
	// output: หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเก้าสิบเก้าสตางค์
}


func ExampleStringToText() {
	text := StringToText("123456789.99")

	fmt.Println(text)
	// output: หนึ่งร้อยยี่สิบสามล้านสี่แสนห้าหมื่นหกพันเจ็ดร้อยแปดสิบเก้าบาทเก้าสิบเก้าสตางค์
}