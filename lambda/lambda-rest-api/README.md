## Notes
* The lambda module had problems with building the application so I did manually.

    ##### Build lambda function 
    * `go get -v all`
    * `GOOS=linux go build -o bin/main cmd/main.go`
    * `zip -jrm bin/main.zip bin/main` (havent tried it)
    
* A sample payloud would be 
    ```json
    {"email" : "bb@aaa.com",
     "firstName" : "Merve",
     "lastName" : "Mörviii" } 
    ```



 