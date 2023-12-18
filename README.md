##Garage Management


#Setup and installations
Pre-requisites:

docker
golang

Note:Make Sure that port 3306 is Not In Use


Setting Up:

1)Database Setup

Open Teminal and run command-
`docker run --name gofr-mysql -e MYSQL_ROOT_PASSWORD=root123 -e MYSQL_DATABASE=garage -p 3306:3306 -d mysql:8.0.30`

This command will download the mysql docker image and will create the database

Then run

`docker exec -it gofr-mysql mysql -uroot -proot123 garage -e "CREATE TABLE person (name VARCHAR(255) NOT NULL, contact VARCHAR(10) PRIMARY KEY, address VARCHAR(255)); CREATE TABLE car (number VARCHAR(255) PRIMARY KEY, name VARCHAR(255) NOT NULL, color VARCHAR(20), person VARCHAR(10), FOREIGN KEY (person) REFERENCES person(contact)); CREATE TABLE job (id INT AUTO_INCREMENT PRIMARY KEY, car VARCHAR(255), fault VARCHAR(255), status VARCHAR(50), FOREIGN KEY (car) REFERENCES car(number));"`

The above command will create the table in the mysql-image

2)Open  a new terminal and run
    `go get gofr.dev`
	`go get github.com/go-sql-driver/mysql`


Then clone the git repo and run:
`go mod tidy`
`go run main.go`

Program will run on localhost:9000

## Database Schema
 


##Files
#constants.go - contains predefined constants
#models.go - contains structure for parsing of data 
#routers
	car_route.go
		/car:{person}:{number}/
			This route has a GET method and fetches data 	corresponding phone number and carnumber
		/car/
			This route has a GET method and fetches all records 
			of car table
		/car/{carNum}
			This is a DELETE method API for removing records with carnumber=carNum
  
garage_route.go
    /garage 
      This route fetches and display all the data related to car in garage
    /garage/:{status}
      This route checks for the cars with some specific status and shows corresponding results
job_Summary.go
    /job_schedule/all
      This route shows all the job assigned in the garage
    /job_schedule:{status}
      This route checks for the jobs with some specific status and shows corresponding results ,if not found it displays all records
    /job_schedule/
      This is a POST request that inserts the data into the database by using a body.
        format for body: JSON format with properly named fields (Field names are mentioned in configs/db_schema)
      /job_schedule/{id}
        This is a DELETE request to delete the records with specific job IDs
      /job_schedule/{id}:{status}
         This is a PUT method route for updating the status of the car according to their job id
  person_route.go
    /person
      This request shows the data of all the car owners like name ,contact and address
    /person/{phone}:{data}
      This is a bit different route.Here we check that if there is an owner with a contact number=phone and  also has some work being performed on its car the this api will fetch that record and display the data
    /person/{phone}
       This route deletes the owner with contact number = phone
  

    


