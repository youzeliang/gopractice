package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {

	secret := "d625392693a3a09def783bf55af2459c"
	timeStamp := "1686226121"
	nonce := "e6e1822f-59a7-42cd-b0aa-91b64d6441d5"
	deviceId := "123456787656"
	signStr := secret + "&X-Genie-Timestamp=" + timeStamp + "&X-Genie-Nonce=" + nonce + "&X-Genie-DeviceId=" + deviceId

	fmt.Println(MD5(signStr))

	start := time.Now()

	// 执行接口调用
	// ...

	elapsed := time.Since(start)
	fmt.Printf("接口执行时间：%s\n", elapsed)

}

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
