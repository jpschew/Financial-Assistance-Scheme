# Financial Assistance Schemes

## Instructions to set up on local development environment
1. Download and install [Go](https://go.dev/doc/install) and [Goland IDE](https://www.jetbrains.com/go/download/#section=mac) based on your operating system
2. Download and install [mySQL](https://dev.mysql.com/downloads/mysql/) and [Redis](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/install-redis-on-mac-os/)
3. Run all the sql queries in *fas.sql* to create the database and tables needed for the application in **mySQL**
4. Download and install the dependencies

EITHER
```azure
go mod tidy
```
OR
```azure
go get github.com/kataras/iris/v12
go get github.com/kataras/iris/v12/middleware/accesslog
go get github.com/go-playground/validator/v10 
go get github.com/dgrijalva/jwt-go 
go get github.com/spf13/viper 
go get github.com/go-redis/redis/v8
go get gorm.io/gorm 
go get gorm.io/driver/mysql
go get golang.org/x/crypto/bcrypt

```
5. Replace **database password** in the *config.yaml* with **your password** for mySQL database connection
6. Run the command below in **Goland IDE Terminal** to start the server
```azure
make run_app
```
7. Import **Financial Assistance Scheme.postman_collection.json** in the postman folder into Postman

## Database Design
1. Applicants Table

| Column            | Data Type                                  | Description                                                                                                                                                                                          | 
|-------------------|--------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                | Integer<br/>Auto Increment<br/>Primary Key | id of applicant                                                                                                                                                                                       |
| first_name        | String                                     | first name of applicant                                                                                                                                                                              |
| last_name         | String                                     | last name of applicatnt                                                                                                                                                                              |
| nric              | String<br/>Unique Key                      | nric of applicant                                                                                                                                                                                    |
| employment_status | Integer                                    | employment status of applicatnt<br/>1 - employed<br/>2 - unemployed                                                                                                                                  |
| martial_status    | Integer                                    | martial status of applicant<br/>1 - single<br/>2 - married<br/>3 - divorced<br/>4 - widowed                                                                                                          |
| sex               | Integer                                    | sex of applicatnt<br/>0 - male<br/>1 - female                                                                                                                                                        |
| date_of_birth     | String                                     | date of birth of applicant<br/>YYYY-MM-DD                                                                                                                                                            |
| househhold        | JSON_ARRAY                                 | household members of applicant<br/><br/>one example shown below<br/>[{"sex": 0, "nric": "T2082722F", "relation": "son", "last_name": "Lee", "first_name": "Vincent", "date_of_birth": "2020-10-09"}] |
| create_time       | UNIX_TIMESTAMP                             | time the record is created                                                                                                                                                                           |
| update_time       | UNIX_TIMESTAMP                             | time the record is updated                                                                                                                                                                           |

2. Schemes Table

| Column            | Data Type                                  | Description                                                                                                                                                  | 
|-------------------|--------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                | Integer<br/>Auto Increment<br/>Primary Key | id of scheme                                                                                                                                                 |
| name              | String<br/>Unique Key                      | name of scheme                                                                                                                                               |
| description       | String                                     | description of scheme                                                                                                                                        |
| employment_status | Integer                                    | employment status criteria<br/>0 - all<br/>1 - employed<br/>2 - unemployed                                                                                   |
| martial_status    | Integer                                    | martial status criteria<br/>0 - all<br/>1 - single<br/>2 - married<br/>3 - divorced<br/>4 - widowed                                                          |
| children_status   | Integer                                    | children status criteria<br/>0 - no schooling children<br/>1 - has schooling children<br/><br/>schooling children - refers to primary to secondary (7 to 16) |
| benefits          | JSON_ARRAY                                 | benefits of scheme<br/><br/>one example shown below<br/>[{"name": "SkillsFuture credits", "amount": 500.0}]                                                  |
| create_time       | UNIX_TIMESTAMP                             | time the record is created                                                                                                                                   |
| update_time       | UNIX_TIMESTAMP                             | time the record is updated                                                                                                                                   |

3. Applications Table

| Column       | Data Type                                  | Description                                                                               | 
|--------------|--------------------------------------------|-------------------------------------------------------------------------------------------|
| id           | Integer<br/>Auto Increment<br/>Primary Key | id of application                                                                         |
| applicant_id | Integer                                    | applicant id mapped to id of applicant table                                              |
| scheme_id    | Integer                                    | scheme id mapped to id of scheme table                                                    |
| status       | Integer                                    | status of application<br/>0 - pending<br/>1 - approved<br/>2 - rejected<br/>4 - withdrawn |
| create_time  | UNIX_TIMESTAMP                             | time the record is created                                                                |
| update_time  | UNIX_TIMESTAMP                             | time the record is updated                                                                |

4. Admins Table

| Column   | Data Type                                  | Description                      | 
|----------|--------------------------------------------|----------------------------------|
| id       | Integer<br/>Auto Increment<br/>Primary Key | id of admin account              |
| name     | String                                     | name of person for admin account |
| username | String<br/>Unique Key                      | username of admin account        |
| password | String                                     | password of admin account        |


## API Endpoints

The API Endpoints below requires jwt token as *Authorization* key in the **Headers** section of the request 

| Method | Path                                    | Purpose                                                                                        |
|--------|-----------------------------------------|------------------------------------------------------------------------------------------------|
| GET    | /api/applicants                         | Get all applicants                                                                             |
| POST   | /api/applicants                         | Create a new applicant                                                                         |
| GET    | /api/schemes                            | Get all schemes                                                                                |
| GET    | /api/schemes/eligible?applicant_id={id} | Get all eligible schemes for an applicant (represented by applicant_id query string parameter) |
| POST   | /api/schemes                            | Create a new scheme                                                                            |
| GET    | /api/applications                       | Get all applications                                                                           |
| POST   | /api/applications                       | Create an application                                                                          |
| POST   | /api/update_application                 | Update an application status                                                                   |

The jwt token can be generated by logging in as an admin using the API Endpoint below

| Method | Path                                    | Purpose                                                                                        |
|--------|-----------------------------------------|------------------------------------------------------------------------------------------------|
| POST   | /api/login                              | Login and get jwt token for admin                                                              |
| POST   | /api/logout                             | Logout and delete jwt token for admin                                                          |
| POST   | /api/create_admin                       | Create a new admin                                                                             |


Refer to the postman file **Financial Assistance Scheme.postman_collection.json** for the details of API endpoints in the postman folder