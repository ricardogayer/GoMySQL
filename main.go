package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Produto struct {
	Codigo    int
	Descricao string
}

func main() {

	db, err := sql.Open("mysql", "golang:S3nh4C0mpl3x4$@tcp(localhost:3306)/golang")
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer db.Close()

	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	rows, err := db.QueryContext(context.Background(),
		"SELECT cd_produto, ds_produto FROM produto")
	if err != nil {
		fmt.Println("Erro ao executar um comando no banco de dados", err)
		return
	}

	defer rows.Close()

	var produtos []Produto

	for rows.Next() {
		var p Produto
		err := rows.Scan(&p.Codigo, &p.Descricao)
		if err != nil {
			fmt.Println("Erro ao recuperar uma linha", err)
		}
		produtos = append(produtos, p)
	}

	for _, p := range produtos {
		fmt.Println(p.Codigo, p.Descricao)
	}

}
