package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

type KeyOutSql struct {
	conn *pgx.Conn
}

func NewKeyOutSql(conn *pgx.Conn) *KeyOutSql {
	return &KeyOutSql{
		conn: conn,
	}
}

func (r *KeyOutSql) Get(instance string) (int, int, error) {
	var start, end int
	log.Printf("Feting configs...\nBegin transaction....")
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		return -999, -999, err
	}
	//select * from configs where assigned is null limit 1;
	//UPDATE configs set assigned = 'instance A' where key_start = 1 and key_end = 10
	rows, err := tx.Query(context.Background(), "select key_start, key_end from configs where assigned is null limit 1;")
	if err != nil {
		return -999, -999, err
	}
	for rows.Next() {
		rows.Scan(&start, &end)
	}
	ct, err := tx.Exec(context.Background(), "UPDATE configs set assigned = $1 where key_start = $2 and key_end = $3", instance, start, end)
	tx.Commit(context.Background())
	if err != nil {
		log.Printf("Error while Updating row.....\n")
		return -999, -999, err
	}
	log.Printf("Transaction complete....\nRows affected %d \n", ct.RowsAffected())

	return start, end, nil
}

func (r *KeyOutSql) Add(start, last int) error {
	log.Printf("Starting to Bring Keys out\nBegin transaction....")
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		return err
	}
	ct, err := tx.Exec(context.Background(), "insert into keys_out (keyval,keyid) select keyval,keyid from keys where keyid >= $1 and keyid <= $2;", start, last)
	if err != nil {
		tx.Commit(context.Background())
		return err
	}
	log.Printf("Rows effected = %d", ct.RowsAffected())
	tx.Commit(context.Background())
	return nil
}
