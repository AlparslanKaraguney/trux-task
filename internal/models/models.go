package models

type SmartModel struct {
	ID         int32          `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"not null" json:"name"`
	Identifier string         `gorm:"uniqueIndex" json:"identifier"`
	Type       string         `gorm:"not null" json:"type"`
	Category   string         `gorm:"not null" json:"category"`
	Features   []SmartFeature `gorm:"foreignKey:SmartModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"features"`
}

type SmartFeature struct {
	ID            int32      `gorm:"primaryKey" json:"id"`
	Name          string     `json:"name"`
	Identifier    string     `gorm:"uniqueIndex" json:"identifier"`
	Functionality string     `json:"functionality"`
	SmartModelID  int32      `gorm:"not null" json:"model_id"`
	SmartModel    SmartModel `gorm:"foreignKey:SmartModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"model"`
}
