package primary

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"

	"lighthouse-volunteer/app/model"
)

// SpiderTarget 爬虫目标表模型
type SpiderTarget struct {
	model.BaseModel
	Name           string          `gorm:"column:name;type:varchar(100);not null;comment:目标名称" json:"name"`
	URL            string          `gorm:"column:url;type:varchar(500);not null;comment:目标URL" json:"url"`
	DataType       string          `gorm:"column:data_type;type:varchar(50);not null;comment:数据类型" json:"data_type"`
	CrawlFrequency int             `gorm:"column:crawl_frequency;type:int(11);default:24;comment:爬取频率(小时)" json:"crawl_frequency"`
	ParseRules     json.RawMessage `gorm:"column:parse_rules;type:json;not null;comment:解析规则" json:"parse_rules"`
	LastCrawlTime  *time.Time      `gorm:"column:last_crawl_time;type:datetime;comment:最后爬取时间" json:"last_crawl_time"`
	Status         int             `gorm:"column:status;type:tinyint(4);default:1;comment:状态：1-启用，0-禁用" json:"status"`
}

// TableName 指定表名
func (SpiderTarget) TableName() string {
	return "spider_target"
}

// SpiderTask 爬虫任务表模型
type SpiderTask struct {
	model.BaseModel
	TargetID     int64      `gorm:"column:target_id;type:bigint(20);not null;comment:目标ID" json:"target_id"`
	TaskID       string     `gorm:"column:task_id;type:varchar(100);not null;comment:Asynq任务ID" json:"task_id"`
	Status       string     `gorm:"column:status;type:varchar(20);default:'pending';comment:状态" json:"status"`
	StartTime    *time.Time `gorm:"column:start_time;type:datetime;comment:开始时间" json:"start_time"`
	EndTime      *time.Time `gorm:"column:end_time;type:datetime;comment:结束时间" json:"end_time"`
	TotalItems   int        `gorm:"column:total_items;type:int(11);default:0;comment:总条目数" json:"total_items"`
	SuccessCount int        `gorm:"column:success_count;type:int(11);default:0;comment:成功条目数" json:"success_count"`
	ErrorMessage string     `gorm:"column:error_message;type:text;comment:错误信息" json:"error_message"`

	// 关联
	Target *SpiderTarget `gorm:"foreignKey:TargetID;references:ID" json:"target,omitempty"`
}

// TableName 指定表名
func (SpiderTask) TableName() string {
	return "spider_task"
}

// BeforeCreate GORM钩子：创建前设置默认值
func (t *SpiderTask) BeforeCreate(tx *gorm.DB) error {
	if t.Status == "" {
		t.Status = "pending"
	}
	return nil
}
