package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/youzeliang/gopractice/service"
	"time"
)

func Crontab(g *gin.Engine) {
	InitCrontab(g)
}

func InitCrontab(g *gin.Engine) () {
	//c := cron.New()
	//c.AddFunc("@every 1ms", service.Test)
	//c.Start()
	loopWorker()

}

func loopWorker() {
	i := 0
	ticker := time.NewTicker(1 * time.Millisecond) // --- A
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			service.Test()
			i++
			doxx(i)

		}
	}
}

func doxx(i int) {
	//time.Sleep(7 * time.Second) // --- B
	fmt.Println("aaa", i, time.Now().Unix())
}

// 时间A和B 以时间长的为间隔周期
