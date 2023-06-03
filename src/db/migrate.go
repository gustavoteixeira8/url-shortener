package db

import (
	"time"

	"github.com/gustavoteixeira8/url-shortener/src/entities"
	"github.com/sirupsen/logrus"
)

func RunMigration() {
	logrus.Infof("Running migrations at %v", time.Now().UnixMilli())

	entities := []interface{}{
		entities.URLShort{},
	}

	for _, entity := range entities {
		err := _DB.AutoMigrate(entity)
		if err != nil {
			logrus.Fatalf(err.Error())
		}
	}
}
