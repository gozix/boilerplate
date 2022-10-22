// Package database provide various implementation works with database.
package database

import (
	gzSQL "github.com/gozix/sql/v3"

	"github.com/iqoption/nap"

	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// CookieRepository implements domain.CookieRepository
type CookieRepository struct {
	db *nap.DB
}

// NewCookie is repository constructor.
func NewCookie(registry *gzSQL.Registry) (_ *CookieRepository, err error) {
	var db *nap.DB
	if db, err = registry.Connection(); err != nil {
		return nil, err
	}

	return &CookieRepository{
		db: db,
	}, nil
}

// Save implementation.
func (r *CookieRepository) Save(e *domain.Cookie) (err error) {
	var query = `
		INSERT INTO cookie (name) VALUES ($1)
		RETURNING id
	`

	return r.db.Master().QueryRow(query, e.Name).Scan(&e.ID)
}

// FindOneByID implementation.
func (r *CookieRepository) FindOneByID(id int64) (cookie *domain.Cookie, err error) {
	var query = `SELECT id, name FROM cookie WHERE id = $1`

	cookie = new(domain.Cookie)
	if err = r.db.QueryRow(query, id).Scan(&cookie.ID, &cookie.Name); err != nil {
		return nil, err
	}

	return cookie, nil
}
