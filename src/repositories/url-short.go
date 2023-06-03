package repositories

import (
	"github.com/gustavoteixeira8/url-shortener/src/db"
	"github.com/gustavoteixeira8/url-shortener/src/entities"
	"gorm.io/gorm"
)

type URLShortRepository struct {
	db *gorm.DB
}

func (r URLShortRepository) Save(urlShort *entities.URLShort) error {
	err := r.db.Save(urlShort).Error
	return err
}

func (r URLShortRepository) Delete(id string) error {
	err := r.db.Delete(entities.URLShort{}, "id = ?", id).Error
	return err
}

func (r URLShortRepository) ExistsWithURL(url string) bool {
	count := new(int64)
	err := r.db.Where("url = ?", url).Count(count).Error

	if err != nil {
		return false
	}

	return *count != 0
}

func (r URLShortRepository) ExistsWithName(name string) bool {
	count := new(int64)
	err := r.db.Where("name = ?", name).Count(count).Error

	if err != nil {
		return false
	}

	return *count != 0
}

func (r URLShortRepository) FindByID(id string) (*entities.URLShort, error) {
	urlShort := new(entities.URLShort)
	err := r.db.First(urlShort, "id = ?", id).Error
	return urlShort, err
}

func (r URLShortRepository) FindByURL(url string) (*entities.URLShort, error) {
	urlShort := new(entities.URLShort)
	err := r.db.First(urlShort, "url = ?", url).Error
	return urlShort, err
}

func (r URLShortRepository) FindByName(name string) (*entities.URLShort, error) {
	urlShort := new(entities.URLShort)
	err := r.db.First(urlShort, "name = ?", name).Error
	return urlShort, err
}

func NewURLShortRepository() *URLShortRepository {
	return &URLShortRepository{db: db.GetDB()}
}
