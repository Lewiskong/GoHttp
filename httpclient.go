package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type HttpMethod string

const (
	httpGet  HttpMethod = "GET"
	httpPost HttpMethod = "POST"
)

type HttpClient struct {
	Error    error
	Response *http.Response

	tasks   map[*interface{}]string
	method  HttpMethod
	request *http.Request
}

func Get(rawurl string) *HttpClient {
	client := CreateDefault()
	client.method = httpGet

	client.request, client.Error = http.NewRequest("GET", "", nil)
	if client.Error != nil {
		return client
	}

	return client.handle(rawurl)
}

func Post(rawurl string) *HttpClient {
	client := CreateDefault()
	client.method = httpPost
	client.request, client.Error = http.NewRequest("POST", "", nil)
	if client.Error != nil {
		return client
	}

	return client.handle(rawurl)
}

func CreateDefault() *HttpClient {
	client := new(HttpClient)
	client.tasks = map[*interface{}]string{}
	return client
}

func (client *HttpClient) handle(rawurl string) *HttpClient {

	if client.Error != nil {
		return client
	}

	client.request.URL, client.Error = url.ParseRequestURI(rawurl)
	if client.Error != nil {
		return client
	}
	// fmt.Println(client.request.URL)
	return client
}

func (client *HttpClient) ContentType(contentType string) *HttpClient {
	if client.Error != nil {
		return client
	}

	client.request.Header.Set("Content-Type", contentType)

	return client
}

func (client *HttpClient) Body(body io.Reader) *HttpClient {
	if client.Error != nil {
		return client
	}

	client.request.Body = ioutil.NopCloser(body)

	return client
}

func (client *HttpClient) AddHeader(key, value string) *HttpClient {
	if client.Error != nil {
		return client
	}

	client.request.Header.Add(key, value)

	return client
}

func (client *HttpClient) AddCookie(name, value string) *HttpClient {
	if client.Error != nil {
		return client
	}

	ck := new(http.Cookie)
	ck.Name = name
	ck.Value = value

	client.request.AddCookie(ck)

	return client
}

func (client *HttpClient) GetString(v interface{}) *HttpClient {
	if client.Error != nil {
		return client
	}
	client.tasks[&v] = "string"
	return client
}

func (client *HttpClient) GetJson(v interface{}) *HttpClient {
	if client.Error != nil {
		return client
	}
	client.tasks[&v] = "json"
	return client
}

func (client *HttpClient) GetJsonp(v interface{}) *HttpClient {
	if client.Error != nil {
		return client
	}
	client.tasks[&v] = "jsonp"
	return client
}

func (client *HttpClient) GetJce(v interface{}) *HttpClient {
	if client.Error != nil {
		return client
	}
	client.tasks[&v] = "jce"
	return client
}

func (client *HttpClient) Execute() (err error) {
	if client.Error != nil {
		return client.Error
	}

	c := http.DefaultClient
	rsp, err := c.Do(client.request)

	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	//debug
	// fmt.Println(string(content[:]))
	for value, tp := range client.tasks {
		switch tp {
		case "string":
			err := setString(value, content)
			if err != nil {
				return err
			}

		case "json":
			err := setJSON(value, content)
			if err != nil {
				return err
			}

		case "jsonp":
			err := setJSONP(value, content)
			if err != nil {
				return err
			}

		case "jce":

		default:
		}
	}
	if err != nil {
		return err
	}
	//debug
	// fmt.Println(string(content[:]))
	return nil
}

func setString(v *interface{}, content []byte) error {
	interf := *v
	vtype := reflect.TypeOf(interf)
	vvalue := reflect.ValueOf(interf)

	if vtype.Kind() != reflect.Ptr {
		return fmt.Errorf("error happened when parse json : the param obj %s must be pointer", vtype.Kind().String())
	}
	vvalue = reflect.Indirect(vvalue)
	vvalue.SetString(string(content[:]))
	return nil
}

// SetJSON hello
func setJSON(v *interface{}, content []byte) error {
	interf := *v
	err := json.Unmarshal(content, interf)
	return err
}

func setJSONP(v *interface{}, content []byte) error {
	interf := *v
	str := string(content[:])
	start := strings.Index(str, "{")
	end := strings.LastIndex(str, "}")
	if start < 0 || end < 0 {
		return fmt.Errorf("Parse jsonp error , wrong jsonp format : %s ", str)
	}
	err := json.Unmarshal(content[start:end+1], interf)
	return err
}

func setJce(v *interface{}) error {
	return nil
}
