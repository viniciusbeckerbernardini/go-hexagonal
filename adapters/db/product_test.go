package db_test

import (
	"database/sql"
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"github.com/viniciusbeckerbernardini/go-hexagonal/adapters/db"
	"github.com/viniciusbeckerbernardini/go-hexagonal/application"
	"log"
	"testing"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := "CREATE TABLE products (id text, name text, price float, status text)"
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := "INSERT INTO products VALUES ('abc', 'Product 1', 10, 'disabled')"
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	//espera o comando ser executado para depois fechar a conexão
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	//espera o comando ser executado para depois fechar a conexão
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 20
	product.Status = application.ENABLED

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}

func TestProductDb_Update(t *testing.T) {
	setUp()
	//espera o comando ser executado para depois fechar a conexão
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.ID = "abc"
	product.Name = "Product 2"
	product.Price = 0
	product.Status = application.DISABLED

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
