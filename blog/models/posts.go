package models

import (
	"database/sql"
	"errors"

	"github.com/wycliff-ochieng/blog/data"
)

type Storage interface {
	CreatePosts(*data.Post) error
	DeletePost(*data.Post) error
	GetPosts(int) error
	UpdatePost(int) error
}

type PostgresPostStore struct {
	db *sql.DB
}

func NewPostgresPostStore() (*PostgresPostStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open(connStr, "postgres")
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresPostStore{
		db: db,
	}, nil
}

func (p *PostgresPostStore) CreatePostTable() error {
	query := `CREATE IF NOT EXISTS TABLE posts(
	id INT PRIMARY KEY,
	author VARCHAR(20),
	title VARCHAR(20),
	body TEXT,
	timecreated TIMESTAMP)`

	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresPostStore) CreatePost(pst *data.Post) error {
	query := `INSERT INTO post(id,author,title,body,timecreated) VALUES(&1,&2,&3,&4,&5)`
	_, err := p.db.Exec(query, &pst.ID, pst.Author, pst.Title, pst.Body, pst.TimeCreated)
	if err != nil {
		return errors.New("unable to insert data into posts table")
	}
	return err
}

func (p *PostgresPostStore) GetPosts(int) error {
	pst := &data.Post{}
	query := `SELECT id,author,title,body,timecreate`
	err := p.db.QueryRow(query).Scan(&pst.ID, &pst.Author, &pst.Title, &pst.Body, &pst.TimeCreated)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresPostStore) DeletePost(pst *data.Post) error {
	return nil
}

func (p *PostgresPostStore) UpdatePost(int) error {
	return nil
}
