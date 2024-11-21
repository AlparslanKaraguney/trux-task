package models

type SmartModel struct {
	ID         int32          `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"not null" json:"name"`
	Identifier string         `gorm:"uniqueIndex" json:"identifier"`
	Type       string         `gorm:"not null" json:"type"`
	Category   string         `gorm:"not null" json:"category"`
	Features   []SmartFeature `gorm:"foreignKey:ModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"features"`
}

type SmartFeature struct {
	ID            int32      `gorm:"primaryKey" json:"id"`
	Name          string     `json:"name"`
	Identifier    string     `gorm:"uniqueIndex" json:"identifier"`
	Functionality string     `json:"functionality"`
	ModelID       int32      `gorm:"index" json:"model_id"`
	Model         SmartModel `gorm:"foreignKey:ModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"model"`
}
