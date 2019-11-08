package wxlogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	appid             = "wxff26f000ea89ecd8"
	secret            = "4a25f27f009a4b2e03d3d9c2bcf50298"
	base_url          = "https://api.weixin.qq.com/"
	get_access_token  = base_url + "sns/oauth2/access_token"
	get_user_info     = base_url + "sns/userinfo"
	get_refresh_token = base_url + "sns/oauth2/refresh_token"
)

func GetAccessToken(code string) (dat map[string]interface{}, err error) {
	dic := make(map[string]string)
	dic["appid"] = appid
	dic["secret"] = secret
	dic["code"] = code
	dic["grant_type"] = "authorization_code"
	urlStr := mapToString(dic, get_access_token+"?")
	result := httpGet(urlStr)
	if err = json.Unmarshal([]byte(result), &dat); err == nil {
		fmt.Println(dat)
	} else {
		fmt.Println(err)
	}
	return
}

func GetRefreshToken(access_token string, openid string) (dat map[string]interface{}, err error) {
	dic := make(map[string]string)
	dic["appid"] = appid
	dic["access_token"] = access_token
	dic["openid"] = openid
	urlStr := mapToString(dic, get_refresh_token+"?")
	result := httpGet(urlStr)
	if err = json.Unmarshal([]byte(result), &dat); err == nil {
		fmt.Println(dat)
	} else {
		fmt.Println(err)
	}
	return
}

func GetUserInfo(access_token string, openid string) (dat map[string]interface{}, err error) {
	dic := make(map[string]string)
	dic["appid"] = appid
	dic["access_token"] = access_token
	dic["openid"] = openid
	urlStr := mapToString(dic, get_user_info+"?")
	result := httpGet(urlStr)
	if err = json.Unmarshal([]byte(result), &dat); err == nil {
		fmt.Println(dat)
	} else {
		fmt.Println(err)
	}
	return
}

// 解析字典
func mapToString(dic map[string]string, url string) string {
	urlStr := bytes.NewBufferString(url)
	index := 1
	for i, v := range dic {
		urlStr.WriteString(i + "=" + v)
		if index < len(dic) {
			urlStr.WriteString("&")
		}
		index++
	}
	return urlStr.String()
}

func httpGet(url string) string {

	response, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}
