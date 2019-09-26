package user

import (
	"log"
	"testing"
)

func TestGenRandString(t *testing.T) {
	//declare string array for a multiple of five instances
	var str [5]string
	//assign values to each instance
	//(i.e) index of the array from the function to test
	for i := range str {
		str[i] = genRandString()
	}

	//iterates through instances
	for i, v := range str {
		for _, k := range str[i+1:] {
			if v == k {
				t.Fail()
				log.Println("strings should not match")
			}
		}
	}
}

// func TestCow(t *testing.T) {
// 	got := Cow()
// 	expected := "Co"
// 	if got != expected {
// 		t.Fatalf("%v | %v", got, expected)
// 	}
// }
