package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func QueryGetProductByID (q string, id string) (product, error) {
	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	var productQuery product

	err := db.QueryRow(q, id).Scan(
		&productQuery.ID,
		&productQuery.SKU,
		&productQuery.Name,
		&productQuery.Price,
		&productQuery.Number,
		&productQuery.Description,
		&productQuery.Cate1,
		&productQuery.Cate2,
		&productQuery.Cate3,
		&productQuery.Cate4,
		&productQuery.Propertises,
	)

	if err != nil {
		return product{}, err
	}
	return productQuery, nil
}

func QueryGetProducts (q, cate1, cate2, cate3, cate4 string) ([]product, error) {
	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	var (
		productQuery  product
		productsQuery []product
	)

	stmt, err := db.Prepare(q)
	if err != nil {
		return []product{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(cate1, cate2, cate3, cate4)
	if err != nil {
		return []product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&productQuery.ID,
			&productQuery.SKU,
			&productQuery.Name,
			&productQuery.Price,
			&productQuery.Number,
			&productQuery.Description,
			&productQuery.Cate1,
			&productQuery.Cate2,
			&productQuery.Cate3,
			&productQuery.Cate4,
			&productQuery.Propertises,
		)
		if err != nil {
			return []product{}, err
		}
		productsQuery = append(productsQuery, productQuery)
	}
	err = rows.Err()
	if err != nil {
		return []product{}, err
	}
	return productsQuery, nil
}

func QueryAddProduct(q string, p product) (product, error) {
	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	var newProduct product

	fmt.Println(p)

	id := time.Now().UnixMilli()
	newProduct.ID = strconv.Itoa(int(id))

	stmt, err := db.Prepare(q)
	if err != nil {
		return product{}, err
	}
	res, err := stmt.Exec(
		p.ID,
		p.SKU,
		p.Name,
		p.Price,
		p.Number,
		p.Description,
		p.Cate1,
		p.Cate2,
		p.Cate3,
		p.Cate4,
		p.Propertises.String(),
	)
	if err != nil {
		return product{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return product{}, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return product{}, err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return p, nil
}

func QueryUpdateProduct(q string, id string, p product) (product, error) {
	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")

	stmt, err := db.Prepare(q)
	if err != nil {
		return product{}, err
	}
	p.ID = id
	res, err := stmt.Exec(
		p.SKU,
		p.Name,
		p.Price,
		p.Number,
		p.Description,
		p.Cate1,
		p.Cate2,
		p.Cate3,
		p.Cate4,
		p.Propertises.String(),
		id,
	)
	if err != nil {
		return product{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return product{}, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return product{}, err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return p, nil
}

func QueryDeleteProduct(q, id string) (error) {
	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(
		id,
	)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return nil
}