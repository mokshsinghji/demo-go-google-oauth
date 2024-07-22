package services

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/mokshsinghji/go-google-oauth/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"time"
)

var googleConfig = &oauth2.Config{
	ClientID:     config.GOOGLE_CLIENT_ID,
	ClientSecret: config.GOOGLE_CLIENT_SECRET,
	Endpoint:     google.Endpoint,
	RedirectURL:  config.REDIRECT_URL,
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
}

func RedirectToGoogleLogin(c *fiber.Ctx) error {
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
	cookie := fiber.Cookie{Name: "oauth-state", Value: state, Expires: expiration}
	c.Cookie(&cookie)

	return state
}
