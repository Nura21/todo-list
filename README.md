# Setup
- mkdir myapp && cd myapp
- go mod init myapp
- go get github.com/labstack/echo/v4

# Package Install
- go get github.com/tkanos/gonfig
- go get -u github.com/go-sql-driver/mysql

# To sync the vendor
- go mod vendor

# To run the project
- go run main.go

# Reference
- https://github.com/tkanos/gonfig
- https://github.com/go-sql-driver/mysql
- https://echo.labstack.com/guide/#installation
- https://www.youtube.com/watch?v=M5JWCzeyYEE&list=PLO2Rv4lKm-K0vR95sfEXznno4421Z67OF&index=5 = #5: Membangun RESTful API dengan Golang Echo Framework - Method GET Untuk Pengambilan Data