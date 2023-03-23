package main

// pprof 的init函数会将pprof里的一些handler注册到http.DefaultServeMux上
// 当不使用http.DefaultServeMux来提供http api时，可以查阅其init函数，自己注册handler
import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	_ "net/http/pprof"
	"strconv"
)

func main() {
	//http.ListenAndServe("0.0.0.0:8080", nil)

	//var a = []int{6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,0,7,0,6,0,0,0,0,6,0,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,6}
	//
	//fmt.Println(len(a))
	//m := 0
	//for i := 0; i <len(a) ; i++ {
	//	if a[i] == 6 {
	//		m++
	//	}
	//}
	//
	//fmt.Println(m)

	i, _ := strconv.ParseFloat("32131321312", 64)

	fmt.Println(cast.ToInt(i))

	fmt.Println(cast.ToInt(math.Abs(math.Round(cast.ToFloat64(i / 60)))))
}
