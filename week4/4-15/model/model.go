package model

import "gorm.io/gorm"

type Employer struct {
	gorm.Model
	Name        string
	CompanyID   uint
	Company     Company
	CreditCards []CreditCard
}

type Company struct {
	gorm.Model
	Name string
}

type CreditCard struct {
	gorm.Model
	Number     string
	EmployerID uint
}
