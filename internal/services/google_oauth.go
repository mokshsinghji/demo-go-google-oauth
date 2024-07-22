package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mokshsinghji/go-google-oauth/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"net/http"
	"time"
)

const oauthStateCookieName = "oauth-state"
const googleOauthApiUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func getGoogleOauthConfig() *oauth2.Config {
	var googleConfig = &oauth2.Config{
		ClientID:     config.GoogleClientId,
		ClientSecret: config.GoogleClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  config.RedirectUrl,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	return googleConfig
}

func RedirectToGoogleLogin(c *fiber.Ctx) error {

	var googleConfig = getGoogleOauthConfig()

	state := CreateGoogleAuthorisationStateCookie(c)

	url := googleConfig.AuthCodeURL(state)

	err := c.Redirect(url, http.StatusTemporaryRedirect)
	if err != nil {
		return err
	}

	return nil
}

func CreateGoogleAuthorisationStateCookie(c *fiber.Ctx) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	state := base64.URLEncoding.EncodeToString(b)
	cookie := fiber.Cookie{Name: oauthStateCookieName, Value: state, Expires: expiration}
	c.Cookie(&cookie)

	fmt.Println("State Cookie:", state)

	return state
}

func GoogleLoginCallback(c *fiber.Ctx) error {
	var oauthStateCookie = c.Cookies(oauthStateCookieName)

	if oauthStateCookie != c.FormValue("state") {
		fmt.Println("oAuthStateCookie is:", oauthStateCookie, "state form value is", c.FormValue("state"))
		err := c.SendStatus(fiber.StatusBadRequest)
		if err != nil {
			return err
		}
		return nil
	}

	data, err := getGoogleUserData(c.FormValue("code"))
	if err != nil {
		return err
	}

	err = c.SendString(data)
	if err != nil {
		return err
	}

	return nil
}

func getGoogleUserData(code string) (string, error) {
	googleConfig := getGoogleOauthConfig()

	userToken, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", err
	}
	response, err := http.Get(googleOauthApiUrl + userToken.AccessToken)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(contents[:]), nil
}
