/*
								Book store management api's

1. Database - Mysql
2. GORM
3. Json marshall, unmarshall
4. Project Structure
5. Gorilla mux

Project structure

cmd
---
main.go

Pkg
---
Config
	APP.GO
Controllers
	BOOK CONTROLLER
Models
	BOOK.GO
Routes
	BOOKSTORE ROUTES
Utils
	UTILS.GO
--------------------------------

Routes
------
Controller functions

GETBOOKS				<----	/book/				<----GET

GETBOOK BY ID			<----	/book/{bookid}		<----GET

CREATE BOOK				<----	/book/				<----POST

UPDATE BOOK				<----	/book/{bookid}		<----PUT

DELETE BOOK				<----	/book/{bookid}		<----DELETE


******************************
PS D:\heleena\golang\projects\go-bookstore> go mod init go-bookstore
go: creating new go.mod: module go-bookstore
PS D:\heleena\golang\projects\go-bookstore> go get github.com/jinzhu/gorm
go: added github.com/jinzhu/inflection v1.0.0
PS D:\heleena\golang\projects\go-bookstore> go get github.com/jinzhu/gorm/dialects/mysql
PS D:\heleena\golang\projects\go-bookstore> go get github.com/gorilla/mux
go: added github.com/gorilla/mux v1.8.1
PS D:\heleena\golang\projects\go-bookstore> go mod tidy
PS D:\heleena\golang\projects\go-bookstore>
***************************************

*/

