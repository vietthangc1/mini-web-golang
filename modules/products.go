package modules

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/vietthangc1/mini-web-golang/models"
)

func QueryGetProductByID (db *sql.DB, q string, id string) (models.Product, error) {
	var productQuery models.Product 

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
		return models.Product{}, err
	}
	return productQuery, nil
}

func QueryGetProducts (db *sql.DB, q, cate1, cate2, cate3, cate4 string) ([]models.Product, error) {
	var (
		productQuery  models.Product
		productsQuery []models.Product
	)

	stmt, err := db.Prepare(q)
	if err != nil {
		return []models.Product{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(cate1, cate2, cate3, cate4)
	if err != nil {
		return []models.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&productQuery.ID,
			&productQuery.Name,
			&productQuery.Price,
			&productQuery.Cate1,
			&productQuery.Cate2,
			&productQuery.Cate3,
			&productQuery.Cate4,
		)
		if err != nil {
			return []models.Product{}, err
		}
		productsQuery = append(productsQuery, productQuery)
	}
	err = rows.Err()
	if err != nil {
		return []models.Product{}, err
	}
	return productsQuery, nil
}

func QueryAddProduct(db *sql.DB, q string, p models.Product) (models.Product, error) {
	var newProduct models.Product

	id := time.Now().UnixMilli()
	newProduct.ID = strconv.Itoa(int(id))

	stmt, err := db.Prepare(q)
	if err != nil {
		return models.Product{}, err
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
		return models.Product{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return models.Product{}, err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return p, nil
}

func QueryUpdateProduct(db *sql.DB, q string, id string, p models.Product) (models.Product, error) {
	stmt, err := db.Prepare(q)
	if err != nil {
		return models.Product{}, err
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
		return models.Product{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return models.Product{}, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return models.Product{}, err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return p, nil
}

func QueryDeleteProduct(db *sql.DB, q, id string) (error) {
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