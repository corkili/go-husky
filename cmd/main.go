package main

import (
	"fmt"
	db2 "go-husky/internal/db"
)

func main() {
	db, err := db2.GetInstance().GetDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	test := db2.Test{}
	//db.Save(&test)
	db.First(&test)
	printTest(&test)

	fmt.Println("----------------------------")
	var strPtr *string = nil
	fmt.Println(strPtr)
	//fmt.Println(*strPtr)
}

func printTest(test *db2.Test) {
	fmt.Println("test_int:[", test.TestInt, "]")
	fmt.Println("test_int_ptr:[", test.TestIntPtr, "]")
	fmt.Println("test_int_ptr value:[", *test.TestIntPtr, "]")
	fmt.Println("test_string:[", test.TestStr, "]")
	fmt.Println("test_string_ptr:[", test.TestStrPtr, "]")
	fmt.Println("test_string_ptr value:[", *test.TestStrPtr, "]")
}
