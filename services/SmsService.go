package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unsafe"
	"autumn/tools/cfg"
	"autumn/tools/i18n"
	"autumn/models"
	"log"
)

type SmsService struct {
}

//发送短信验证码
func (this *SmsService) Send(mobile string, code string, country string, language string) error {
	msg := ""

	msg = fmt.Sprintf(i18n.Get("sms", "verify",language),
		code, strconv.Itoa(int(cfg.Get("sms", "expire").Int())/60))

	log.Println("Service.Sms.Send:", msg)

	params := make(map[string]string)

	params["zone"] = country
	params["msg"] = msg
	params["phone"] = mobile
	params["smsType"] = "0"

	bytesData, err := json.Marshal(params)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", cfg.Get("sms", "api").String(), reader)
	if err != nil {
		return err

	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err

	}
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err

	}

	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))

	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(*str), &dat); err == nil {

		if dat["code"].(float64) == 0 {
			return nil
		}

		return errors.New(dat["errorMsg"].(string))
	}

	return errors.New("send-error")

}

func (*SmsService) Verify(mobile string, code string, usage bool) int {

	//验证码输入错误
	if len(code) != 6 {
		return 11000
	}

	//短信测试开启
	if code == cfg.Get("sms", "test").String() {
		return 0
	}

	info := new(models.Sms).Info(mobile, code)

	if info.Id == 0 {
		return 11001
	}

	if time.Now().Unix() > info.Expire {
		return 11002
	}

	if info.Status == 1 {
		return 11003
	}

	if usage {
		info.Status = 1
		info.Update()
	}

	return 0
}

//验证短信
func (this *SmsService) VerifySms(mobile string, verifyCode string, isUsage bool) int {
	if mobile != "" {
		if verifyCode == "" {
			return 11010
		} else {
			sms_code := this.Verify(mobile, verifyCode, isUsage)
			if sms_code > 0 {
				return sms_code
			}
		}
		return 0
	}
	return -1
}
