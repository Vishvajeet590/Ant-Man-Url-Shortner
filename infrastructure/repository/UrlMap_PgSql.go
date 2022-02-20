package repository

import (
	"Ant-Man-Url/entity"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"log"
)

type UrlMapSql struct {
	conn *pgx.Conn
}

func NewUrlMapSql(conn *pgx.Conn) *UrlMapSql {
	return &UrlMapSql{
		conn: conn,
	}
}

func (r *UrlMapSql) Link(LongUrl *entity.Url, isJwt bool, userId *int) (*entity.Url, error) {
	//select * from keys_out where url is null limit 1;
	//UPDATE keys_out set url = 'www.google.com' where keyval = ABCD
	var key_id int
	var keyval string
	var ct pgconn.CommandTag
	log.Printf("Linking URL...\nBegin transaction...")
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		log.Printf("Error : Couldn't start transaction\n")
		return nil, err
	}
	rows, err := tx.Query(context.Background(), "select key_id, keyval from keys_out where url is null limit 1;")
	if err != nil {
		log.Printf("Error : Couldn't fetch empty key...")
	}
	for rows.Next() {
		rows.Scan(&key_id, &keyval)
	}
	if key_id < 0 && keyval == "" {
		return nil, nil
	}
	fmt.Printf("\n\n %s , %d \n\n", keyval, key_id)
	if isJwt {
		fmt.Printf("\nisJwt : %v \n", isJwt)
		ct, err = tx.Exec(context.Background(), "UPDATE keys_out set url = $1, user_id = $2, redirect_count = $3 where key_id = $4", LongUrl.Long_link, *userId, 0, key_id)
	} else {
		ct, err = tx.Exec(context.Background(), "UPDATE keys_out set url = $1, redirect_count = $2 where key_id = $3", LongUrl.Long_link, 0, key_id)
	}

	tx.Commit(context.Background())

	if err != nil {
		log.Printf("Fething error....")
		log.Printf("Error : %s", err)
		return nil, err
	}

	log.Printf("Transaction complete...\nRows affected : %d \n", ct.RowsAffected())
	LongUrl.Keyval = keyval
	return LongUrl, nil // returning LongUrl as shortUrl by adding keyval
}

func (r *UrlMapSql) Resolve(shortUrl *entity.Url) (*entity.Url, error) {
	log.Printf("Fetching URL...\n Begin Transaction\n")
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		log.Printf("Error : Couldn't start transaction\n")
		return nil, err
	}
	rows, err := tx.Query(context.Background(), "select url from keys_out where keyval = $1", shortUrl.Keyval)

	if err != nil {
		log.Printf("Error : Couldn't fetch keyval...")
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&shortUrl.Long_link)
	}
	tx.Commit(context.Background())
	log.Printf("Transaction complete...\n")
	return shortUrl, nil

}

func (r *UrlMapSql) Delete(ShortUrl *entity.Url, isJwt bool, userId *int) (bool, error) {
	var urlOwner = 0
	//DELETE FROM keys_out WHERE key_id = 1
	log.Printf("Deleting URL...\nBegin Transaction\n")
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		log.Printf("Error : Couldn't start transaction\n")
		return false, err
	}

	row := tx.QueryRow(context.Background(), "select user_id from keys_out where keyval = $1 limit 1;", ShortUrl.Keyval)
	err = row.Scan(&urlOwner)
	if err != nil {
		return false, err
	}

	if *userId != urlOwner {
		return false, fmt.Errorf("You don't own this url.")
	}

	ct, err := tx.Exec(context.Background(), "DELETE FROM keys_out WHERE keyval = $1", ShortUrl.Keyval)
	tx.Commit(context.Background())
	if ct.Delete() == false {
		fmt.Printf("Error : \n", err)
		return false, nil
	}
	fmt.Printf("Transaction complete...\n")
	return true, nil

}
