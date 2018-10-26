// Package domain provide interfaces and types represented domain logic.
package domain

// Cookie entity.
//easyjson:json
type Cookie struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// CookieRepository repository interface.
type CookieRepository interface {
	Save(e *Cookie) error
	FindOneByID(id int64) (cookie *Cookie, err error)
}
