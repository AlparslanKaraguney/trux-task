package models

import "time"

type SmartModel struct {
	ID         int32     `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	Identifier string    `gorm:"uniqueIndex" json:"identifier"`
	Type       string    `gorm:"not null" json:"type"`
	Category   string    `gorm:"not null" json:"category"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Features []SmartFeature `gorm:"foreignKey:SmartModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"features"`
}

type SmartFeature struct {
	ID            int32      `gorm:"primaryKey" json:"id"`
	Name          string     `json:"name"`
	Identifier    string     `gorm:"uniqueIndex" json:"identifier"`
	Functionality string     `json:"functionality"`
	SmartModelID  int32      `gorm:"not null" json:"model_id"`
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	SmartModel    SmartModel `gorm:"foreignKey:SmartModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"model"`
}
