# WxLogin-Go 微信登录GO版本

WXSDK.cs文件中
```go
const (
	appid             = "appid"
	secret            = "secret"
	base_url          = "https://api.weixin.qq.com/"
	get_access_token  = base_url + "sns/oauth2/access_token"
	get_user_info     = base_url + "sns/userinfo"
	get_refresh_token = base_url + "sns/oauth2/refresh_token"
)
```
appid和secret，需要替换成对应的参数；

----
本项目只实现了
1. AccessToken: access_token的获取
2. RefreshToken: access_token的刷新
3. UserInfo: 用户信息的获取
----
可以直接使用但愿测试进行开箱体验；

在使用的过程中，需要先请求以下地址：
```
https://open.weixin.qq.com/connect/qrconnect?appid=appid&redirect_uri=redirect_uri&response_type=code&scope=snsapi_login&state=STATE#wechat_redirect
```

1. appid:微信开放平台申请完应用之后会提供，可在微信开放平台处查看
1. redirect_uri:扫码后回调地址
1. response_type:填code
1. scope:应用授权作用域，拥有多个作用域用逗号（,）分隔，网页应用目前仅填写snsapi_login即
1. state:用于保持请求和回调的状态，授权请求后原样带回给第三方。该参数可用于防止csrf攻击（跨站请求伪造攻击），建议第三方带上该参数，可设置为简单的随机数加session进行校验

其次，就可以使用本项目实现的三个接口；

## 开箱即用！

在使用中遇到任何问题欢迎留言！
