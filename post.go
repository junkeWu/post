package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var GetCsrfTokenUrl = "https://xskydata.jobs.feishu.cn/api/v1/csrf/token"
var GetCsrfTokenErrUrl = "https://xskydata.jobs.feishu.cn/api/v1/csrf"
var GetPostUrl = "https://xskydata.jobs.feishu.cn/api/v1/search/job/posts"

func main() {
	log.Println("start get posts...")
	respData, err := getCsrfToken(GetCsrfTokenUrl, map[string]int{"portal_entrance": 1})
	if err != nil {
		log.Println("get token failed!")
		return
	}
	postReq := GetPostDataRequest{
		Limit:             1,
		Offset:            0,
		PortalType:        6,
		JobFunctionIdList: nil,
		PortalEntrance:    1,
	}
	resp, err := post(GetPostUrl, respData.Data.Token, postReq)
	if err != nil {
		log.Printf("get post data failed, err= %v", err)
	}
	file, err := os.OpenFile("./data/post.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		log.Println("create file failed")
		return
	}
	var postData GetPostDataResp
	err = json.Unmarshal([]byte(resp), &postData)
	if err != nil {
		log.Println("resp marshal failed")
		return
	}
	writer := bufio.NewWriter(file)
	for i := 0; i < postData.Data.Count/10+1; i++ {
		respList, err := post(GetPostUrl, respData.Data.Token, GetPostDataRequest{
			Limit:             10,
			Offset:            i * 10,
			PortalType:        6,
			JobFunctionIdList: nil,
			PortalEntrance:    1,
		})
		if err != nil {
			log.Printf("get post faild, err =%v", err)
			continue
		}
		writer.WriteString(respList)
	}
	writer.Flush()
}

func getCsrfToken(url string, param map[string]int) (*GetTokenRespData, error) {

	paramBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(paramBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respData GetTokenRespData
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &respData)
	if err != nil {
		return nil, err
	}
	if respData.Code != 0 {
		return nil, errors.New("resp data' code is not equal zero")
	}
	return &respData, nil
}

func post(url, token string, body GetPostDataRequest) (string, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("accept-encoding", "deflate")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36\",\"application/json")
	cookie := fmt.Sprintf("atsx-csrf-token=%s", token)
	req.Header.Add("cookie", cookie)
	req.Header.Add("x-csrf-token", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var postResp GetPostDataResp
	err = json.Unmarshal(data, &postResp)
	if err != nil {
		return "", err
	}
	if postResp.Code != 0 {
		return "", errors.New("get post failed ")
	}
	return string(data), nil
}
