package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	//data := []string{"element1", "element2", "element3", "element4", "element5"}
	data := make([]string, 50000, 50000)

	// 设置并发控制信号量
	concurrencyLimit := 5000
	semaphore := make(chan struct{}, concurrencyLimit)

	// 创建一个等待组以等待所有goroutine完成
	var wg sync.WaitGroup

	errorFile := "errors.txt"
	successFile := "success.txt"

	// 创建互斥锁来保护文件写入操作
	var fileLock sync.Mutex

	// 创建一个映射来跟踪成功的元素
	successElements := make(map[string]bool)

	// 读取已成功的数据
	successData, err := readSuccessData(successFile)
	if err != nil {
		fmt.Printf("Failed to read success data: %v\n", err)
	}

	for _, elem := range successData {
		successElements[elem] = true
	}

	startTime := time.Now()

	for i := range data {
		wg.Add(1)
		semaphore <- struct{}{} // 获取信号量，限制并发数

		go func(elem string) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			// 检查是否已经在成功文件中
			fileLock.Lock()
			_, alreadySuccessful := successElements[elem]
			fileLock.Unlock()

			if !alreadySuccessful {
				err := callThirdPartyAPI(elem)

				fileLock.Lock()
				if err != nil {
					appendToFile(errorFile, elem)
				} else {
					appendToFile(successFile, elem)
					// 标记元素为成功
					successElements[elem] = true
				}
				fileLock.Unlock()
			}
		}(strconv.Itoa(i))
	}

	wg.Wait()
	close(semaphore) // 关闭信号量通道

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Total elapsed time: %s\n", elapsedTime)

	fmt.Println("All tasks completed.")
}

func callThirdPartyAPI(element string) error {
	// 模拟调用第三方接口
	// 如果发生错误，返回相应的错误
	return nil
}

func appendToFile(filename, element string) {
	// 打开文件，如果文件不存在则创建，如果文件存在则追加写入
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to open or create file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(element + "\n")
	if err != nil {
		fmt.Printf("Failed to write to file: %v\n", err)
	}
}

func readSuccessData(filename string) ([]string, error) {
	// 检查文件是否存在
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, nil // 文件不存在，返回空数据
	} else if err != nil {
		return nil, err // 出现其他错误
	}

	// 读取已成功的数据
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// 分割数据并去除末尾的空行
	successData := strings.Split(string(data), "\n")
	successData = successData[:len(successData)-1]

	return successData, nil
}
