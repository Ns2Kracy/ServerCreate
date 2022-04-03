//接口类型：互亿无线语音通知接口。
//账户注册：请通过该地址开通账户http://sms.ihuyi.com/register.html
//注意事项：
//（1）调试期间，请仔细阅读接口文档；
//（2）请使用APIID（查看APIID请登录用户中心->语音通知->产品总览->APIID）及APIkey来调用接口；
//（3）该代码仅供接入互亿无线语音通知接口参考使用，客户可根据实际需要自行编写；

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
func main() {
	v := url.Values{}
	_now := strconv.FormatInt(time.Now().Unix(), 10)
	//fmt.Printf(_now)
	_account := "VM45617014"                        //用户名是登录用户中心->语音通知->产品总览->APIID
	_password := "b55dc062bd7902dacc6d0c3aca814e50" //查看密码请登录用户中心->语音通知->产品总览->APIKEY
	_mobile := "18723849816"
	_content := "您的订单号是：0648。已由顺风快递发出，请注意查收。"
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	//body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	body := strings.NewReader(v.Encode()) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://api.vm.ihuyi.com/webservice/voice.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}
