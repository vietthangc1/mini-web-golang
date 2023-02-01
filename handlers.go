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
	SELECT id, sku, name, price, number, description, cate1, cate2, color, size
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
		&productQuery.Color,
		&productQuery.Size,
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

	db, _ := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")

	query := `
	SELECT id, sku, name, price, number, description, cate1, cate2, color, size
	FROM products 
	WHERE 1=1
	AND cate1 like ?
	AND cate2 like ?
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(cate1, cate2)
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
			&productQuery.Color,
			&productQuery.Size,
		)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		}
		productsQuery = append(productsQuery, productQuery)
	}
	err = rows.Err()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
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

	query := "INSERT INTO products ( id, sku, name, price, number, description, cate1, cate2, color, size) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
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
		newProduct.Color,
		newProduct.Size,
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
	SET sku = ?, name = ?, price = ?, number = ?, description = ?, cate1 = ?, cate2 = ?, color = ?, size = ?
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
		updateProduct.Color,
		updateProduct.Size,
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