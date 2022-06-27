package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FieldRelation struct {
	gorm.Model
	Flag       string
	Attributes datatypes.JSON
}
