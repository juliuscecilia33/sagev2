SageV2 is a Go-based application that uses Docker, GORM, and PostgreSQL for building a scalable API. This guide provides instructions for setting up the development environment using Docker Compose.

<img width="1333" alt="image" src="https://github.com/user-attachments/assets/9dc4d943-713f-4f4d-b727-5ca9b8a76c72">


# Prerequisites

### Make sure you have the following installed:

Docker: Install Docker (https://docs.docker.com/engine/install/)

Getting Started

1. Clone the Repository

   - Clone the repository to your local machine:
   - `git clone https://github.com/juliuscecilia33/sagev2.git`

2. Set Up Environment Variables

   - Create a .env file in the root directory of the project with the following contents:
   - Replace the placeholders with your actual database credentials.
     - `DB_HOST=localhost`
     - `DB_NAME=your_database_name`
     - `DB_USER=your_database_user`
     - `DB_PASSWORD=your_database_password`
     - `DB_SSLMODE=disable`
     - `JWT_SECRET=your_jwt_secret`

3. Docker Compose
   The docker-compose.yaml file is already configured to build and run the application and database services.

- Use the Makefile to start the application. Run:

  `make start`

4. Access the Application

   - After the services have started, you can access the application on http://localhost:8082.

5. Database Access
   - The PostgreSQL database is accessible on port 5433 (mapped from the container's internal port 5432). You can connect to it using any PostgreSQL client with the credentials provided in the .env file.
