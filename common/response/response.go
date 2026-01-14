package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Success 成功响应
func Success(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, map[string]interface{}{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// SuccessWithMsg 成功响应带消息
func SuccessWithMsg(w http.ResponseWriter, msg string, data interface{}) {
	httpx.OkJson(w, map[string]interface{}{
		"code": 200,
		"msg":  msg,
		"data": data,
	})
}

// Error 错误响应
func Error(w http.ResponseWriter, code int, msg string) {
	httpx.Error(w, &struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{
		Code: code,
		Msg:  msg,
	})
}

// BadRequest 参数错误
func BadRequest(w http.ResponseWriter, msg string) {
	Error(w, http.StatusBadRequest, msg)
}

// Unauthorized 未授权
func Unauthorized(w http.ResponseWriter, msg string) {
	Error(w, http.StatusUnauthorized, msg)
}

// Forbidden 禁止访问
func Forbidden(w http.ResponseWriter, msg string) {
	Error(w, http.StatusForbidden, msg)
}

// NotFound 未找到
func NotFound(w http.ResponseWriter, msg string) {
	Error(w, http.StatusNotFound, msg)
}

// InternalServerError 服务器内部错误
func InternalServerError(w http.ResponseWriter, msg string) {
	Error(w, http.StatusInternalServerError, msg)
}

// TooManyRequests 请求过于频繁
func TooManyRequests(w http.ResponseWriter, msg string) {
	Error(w, http.StatusTooManyRequests, msg)
}

// PageResponse 分页响应
type PageResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Total     int64       `json:"total"`
	Page      int         `json:"page"`
	PageSize  int         `json:"page_size"`
	Pages     int         `json:"pages"`
}

// SuccessPage 分页成功响应
func SuccessPage(w http.ResponseWriter, data interface{}, total int64, page, pageSize int) {
	pages := 0
	if pageSize > 0 {
		pages = int((total + int64(pageSize) - 1) / int64(pageSize))
	}

	httpx.OkJson(w, PageResponse{
		Code:     200,
		Msg:      "success",
		Data:     data,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		Pages:    pages,
	})
}

// ValidateError 验证错误响应
func ValidateError(w http.ResponseWriter, field, msg string) {
	Error(w, http.StatusBadRequest, field+": "+msg)
}

// FieldError 字段错误响应
func FieldError(w http.ResponseWriter, field, msg string) {
	Error(w, http.StatusBadRequest, field+": "+msg)
}

// SystemError 系统错误响应
func SystemError(w http.ResponseWriter, msg string) {
	InternalServerError(w, "系统错误: "+msg)
}

// DBError 数据库错误响应
func DBError(w http.ResponseWriter, msg string) {
	InternalServerError(w, "数据库错误: "+msg)
}

// CacheError 缓存错误响应
func CacheError(w http.ResponseWriter, msg string) {
	InternalServerError(w, "缓存错误: "+msg)
}

// APIError API调用错误响应
func APIError(w http.ResponseWriter, msg string) {
	InternalServerError(w, "API调用错误: "+msg)
}

// TimeoutError 超时错误响应
func TimeoutError(w http.ResponseWriter, msg string) {
	Error(w, http.StatusRequestTimeout, "请求超时: "+msg)
}

// JSONResponse JSON响应
type JSONResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// WriteJSON 写入JSON响应
func WriteJSON(w http.ResponseWriter, code int, msg string, data interface{}) {
	httpx.OkJson(w, JSONResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// OK 200响应
func OK(w http.ResponseWriter, data interface{}) {
	WriteJSON(w, http.StatusOK, "success", data)
}

// Created 201响应
func Created(w http.ResponseWriter, data interface{}) {
	WriteJSON(w, http.StatusCreated, "created", data)
}

// NoContent 204响应
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// Redirect 重定向响应
func Redirect(w http.ResponseWriter, url string, code int) {
	http.Redirect(w, nil, url, code)
}

// PermanentRedirect 永久重定向
func PermanentRedirect(w http.ResponseWriter, url string) {
	Redirect(w, url, http.StatusMovedPermanently)
}

// TemporaryRedirect 临时重定向
func TemporaryRedirect(w http.ResponseWriter, url string) {
	Redirect(w, url, http.StatusFound)
}
