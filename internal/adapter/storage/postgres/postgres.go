package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/evrintobing17/go-hexagonal-arch/internal/adapter/config"
	_ "github.com/lib/pq"
)

/**
 * DB is a wrapper for PostgreSQL database connection
 * that uses pgxpool as database driver.
 * It also holds a reference to squirrel.StatementBuilderType
 * which is used to build SQL queries that compatible with PostgreSQL syntax
 */

type DB struct {
	DB *sql.DB
}

func New(ctx context.Context, config *config.DB) (*DB, error) {

	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}

func (db *DB) Close() {
	db.DB.Close()
}
