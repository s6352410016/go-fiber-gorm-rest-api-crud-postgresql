package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustFullname string  `json:"custFullname"`
	CustAddress  string  `json:"custAddress"`
	CustTel      string  `json:"custTel"`
	CustSalary   float32 `json:"custSalary"`
}
