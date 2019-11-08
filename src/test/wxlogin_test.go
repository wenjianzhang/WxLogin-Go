package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"wxlogin-go/src/wxlogin"
)

func TestGetAccessToken(t *testing.T) {
	if ref, err := wxlogin.GetAccessToken("021sHMYH17ZBE00mSLZH1HfKYH1sHMYj"); err != nil {
		t.Error(err)
	} else {
		t.Log(ref)
	}
}

func TestGetRefreshToken(t *testing.T) {
	if ref, err := wxlogin.GetRefreshToken("27_9jsRNYzF4m6NcXN7_QYqKHzR31M4LUAZakWqpLH6YvC5p06Iy7e8gdaopuYKYTuZ5GpGo5SQJtLU9LRdLgUncQww8vBcvjx58M6Q5MNT0_Q", "oE2RCwsWY3WM0Xr9CWfvXJuazzB0"); err != nil {
		t.Error(err)
	} else {
		t.Log(ref)
	}
}

func TestGetUserInfo(t *testing.T) {
	if ref, err := wxlogin.GetUserInfo("27_9jsRNYzF4m6NcXN7_QYqKHzR31M4LUAZakWqpLH6YvC5p06Iy7e8gdaopuYKYTuZ5GpGo5SQJtLU9LRdLgUncQww8vBcvjx58M6Q5MNT0_Q", "oE2RCwsWY3WM0Xr9CWfvXJuazzB0"); err != nil {
		t.Error(err)
	} else {
		t.Log(ref)
	}
}

func TestJson(t *testing.T) {
	s := "{\"errcode\":40029,\"errmsg\":\"invalid code, hints: [ req_id: A.Df2qLnRa-f8ItkA ]\"}"

	var errorInfo ErrorInfo
	if err := json.Unmarshal([]byte(s), &errorInfo); err == nil {
		fmt.Println(errorInfo)
		fmt.Println(errorInfo.Errcode)
	} else {
		fmt.Println(err)
	}
}

type AccessToken struct {
	RefreshToken
	Unionid string `json:"unionid"`
}

type RefreshToken struct {
	AccessToken  string `json:"access_token"`
	Expires      int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

type UserInfo struct {
	Openid     string
	Nickname   string
	Sex        int
	Province   string
	City       string
	Country    string
	Headimgurl string
	Privilege  []string
	Unionid    string
}

type ErrorInfo struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
