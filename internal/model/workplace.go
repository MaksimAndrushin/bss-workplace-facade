package model

import (
	"time"
)

type Workplace struct {
	ID      uint64    `db:"id"`
	Name    string    `db:"name"`
	Size    uint32    `db:"size"`
	Removed bool      `db:"removed"`
	Created time.Time `db:"created"`
	Updated time.Time `db:"updated"`
}
