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
            Execute()
```


