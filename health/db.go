package health

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	config "github.com/piyakker/GoMall/configs"
)

type PostgresChecker struct {
	cfg *config.Config
}

func NewPostgresChecker(cfg *config.Config) *PostgresChecker {
	return &PostgresChecker{cfg: cfg}
}

func (p *PostgresChecker) Check(ctx context.Context) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.cfg.DBHost, p.cfg.DBPort, p.cfg.DBUser, p.cfg.DBPass, p.cfg.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.PingContext(ctx)
}
