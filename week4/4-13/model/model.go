package model

import "gorm.io/gorm"

type Employer struct {
	gorm.Model
	Name      string
	CompanyID uint
	Company   Company
}

type Company struct {
	gorm.Model
	Name string
}
