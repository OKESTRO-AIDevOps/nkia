package omodules

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var CONFIG_JSON = GetConfig()

type OAuthStruct struct {
	ID             string `json:"id"`
	EMAIL          string `json:"email"`
	VERIFIED_EMAIL bool   `json:"verified_email"`
	PICTURE        string `json:"picture"`
}

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:1337/oauth2/google/callback",
	ClientID:     CONFIG_JSON["GOOGLE_OAUTH_CLIENT_ID"],
	ClientSecret: CONFIG_JSON["GOOGLE_OAUTH_CLIENT_SECRET"],
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func GenerateStateOauthCookie(c *gin.Context) string {

	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	session := sessions.Default(c)

	session.Set("OSID", state)
	session.Save()

	return state
}

func GetUserDataFromGoogle(code string) ([]byte, error) {

	token, err := GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}
