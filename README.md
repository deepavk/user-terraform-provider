This custom terraform provider acts as a bridge between terraform and a simple mock users http service.
The http service exposes APIs to allow CRUD operations on a user resource.


### Requirements
   * [Golang](https://golang.org/doc/install)
   * [Terraform CLI](https://learn.hashicorp.com/tutorials/terraform/install-cli) 


### Starting the go server 

``` 
    cd gohttpsrv
    go run main.go
```
* The server can be accessed at http://localhost:8000   
   
    
### Building the provider
* The makefile is used to install provider
    ```
    $ make install
    ```


### Applying resource changes

* On the root folder these commands are run to apply changes:
    ```
    
    $ terraform init 
    $ terraform plan
    $ terraform apply
    

    ```
