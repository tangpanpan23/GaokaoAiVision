package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// BaseModel 基础模型，包含公共字段
type BaseModel struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;comment:主键ID" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删除时间" json:"deleted_at,omitempty"`
}

// JSONTime 自定义JSON时间类型
type JSONTime time.Time

// MarshalJSON 实现json.Marshaler接口
func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON 实现json.Unmarshaler接口
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = JSONTime(now)
	return err
}

// String 实现Stringer接口
func (t JSONTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

// Value 实现driver.Valuer接口
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if time.Time(t).IsZero() || time.Time(t).Equal(zeroTime) {
		return nil, nil
	}
	return time.Time(t), nil
}

// Scan 实现sql.Scanner接口
func (t *JSONTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = JSONTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Time 转换为time.Time
func (t JSONTime) Time() time.Time {
	return time.Time(t)
}

// PageRequest 分页请求
type PageRequest struct {
	Page     int `json:"page" form:"page" validate:"min=1"`         // 页码，从1开始
	PageSize int `json:"page_size" form:"page_size" validate:"min=1,max=100"` // 每页大小
}

// PageResponse 分页响应
type PageResponse struct {
	Total    int64 `json:"total"`     // 总记录数
	Page     int   `json:"page"`      // 当前页码
	PageSize int   `json:"page_size"` // 每页大小
	Pages    int   `json:"pages"`     // 总页数
}

// CalculatePages 计算总页数
func (p *PageResponse) CalculatePages() {
	if p.PageSize > 0 {
		p.Pages = int((p.Total + int64(p.PageSize) - 1) / int64(p.PageSize))
	}
}

// Offset 获取偏移量
func (p *PageRequest) Offset() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
	return (p.Page - 1) * p.PageSize
}

// Limit 获取限制数
func (p *PageRequest) Limit() int {
	if p.PageSize <= 0 {
		p.PageSize = 20
	}
	return p.PageSize
}
