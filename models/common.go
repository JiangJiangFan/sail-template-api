package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/plugin/soft_delete"
)

// Model 自增主键ID,创建、更新时间,软删除
type Model struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt LocalTime `json:"created_at,omitempty" gorm:"not null; default:current_timestamp"`
	UpdatedAt LocalTime `json:"updated_at,omitempty"`
	// DeletedAt time.Time `json:"deleted_at,omitempty" sql:"index"`
	// 软删除
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

type LocalTime struct {time.Time}

// 序列化时间格式
func ( t *LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%v\"", t.Format("2006-01-02 15:04:05"))), nil
}

// 反序列化
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	timeStr := strings.Trim(string(data),"\"")
	lt, err := time.ParseInLocation("2006-01-02 15:04:05",timeStr, time.Local)

	*t = LocalTime{Time: lt}

	if err != nil {
		return errors.New("时间格式有误，转换失败")
	}
	return nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zero time.Time
	// 判定给定时间是否和默认零时时间戳相同
	if t.Time.UnixMicro() == zero.UnixMicro() {
		return nil, nil
	}
	return t.Time,nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("不能转换 %V 为时间戳", v)
}
