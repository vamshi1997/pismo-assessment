# Pismo Assessment #

## Introduction ##

This is a simple API server built with Go. It Contains three endpoints. It uses mysql as database


## Project Structure ##

```
├── README.md
├── cmd
│   └── main.go
├── configs
│   ├── default.toml
│   
├── internal
│   ├── models
│   │    └── account.go
│   │    └── transaction.go
│   |── boot
│   |    └── boot.go
│   │    └── config.go
│   |── repo
│   |    └── account.go
│   |    └── transaction.go
│   └── controllers
│   |    └── app.go
│   └── router
│       └── router.go
│       └── routes.go
|── go.mod
├── go.sum
|-- Dockerfile
|-- docker-compose.yaml
|-- run.sh
```


## Project Details ##


1. cmd/main file contains the main.go file, which is the starting point of the application
2. configs/default.toml file contains application configurations, such as database parameters and application host and port numbers.
3. internal file contains multiple folder related to application logic such as router, controller, boot & model
    i. router folder contains files which has endpoints and their associated names and functions
    ii. controllers contains files which process input requests such as retrieving requests body, validating and sending it to database.
    iii. model contains object entities for application. For example this containd account and transaction
    iv. repo contains common database functions (CRUD) respective to the model entities
    v. boot contain exteranal entities initialization such as database etc..,
4. go.mod and go.sum contains external go modules that at required for this application
5. docker-compose.yaml file contains application and database definations and volumes that are required for application to run in docker environment
6. Dockerfile contains set of commands for build image for the application and running them in docker
7. run.sh file contains commands which are useful if just execute run.sh instead of running all required commands.

## Prerequisites ##

1. Docker

## How to Run the Application in Docker ##

1. Run below commands in terminal
    ```bash run.sh``` or ```./run.sh```
2. Incase of permission issue run below command first
    ```chmod +x run.sh```
3. Once these are run successfully both database and application containers will be up and running. It may take some time to run them.
4. We can then hit endpoints accordingly to test the application

## How to Run the TestCases in Docker ##

1. Run below commands in terminal
    ```bash run-tests.sh``` or ```./run.sh```
2. Incase of permission issue run below command first
    ```chmod +x run-tests.sh```
3. Need to wait until build is done and you can see files with test percentage and test cases passed or not.


## Endpoints for the Applicatoin ##

### 1. For creating the account, we can use below curl. document number should be 11 character else we get error. ###

```
account create endpoint & curl:

curl --location 'http://localhost:8080/accounts' \
--header 'Content-Type: application/json' \
--data '{
    "document_number": "12345678901"
}'

multiple scenarios:

i. If valid document number is provided

response:

200 success
{
    "account_id": 3,
    "document_number": "12345678903",
    "msg": "Account created successfully"
}

ii. If invalid document number is provided, like document number len is not equal to 11

response:

400 bad request
{
    "error_msg": "Document number given is not valid",
    "msg": "Not able to create account"
}
```



### 2. For fetching account details, we can use below curl. Need to provide valid account number in top endpoint and it will fetch the result ###

```
account details fetch & curl:

curl --location 'http://localhost:8080/accounts/1'

multiple scenarios:

i. if valid account id is provided

response:

200 success
{
    "account_id": 1,
    "document_number": "12345678901",
    "msg":             "Account details fetched successfully"
}

ii. if invalid account id is provided which not exists in database, eg: provide 10 or something which is not in database

response:

404 Not Found
{
    "error": "record not found",
    "error_msg": "Account not found",
    "msg": "Not able to fetch account details"
}
```



### 3. For creating transaction, we need to use below curl. Need to provide valid account id, valid operantion_type_id and amount. ###

```
transaction create endpoint & curl:

curl --location 'http://localhost:8080/transactions' \
--header 'Content-Type: application/json' \
--data '{
    "account_id": 1,
    "operation_type_id": 3,
    "amount": -50
}'

multiple scenarios:

i. For valid details

request:

{
    "account_id": 1,
    "operation_type_id": 3,
    "amount": -50
}

response:

200 success
{
    "account_id": 1,
    "amount": -50,
    "msg": "transaction created successfully",
    "operation_type_id": 3,
    "transaction_id": 2
}

ii. For invalid account, providing account_id as 2 which not present in database

request:

{
    "account_id": 2,
    "operation_type_id": 5,
    "amount": -50
}

response:

404 Not Found
{
    "error": "record not found",
    "error_msg": "Account not found",
    "msg": "Not able to create transaction"
}

iii. Invalid Operation Type, providing 5 value which is not present

request:

{
    "account_id": 1,
    "operation_type_id": 5,
    "amount": -50
}

response:

400 Bad Request
{
    "error_msg": "Invalid operation type",
    "msg": "Not able to create transaction"
}


iv. Invalid Amount, for operation 4 credit voucher amount should be positive

request:

{
    "account_id": 1,
    "operation_type_id": 4,
    "amount": -50
}

response:

400 Bad Request
{
    "error_mgs": "Amount can not be negative for this operation type",
    "msg": "Not able to create transaction"
}
```

New Features changes Screenshot

<img width="1710" alt="Screenshot 2025-02-12 at 7 58 03 PM" src="https://github.com/user-attachments/assets/92fbb718-a93f-4e98-8364-75ad7de9e921" />










