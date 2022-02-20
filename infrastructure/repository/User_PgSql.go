package repository

import (
	"Ant-Man-Url/api/middleware"
	"Ant-Man-Url/entity"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

type UserSql struct {
	conn *pgx.Conn
}

func NewUserSql(conn *pgx.Conn) *UserSql {
	return &UserSql{
		conn: conn,
	}
}

func (r *UserSql) SignUp(userName, userPassword, userEmail, userRole string) (*entity.User, error) {
	//INSERT INTO users (username, email, hash_password, roles, url_count) VALUES ('Visjwajeet590','vishwajeet878@gmail.com','fdnfejfnjdnfjdsfsdfjsdfjsdfjnfnsfj','user',0) on conflict (email) do nothing RETURNING user_id
	hashedPassword, err := middleware.HashPassword(userPassword)
	if err != nil {
		return nil, fmt.Errorf("Error : %s", err)
	}
	log.Printf("Creating new user...\nBegin Transaction...")
	var id = -999
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		log.Printf("Error : Couldn't start transaction\n")
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := tx.Query(context.Background(), "INSERT INTO users (username, email, hash_password, roles, url_count) VALUES ($1,$2,$3,$4,$5) ON CONFLICT (email) DO NOTHING RETURNING user_id", userName, userEmail, hashedPassword, userRole, 0)

	fmt.Printf("INSETR DEKHO : %s", err)
	//	ct, err := tx.Exec(context.Background(), "INSERT INTO users (username, email, hash_password, roles, url_count) VALUES ($1,$2,$3,$4,$5) ON CONFLICT (email) DO NOTHING RETURNING user_id", username, email, hashedPassword, role, 0)

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
	}
	fmt.Printf("ID : %d", id)
	if id < 0 {
		defer rows.Close()
		tx.Commit(context.Background())
		return nil, fmt.Errorf("user is already signedUp")
	}
	defer rows.Close()
	tx.Commit(context.Background())

	return &entity.User{
		Username: userName,
		Id:       int32(id),
		Roles:    userRole,
		Token:    "",
	}, nil

}

func (r *UserSql) Login(email, password string) (*entity.User, error) {
	log.Printf("Loging in user...\nBegin Connection\n")
	var dbPass, username string
	var userId int
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		log.Printf("Error : Couldn't start transaction\n")
		return nil, err
	}

	row := tx.QueryRow(context.Background(), "SELECT user_id, username,hash_password FROM users WHERE email = $1 LIMIT 1", email)
	err = row.Scan(&userId, &username, &dbPass)
	if err != nil {
		//log.Printf("Error : Couldn't fetch dbPass")
		return nil, err
	}
	err = middleware.CheckPassword(password, dbPass)
	if err != nil {
		//log.Printf("Incorrect Password")
		return nil, err
	}

	jwt, err := middleware.GenerateToken(userId, "user")
	if err != nil {
		return nil, err
	}
	ct, err := tx.Exec(context.Background(), "UPDATE users SET jwt_token = $1 WHERE user_id = $2", jwt, userId)
	tx.Commit(context.Background())
	fmt.Printf("Rows affected : %d\n", ct.RowsAffected())
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Username: username,
		Id:       int32(userId),
		Token:    jwt,
		Roles:    "user",
	}, nil
}

func (r *UserSql) Get(user_id int, keyval string) (*entity.Url, error) {
	var longUrl string
	var createdAt time.Time
	var userId, redirectCount int

	log.Printf("Fetching stat for key : %s\nBegin transaction...\n", keyval)
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		log.Printf("Error : Couldn't start transaction\n")
		return nil, err
	}

	//ownerRow := tx.QueryRow(context.Background(), "select user_id from keys_out where keyval = $1 limit 1;", keyval)
	//err = ownerRow.Scan(&urlOwner)
	//if err != nil {
	//	return nil,
	//}
	//
	//if user_id != urlOwner {
	//	return nil, fmt.Errorf("You don't own this url.")
	//}

	row := tx.QueryRow(context.Background(), "SELECT url,redirect_count,created_at,user_id from keys_out where keyval = $1 limit 1", keyval)
	err = row.Scan(&longUrl, &redirectCount, &createdAt, &userId)
	tx.Commit(context.Background())
	if err != nil {
		fmt.Printf("Error : %s", err)
		return nil, fmt.Errorf("Error 404")
	}
	if userId != user_id {
		return nil, fmt.Errorf("You don't own this url.\n")
	}
	return &entity.Url{
		Long_link:  longUrl,
		Keyval:     keyval,
		Redirects:  redirectCount,
		OwnerId:    userId,
		Created_at: createdAt,
	}, nil
}

func (r *UserSql) List(user_id int) ([]*entity.Url, error) {

	log.Printf("Fetching url list...\nBegin Transaction...\n")
	rowCount := 0
	var long, key string
	var created_at time.Time
	var red, oId int
	var counter = 0
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		log.Printf("Error : Couldn't start transaction\n")
		return nil, err
	}

	coutRow := tx.QueryRow(context.Background(), "SELECT COUNT(*) from keys_out WHERE user_id = $1", user_id)
	err = coutRow.Scan(&rowCount)
	if err != nil {
		return nil, err
	}

	//To comment
	fmt.Printf("ROW COUNT : %d", rowCount)

	rows, err := tx.Query(context.Background(), "SELECT url,redirect_count,created_at,user_id,keyval from keys_out WHERE user_id = $1", user_id)
	if err != nil {
		return nil, err
	}
	if rowCount == 0 {
		return nil, fmt.Errorf("No url found against the user...")
	}
	urlList := make([]*entity.Url, rowCount)
	for rows.Next() {
		curr := new(entity.Url)

		//curr := urlList[counter]
		err = rows.Scan(&long, &red, &created_at, &oId, &key)
		if err != nil {
			return nil, err
		}
		curr.Long_link = long
		curr.Keyval = key
		curr.Created_at = created_at
		curr.Redirects = red
		curr.OwnerId = oId

		urlList[counter] = curr
		counter++
	}
	tx.Commit(context.Background())

	return urlList, nil

}
