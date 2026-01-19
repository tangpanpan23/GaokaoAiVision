package primary

import (
	"lighthouse-volunteer/app/model"
)

// AdmissionScore 录取分数线表模型
type AdmissionScore struct {
	model.BaseModel
	Year             int    `gorm:"column:year;type:int(11);not null;comment:年份" json:"year"`
	Province         string `gorm:"column:province;type:varchar(20);not null;comment:省份" json:"province"`
	CollegeCode      string `gorm:"column:college_code;type:varchar(20);not null;comment:学校代码" json:"college_code"`
	CollegeName      string `gorm:"column:college_name;type:varchar(100);not null;comment:学校名称" json:"college_name"`
	MajorCode        string `gorm:"column:major_code;type:varchar(20);comment:专业代码" json:"major_code"`
	MajorName        string `gorm:"column:major_name;type:varchar(100);comment:专业名称" json:"major_name"`
	Batch            string `gorm:"column:batch;type:varchar(50);not null;comment:录取批次" json:"batch"`
	ScoreType        int    `gorm:"column:score_type;type:tinyint(4);not null;comment:分数类型：1-文科，2-理科，3-综合改革" json:"score_type"`
	MinScore         int    `gorm:"column:min_score;type:int(11);comment:最低分" json:"min_score"`
	MinRank          int    `gorm:"column:min_rank;type:int(11);comment:最低位次" json:"min_rank"`
	AvgScore         int    `gorm:"column:avg_score;type:int(11);comment:平均分" json:"avg_score"`
	EnrollmentCount  int    `gorm:"column:enrollment_count;type:int(11);comment:录取人数" json:"enrollment_count"`
	DataSource       string `gorm:"column:data_source;type:varchar(200);not null;comment:数据来源URL" json:"data_source"`
	DataQuality      int    `gorm:"column:data_quality;type:tinyint(4);default:1;comment:数据质量评分：1-5" json:"data_quality"`
}

// TableName 指定表名
func (AdmissionScore) TableName() string {
	return "admission_score"
}

// CollegeInfo 学校信息表模型
type CollegeInfo struct {
	model.BaseModel
	CollegeCode      string  `gorm:"column:college_code;type:varchar(20);not null;uniqueIndex;comment:学校代码" json:"college_code"`
	CollegeName      string  `gorm:"column:college_name;type:varchar(100);not null;comment:学校名称" json:"college_name"`
	Province         string  `gorm:"column:province;type:varchar(20);not null;comment:所在省份" json:"province"`
	City             string  `gorm:"column:city;type:varchar(20);comment:所在城市" json:"city"`
	Level            string  `gorm:"column:level;type:varchar(20);comment:学校等级" json:"level"`
	Type             string  `gorm:"column:type;type:varchar(20);comment:学校类型" json:"type"`
	Nature           string  `gorm:"column:nature;type:varchar(20);comment:办学性质" json:"nature"`
	Website          string  `gorm:"column:website;type:varchar(200);comment:学校官网" json:"website"`
	Description      string  `gorm:"column:description;type:text;comment:学校简介" json:"description"`
	Ranking          int     `gorm:"column:ranking;type:int(11);comment:学校排名" json:"ranking"`
	EmploymentRate   float64 `gorm:"column:employment_rate;type:decimal(5,2);comment:就业率" json:"employment_rate"`
}

// TableName 指定表名
func (CollegeInfo) TableName() string {
	return "college_info"
}

// MajorInfo 专业信息表模型
type MajorInfo struct {
	model.BaseModel
	MajorCode         string `gorm:"column:major_code;type:varchar(20);not null;uniqueIndex;comment:专业代码" json:"major_code"`
	MajorName         string `gorm:"column:major_name;type:varchar(100);not null;comment:专业名称" json:"major_name"`
	Category          string `gorm:"column:category;type:varchar(50);comment:专业大类" json:"category"`
	Subcategory       string `gorm:"column:subcategory;type:varchar(50);comment:专业小类" json:"subcategory"`
	Degree            string `gorm:"column:degree;type:varchar(20);comment:学位类型" json:"degree"`
	Duration          int    `gorm:"column:duration;type:int(11);comment:学制(年)" json:"duration"`
	Description       string `gorm:"column:description;type:text;comment:专业简介" json:"description"`
	EmploymentDirection string `gorm:"column:employment_direction;type:text;comment:就业方向" json:"employment_direction"`
	SalaryRange       string `gorm:"column:salary_range;type:varchar(50);comment:薪资范围" json:"salary_range"`
	DemandLevel       string `gorm:"column:demand_level;type:varchar(20);comment:需求程度" json:"demand_level"`
}

// TableName 指定表名
func (MajorInfo) TableName() string {
	return "major_info"
}

