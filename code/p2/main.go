package main

//
//import "fmt"
//import "github.com/xuri/excelize/v2"
//
//func main() {
//	E([]string{"321", "2", "3"})
//}
//
//func E(s []string) {
//	// 创建一个新的xlsx文件
//	f := excelize.NewFile()
//
//	// 在工作表中创建一个名为"Sheet1"的新工作表
//	index, _ := f.NewSheet("Sheet1")
//
//	for i := 0; i < len(s); i++ {
//		cell := fmt.Sprintf("A%d", i)
//		f.SetCellValue("Sheet1", cell, s[i])
//	}
//
//	// 将新工作表设为活动工作表
//	f.SetActiveSheet(index)
//
//	// 保存文件
//	if err := f.SaveAs("test666.xlsx"); err != nil {
//		fmt.Println(err)
//	}
//}
