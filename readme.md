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

![Database models]("https://github.com/vietthangc1/mini-web-golang/blob/main/database-map.png")