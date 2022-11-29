package utils_tool

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// GetRequest 		发送消息
// url				地址
// token			token
func GetRequest(url, token string) ([]byte, error) {
	log.Println("-----------------------------> [GET]: ", url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// 添加请求头
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Set("SOC-TOKEN", token)

	// 添加cookie
	//cookie1 := &http.Cookie{
	//	Name:  "aaa",
	//	Value: "aaa-value",
	//}
	//req.AddCookie(cookie1)

	// 发送请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("-----------------------------> 响应信息: ", string(response))
	log.Println("-----------------------------> END \n\r")

	return response, nil
}

// PostRequest 	post方式发送消息
// url 			请求地址
// token 		token
// param 		参数
func PostRequest(url, token string, param interface{}) ([]byte, error) {
	log.Println("-----------------------------> [POST]: ", url)
	client := http.Client{}
	// 带数据 json 类型
	b1, _ := json.Marshal(&param)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b1))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// 添加请求头
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Set("SOC-TOKEN", token)

	//发送
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	//关闭
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(resp.Body)

	//读取
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	log.Println("-----------------------------> 响应信息: ", string(response))
	log.Println("-----------------------------> END \n\r")

	return response, nil
}
