package utils

import (
	"github.com/spaceCh1mp/pow/server/db"
)

// Collection sets the db configuration for a new connection
func Collection(p *db.LiveSession, str string) error {
	err := p.SetCollection(str)
	if err != nil {
		return ErrFC
	}

	return nil
}
