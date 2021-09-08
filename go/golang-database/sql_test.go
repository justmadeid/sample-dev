package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	//can insert update delet
	script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new customer")

}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
	}

	defer rows.Close()

}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("========================")
		fmt.Println("Id : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		fmt.Println("Birth Date : ", birthDate)
		fmt.Println("Married : ", married)
		fmt.Println("Created at : ", createdAt)
	}

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	sqlQuery := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("Login Gagal")
	}

}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	} else {
		fmt.Println("Login Gagal")
	}

}

func TestExecSqlSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "made"
	password := "made"

	//can insert update delet
	script := "INSERT INTO user(username,password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Succes insert new user")

}

func TestExecSqlAi(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "made@gmail.com"
	coment := "made"

	//can insert update delet
	script := "INSERT INTO coments(email,coment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, coment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last insert id", insertId)

	fmt.Println("Succes insert new coment")

}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO coments(email,coment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 10; i++ {
		email := "baru" + strconv.Itoa(i) + "@gmail.com"
		coment := "hahah lucu" + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, coment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("coment id", id)
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO coments(email,coment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "baru" + strconv.Itoa(i) + "@gmail.com"
		coment := "hahah lucu" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, coment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("coment id", id)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
