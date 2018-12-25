package utils

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/1046102779/base/pkg/timer"
	"github.com/segmentio/ksuid"
)

var chars = []byte("0123456789")

//Precision,取float 精度
//生成指定长度的随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(timer.Now.UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 版本对比
func CompareVersion(newVerCode string, curVerCode string) bool {
	newVers := strings.Split(newVerCode, ",")
	curVers := strings.Split(curVerCode, ",")
	for i := 0; i < len(newVers) && i < len(curVers); i++ {
		if newVers[i] > curVers[i] {
			return true
		} else if newVers[i] < curVers[i] {
			return false
		}
	}
	if len(newVers) > len(curVers) {
		return true
	}
	return false
}

// src="1;2;3" -> dest = [1, 2, 3]
func ConvertStrToInts(src string) (dest []int) {
	if src == "" {
		return
	}
	tmps := strings.Split(src, ";")
	for _, tmp := range tmps {
		s, _ := strconv.Atoi(tmp)
		dest = append(dest, s)
	}
	return
}

func GenerateUUID() string {
	return ksuid.New().String()
}

func ConvertToGraphqlList(tmps []string) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for _, tmp := range tmps {
		buffer.WriteString(fmt.Sprintf(" \"%s\" ", tmp))
	}
	buffer.WriteString("]")
	return buffer.String()
}

//GenerateVerifyCode generate verify code
func GenerateVerifyCode() string {
	bs := []byte{}
	for i := 0; i < 4; i++ {
		bs = append(bs, chars[rand.Intn(10)])
	}
	return string(bs)
}
