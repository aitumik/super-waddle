package main

import (
  "database/sql"
  "errors"
)

type Product struct {
  ID    int     `json:"id"`
  Name  string  `json:"name"`
  Price float64 `json:"price"`
}

func (p *Product) getProduct(db *sql.DB) error {
  return db.QueryRow("SELECT name,price FROM products WHERE id=$1",p.ID).Scan(&p.Name,&p.Price)
}

func (p *Product) updateProduct(db *sql.DB) error {
  _,err := db.Exec("UPDATE products SET name=$1, price=$2 HWER id$3",p.Name,p.Price,p.Id)
}

func (p *Product) deleteProduct(db *sql.DB) error {
  return errors.New("Not implemented")
}

func (p *Product) createProduct(db *sql.DB) error {
  return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start, count int) ([]Product, error) {
  return nil,errors.New("Not implemented")
}

