# go-basic-rest-api
My pratices build rest api using Go-Lang
reference : [https://github.com/kkamdooong/go-restful-api-example](https://github.com/kkamdooong/go-restful-api-example)

## how to install
```
git clone https://github.com/ikhlashmulya/go-basic-rest-api.git
cd go-basic-rest-api
go mod tidy
go run server.go
```

## endpoint
[http://localhost:8080/api/students](http://localhost:8080/api/students)
* method `GET` : get list of students
* method `POST` : add new data student

[http://localhost:8080/api/students/{id}](http://localhost:8080/api/students/{id})
* method `GET` : get data student by id
* method `PUT` : update data student
* method `DELETE` : delete data student

## JSON
```json
{
    "id" : "string", 
    "name" : "string", 
    "email" : "string" 
} 
```
