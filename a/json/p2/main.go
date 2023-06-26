package main

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"strings"
)

type ValidFunc func(context.Context, []byte) ([]byte, error)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

type GetAppCatalogInfoResponse struct {
	ID                  string            `json:"id"`
	AppID               string            `json:"app_id"`
	Name                string            `json:"name"`
	YearID              string            `json:"year_id"`
	TermID              string            `json:"term_id"`
	GradeID             string            `json:"grade_id"`
	BaseInfoExtendField string            `json:"base_info_extend_field"`
	ParentID            string            `json:"parent_id"`
	OtherExtendInfo     string            `json:"other_extend_info"`
	Status              string            `json:"status"`
	CreateRealName      string            `json:"create_real_name"`
	ModifyName          string            `json:"modify_name"`
	ModifyRealName      string            `json:"modify_real_name"`
	CatalogInfo         CatalogInfoDetail `json:"catalog_info"`
}

type CatalogInfoDetail struct {
	Catalog []map[string]interface{} `json:"catalog"`
}

type AA struct {
	Status int `json:"status"`
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func GetSign(stuId, stuCouId, planId string) (token string) {
	sign := "Xes-Report-Token-" + stuId + "-" + stuCouId + "-" + planId
	return Md5(sign)
}

func main() {

	garr := strings.Split("", ",")

	fmt.Println(len(garr))

	to := GetSign("[{\"stuId\":59616,\"stuCouId\":10056662}]", "173937", "1477478")
	fmt.Println(to)

	//var a = "{\"status\":100,\"data\":{\"id\":\"1034083\",\"app_id\":\"45\",\"name\":\"321\",\"year_id\":\"21\",\"term_id\":\"1\",\"grade_id\":\"7\",\"base_info_extend_field\":\"{}\",\"parent_id\":\"0\",\"other_extend_info\":\"{}\",\"sort\":\"1\",\"status\":\"3\",\"unique_video_id\":\"0\",\"school_id\":\"1\",\"catalog_info\":[]}}"

	a := `{"status": 1}`

	cc, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	type BasicResp struct {
		Status int `json:"status"`
		//Data jsoniter.RawMessage `json:"data"`
	}
	var resp BasicResp
	err = json.Unmarshal(cc, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cc)
	data, err := CheckValid(context.Background(), cc, CheckStatus100, CheckStat1EmptyData)
	if err != nil {
		fmt.Println(err)
		return
	}

	var s GetAppCatalogInfoResponse
	err = jsoniter.Unmarshal(data, &s)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CheckStatus100(ctx context.Context, data []byte) (ret []byte, err error) {
	type BasicResp struct {
		Stat int `json:"status"`
		//Data jsoniter.RawMessage `json:"data"`
	}
	var resp BasicResp
	err = jsoniter.Unmarshal(data, &resp)
	if err != nil {
		return
	}
	if resp.Stat != 100 {
		fmt.Println(err)
		return
	}
	//return resp.Data, nil

	return nil, err
}

func CheckValid(ctx context.Context, data []byte, fns ...ValidFunc) (ret []byte, err error) {
	ret = data
	for _, fn := range fns {
		ret, err = fn(ctx, ret)
		if err != nil {
			return
		}
	}
	return
}

func CheckStat1EmptyData(ctx context.Context, data []byte) (ret []byte, err error) {
	ret = data
	if bytes.HasPrefix(data, []byte("[]")) {
		ret = []byte("{}")
	}
	return
}
