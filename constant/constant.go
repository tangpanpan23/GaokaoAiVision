package constant

// 服务名称常量
const (
	ServiceNameAPI     = "api"
	ServiceNameSpider  = "spider-rpc"
	ServiceNameScore   = "score-rpc"
	ServiceNameAI      = "ai-rpc"
	ServiceNameUser    = "user-rpc"
	ServiceNameModel   = "model-rpc"
	ServiceNameCommand = "command"
)

// 数据库表名常量
const (
	TableAdmissionScore    = "admission_score"
	TableCollegeInfo       = "college_info"
	TableMajorInfo         = "major_info"
	TableUser              = "user"
	TableUserProfile       = "user_profile"
	TableSeniorShare       = "senior_share"
	TableSeniorShareLike   = "senior_share_like"
	TableCareerAssessment  = "career_assessment"
	TableSubjectPlan       = "subject_plan"
	TableSpiderTask        = "spider_task"
	TableSpiderTarget      = "spider_target"
)

// Redis键前缀常量
const (
	RedisKeyScorePrefix       = "score:"
	RedisKeyCollegePrefix     = "college:"
	RedisKeyMajorPrefix       = "major:"
	RedisKeyUserPrefix        = "user:"
	RedisKeyAISuggestion      = "ai:suggestion:"
	RedisKeySpiderTask        = "spider:task:"
	RedisKeySpiderResult      = "spider:result:"
)

// 分数类型常量
const (
	ScoreTypeLiberalArts = 1 // 文科
	ScoreTypeScience     = 2 // 理科
	ScoreTypeComprehensive = 3 // 综合改革
)

// 录取批次常量
const (
	BatchTypeFirst    = "一本"
	BatchTypeSecond   = "二本"
	BatchTypeThird    = "三本"
	BatchTypeSpecial  = "专科"
	BatchTypeArt      = "艺术类"
	BatchTypeSports   = "体育类"
)

// 省份代码常量
const (
	ProvinceBeijing   = "北京"
	ProvinceShanghai  = "上海"
	ProvinceTianjin   = "天津"
	ProvinceChongqing = "重庆"
	// ... 其他省份
)

// 爬虫状态常量
const (
	SpiderStatusPending    = "pending"
	SpiderStatusRunning    = "running"
	SpiderStatusCompleted  = "completed"
	SpiderStatusFailed     = "failed"
	SpiderStatusCancelled  = "cancelled"
)

// AI分析类型常量
const (
	AIAnalysisTypeSuggestion = "suggestion" // 志愿推荐
	AIAnalysisTypeAdvice     = "advice"     // 政策咨询
	AIAnalysisTypeAssessment = "assessment" // 职业评估
)

// 用户状态常量
const (
	UserStatusActive   = 1 // 活跃
	UserStatusInactive = 0 // 未激活
	UserStatusBanned   = 2 // 封禁
)

// 分享类型常量
const (
	ShareTypeExperience = "experience" // 学习经历
	ShareTypeAdvice     = "advice"     // 填报建议
	ShareTypeInterview  = "interview"  // 面试经验
)

// 缓存过期时间常量 (秒)
const (
	CacheExpireScore        = 3600 * 24 * 7   // 7天
	CacheExpireCollege      = 3600 * 24 * 30  // 30天
	CacheExpireAISuggestion = 3600 * 24 * 1   // 1天
	CacheExpireUser         = 3600 * 24 * 1   // 1天
)

// API响应码常量
const (
	CodeSuccess        = 200
	CodeBadRequest     = 400
	CodeUnauthorized   = 401
	CodeForbidden      = 403
	CodeNotFound       = 404
	CodeInternalError  = 500
	CodeServiceUnavailable = 503
)

// 分页常量
const (
	DefaultPageSize = 20
	MaxPageSize     = 100
)

// 文件上传常量
const (
	MaxFileSize        = 10 * 1024 * 1024 // 10MB
	AllowedImageTypes  = "jpg,jpeg,png,gif"
	AllowedDocTypes    = "pdf,doc,docx"
)

// 消息队列主题常量
const (
	KafkaTopicSpiderResult  = "spider.result"
	KafkaTopicUserAction    = "user.action"
	KafkaTopicAIAnalysis    = "ai.analysis"
)

// Elasticsearch索引常量
const (
	ESIndexSeniorShare = "senior_share"
	ESIndexCollege     = "college"
	ESIndexMajor       = "major"
)

