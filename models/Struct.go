package models

import (
	"time"

	"gorm.io/gorm"
)

type Agency struct {
	gorm.Model
	Name     string
	Location []Location `gorm:"foreignKey:AgencyID"`
}

type Auditor struct {
	gorm.Model
	FirstName string
	LastName  string
}

type Category struct {
	gorm.Model
	Name string
	// Equipment []Equipment `gorm:"foreignKey:CategoryID"`
}

type Consider struct {
	gorm.Model
	ReqlacementID uint        `gorm:"not null"`
	Reqlacement   Reqlacement `gorm:"constraint:OnDelete:CASCADE"`
	EquipmentID   uint        `gorm:"not null"`
	Equipment     Equipment   `gorm:"constraint:OnDelete:CASCADE"`
	CanUse        int
	Missing       int
	All           int
}

type Equipment struct {
	gorm.Model
	Name       string
	CategoryID uint       `gorm:"not null"` // Ensure foreign keys are not null
	Category   Category   `gorm:"constraint:OnDelete:CASCADE"`
	UnitID     uint       `gorm:"not null"` // Ensure foreign keys are not null
	Unit       Unit       `gorm:"constraint:OnDelete:CASCADE"`
	Consider   []Consider `gorm:"foreignKey:EquipmentID"`
}

type File struct {
	gorm.Model
	NameQuotation string
	QuotationPath string

	NamePhoto string
	PhotoPath string

	NameRepait string
	RepairPath string

	NameDiagram string
	DiagramPath string

	NameDetails string
	DetailsPath string

	ReqlacementID uint        `gorm:"not null"`
	Reqlacement   Reqlacement `gorm:"constraint:OnDelete:CASCADE"`
}

type Location struct {
	gorm.Model
	Name       string
	AgencyID   uint     `gorm:"not null"`
	Agency     Agency   `gorm:"constraint:OnDelete:CASCADE"`
	ProvinceID uint     `gorm:"not null"`
	Province   Province `gorm:"constraint:OnDelete:CASCADE"`
}

type Money struct {
	gorm.Model
	Name        string
	Reqlacement []Reqlacement `gorm:"foreignKey:MoneyID"`
}

type Necessary struct {
	gorm.Model
	LocationID uint     `gorm:"not null"`
	Location   Location `gorm:"constraint:OnDelete:CASCADE"`
	Frequency  string   `gorm:"type:string;check:frequency IN ('ใช้บ่อย','ไม่บ่อย');default:'ไม่บ่อย'"`
	Province   string   `gorm:"type:string;check:province IN ('Y','N');default:'N'"`
	Other      string
}

type Price struct {
	gorm.Model
	EquipmentID uint      `gorm:"not null"`
	Equipment   Equipment `gorm:"constraint:OnDelete:CASCADE"`
	InMarket    int
	OutMarket   int
}

type Province struct {
	gorm.Model
	Name     string
	Location []Location `gorm:"foreignKey:ProvinceID"`
}

type Reqlacement struct {
	gorm.Model
	EquipmentID uint       `gorm:"not null"`
	Equipment   Equipment  `gorm:"constraint:OnDelete:CASCADE"`
	SourceID    uint       `gorm:"not null"`
	Source      Source     `gorm:"constraint:OnDelete:CASCADE"`
	MoneyID     uint       `gorm:"not null"`
	Money       Money      `gorm:"constraint:OnDelete:CASCADE"`
	Consider    []Consider `gorm:"foreignKey:ReqlacementID"`
	File        []File     `gorm:"foreignKey:ReqlacementID"`
	Day         time.Time  `gorm:"type:DATE;"`
}

type Source struct {
	gorm.Model
	Name        string
	Reqlacement []Reqlacement `gorm:"foreignKey:SourceID"`
}

type Unit struct {
	gorm.Model
	Name      string
	Equipment []Equipment `gorm:"foreignKey:UnitID"`
}

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	LevelUser    string
	AccountUser  string
	PasswordUser string
}
