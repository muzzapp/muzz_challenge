package testcontainerx

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

func StartMySqlContainer() (*sqlx.DB, *mariadb.MariaDBContainer, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	var migrations []string
	execDir := filepath.Dir(filename)
	err := filepath.Walk(
		filepath.Join(execDir, "sql"),
		func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".sql" {
				migrations = append(migrations, path)
			}
			return nil
		},
	)
	if err != nil {
		panic(err)
	}

	img := "mariadb:10.6"
	ctx := context.Background()
	msqlContainer, err := mariadb.Run(
		ctx,
		img,
		mariadb.WithDatabase("foo"),
		mariadb.WithUsername("root"),
		mariadb.WithPassword("password"),
		mariadb.WithScripts(migrations...),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("run container: %w", err)
	}

	str, err := msqlContainer.ConnectionString(
		ctx,
		"parseTime=true",
		"charset=utf8mb4",
		"loc=UTC",
	)
	if err != nil {
		return nil, nil, fmt.Errorf("connection string: %w", err)
	}

	open, err := sql.Open("mysql", str)
	if err != nil {
		return nil, nil, fmt.Errorf("open: %w", err)
	}
	db := sqlx.NewDb(open, "mysql")
	// check DB connection
	for i := 0; i < 10; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("ping: %w", err)
	}

	return db, msqlContainer, nil
}

func StopMySqlContainer(ctx context.Context, db *sqlx.DB, msqlContainer *mariadb.MariaDBContainer) {
	if db != nil {
		_ = db.Close()
	}

	if msqlContainer != nil {
		_ = msqlContainer.Terminate(ctx)
	}
}
