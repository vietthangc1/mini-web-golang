# Mini web project

User Authentication and perform CRUD method on products.

## Tools

- Framework: Golang ([gin](https://github.com/gin-gonic/gin) framework)
- Database: Mysql (use [gorm](https://gorm.io/gorm) package to connect)
- Cache: Redis ([go-redis](https://github.com/redis/go-redis/v9) package)
- Authentication: [JWT-go](https://github.com/dgrijalva/jwt-go)
- Crypto: [Bcrypt](https://golang.org/x/crypto/bcrypt)

## Running

Run code in the terminal, the server will run on port 8080

```
go run main.go
```

## Database

- Users: Store Email and Password of users
- Products: Store products information
- Properties: Store other properties (brand, color, size) of products

<img alt="Model Database" src="https://www.linkpicture.com/q/database-map.png">

## Routes

Note: (*): require authorization, (**): use/update to cache

### User routes

- POST "/user": Register user. Require email, password in body. Hash registered password and save to database
- POST "/login": Log in user. Require email, password in body.
- GET "/user" (*): Get information of logged in user
- DELETE "/user" (*): Delete current user and log out


### Product routes

- POST "/product" (*, **): Add a product to database
- GET "/product/:_id" (**): Get information of product with id of _id
- GET "/products": Get all products with filter query (filter by cate1, cate2, cate3, cate4, color, brand, size)
- PUT "/products/:_id" (*, **): Update product with id of _id with the request body
- DELETE "/product/:_id" (*, **): Delete product with id of _id