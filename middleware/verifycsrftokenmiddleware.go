package middleware

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qclaogui/goforum/model"
)

const (
	csrfSecret = "forum-csrfSecret"
	csrfSalt   = "forum-csrfSalt"
	csrfToken  = "forum-csrfToken"
)

/**
 * Determine if the HTTP request uses a ‘read’ verb.
 *
 * @return bool
 */
func isReading(c *gin.Context) bool {

	for _, v := range []string{"HEAD", "GET", "OPTIONS"} {
		if v == c.Request.Method {
			return true
		}
	}

	return false
}

func VerifyCsrfToken() gin.HandlerFunc {

	return func(c *gin.Context) {

		session := sessions.Default(c)
		var salt string
		if s, ok := session.Get(csrfSalt).(string); !ok || len(s) == 0 {
			CsrfToken(c)
			c.Next()
			return
		} else {
			salt = s
		}
		session.Delete(csrfSalt)

		if isReading(c) || tokensMatch(c, salt) {
			CsrfToken(c)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_code": 10009,
			"message":    "CSRF token mismatch",
		})
		return
	}
}

/**
 * Determine if the session and input CSRF tokens match.
 *
 * @return bool
 */
func tokensMatch(c *gin.Context, salt string) bool {

	return CreateCsrfToken(salt) == getCsrfTokenFromRequest(c)
}
func getCsrfTokenFromRequest(c *gin.Context) string {

	r := c.Request
	if cToken := r.FormValue("_token"); len(cToken) > 0 {
		return cToken
	} else if cToken := r.URL.Query().Get("_token"); len(cToken) > 0 {
		return cToken
	} else if cToken := r.Header.Get("X-CSRF-TOKEN"); len(cToken) > 0 {
		return cToken
	} else if cToken := r.Header.Get("X-XSRF-TOKEN"); len(cToken) > 0 {
		return cToken
	}
	return ""
}

func CreateCsrfToken(salt string) string {
	h := sha1.New()
	io.WriteString(h, csrfSecret+salt)
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func CsrfToken(c *gin.Context) string {

	session := sessions.Default(c)
	if t, ok := c.Get(csrfToken); ok {
		return t.(string)
	}

	salt := model.RandomString(16)
	session.Set(csrfSalt, salt)
	session.Save()

	token := CreateCsrfToken(salt)
	c.Set(csrfToken, token)

	return token
}
