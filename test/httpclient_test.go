package httptest

import (
	"fmt"
	httpclient "httplocal"
	"io/ioutil"
	"net/http"
	// "strings"
	"testing"
)

func TestGetMethod(t *testing.T) {
	err :=
		httpclient.Get("https://www.baidu.com").
			Execute()
	if err != nil {
		t.Fail()
	}
}

func TestEverything(t *testing.T) {
	// rsp, err := http.Get("https://www.okcoin.cn/api/v1/ticker.do?symbol=btc_cny")
	req, _ := http.NewRequest("GET", "https://www.okcoin.cn/api/v1/ticker.do?symbol=btc_cny", nil)
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bts, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(bts[:]))
}

func TestGetString(t *testing.T) {
	str := ""
	err :=
		httpclient.Get("https://www.okcoin.cn/api/v1/ticker.do?symbol=btc_cny").
			GetString(&str).
			Execute()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(str)
}

/**
 *
 * 测试GetJson
 *
 */
func TestGetJson(t *testing.T) {
	obj := JsonObject{}
	err :=
		httpclient.Get("https://www.okcoin.cn/api/v1/ticker.do?symbol=btc_cny").
			GetJson(&obj).
			Execute()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(obj)
}
func TestGetJsonp(t *testing.T) {
	obj := JsonpObject{}
	err :=
		httpclient.Get("http://10.240.109.25:8080/upload/reportasnytasks?type=41&vid=j05133ah8u0&format=320088&ip=10.177.143.50&otype=json").
			GetJsonp(&obj).
			Execute()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(obj)
}

func TestMixed(t *testing.T) {
	obj := JsonObject{}
	str := ""
	err :=
		httpclient.Get("https://www.okcoin.cn/api/v1/ticker.do?symbol=btc_cny").
			GetJson(&obj).
			GetString(&str).
			Execute()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	fmt.Println(obj)
	fmt.Println(str)
}

// func TestHeaderAndCookie(t *testing.T) {
// 	err :=
// 		httpclient.Post("http://localhost:8888/upload").
// 			ContentType("application/json").
// 			Body(strings.NewReader("{\"a\":1}")).
// 			AddCookie("name", "lewiskong").
// 			AddHeader("from", "localhost").
// 			Execute()
// 	if err != nil {
// 		panic(err)
// 		t.Fail()
// 	}
// }

func TestWrongUrl(t *testing.T) {
	err :=
		httpclient.Post("http://localhost:1231/nil").
			Execute()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestName(t *testing.T) {
	str := ""
	err := httpclient.Post("zkname://interface.vbasic.cm.com/upload/reportasnytasks?type=41&vid=p0016v94qse&format=320085&ip=10.177.143.50&otype=json").
		GetString(&str).
		Execute()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(str)
}

func TestL5(t *testing.T) {
	err := httpclient.Post("l5://371329:720896/").
		Execute()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

// func TestAlert(t *testing.T) {
// 	str := ""
// 	err := httpclient.Post("l5://371329:720896/").
// 		SetMonitor(&httpclient.Monitor{1, 2, 3, 4}).
// 		SetModCall(&httpclient.ModuleCall{1111, 2222, 444, 333}).
// 		GetString(&str).
// 		Execute()
// 	if err != nil {
// 		t.Log(err)
// 		t.Fail()
// 	}

// }

type JsonObject struct {
	Date   string     `json:"date"`
	Ticker TickerItem `json:"ticker"`
}
type TickerItem struct {
	Low  string `json:"low"`
	Sell string `json:"sell"`
	Vol  string `json:"vol"`
	Buy  string `json:"buy"`
	High string `json:"high"`
	Last string `json:"last"`
}

type JsonpObject struct {
	Result ResultItem `json:"result"`
}
type ResultItem struct {
	Code float64 `json:"code"`
	Msg  string  `json:"msg"`
}
