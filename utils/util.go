package utils

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"

	"github.com/sony/sonyflake"
)

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

func GetUuidString() string {
	return strconv.FormatUint(GetUuidUInt64(), 10)
}

func GetUuidUInt64() uint64 {
	uid, _ := flake.NextID()
	return uid
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

func GetGreeting(name string) (greeting string) {
	greeting = "Hello " + name + "."
	return
}
