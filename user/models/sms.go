package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/microsvs/base/cmd/discovery"
	"github.com/microsvs/base/pkg/errors"
	"github.com/microsvs/base/pkg/log"
)

// 第三方短信服务
func SendShortMessage(phone string, verifyCode string) (bool, error) {
	// curl -d "templateId=freego02&to=18098919046&company=jiwei&channel=dx9636&context=1234" "http://192.168.3.83:9401/now"
	sendUrl := discovery.KVRead("params/shortmessage/sendurl", "http://10.2.40.67:9401/now")
	company := discovery.KVRead("params/shortmessage/company", "jiwei")
	channel := discovery.KVRead("params/shortmessage/channel", "of8679")
	template := discovery.KVRead("params/shortmessage/template", "ofo_rompers_verifycode")
	tokenname := discovery.KVRead("params/shortmessage/tokenname", "x-ofo-token")
	token := discovery.KVRead("params/shortmessage/token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxODYzNjY0Nzk2MiIsIm5hbWUiOiLmt7HlnLMifQ.EjXJEjEWGKcsI896Mx6BUCbtnlq_gcnQ2NjpQaZSLkE")
	data := make(url.Values)
	data["templateId"] = []string{template}
	data["to"] = []string{phone}
	data["company"] = []string{company}
	data["channel"] = []string{channel}
	data["context"] = []string{verifyCode}
	usrParam := data.Encode()
	smsClient := &http.Client{}
	req, err := http.NewRequest("POST", sendUrl, strings.NewReader(usrParam))
	if err != nil {
		log.ErrorRaw("[SendSMS] new request failed. err=%s", err)
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(tokenname, token)
	resp, err := smsClient.Do(req)
	if err != nil {
		log.ErrorRaw("[SendSMS] do request failed. err=%s", err)
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, errors.FGESendShortMessage
	} else {
		postRT := OfoMessageRT{}
		json.Unmarshal(body, &postRT)
		if postRT.Code == 200 {
			return true, nil
		} else {
			return false, errors.FGESendShortMessage
		}
	}
	return false, errors.FGESendShortMessage
}

type OfoMessageRT struct {
	Code int
	Err  string
}
