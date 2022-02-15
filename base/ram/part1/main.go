package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Printf("bool size: %d\n", unsafe.Sizeof(bool(true)))
	//fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
	//fmt.Printf("int8 size: %d\n", unsafe.Sizeof(int8(0)))
	//fmt.Printf("int64 size: %d\n", unsafe.Sizeof(int64(0)))
	//fmt.Printf("byte size: %d\n", unsafe.Sizeof(byte(0)))
	//fmt.Printf("string size: %d\n", unsafe.Sizeof("EDDYCJY"))
	//
	//part1 := Part1{}
	//
	//fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))

	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	date := t.AddDate(0, 0, 0)

	fmt.Println(date)

	fmt.Println(int64(-4))

}
