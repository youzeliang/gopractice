package main

import (
	"fmt"
	"log"
	"os/user"
	"time"
)

func main() {

	//r := gin.New()
	//
	//helper.InitMysql()
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//go router.Crontab(r)
	//
	//endless.ListenAndServe(":"+conf.Conf.Port, r)

	//// 尝试与 google.com:80 建立 tcp 连接
	//conn, err := net.Dial("tcp", "google.com:80")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer conn.Close() // 退出关闭连接
	//
	//// 通过连接发送 GET 请求，访问首页
	//_, err = fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//dat, err := ioutil.ReadAll(conn)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(string(dat))

	fmt.Println("--------")
	t()
}

func test1() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--------")

	fmt.Println("Home dir:", u.HomeDir)
	fmt.Println("Home dir:", u.Gid)
	fmt.Println("Home dir:", u.Name)
	fmt.Println("Home dir:", u.Uid)
	fmt.Println("Home dir:", u.Username)
}

func getCurTime() {
	// 本地时间（如果是在中国，获取的是东八区时间）

	curLocalTime := time.Now()
	// UTC时间
	curUTCTime := time.Now().UTC()
	fmt.Println(curLocalTime, curUTCTime)

}

func t() {
	timer1 := time.NewTimer(time.Second * 3)

	go func() {
		//等触发时的信号
		<-timer1.C
		fmt.Println("do the thing")
	}()
	//由于上面的等待信号是在新线程中，所以代码会继续往下执行，停掉计时器
	time.Sleep(time.Second * 1)




}
