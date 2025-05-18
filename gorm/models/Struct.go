package models

import (
	"time"

	"gorm.io/gorm"
)

type MyModels struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Agency struct {
	MyModels
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"Name"`
}

type Auditor struct {
	MyModels
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type Category struct {
	MyModels
	Name string
}

type Consider struct {
	MyModels
	ReqlacementID uint        `gorm:"not null"`
	Reqlacement   Reqlacement `gorm:"constraint:OnDelete:CASCADE"`
	EquipmentID   uint        `gorm:"not null"`
	Equipment     Equipment   `gorm:"constraint:OnDelete:CASCADE"`
	CanUse        int
	Missing       int
	All           int
}

type Equipment struct {
	MyModels
	Name       string
	CategoryID uint       `gorm:"not null"` // Ensure foreign keys are not null
	Category   Category   `gorm:"constraint:OnDelete:CASCADE"`
	UnitID     uint       `gorm:"not null"` // Ensure foreign keys are not null
	Unit       Unit       `gorm:"constraint:OnDelete:CASCADE"`
	Consider   []Consider `gorm:"foreignKey:EquipmentID"`
}

type File struct {
	MyModels
	NameQuotation string
	QuotationPath string

	NamePhoto string
	PhotoPath string

	NameRepair string
	RepairPath string

	NameDiagram string
	DiagramPath string

	NameDetails string
	DetailsPath string

	ReqlacementID uint        `gorm:"not null"`
	Reqlacement   Reqlacement `gorm:"constraint:OnDelete:CASCADE"`
}

type Location struct {
	MyModels
	Name       string
	AgencyID   uint     `gorm:"not null"`
	Agency     Agency   `gorm:"constraint:OnDelete:CASCADE"`
	ProvinceID uint     `gorm:"not null"`
	Province   Province `gorm:"constraint:OnDelete:CASCADE"`
}

type Money struct {
	MyModels
	Name string
}

type Necessary struct {
	MyModels
	LocationID uint     `gorm:"not null"`
	Location   Location `gorm:"constraint:OnDelete:CASCADE"`
	Frequency  string   //`gorm:"type:enum('ใช้บ่อย','ไม่บ่อย');default:'ไม่บ่อย'"`
	Province   string   //`gorm:"type:enum('Y','N');default:'N'"`
	Other      string
}

type Price struct {
	MyModels
	EquipmentID uint      `gorm:"not null"`
	Equipment   Equipment `gorm:"constraint:OnDelete:CASCADE"`
	InMarket    int
	OutMarket   int
}

type Province struct {
	MyModels
	Name string
}

type Reqlacement struct {
	MyModels
	EquipmentID uint      `gorm:"not null"`
	Equipment   Equipment `gorm:"constraint:OnDelete:CASCADE"`
	SourceID    uint      `gorm:"not null"`
	Source      Source    `gorm:"constraint:OnDelete:CASCADE"`
	MoneyID     uint      `gorm:"not null"`
	Money       Money     `gorm:"constraint:OnDelete:CASCADE"`
	Day         time.Time `gorm:"type:DATE;"`
}

type Requests struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	AgencyID    uint
	Agency      Agency `gorm:"foreignKey:AgencyID"`
	LocationID  uint
	Location    Location `gorm:"foreignKey:LocationID"`
	EquipmentID uint
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	Replace     string
	Request     string
	PriceID     uint
	Price       Price `gorm:"foreignKey:PriceID"`
	AuditorID   uint
	Auditor     Auditor `gorm:"foreignKey:AuditorID"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
	Remark      string
	Date        time.Time
	Necessary
}

type Source struct {
	MyModels
	Name string
}

type Unit struct {
	MyModels
	Name string
}

type User struct {
	MyModels
	FirstName    string
	LastName     string
	LevelUser    string
	Email        string
	PasswordUser string
}
