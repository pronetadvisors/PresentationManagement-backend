package models

import (
	"time"
)

type Presentation struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time ``
	UpdatedAt   time.Time ``
	SessionID   string    `json:"session_id"`
	Time        time.Time `json:"time"`
	EndTime     time.Time `json:"endtime"`
	Location    string    `gorm:"size:255; not null;" json:"location"`
	Speaker     string    `gorm:"size:255; not null;" json:"speaker"`
	Title       string    `gorm:"size:255; not null;" json:"title"`
	Description string    `gorm:"size:1000; not null;" json:"description"`
	Powerpoint  string    `gorm:"size:255;" json:"powerpoint"`
}

func (p *Presentation) CreatePresentation() (*Presentation, error) {
	if err := DB.Create(&p).Error; err != nil {
		return &Presentation{}, err
	}
	return p, nil
}

func GetPresentation() ([]Presentation, error) {
	var presentations []Presentation
	if err := DB.Find(&presentations).Error; err != nil {
		return presentations, err
	}
	return presentations, nil
}

func (p *Presentation) UpdatePresentation() (*Presentation, error) {
	var presentation Presentation
	if err := DB.Where("id = ?", p.ID).First(&presentation).Error; err != nil {
		return &Presentation{}, err
	}

	if err := DB.Model(&presentation).Updates(&p).Error; err != nil {
		return &Presentation{}, err
	}

	return p, nil
}

func (p *Presentation) DeletePresentation() (*Presentation, error) {
	var presentation Presentation
	if err := DB.Where("id = ?", p.ID).First(&presentation).Error; err != nil {
		return &Presentation{}, err
	}

	if err := DB.Delete(&presentation).Error; err != nil {
		return &Presentation{}, err
	}

	return p, nil
}

func (p *Presentation) UpdatePowerpoint() (*Presentation, error) {
	var presentation Presentation
	if err := DB.Where("session_id = ?", p.SessionID).First(&presentation).Error; err != nil {
		return &Presentation{}, err
	}

	if err := DB.Model(&presentation).Update("powerpoint", p.Powerpoint).Error; err != nil {
		return &Presentation{}, err
	}

	return p, nil
}

func (p *Presentation) DeletePowerpoint() (*Presentation, error) {
	var presentation Presentation
	if err := DB.Where("id = ?", p.ID).First(&presentation).Error; err != nil {
		return &Presentation{}, err
	}

	if err := DB.Model(&presentation).Update("powerpoint", "").Error; err != nil {
		return &Presentation{}, err
	}

	return p, nil
}

func GetPowerpointbyRoom(room string) ([]Presentation, error) {
	var presentations []Presentation
	if err := DB.Where("location = ?", room).Find(&presentations).Error; err != nil {
		return presentations, err
	}

	return presentations, nil
}
