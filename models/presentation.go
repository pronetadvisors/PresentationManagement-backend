package models

import (
	"time"
)

type Presentation struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time ``
	UpdatedAt  time.Time ``
	Time       time.Time `json:"time"`
	Location   string    `gorm:"size:255; not null; unique" json:"location"`
	Speaker    string    `gorm:"size:255; not null; unique" json:"speaker"`
	Title      string    `gorm:"size:255; not null; unique" json:"title"`
	Powerpoint string    `gorm:"size:255;" json:"powerpoint"`
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

//func getPresentationByDate() ([]Presentation, error) {}
//func getPresentationByLoc() ([]Presentation, error) {}
//func getPresentationByDateLoc() ([]Presentation, error) {}

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
	if err := DB.Where("id = ?", p.ID).First(&presentation).Error; err != nil {
		return &Presentation{}, err
	}

	if err := DB.Model(&presentation).Update("powerpoint", p.Powerpoint).Error; err != nil {
		return &Presentation{}, err
	}

	return p, nil
}
