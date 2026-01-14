package ctxdata

import (
	"context"
	"strconv"

	"lighthouse-volunteer/constant"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// ContextKeyUserID 用户ID上下文键
	ContextKeyUserID = "user_id"
	// ContextKeyOpenID OpenID上下文键
	ContextKeyOpenID = "open_id"
	// ContextKeyRequestID 请求ID上下文键
	ContextKeyRequestID = "request_id"
)

// SetUserID 设置用户ID到上下文
func SetUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, ContextKeyUserID, userID)
}

// GetUserID 从上下文获取用户ID
func GetUserID(ctx context.Context) int64 {
	if ctx == nil {
		return 0
	}
	if userID, ok := ctx.Value(ContextKeyUserID).(int64); ok {
		return userID
	}
	return 0
}

// GetUserIDString 从上下文获取用户ID字符串
func GetUserIDString(ctx context.Context) string {
	userID := GetUserID(ctx)
	if userID > 0 {
		return strconv.FormatInt(userID, 10)
	}
	return ""
}

// SetOpenID 设置OpenID到上下文
func SetOpenID(ctx context.Context, openID string) context.Context {
	return context.WithValue(ctx, ContextKeyOpenID, openID)
}

// GetOpenID 从上下文获取OpenID
func GetOpenID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if openID, ok := ctx.Value(ContextKeyOpenID).(string); ok {
		return openID
	}
	return ""
}

// SetRequestID 设置请求ID到上下文
func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestID, requestID)
}

// GetRequestID 从上下文获取请求ID
func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if requestID, ok := ctx.Value(ContextKeyRequestID).(string); ok {
		return requestID
	}
	return ""
}

// GetAIAPIKey 获取AI API密钥
func GetAIAPIKey(provider string) string {
	// 从环境变量或配置文件获取API密钥
	// 这里应该从配置中心或环境变量获取
	switch provider {
	case "qwen":
		return getEnvOrConfig("AI_QWEN_API_KEY")
	case "ernie":
		return getEnvOrConfig("AI_ERNIE_API_KEY")
	case "gpt":
		return getEnvOrConfig("AI_GPT_API_KEY")
	default:
		return ""
	}
}

// GetWechatAppConfig 获取微信小程序配置
func GetWechatAppConfig() (appID, appSecret string) {
	appID = getEnvOrConfig("WECHAT_APP_ID")
	appSecret = getEnvOrConfig("WECHAT_APP_SECRET")
	return
}

// GetDatabaseConfig 获取数据库配置
func GetDatabaseConfig() (host string, port int, user, password, database string) {
	host = getEnvOrConfig("DB_HOST")
	portStr := getEnvOrConfig("DB_PORT")
	if portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}
	user = getEnvOrConfig("DB_USER")
	password = getEnvOrConfig("DB_PASSWORD")
	database = getEnvOrConfig("DB_DATABASE")
	return
}

// GetRedisConfig 获取Redis配置
func GetRedisConfig() (host, password string, db int) {
	host = getEnvOrConfig("REDIS_HOST")
	password = getEnvOrConfig("REDIS_PASSWORD")
	dbStr := getEnvOrConfig("REDIS_DB")
	if dbStr != "" {
		if d, err := strconv.Atoi(dbStr); err == nil {
			db = d
		}
	}
	return
}

// GetKafkaConfig 获取Kafka配置
func GetKafkaConfig() (brokers []string, topic string) {
	brokersStr := getEnvOrConfig("KAFKA_BROKERS")
	if brokersStr != "" {
		// 简单分割，实际应该更复杂的解析
		brokers = []string{brokersStr}
	}
	topic = getEnvOrConfig("KAFKA_TOPIC")
	return
}

// GetElasticsearchConfig 获取Elasticsearch配置
func GetElasticsearchConfig() (hosts []string, index string) {
	hostsStr := getEnvOrConfig("ES_HOSTS")
	if hostsStr != "" {
		hosts = []string{hostsStr}
	}
	index = getEnvOrConfig("ES_INDEX")
	return
}

// getEnvOrConfig 从环境变量获取配置（简化实现）
func getEnvOrConfig(key string) string {
	// 这里应该实现从环境变量或配置中心获取的逻辑
	// 暂时返回空值，实际项目中需要实现
	logx.Infof("Getting config for key: %s", key)
	return ""
}

// IsInternalRequest 判断是否为内部请求
func IsInternalRequest(ctx context.Context) bool {
	// 通过检查特定的header或token来判断
	// 这里简化为检查用户ID是否存在
	return GetUserID(ctx) > 0
}

// GetClientIP 获取客户端IP
func GetClientIP(ctx context.Context) string {
	// 从上下文或header中获取客户端IP
	// 这里需要根据实际的上下文实现
	return ""
}

// GetUserAgent 获取用户代理
func GetUserAgent(ctx context.Context) string {
	// 从上下文或header中获取用户代理
	return ""
}

// GetRequestPath 获取请求路径
func GetRequestPath(ctx context.Context) string {
	// 从上下文获取请求路径
	return ""
}

// LogRequest 记录请求日志
func LogRequest(ctx context.Context, method, path string, statusCode int, duration int64) {
	userID := GetUserIDString(ctx)
	requestID := GetRequestID(ctx)

	logx.WithContext(ctx).Infof("Request: method=%s, path=%s, user_id=%s, request_id=%s, status=%d, duration=%dms",
		method, path, userID, requestID, statusCode, duration)
}

// LogError 记录错误日志
func LogError(ctx context.Context, err error, msg string) {
	userID := GetUserIDString(ctx)
	requestID := GetRequestID(ctx)

	logx.WithContext(ctx).Errorf("Error: user_id=%s, request_id=%s, msg=%s, error=%v",
		userID, requestID, msg, err)
}

// CreateRequestContext 创建请求上下文
func CreateRequestContext(userID int64, openID, requestID string) context.Context {
	ctx := context.Background()
	ctx = SetUserID(ctx, userID)
	ctx = SetOpenID(ctx, openID)
	ctx = SetRequestID(ctx, requestID)
	return ctx
}

// ValidateContext 验证上下文
func ValidateContext(ctx context.Context) error {
	if ctx == nil {
		return constant.ErrContextNil
	}

	userID := GetUserID(ctx)
	if userID <= 0 {
		return constant.ErrUserNotFound
	}

	openID := GetOpenID(ctx)
	if openID == "" {
		return constant.ErrInvalidOpenID
	}

	return nil
}
