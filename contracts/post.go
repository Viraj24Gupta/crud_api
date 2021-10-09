package contracts

import (
	"time"
)

type Posts struct {
	Id      int64     `bson:"id"`
	Caption string    `bson:"caption"`
	Image   string    `bson:"image"`
	Posted  time.Time `bson:"posted"`
}
