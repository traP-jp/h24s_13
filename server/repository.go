package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/traP-jp/h24s_13/server/utils/ds"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository() (*Repository, error) {
	conf := mysql.Config{
		User:   getEnv("root", "MARIADB_USER", "NS_MARIADB_USER"),
		Passwd: getEnv("", "MARIADB_PASSWORD", "NS_MARIADB_PASSWORD"),
		Net:    "tcp",
		Addr: getEnv("localhost", "MARIADB_HOSTNAME", "NS_MARIADB_HOSTNAME") +
			":" + getEnv("3306", "MARIADB_PORT", "NS_MARIADB_PORT"),
		DBName:               getEnv("", "MARIADB_DATABASE", "NS_MARIADB_DATABASE"),
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}

type repositoryUser struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}

type repositoryUserGroup struct {
	UserID string `db:"id"`
	Name   string `db:"name"`
}

func (r *Repository) GetUser(id string) (*User, error) {
	var ru repositoryUser
	err := r.db.Get(&ru, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	var rug []*repositoryUserGroup
	err = r.db.Select(&rug, "SELECT * FROM user_groups WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:     ru.ID,
		Name:   ru.Name,
		Groups: ds.Map(rug, func(rug *repositoryUserGroup) string { return rug.Name }),
	}, nil
}

func (r *Repository) GetUserIDs() ([]string, error) {
	var ids []string
	err := r.db.Select(&ids, "SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	return ids, nil
}

type repositoryUserConnection struct {
	ID1      string  `db:"id_1"`
	ID2      string  `db:"id_2"`
	Strength float64 `db:"strength"`
}

func (r *Repository) GetConnections(id string) (map[string]float64, error) {
	var rug []*repositoryUserConnection
	err := r.db.Select(&rug, "SELECT * FROM user_connections WHERE id_1 = ?", id)
	if err != nil {
		return nil, err
	}

	connections := make(map[string]float64, len(rug))
	for _, c := range rug {
		connections[c.ID2] = c.Strength
	}

	return connections, nil
}

func (r *Repository) GetFriends(ids []string) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var query = `SELECT DISTINCT id_2 FROM (
    SELECT
        *,
        ROW_NUMBER() OVER (PARTITION BY id_1 ORDER BY strength DESC) AS num
    FROM user_connections
) t
JOIN (
    SELECT id_1 AS id, COUNT(*) / 2 AS border
    FROM user_connections
    GROUP BY id_1
) border ON t.id_1 = border.id
WHERE id_1 IN (?)
    AND num <= border`

	args := ds.Map(ids, func(s string) any { return s })
	query, args, err := sqlx.In(query, args)
	if err != nil {
		return nil, err
	}

	var result []string
	err = r.db.Select(&result, query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
