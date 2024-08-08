package dao

import (
	"github.com/bitwormhole/naspan-client/app/data/dxo"
	"github.com/bitwormhole/naspan-client/app/data/entity"
	"gorm.io/gorm"
)

// ExampleDAO ...
type ExampleDAO interface {
	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)
}
