//接口类型：互亿无线触发短信接口，支持发送验证码短信、订单通知短信等。
// 账户注册：请通过该地址开通账户http://user.ihuyi.com/register.html
// 注意事项：
//（1）调试期间，请使用用系统默认的短信内容：您的验证码是：【变量】。请不要把验证码泄露给其他人。
//（2）请使用 APIID 及 APIKEY来调用接口，可在会员中心获取；
//（3）该代码仅供接入互亿无线短信接口参考使用，客户可根据实际需要自行编写；

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
	_account := "C94969796"                         //查看用户名 登录用户中心->验证码通知短信>产品总览->API接口信息->APIID
	_password := "08cb6d54a79fe8bc0432e0d6484be982" //查看密码 登录用户中心->验证码通知短信>产品总览->API接口信息->APIKEY
	_mobile := "0648"
	_content := "您的验证码是：9552。请不要把验证码泄露给其他人。"
	v.Set("account", _account)
	v.Set("password", GetMd5String(_account+_password+_mobile+_content+_now))
	v.Set("mobile", _mobile)
	v.Set("content", _content)
	v.Set("time", _now)
	body := strings.NewReader(v.Encode()) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://106.ihuyi.com/webservice/sms.php?method=Submit&format=json", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//fmt.Printf("%+v\n", req) //看下发送的结构

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
}
