package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func HttpPostBody(uri string, accept string, bodyData []byte) (retBody []byte, err error) {
	var (
		resp *http.Response
	)
	if resp, err = http.Post(uri, accept, bytes.NewBuffer(bodyData)); err != nil {
		return
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if resp.StatusCode != 200 {
		err = errors.New("http status not ok." + resp.Status)
		return
	}
	retBody, err = ioutil.ReadAll(resp.Body)
	return
}

func HttpPostJson(uri string, accept string, bodyData []byte) (retJson map[string]interface{}, err error) {
	var (
		retBody []byte
	)
	if retBody, err = HttpPostBody(uri, accept, bodyData); err == nil {
		err = json.Unmarshal(retBody, &retJson)
	}
	return
}
