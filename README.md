# mysmall-wallet
My small wallet is a simple project initially written for a challenge. This project, contains two main APIs: 

● get-balance:
This API should return a JSON to show the current balance of a user. The parameter which is
needed for this API is user_id and the output should be like the below sample:
```
Input : user_id int
Output : {"balance":4000}
```
● add-money:
This API should add money to the wallet of a user and at the end return the transaction reference
number. The parameter which is needed for this API is user_id and amount and the output should
be like the below sample:
```
Input: user_id int amount int (this parameter can be negative)
Output: {"reference_id":12312312312}
```
There is also an scheduling feature implemented which calculates total amount of transaction done in a day and print it out in the terminal.

## How to run? 
This application is going to be dockerized soon but till then, if you have go 1.18 installed, simply clone the project, cd to the project dir
and type out these following instructions:
```Batchfile
go mod tidy
go mod vendor
```
after this, you will have all dependancies installed. Now, for running the application simply type out the following command
```Batchfile
go run cmd/main.go
```

You also need an sql database to run migrations on. You can set up your database and schema preferences in `envs/evs.json`
