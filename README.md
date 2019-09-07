# Simple Restful Service
It is a simple restful service in GO lang which handle basic operations for a User like retrieve user/s, 
create, update and delete user.


Installation: 
1. Need [go lang installation](https://golang.org/dl/).
1. Download mux package `go get -u github.com/gorilla/mux`

Run application:
1. Go to `/src`
1. Run command `go run main.go`
1. It will run the server on `8001` port

Examples:

1. Retrieve all users, endpoint - `/users` (method: GET)
 
    ```json
        [{"Id":1,"Name":"First User","Address":"First Country"},{"Id":2,"Name":"Second User","Address":"Second Country"}]

    ```

1. Retrieve user with id `1`, endpoint - `/user/1` (method: GET)  
    
    ```json
     {"Id":1,"Name":"First User","Address":"First Country"}
    ```