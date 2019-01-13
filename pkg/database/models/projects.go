package models

import (
	null "gopkg.in/guregu/null.v3"
)

type Project struct {
	Id          null.Int `gorm:"column:Id;PRIMARY_KEY"`
	Company     string
	Background  string
	Image       string
	Title       string
	Description string
}
