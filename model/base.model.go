package model

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Timestamp ... ( yyyy-MM-dd HH:mm:ss )
type Timestamp time.Time

// Date ... ( yyyy-MM-dd )
type Date time.Time

// BaseModel ...
type BaseModel struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	CreatedAt *Timestamp `gorm:"column:created_at; type:timestamp without time zone" json:"createdAt"`
	CreatedBy string     `gorm:"column:created_by" json:"createdBy"`
	UpdatedAt *Timestamp `gorm:"column:updated_at; type:timestamp without time zone;" json:"updatedAt"`
	UpdatedBy *string    `gorm:"column:updated_by" json:"updatedBy"`
}

// BeforeCreate ...
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	base.ID = uuid.String()
	return nil
}

// UnmarshalJSON ...
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var timeStr = string(data)
	if timeStr == "null" || timeStr == `""` {
		return nil
	}
	if len(timeStr) > 0 && timeStr[0] == '"' {
		timeStr = timeStr[1:]
	}
	if len(timeStr) > 0 && timeStr[len(timeStr)-1] == '"' {
		timeStr = timeStr[:len(timeStr)-1]
	}

	layout := "2006-01-02 15:04:05"

	ts, err := time.Parse(layout, timeStr)
	*t = Timestamp(ts)

	log.Println("ts:", ts)
	log.Println("timeStr", timeStr)
	log.Println("t", t)
	return err
}

// UnmarshalJSON ...
func (t *Date) UnmarshalJSON(data []byte) error {
	var timeStr = string(data)
	if timeStr == "null" || timeStr == `""` {
		return nil
	}
	if len(timeStr) > 0 && timeStr[0] == '"' {
		timeStr = timeStr[1:]
	}
	if len(timeStr) > 0 && timeStr[len(timeStr)-1] == '"' {
		timeStr = timeStr[:len(timeStr)-1]
	}

	layout := "2006-01-02"

	ts, err := time.Parse(layout, timeStr)
	*t = Date(ts)

	log.Println("ts:", ts)
	log.Println("timeStr", timeStr)
	log.Println("t", t)
	return err
}

// MarshalJSON ...
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// MarshalJSON ...
func (t *Date) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02"))
	return []byte(stamp), nil
}
