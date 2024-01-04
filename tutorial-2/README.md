# MyWebApp with Golang GIN JWT

Gin JWT Auth is a Golang web application featuring JWT authentication with the GIN framework. This project provides a quick start for building secure web applications.

For detailed instructions and explanations, please refer to [my guide](https://www.fastdt.app/2023/12/20/creating-a-gin-project-with-basic-jwt-authentication-a-step-by-step-guide/).

## Installation and Setup

### Prerequisites

Make sure you have the following tools installed:

- [Go](https://golang.org/dl/)
- [MySQL](https://dev.mysql.com/downloads/)

### Quick Start

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/joequah1/go-gin-jwt-auth.git
   cd go-gin-jwt-auth
   ```

2. **Set Up Environment Variables:**

    Create a .env file in the project root with your database configurations and secret key:

	```bash
	DB_DRIVER=mysql
	DB_HOST=localhost
	DB_USER=root
	DB_PASSWORD=your_password
	DB_NAME=mywebapp
	DB_PORT=3306
	SECRET_KEY=your_secret_string
	TOKEN_LIFESPAN=1
	```

3. **Install Dependencies:**

	```bash
	go get -u github.com/gin-gonic/gin
	go get golang.org/x/crypto/bcrypt
	go get github.com/dgrijalva/jwt-go
	go get github.com/jinzhu/gorm
	go get -u github.com/jinzhu/gorm/dialects/mysql
	go get github.com/joho/godotenv
	```

4. **Run the Application:**

	```bash
	go run main.go
	```

## Run and Test

1. **Register a User:**

	```bash
	curl -X POST -H "Content-Type: application/json" -d '{"username":"yourusername", "password":"yourpassword", "name":"Your Name", "email":"your.email@example.com"}' http://localhost:8080/api/auth/register
	```

2. **Login to Get a Token:**

    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"username":"yourusername", "password":"yourpassword", "name": "name", "email": "email"}' http://localhost:8080/api/auth/login
    ```

    You will receive a token in the response.

3. **Access the Profile Using the Token:**

    ```bash
    export TOKEN="your_received_token"
    curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/profile
    ```

For more [my guide](https://www.fastdt.app/2023/12/20/creating-a-gin-project-with-basic-jwt-authentication-a-step-by-step-guide/), check out my guide where I explain the project in depth.
