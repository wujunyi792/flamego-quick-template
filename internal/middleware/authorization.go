package middleware

import (
	"github.com/flamego/flamego"
	"github.com/wujunyi792/gin-template-new/pkg/jwt"
)

func Authorization(c flamego.Context, r flamego.Render) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		UnAuthorization(r)
		return
	}
	entry, err := jwt.ParseToken(token)
	if err != nil {
		UnAuthorization(r)
		return
	}
	c.Map(entry.Info)
}
