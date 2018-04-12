# GoHttp
A Chain call http lib for go.Support JSON &amp;&amp; JSONP &amp;&amp; String get directly

# Usage 
```
    type JsonObj struct {
        name string
    }
    str := ""
    obj := JsonObj{}
    err:=httpclient.Post("http://localhost:8080/upload").
            ContentType("application/json").
            Body(`{"name":"lewiskong"}`).
            AddHeader("Referer","localhost").
            AddCookie("uin","guest").
            GetString(&str).
            GetJson(&obj). //when the result is json
            GetJsonp(&obj). // when the result is jsonp
            GetJce(&obj). //when the result is jce bytes 待实现
            Execute()
```

# URL路由支持
httpclient除支持`http`以及`https`协议外
还支持使用`l5`和`zkname`来获取需访问的ip和端口，格式如下：

```
    httpclient.Post("l5://12345:23412/upload").Execute()
    
    httpClient.Post("zkname://test.example.com/test").Execute()
```

使用l5以及zkname的请求默认使用http协议，如需使用https协议:

```
    httpClient.Post("zkname://test.example.com/upload").
        UseHTTPS(true).
        Execute()
```

# 监控及模块调用

### 代码中配置：
httpclient内部封装了监控和模调的类型：

```
    // 监控
    type Monitor struct {
    	Success         int `json:"success"` //成功
    	Fail            int `json:"fail"` //失败
    	RequestFail     int `json:"request_fail"` //请求失败
    	ResultParseFail int `json:"result_parse_fail"` //请求解析失败
    }

    //ModuleCall 模调
    type ModuleCall struct {
    	PositiveModID    int `json:"pos_mod_id"` //主调模块id
    	PositiveInterfID int `json:"pos_interf_id"` //主调接口id
    	PassiveModID     int `json:"pas_mod_id"` //被调模块id
    	PassiveInterfID  int `json:"pas_interf_id"` //被调接口id
    }
```

需要设置监控时，httpclient提供了两个接口:

```
    //SetMonitor 设置监控
    func (client *HttpClient) SetMonitor(monitor *Monitor) *HttpClient 

    //SetModCall 设置模调
    func (client *HttpClient) SetModCall(mcall *ModuleCall) *HttpClient 
```

模调用和监控可以单独进行设置，如果不设置则不尽兴上报,示例:

```
    httpclient.Post("zkname://test.lewis.com/upload").
        SetMonitor(&httpclient.Monitor{34001,34002,34003,34004}).
        Execute()
```

### 配置文件读取：
除了上述方式，httpclient还提供了从文件加载监控配置的选项。使用方法如下:

**main.go**

```
    httpclient.Post("zkname://test.lewis.com/upload").
        LoadConfig("../conf/conf.json","uploadConfig").
        Execute()
```

**conf.json**

```
    {
        "uploadConfig":{
            "success":210001,
            "fail":210002,
            "request_fail":210003,
            "result_parse_fail":210004,
            "pos_mod_id":123874,
            "pos_interf_id":215123,
            "pas_mod_id":129872,
            "pas_interf_id":245981
        }
        // other configs ...
    }
```

# 错误处理
**目前未对错误进行具体类型细化，用户还不能知道具体是中间哪一步出错，只能进行统一处理。细化错误待实现**

