package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository() (*Repository, error) {
	conf := mysql.Config{
		User:   getEnv("root", "MARIADB_USER", "NS_MARIADB_USER"),
		Passwd: getEnv("", "MARIADB_PASSWORD", "NS_MARIADB_PASSWORD"),
		Net:    "tcp",
		Addr: getEnv("localhost", "MARIADB_HOST", "NS_MARIADB_HOST") +
			":" + getEnv("3306", "MARIADB_PORT", "NS_MARIADB_PORT"),
		DBName:    getEnv("", "MARIADB_NAME", "NS_MARIADB_NAME"),
		ParseTime: true,
	}
	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}

type repositoryUserGroup struct {
	UserID string `db:"id"`
	Name   string `db:"name"`
}

func (r *Repository) GetUserGroups(id string) ([]*repositoryUserGroup, error) {
	var userGroups []*repositoryUserGroup
	err := r.db.Select(&userGroups, "SELECT * FROM user_group WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return userGroups, nil
}
