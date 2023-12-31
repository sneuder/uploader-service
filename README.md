# Uploader Service

This project is an uploader service built using Go programming language, Echo web framework, GORM as the ORM (Object-Relational Mapping) tool, and JWT for authentication. The service allows registered users to upload files and store associated information in a database.

## Feature
- **User Registration**: Users can register to the service to gain access to file upload functionalities.
- **File Upload**: Registered users can upload files to the service.
- **Data Storage**: Along with files, metadata associated with the uploaded files are stored in the database.
- **Authentication**: JWT (JSON Web Tokens) are used for authentication, ensuring secure access to the service.

## Technologies Used
- **Go**: The programming language used to develop the service (1.21.1).
- **Echo**: A web framework used to create the API endpoints and handle HTTP requests.
- **GORM**: An ORM library used for database operations, simplifying interactions with the database.
- **JWT**: JSON Web Tokens used for user authentication and secure access.

## Set Up Project

1. **Download**:
`git clone git@github.com:sneuder/uploader-service.git`

3. **Install Dependencies**:
`go download mod`

3. **Set Database**:
You can use the default LiteSQL Setting or intall the driver based on your preferences

4. **Run Project**:
`go run cmd/uploader-service/main.go`

5. **Access to project**:
`http://localhost:3000`
