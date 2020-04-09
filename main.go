package main

import (
<<<<<<< HEAD
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




=======
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/imroc/req"
	"go.uber.org/zap"
)

type ApiRequest struct {
	SystemID  string  `json:"systemId"`
	Timestamp string  `json:"timestamp"`
	Sign      string  `json:"sign"`
	BizData   BizData `json:"bizData"`
}

type BizData struct {
	EventID      string            `json:"eventId"`
	EventName    string            `json:"eventName"`
	LanguageCode string            `json:"languageCode"`
	CountryCode  string            `json:"countryCode"`
	Channel      string            `json:"channel"`
	ContactID    string            `json:"contactId"`
	FirstName    string            `json:"firstName"`
	LastName     string            `json:"lastName"`
	MobileNo     string            `json:"mobileNo"`
	Email        string            `json:"email"`
	WeChatOpenID string            `json:"weChatOpenId"`
	EventParams  map[string]string `json:"eventParams"`
}

type CoreRequest struct {
	SystemId    string                 `json:"system_id"`
	SequenceId  string                 `json:"sequence_id"`
	CountryCode string                 `json:"country_code"`
	BizData     map[string]interface{} `json:"biz_data"`
}

type ApiResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type stdRepository struct {
	logger        *zap.Logger
	serverAddress string
}

func NewStdRepository(logger *zap.Logger) api.Repository {
	serverAddress := os.Getenv("MESSAGE_HUB_CORE_SERVER_ADDRESS")

	return &stdRepository{
		logger:        logger,
		serverAddress: serverAddress,
	}
}

func (repo *stdRepository) SendMessage(apiRequest ApiRequest) (apiResponse ApiResponse, err error) {
	url := repo.serverAddress + "/message/send"

	header := repo.prepareReqHeader()

	// coreRequest := repo.prepareCommonParams(apiRequest)

	coreRequest := CoreRequest{}
	coreRequest.BizData = make(map[string]interface{})
	coreRequest = repo.prepareCommonParams(apiRequest)
	bizData := make(map[string]interface{})

	bizData["event_id"] = apiRequest.BizData.EventID
	bizData["event_name"] = apiRequest.BizData.EventName
	bizData["language_code"] = apiRequest.BizData.LanguageCode
	bizData["country_code"] = apiRequest.BizData.CountryCode
	channel := strings.ToLower(apiRequest.BizData.Channel)
	bizData["channel"] = channel
	switch channel {
	case "sms":
		bizData["mobile_no"] = apiRequest.BizData.MobileNo
	case "email":
		bizData["email"] = apiRequest.BizData.Email
	case "wechat":
		bizData["wechat_open_id"] = apiRequest.BizData.WeChatOpenID
	default:
		apiResponse = ApiResponse{
			Code:    http.StatusNotFound,
			Message: "Resource not found",
		}
		return
	}
	bizData["contact_id"] = apiRequest.BizData.ContactID
	bizData["first_name"] = apiRequest.BizData.FirstName
	bizData["last_name"] = apiRequest.BizData.LastName
	eventParams := apiRequest.BizData.EventParams
	for k, v := range eventParams {
		bizData[k] = v
	}
	coreRequest.BizData = bizData

	t, _ := json.Marshal(coreRequest)
	b := string(t)
	fmt.Println(b)

	params := req.BodyJSON(&coreRequest)

	r, err := req.Post(url, header, params)
	if err != nil {
		repo.logger.Error(err.Error())
	}

	resp := r.Response()
	defer resp.Body.Close()

	result := ApiResponse{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		repo.logger.Error(err.Error())
	}

	apiResponse = ApiResponse{
		Code:    result.Code,
		Message: result.Message,
		Data:    result.Data,
	}

	return
}

func (repo *stdRepository) prepareReqHeader() (header req.Header) {
	header = req.Header{
		"Accept": "application/json",
	}

	return header
}

func (repo *stdRepository) prepareCommonParams(apiRequest ApiRequest) (param CoreRequest) {
	return CoreRequest{
		SystemId:    apiRequest.SystemID,
		SequenceId:  apiRequest.SystemID + "_" + strconv.FormatInt(time.Now().UnixNano(), 10),
		CountryCode: apiRequest.BizData.CountryCode,
	}
>>>>>>> c812791a34cfcf89a2154daefae37c04ef19b7d8
}
