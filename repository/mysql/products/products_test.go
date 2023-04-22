package products

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/repository/mysql/db"
)

func TestAddProduct(t *testing.T) {
	art := assert.New(t)
	product := models.Product{
	}

	jsonFile, err := os.Open("addProduct.json")
	art.Nil(err, "error in open json mock")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &product)
	art.Nil(err, "error in unmarshalling")

	dbTest, err := db.ConnectDatabaseORMTest(true)
	art.Nil(err, "error in connecting database")

	impl := NewProductRepo(dbTest)
	newProduct, err := impl.AddProduct(product)

	art.Nil(err, "error in add product")
	art.Equal(product.SKU, newProduct.SKU, "wrong in response add product")

	productId := newProduct.ID
	productQuery, err := impl.GetProductByID(productId)

	art.Nil(err, "error in get added product")
	art.Equal(product.SKU, productQuery.SKU, "conflict information between input and output")
}
