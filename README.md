# GoHttp
A Chain call http lib for go.Support JSON &amp;&amp; JSONP &amp;&amp; String get directly

#Usage 
```
    str := ""
    err:=httpclient.Post("http://localhost:8080/upload").
            ContentType("application/json").
            Body("{\"name\":\"lewis\"}").
            AddHeader("Referer","localhost").
            AddCookie("uin","guest").
            GetString(&str).
            Execute()
```