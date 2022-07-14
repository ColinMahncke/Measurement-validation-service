package objects

import (
	"context"
	"errors"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GormGeometry /*
//The Gorm Geometry Type should be used for entities and their geometry.
//It can handle the parsing from geometric types out of the box.
type GormGeometry struct {
	Geometry orb.Geometry
}

func (geo *GormGeometry) GormDataType() string {
	return "bytea"
}

func (geo GormGeometry) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{ewkb.Value(geo.Geometry, 4326)},
	}
}

func (geo *GormGeometry) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("Object in databases is not an byte area: %v.", value))
	}
	unmarshal, _, err := ewkb.Unmarshal(bytes)
	if err != nil {
		return err
	}
	geo.Geometry = unmarshal
	return nil
}
