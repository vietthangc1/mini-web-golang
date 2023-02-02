package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	var (
		productQuery product
	)
	query := `
	SELECT id, sku, name, price, number, description, cate1, cate2, coalesce(cate3, '') as cate3, coalesce(cate4, '') as cate4, propertises
	FROM products 
	WHERE id = ?
	`
	err := db.QueryRow(query, id).Scan(
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
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusFound, productQuery)
}

func GetProducts(c *gin.Context) {
	var (
		productQuery  product
		productsQuery []product
	)

	filterProducts := c.Request.URL.Query()

	cate1 := filterProducts.Get("cate1")
	if cate1 == "" {
		cate1 = "%%"
	}
	cate2 := filterProducts.Get("cate2")
	if cate2 == "" {
		cate2 = "%%"
	}
	cate3 := filterProducts.Get("cate3")
	if cate3 == "" {
		cate3 = "%%"
	}
	cate4 := filterProducts.Get("cate4")
	if cate4 == "" {
		cate4 = "%%"
	}

	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")

	query := `
	SELECT id, sku, name, price, number, description, cate1, cate2, coalesce(cate3, '') as cate3, coalesce(cate4, '') as cate4, propertises
	FROM products 
	WHERE 1=1
	AND cate1 like ?
	AND cate2 like ?
	AND cate3 like ?
	AND cate4 like ?
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(cate1, cate2, cate3, cate4)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
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
			log.Fatal(err)
			// c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
			return
		}
		productsQuery = append(productsQuery, productQuery)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		// c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	// Pagination
	
	productPerPage := 2
	pageNum := 1
	
	page := filterProducts.Get("page")	
	if (page != "") {
		pageNum, _ = strconv.Atoi(page) 
	}

	positionStart := (pageNum-1)*productPerPage
	positionEnd := int(math.Min(float64(pageNum*productPerPage), float64(len(productsQuery))))

	productsPagination := []product{}
	if (positionStart < positionEnd) {
		productsPagination = productsQuery[positionStart:positionEnd]
	}

	c.IndentedJSON(http.StatusFound, productsPagination)
}

func AddProduct(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	id := time.Now().UnixMilli()
	newProduct.ID = strconv.Itoa(int(id))

	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")

	query := "INSERT INTO products ( id, sku, name, price, number, description, cate1, cate2, cate3, cate4, propertises) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(
		newProduct.ID,
		newProduct.SKU,
		newProduct.Name,
		newProduct.Price,
		newProduct.Number,
		newProduct.Description,
		newProduct.Cate1,
		newProduct.Cate2,
		newProduct.Cate3,
		newProduct.Cate4,
		newProduct.Propertises.String(),
	)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

func UpdateProduct(c *gin.Context) {
	var updateProduct product

	id := c.Param("id")

	if err := c.BindJSON(&updateProduct); err != nil {
		return
	}

	query := `
	UPDATE products 
	SET sku = ?, name = ?, price = ?, number = ?, description = ?, cate1 = ?, cate2 = ?, cate3 = ?, cate4 = ?, propertises = ?
	WHERE id = ?
	`

	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(
		updateProduct.SKU,
		updateProduct.Name,
		updateProduct.Price,
		updateProduct.Number,
		updateProduct.Description,
		updateProduct.Cate1,
		updateProduct.Cate2,
		updateProduct.Cate3,
		updateProduct.Cate4,
		updateProduct.Propertises.String(),
		id,
	)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	c.IndentedJSON(http.StatusCreated, updateProduct)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	query := `
	DELETE FROM products 
	WHERE id = ?
	`

	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(
		id,
	)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Deleted!"})
}