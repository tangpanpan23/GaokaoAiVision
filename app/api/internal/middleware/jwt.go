package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type JwtAuthMiddleware struct {
	secret string
}

func NewJwtAuthMiddleware(secret string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		secret: secret,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			httpx.ErrorCtx(r.Context(), w, &struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: 401,
				Msg:  "未授权访问",
			})
			return
		}

		// 解析Bearer Token
		if len(auth) <= 7 || auth[:7] != "Bearer " {
			httpx.ErrorCtx(r.Context(), w, &struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: 401,
				Msg:  "无效的授权格式",
			})
			return
		}

		tokenString := auth[7:]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.secret), nil
		})

		if err != nil || !token.Valid {
			httpx.ErrorCtx(r.Context(), w, &struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: 401,
				Msg:  "无效的访问令牌",
			})
			return
		}

		// 将用户信息添加到请求上下文中
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			r.Header.Set("X-User-ID", claims["user_id"].(string))
			r.Header.Set("X-Open-ID", claims["open_id"].(string))
		}

		next(w, r)
	}
}

