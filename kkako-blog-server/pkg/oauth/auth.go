package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"kkako-blog/internal/model"
	"kkako-blog/pkg/errcode"
	"net/http"
)

type githubConf struct {
	ClientId    string `json:"client_id"`
	RedirectUrl string `json:"redirect_url"`
	Secret      string `json:"secret"`
}

type githubToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type AuthUserInfo struct {
	ID int `json:"id"`
	Username string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
}

var GithubConf = &githubConf{}

func InitOauth(viper *viper.Viper) {
	err := viper.Sub("github").Unmarshal(GithubConf)
	if err != nil {
		panic(err)
	}
}

func GetUserByAuth(authType int, code string) (*AuthUserInfo, error) {
	switch authType {
	case model.IdentityTypeGithub:
		token, err := GetGithubToken(code)
		if err != nil {
			return nil, err
		}
		userInfo, err := GetGithubUserInfo(token)
		if err != nil {
			return nil, err
		}
		return userInfo, nil
	default:
		return nil, errcode.NewErr("第三方登录失败")
	}
}

func GetGithubToken(code string) (githubToken, error) {
	var token githubToken
	authUrl := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		GithubConf.ClientId, GithubConf.Secret, code,)
	req, err := http.NewRequest(http.MethodGet, authUrl, nil)
	if err != nil {
		return token, errors.Wrap(err, "github req生成登录失败")
	}
	req.Header.Set("accept", "application/json")
	var httpClient = http.Client{}
	res, err := httpClient.Do(req);
	if err != nil {
		return token, errors.Wrap(err, "github登录失败")
	}

	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return token, errors.Wrap(err, "github登录失败")
	}
	return token, nil
}

func GetGithubUserInfo(token githubToken) (*AuthUserInfo, error) {
	var userInfoUrl = "https://api.github.com/user"	// github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, errors.Wrap(err, "github登录失败")
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token.AccessToken))

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, errors.Wrap(err, "github登录失败")
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = &AuthUserInfo{}
	if err = json.NewDecoder(res.Body).Decode(userInfo); err != nil {
		return nil, errors.Wrap(err, "github登录失败")
	}
	return userInfo, nil
}