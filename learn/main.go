package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache" // 使用前先import包
)

func main() {
	// 创建一个cache对象，默认ttl 5分钟，每10分钟对过期数据进行一次清理
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set一个KV，key是"foo"，value是"bar"
	// TTL是默认值（上面创建对象的入参，也可以设置不同的值）5分钟
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set了一个没有TTL的KV，只有调用delete接口指定key时才会删除
	c.Set("baz", 42, cache.NoExpiration)

	// 从cache中获取key对应的value
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	// 如果想提高性能，存储指针类型的值
	//c.Set("foo", &MyStruct, cache.DefaultExpiration)
	//if x, found := c.Get("foo"); found {
	//	foo := x.(*MyStruct)
	//	// ...
	//}

	// 1
	// 2

}
