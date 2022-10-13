package models

import (
	"time"
)

type Settings struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time ``
	UpdatedAt time.Time ``
	Company   string    `json:"company"`
	EventName string    `json:"event"`
}

func (s *Settings) CreateSettings() (*Settings, error) {
	if err := DB.Create(&s).Error; err != nil {
		return &Settings{}, err
	}
	return s, nil
}

func GetSettings() ([]Settings, error) {
	var settings []Settings
	if err := DB.Find(&settings).Error; err != nil {
		return settings, err
	}
	return settings, nil
}

func (s *Settings) UpdateSettings() (*Settings, error) {
	var setting Settings
	if err := DB.Where("id = ?", s.ID).First(&setting).Error; err != nil {
		return &Settings{}, err
	}

	if err := DB.Model(&setting).Updates(&s).Error; err != nil {
		return &Settings{}, err
	}

	return s, nil
}

func (s *Settings) DeleteSettings() (*Settings, error) {
	var presentation Presentation
	if err := DB.Where("id = ?", s.ID).First(&presentation).Error; err != nil {
		return &Settings{}, err
	}

	if err := DB.Delete(&presentation).Error; err != nil {
		return &Settings{}, err
	}

	return s, nil
}