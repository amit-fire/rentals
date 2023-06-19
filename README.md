# rentals

To run project: go run main.go
To run tests: go test ./... -cover (from the project's root folder)

Query Paramaters validations:
* id(s), limit, sort must be positive integer.
* min/max price should be a positive number. min price must be less than max price.
* sort can only be price, year, length, sleep.
* no validation is provided if the id doesn't exist. null is returned. should be handled later.
* in case the list of ids contains no existing id, the api returns ids that are found.
ideally, there should be a query parameter that allows the user to decide how to handle this scenario. return list of found ids or return an error. 

Libraries:
* Web Framework: Gin (https://github.com/gin-gonic/gin)
* Database: pq, PostgreSQL driver (https://github.com/lib/pq)
  * Note: Wanted to use GoRM (https://gorm.io), but had trouble collecting data (join) from the 2 tables.
  * For the sake of completing the task in time, the SQL package is used.
Switching to GoRM can be done later.
* Tests: Testify (https://github.com/stretchr/testify)
