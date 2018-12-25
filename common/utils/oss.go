package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	ACCESS_KEY_ID     = "LTAIt6IV8eSjEwOG"
	ACCESS_KEY_SECRET = "LjIIug1AJfPAErsbeDPi0hKgOkohVM"
	END_POINT         = "oss-cn-shanghai.aliyuncs.com"
	BUCKET            = "gewuwei"
	ROLE_ARN          = "acs:ram::1824796786346257:role/role-gewuwei"
	SESSION_NAME      = "gewuwei"
	DURATION_TIME     = "3600"
)

//--------------------------------------阿里云oss相关结构--------------------------------------//
type OssStsClientInfo struct {
	ChildAccountKeyId  string
	ChildAccountSecret string
	RoleAcs            string
}

func (cli *OssStsClientInfo) GenerateSignatureUrl(sessionName, durationSeconds string) (string, error) {
	assumeUrl := "SignatureVersion=1.0"
	assumeUrl += "&Format=JSON"
	assumeUrl += "&Timestamp=" + url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	assumeUrl += "&RoleArn=" + url.QueryEscape(cli.RoleAcs)
	assumeUrl += "&RoleSessionName=" + sessionName
	assumeUrl += "&AccessKeyId=" + cli.ChildAccountKeyId
	assumeUrl += "&SignatureMethod=HMAC-SHA1"
	assumeUrl += "&Version=2015-04-01"
	assumeUrl += "&Action=AssumeRole"
	tmp := uuid.NewV4()
	assumeUrl += "&SignatureNonce=" + tmp.String()
	assumeUrl += "&DurationSeconds=" + durationSeconds

	// 解析成V type
	signToString, err := url.ParseQuery(assumeUrl)
	if err != nil {
		return "", err
	}

	// URL顺序化
	result := signToString.Encode()

	// 拼接
	StringToSign := "GET" + "&" + "%2F" + "&" + url.QueryEscape(result)

	// HMAC
	hashSign := hmac.New(sha1.New, []byte(cli.ChildAccountSecret+"&"))
	hashSign.Write([]byte(StringToSign))

	// 生成signature
	signature := base64.StdEncoding.EncodeToString(hashSign.Sum(nil))

	// Url 添加signature
	assumeUrl = "https://sts.aliyuncs.com/?" + assumeUrl + "&Signature=" + url.QueryEscape(signature)

	return assumeUrl, nil
}

// 请求构造好的URL,获得授权信息
// TODO: 安全认证 HTTPS
func (cli *OssStsClientInfo) GetStsResponse(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}
