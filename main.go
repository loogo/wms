package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const base_url = "http://192.168.0.218/bitbucket/rest/api/1.0/projects"

type bitbuchet struct {
	name     string
	password string
	url      string
	body     string
	client   *http.Client
	method   string
}

func (bb *bitbuchet) basicAuth(request *http.Request) {
	request.SetBasicAuth(bb.name, bb.password)
}

func (bb *bitbuchet) doRequest() (*http.Response, error) {
	var jsonBody io.Reader
	if bb.body != "" {
		jsonBody = bytes.NewBuffer([]byte(bb.body))
	}
	req, err := http.NewRequest(bb.method, bb.url, jsonBody)
	if err != nil {
		return nil, err
	}
	bb.basicAuth(req)
	req.Header.Set("Content-Type", "application/json")
	resp, err := bb.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/logs", 301)
	})
	http.HandleFunc("/pull", pullHandler)
	http.HandleFunc("/logs", jenkinsLogHandler)
	log.Println("listening on port: 8087")
	log.Fatal(http.ListenAndServe(":8087", nil))
}

func pullHandler(w http.ResponseWriter, req *http.Request) {
	var bb = &bitbuchet{
		name:     "1009751265",
		password: "wangcl",
		client:   &http.Client{},
	}
	bb.url = fmt.Sprintf("%s/%s/repos/%s/pull-requests", base_url, "SCM-DEV", "dyedp-wms")
	bb.method = "POST"
	bb.body = `{
				"title": "auto pull request by api",
				"description": "自动创建的合并",
				"state": "OPEN",
				"open": true,
				"closed": false,
				"fromRef": {
					"id": "refs/heads/master",
					"repository":{
						"slug":"dyedp-wms",
						"name":null,
						"project":{
							"key":"SCM-DEV"
						}
					}
				},
				"toRef": {
					"id": "refs/heads/jenkins",
					"repository":{
						"slug":"dyedp-wms",
						"name":null,
						"project":{
							"key":"SCM-DEV"
						}
					}
				},
				"locked": false,
				"reviewers": [
					{
						"user": {
							"name": "1009751265"
						}
					}
				]
			}`
	resp, err := bb.doRequest()
	defer resp.Body.Close()
	if err != nil {
		io.WriteString(w, err.Error())
	} else {
		var data map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&data)
		if _, ok := data["errors"]; ok {
			io.WriteString(w, fmt.Sprintf("%s", data["errors"]))
		} else {
			pullId := int(data["id"].(float64))
			bb.url = fmt.Sprintf("%s/%s/repos/%s/pull-requests/%d/merge?version=0", base_url, "SCM-DEV", "dyedp-wms", pullId)
			bb.method = "POST"
			bb.body = ""
			resp, err = bb.doRequest()
			defer resp.Body.Close()
			if err != nil {
				io.WriteString(w, err.Error())
			} else {
				request, _ := http.NewRequest("POST", "http://172.16.105.146:8080/job/yiritong.dyedp.wms/build", nil)
				request.SetBasicAuth("admin", "admin2018")
				resp, _ = bb.client.Do(request)
				defer resp.Body.Close()
				//io.WriteString(w, "执行成功，等待Jenkins构建完成，大概5分钟左右")
				//io.Copy(w, resp.Body)
				http.Redirect(w, req, "/logs", 301)
			}
		}
	}
}

func jenkinsLogHandler(w http.ResponseWriter, req *http.Request) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://172.16.105.146:8080/job/yiritong.dyedp.wms/lastBuild/consoleText", nil)
	request.SetBasicAuth("admin", "admin2018")
	resp, _ := client.Do(request)
	defer resp.Body.Close()
	//io.WriteString(w, "执行成功，等待Jenkins构建完成，大概5分钟左右")
	io.Copy(w, resp.Body)
}
