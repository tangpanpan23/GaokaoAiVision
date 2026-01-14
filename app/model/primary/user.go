package primary

import (
	"encoding/json"
	"time"

	"lighthouse-volunteer/app/model"
)

// User 用户表模型
type User struct {
	model.BaseModel
	OpenID        string     `gorm:"column:open_id;type:varchar(100);not null;uniqueIndex;comment:微信OpenID" json:"open_id"`
	UnionID       string     `gorm:"column:union_id;type:varchar(100);comment:微信UnionID" json:"union_id"`
	Nickname      string     `gorm:"column:nickname;type:varchar(50);comment:昵称" json:"nickname"`
	Avatar        string     `gorm:"column:avatar;type:varchar(500);comment:头像URL" json:"avatar"`
	Gender        int        `gorm:"column:gender;type:tinyint(4);comment:性别：0-未知，1-男，2-女" json:"gender"`
	Province      string     `gorm:"column:province;type:varchar(20);comment:省份" json:"province"`
	City          string     `gorm:"column:city;type:varchar(20);comment:城市" json:"city"`
	Status        int        `gorm:"column:status;type:tinyint(4);default:1;comment:状态：0-未激活，1-活跃，2-封禁" json:"status"`
	LastLoginTime *time.Time `gorm:"column:last_login_time;type:datetime;comment:最后登录时间" json:"last_login_time"`

	// 关联表
	Profile *UserProfile `gorm:"foreignKey:UserID;references:ID" json:"profile,omitempty"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

// UserProfile 用户档案表模型
type UserProfile struct {
	model.BaseModel
	UserID          int64   `gorm:"column:user_id;type:bigint(20);not null;uniqueIndex;comment:用户ID" json:"user_id"`
	GraduationYear  int     `gorm:"column:graduation_year;type:int(11);comment:毕业年份" json:"graduation_year"`
	Province        string  `gorm:"column:province;type:varchar(20);comment:高考省份" json:"province"`
	ScoreType       int     `gorm:"column:score_type;type:tinyint(4);comment:分数类型" json:"score_type"`
	Subjects        string  `gorm:"column:subjects;type:varchar(100);comment:选考科目" json:"subjects"`
	TotalScore      int     `gorm:"column:total_score;type:int(11);comment:总分" json:"total_score"`
	Rank            int     `gorm:"column:rank;type:int(11);comment:位次" json:"rank"`
	TargetCollege   string  `gorm:"column:target_college;type:varchar(100);comment:目标学校" json:"target_college"`
	TargetMajor     string  `gorm:"column:target_major;type:varchar(100);comment:目标专业" json:"target_major"`
	InterestTags    string  `gorm:"column:interest_tags;type:varchar(500);comment:兴趣标签" json:"interest_tags"`
	PersonalityType string  `gorm:"column:personality_type;type:varchar(20);comment:性格类型" json:"personality_type"`
}

// TableName 指定表名
func (UserProfile) TableName() string {
	return "user_profile"
}

// SeniorShare 学长分享表模型
type SeniorShare struct {
	model.BaseModel
	UserID       int64           `gorm:"column:user_id;type:bigint(20);not null;comment:分享者用户ID" json:"user_id"`
	CollegeCode  string          `gorm:"column:college_code;type:varchar(20);not null;comment:学校代码" json:"college_code"`
	CollegeName  string          `gorm:"column:college_name;type:varchar(100);not null;comment:学校名称" json:"college_name"`
	MajorCode    string          `gorm:"column:major_code;type:varchar(20);comment:专业代码" json:"major_code"`
	MajorName    string          `gorm:"column:major_name;type:varchar(100);comment:专业名称" json:"major_name"`
	ShareType    string          `gorm:"column:share_type;type:varchar(20);not null;comment:分享类型" json:"share_type"`
	Title        string          `gorm:"column:title;type:varchar(200);not null;comment:标题" json:"title"`
	Content      string          `gorm:"column:content;type:text;not null;comment:内容" json:"content"`
	Tags         string          `gorm:"column:tags;type:varchar(500);comment:标签" json:"tags"`
	IsAnonymous  int             `gorm:"column:is_anonymous;type:tinyint(4);default:0;comment:是否匿名" json:"is_anonymous"`
	ViewCount    int             `gorm:"column:view_count;type:int(11);default:0;comment:浏览数" json:"view_count"`
	LikeCount    int             `gorm:"column:like_count;type:int(11);default:0;comment:点赞数" json:"like_count"`
	CommentCount int             `gorm:"column:comment_count;type:int(11);default:0;comment:评论数" json:"comment_count"`
	Status       int             `gorm:"column:status;type:tinyint(4);default:1;comment:状态" json:"status"`
	PublishedAt  *time.Time      `gorm:"column:published_at;type:datetime;comment:发布时间" json:"published_at"`

	// 关联表
	User         *User           `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
}

// TableName 指定表名
func (SeniorShare) TableName() string {
	return "senior_share"
}

// SeniorShareLike 学长分享点赞表模型
type SeniorShareLike struct {
	model.BaseModel
	UserID  int64 `gorm:"column:user_id;type:bigint(20);not null;comment:用户ID" json:"user_id"`
	ShareID int64 `gorm:"column:share_id;type:bigint(20);not null;comment:分享ID" json:"share_id"`
}

// TableName 指定表名
func (SeniorShareLike) TableName() string {
	return "senior_share_like"
}

// CareerAssessment 职业测评表模型
type CareerAssessment struct {
	model.BaseModel
	UserID           int64           `gorm:"column:user_id;type:bigint(20);not null;comment:用户ID" json:"user_id"`
	AssessmentType   string          `gorm:"column:assessment_type;type:varchar(20);not null;comment:测评类型" json:"assessment_type"`
	Result           json.RawMessage `gorm:"column:result;type:json;not null;comment:测评结果" json:"result"`
	ScoreDetails     json.RawMessage `gorm:"column:score_details;type:json;comment:分数详情" json:"score_details"`
	Recommendations  string          `gorm:"column:recommendations;type:text;comment:职业建议" json:"recommendations"`
}

// TableName 指定表名
func (CareerAssessment) TableName() string {
	return "career_assessment"
}

// SubjectPlan 选科规划表模型
type SubjectPlan struct {
	model.BaseModel
	Province            string  `gorm:"column:province;type:varchar(20);not null;comment:省份" json:"province"`
	ScoreType           int     `gorm:"column:score_type;type:tinyint(4);not null;comment:分数类型" json:"score_type"`
	SubjectCombination  string  `gorm:"column:subject_combination;type:varchar(100);not null;comment:科目组合" json:"subject_combination"`
	RecommendedMajors   string  `gorm:"column:recommended_majors;type:text;not null;comment:推荐专业" json:"recommended_majors"`
	SuccessRate         float64 `gorm:"column:success_rate;type:decimal(5,2);comment:成功率" json:"success_rate"`
	AvgScoreRequired    int     `gorm:"column:avg_score_required;type:int(11);comment:平均所需分数" json:"avg_score_required"`
}

// TableName 指定表名
func (SubjectPlan) TableName() string {
	return "subject_plan"
}
