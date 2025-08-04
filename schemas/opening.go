package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Cargo string
	Empresa string
	Localizacao string
	Remoto bool
	Link string
	Salario int64
}