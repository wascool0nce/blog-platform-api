package sqlite

import (
	"database/sql"

	"github.com/wascool0nce/blog-platform-api/config"
)

func ConnectDB(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.Path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
