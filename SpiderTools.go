package spidertools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func httpClient() *http.Client {

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	return httpClient
}

// 创建请求头
func createHeaders() http.Header {
	return http.Header{
		"User-Agent": []string{"your-user-agent-here"}, // 这里替换成你的User-Agent
	}
}

// 发送GET请求
func Get(url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, values := range headers {
		req.Header[key] = values
	}
	return httpClient().Do(req)
}

// 发送POST请求
func Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	var jsonBytes []byte
	var err error
	if body != nil {
		jsonBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for key, values := range headers {
		req.Header[key] = values
	}
	return httpClient().Do(req)
}

func Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	var jsonBytes []byte
	var err error
	if body != nil {
		jsonBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for key, values := range headers {
		req.Header[key] = values
	}
	return httpClient().Do(req)
}

// Delete - 发送DELETE请求
func Delete(url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	for key, values := range headers {
		req.Header[key] = values
	}
	return httpClient().Do(req)
}

// Options - 发送OPTIONS请求
func Options(url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		return nil, err
	}
	for key, values := range headers {
		req.Header[key] = values
	}
	return httpClient().Do(req)
}

// Head - 发送HEAD请求
func Head(url string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}
	for key, values := range headers {
		req.Header[key] = values
	}
	return httpClient().Do(req)
}

// Patch - 发送PATCH请求
func Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	var jsonBytes []byte
	var err error
	if body != nil {
		jsonBytes, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for key, values := range headers {
		req.Header[key] = values
	}
	return httpClient().Do(req)
}

// HandleResponse - 处理HTTP响应
func HandleResponse(res *http.Response, err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("Response Status: %s\n", res.Status)
	fmt.Printf("Response Body: %s\n", body)
}
