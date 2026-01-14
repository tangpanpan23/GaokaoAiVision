package types

import "lighthouse-volunteer/app/model"

// 通用响应
type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// 分页请求
type PageRequest struct {
	Page     int `json:"page" form:"page" validate:"min=1"`         // 页码，从1开始
	PageSize int `json:"page_size" form:"page_size" validate:"min=1,max=100"` // 每页大小
}

// 分页响应
type PageResponse struct {
	Total    int64 `json:"total"`     // 总记录数
	Page     int   `json:"page"`      // 当前页码
	PageSize int   `json:"page_size"` // 每页大小
	Pages    int   `json:"pages"`     // 总页数
	List     interface{} `json:"list"`     // 数据列表
}

// 录取分数查询请求
type ScoreQueryReq struct {
	Province  string `json:"province" form:"province" validate:"required"`   // 省份
	ScoreType int    `json:"score_type" form:"score_type" validate:"required,min=1,max=3"` // 分数类型：1-文科，2-理科，3-综合改革
	Score     int    `json:"score" form:"score" validate:"min=0,max=750"`    // 分数
	Rank      int    `json:"rank" form:"rank" validate:"min=0"`              // 位次
	Subjects  string `json:"subjects" form:"subjects"`                       // 选考科目
}

// 录取分数查询响应
type ScoreQueryResp struct {
	CollegeCode string `json:"college_code"` // 学校代码
	CollegeName string `json:"college_name"` // 学校名称
	MajorCode   string `json:"major_code"`   // 专业代码
	MajorName   string `json:"major_name"`   // 专业名称
	Batch       string `json:"batch"`        // 录取批次
	MinScore    int    `json:"min_score"`    // 最低分
	MinRank     int    `json:"min_rank"`     // 最低位次
	AvgScore    int    `json:"avg_score"`    // 平均分
	Year        int    `json:"year"`         // 年份
}

// 志愿推荐请求
type VolunteerSuggestionReq struct {
	Province     string   `json:"province" validate:"required"`               // 省份
	ScoreType    int      `json:"score_type" validate:"required,min=1,max=3"` // 分数类型
	Score        int      `json:"score" validate:"min=0,max=750"`             // 分数
	Rank         int      `json:"rank" validate:"min=0"`                       // 位次
	Subjects     string   `json:"subjects"`                                    // 选考科目
	InterestTags []string `json:"interest_tags"`                               // 兴趣标签
}

// 志愿推荐响应
type VolunteerSuggestionResp struct {
	Category string               `json:"category"` // 分类：冲/稳/保
	Colleges []CollegeSuggestion  `json:"colleges"` // 推荐学校列表
	Reason   string               `json:"reason"`   // 推荐理由
}

// 学校推荐详情
type CollegeSuggestion struct {
	CollegeCode string `json:"college_code"` // 学校代码
	CollegeName string `json:"college_name"` // 学校名称
	MajorCode   string `json:"major_code"`   // 专业代码
	MajorName   string `json:"major_name"`   // 专业名称
	Batch       string `json:"batch"`        // 录取批次
	MinScore    int    `json:"min_score"`    // 最低分
	MinRank     int    `json:"min_rank"`     // 最低位次
	Year        int    `json:"year"`         // 年份
}

// AI咨询请求
type AIAdviceReq struct {
	Query     string `json:"query" validate:"required,max=1000"` // 咨询内容
	UserID    int64  `json:"user_id"`                             // 用户ID
	SessionID string `json:"session_id"`                          // 会话ID
}

// AI咨询响应
type AIAdviceResp struct {
	Answer    string `json:"answer"`     // 回答内容
	SessionID string `json:"session_id"` // 会话ID
	Sources   []string `json:"sources"`  // 信息来源
}

// 学校信息查询请求
type CollegeQueryReq struct {
	PageRequest
	Province string `json:"province" form:"province"`     // 省份
	Level    string `json:"level" form:"level"`           // 等级：985,211,双一流
	Type     string `json:"type" form:"type"`             // 类型：综合,理工,师范
	Name     string `json:"name" form:"name"`             // 学校名称关键词
}

// 学校信息查询响应
type CollegeQueryResp struct {
	PageResponse
	List []CollegeInfo `json:"list"`
}

// 学校信息详情
type CollegeInfo struct {
	CollegeCode     string  `json:"college_code"`      // 学校代码
	CollegeName     string  `json:"college_name"`      // 学校名称
	Province        string  `json:"province"`          // 省份
	City            string  `json:"city"`              // 城市
	Level           string  `json:"level"`             // 等级
	Type            string  `json:"type"`              // 类型
	Nature          string  `json:"nature"`            // 办学性质
	Website         string  `json:"website"`           // 官网
	Description     string  `json:"description"`       // 简介
	Ranking         int     `json:"ranking"`           // 排名
	EmploymentRate  float64 `json:"employment_rate"`   // 就业率
}

// 专业信息查询请求
type MajorQueryReq struct {
	PageRequest
	Category    string `json:"category" form:"category"`       // 专业大类
	DemandLevel string `json:"demand_level" form:"demand_level"` // 需求程度
	Name        string `json:"name" form:"name"`                // 专业名称关键词
}

// 专业信息查询响应
type MajorQueryResp struct {
	PageResponse
	List []MajorInfo `json:"list"`
}

// 专业信息详情
type MajorInfo struct {
	MajorCode           string `json:"major_code"`            // 专业代码
	MajorName           string `json:"major_name"`            // 专业名称
	Category            string `json:"category"`              // 大类
	Subcategory         string `json:"subcategory"`           // 小类
	Degree              string `json:"degree"`                // 学位
	Duration            int    `json:"duration"`              // 学制
	Description         string `json:"description"`           // 简介
	EmploymentDirection string `json:"employment_direction"`  // 就业方向
	SalaryRange         string `json:"salary_range"`          // 薪资范围
	DemandLevel         string `json:"demand_level"`          // 需求程度
}

// 用户登录请求
type UserLoginReq struct {
	Code string `json:"code" validate:"required"` // 微信授权码
}

// 用户登录响应
type UserLoginResp struct {
	UserID      int64  `json:"user_id"`      // 用户ID
	OpenID      string `json:"open_id"`      // 微信OpenID
	Token       string `json:"token"`        // JWT Token
	NeedProfile bool   `json:"need_profile"` // 是否需要完善档案
}

// 用户档案更新请求
type UserProfileUpdateReq struct {
	UserID          int64   `json:"user_id" validate:"required"`              // 用户ID
	GraduationYear  int     `json:"graduation_year" validate:"min=2020,max=2030"` // 毕业年份
	Province        string  `json:"province"`                                 // 高考省份
	ScoreType       int     `json:"score_type" validate:"min=1,max=3"`       // 分数类型
	Subjects        string  `json:"subjects"`                                // 选考科目
	TotalScore      int     `json:"total_score" validate:"min=0,max=750"`    // 总分
	Rank            int     `json:"rank" validate:"min=0"`                   // 位次
	TargetCollege   string  `json:"target_college"`                          // 目标学校
	TargetMajor     string  `json:"target_major"`                            // 目标专业
	InterestTags    []string `json:"interest_tags"`                          // 兴趣标签
	PersonalityType string  `json:"personality_type"`                        // 性格类型
}

// 学长分享列表请求
type SeniorShareListReq struct {
	PageRequest
	ShareType   string `json:"share_type" form:"share_type"`     // 分享类型
	CollegeCode string `json:"college_code" form:"college_code"` // 学校代码
	MajorCode   string `json:"major_code" form:"major_code"`     // 专业代码
	Keyword     string `json:"keyword" form:"keyword"`           // 关键词搜索
	SortBy      string `json:"sort_by" form:"sort_by"`           // 排序方式：time,hot
}

// 学长分享列表响应
type SeniorShareListResp struct {
	PageResponse
	List []SeniorShareInfo `json:"list"`
}

// 学长分享详情
type SeniorShareInfo struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	CollegeCode  string    `json:"college_code"`
	CollegeName  string    `json:"college_name"`
	MajorCode    string    `json:"major_code"`
	MajorName    string    `json:"major_name"`
	ShareType    string    `json:"share_type"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Tags         []string  `json:"tags"`
	IsAnonymous  bool      `json:"is_anonymous"`
	ViewCount    int       `json:"view_count"`
	LikeCount    int       `json:"like_count"`
	CommentCount int       `json:"comment_count"`
	PublishedAt  *model.JSONTime `json:"published_at"`
	User         *ShareUser `json:"user,omitempty"`
	IsLiked      bool      `json:"is_liked"` // 当前用户是否点赞
}

// 分享用户简要信息
type ShareUser struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// 学长分享创建请求
type SeniorShareCreateReq struct {
	UserID      int64    `json:"user_id" validate:"required"`      // 用户ID
	CollegeCode string   `json:"college_code" validate:"required"` // 学校代码
	CollegeName string   `json:"college_name" validate:"required"` // 学校名称
	MajorCode   string   `json:"major_code"`                       // 专业代码
	MajorName   string   `json:"major_name"`                       // 专业名称
	ShareType   string   `json:"share_type" validate:"required"`   // 分享类型
	Title       string   `json:"title" validate:"required,max=200"` // 标题
	Content     string   `json:"content" validate:"required"`      // 内容
	Tags        []string `json:"tags"`                             // 标签
	IsAnonymous int      `json:"is_anonymous"`                     // 是否匿名
}

// 职业测评请求
type CareerAssessmentReq struct {
	UserID         int64             `json:"user_id" validate:"required"` // 用户ID
	AssessmentType string            `json:"assessment_type" validate:"required"` // 测评类型
	Answers        map[string]interface{} `json:"answers" validate:"required"`      // 答案
}

// 职业测评响应
type CareerAssessmentResp struct {
	AssessmentType  string                 `json:"assessment_type"`   // 测评类型
	Result          map[string]interface{} `json:"result"`            // 测评结果
	ScoreDetails    map[string]interface{} `json:"score_details"`     // 分数详情
	Recommendations string                 `json:"recommendations"`   // 职业建议
	CreatedAt       *model.JSONTime        `json:"created_at"`        // 测评时间
}

// 选科规划查询请求
type SubjectPlanQueryReq struct {
	Province string `json:"province" form:"province" validate:"required"` // 省份
	ScoreType int   `json:"score_type" form:"score_type" validate:"required,min=1,max=3"` // 分数类型
	Subjects  string `json:"subjects" form:"subjects"`                     // 科目组合
}

// 选科规划查询响应
type SubjectPlanQueryResp struct {
	SubjectCombination string   `json:"subject_combination"` // 科目组合
	RecommendedMajors  []string `json:"recommended_majors"`  // 推荐专业
	SuccessRate        float64  `json:"success_rate"`        // 成功率
	AvgScoreRequired   int      `json:"avg_score_required"`  // 平均所需分数
}
