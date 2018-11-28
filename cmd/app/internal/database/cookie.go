// Package database provide various implementation works with database.
package database

import (
	"github.com/gozix/sql"
	"github.com/iqoption/nap"
	"github.com/sarulabs/di"

	"github.com/gozix/boilerplate/cmd/app/internal/domain"
)

// CookieRepository implements domain.CookieRepository
type CookieRepository struct {
	db *nap.DB
}

// DefCookieRepositoryName is container name.
const DefCookieRepositoryName = "database.repository.cookie"

// DefCookieRepository register repository in di container.
func DefCookieRepository() di.Def {
	return di.Def{
		Name: DefCookieRepositoryName,
		Build: func(ctn di.Container) (_ interface{}, err error) {
			var registry *sql.Registry
			if err = ctn.Fill(sql.BundleName, &registry); err != nil {
				return nil, err
			}

			var db *nap.DB
			if db, err = registry.Connection(); err != nil {
				return nil, err
			}

			return NewCookieRepository(db), nil
		},
	}
}

// NewCookieRepository constructor
func NewCookieRepository(db *nap.DB) *CookieRepository {
	return &CookieRepository{
		db: db,
	}
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
