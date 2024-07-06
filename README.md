# QRThrough: Back-end

**<p style ="text-align: center;">A back-end server for QRThrough system to comunicate with the database, device and the front-end.</p>**

# Description

The server is responsible for handling the requests from the front-end and the device, and communicate with [LINE Messaging API](https://developers.line.biz/en/services/messaging-api/) and the database. The server is built with [Fiber](https://docs.gofiber.io/) framework, a web framework for Go.

## Getting Started

To develop and run this project, you need to have Go installed on your machine. You can download Go from [here](https://go.dev/dl/).

### Prerequisites

- Go 1.16 or higher
- Docker
- Make

### Installation

1. Install the dependencies by running the following command:

   ```bash
    go mod download
   ```

2. Create a `.env.dev` for development environment file in the config directory. The file should contain the following environment variables:

   ```
   # Application Config
   APP_NAME="qr-through-api" // Application name
   PORT="9000" // Port number for the server

   # Line Config

   LINE_BOT_CHANNEL_SECRET="" // Line bot channel secret
   LINE_BOT_CHANNEL_TOKEN="" // Line bot channel token
   LINE_CHANNEL_ID="" // Line channel id for user application
   LINE_DASHBOARD_CHANNEL_ID="" // Line channel id for dashboard application

   # Postgres Config

   POSTGRES_HOST="" // Postgres host
   POSTGRES_NAME="" // Postgres database name
   POSTGRES_USER="" // Postgres user
   POSTGRES_PORT="" // Postgres port
   POSTGRES_PASSWORD="" // Postgres password

   #ThaiBulksms
   THAIBULKSMS_SECRET ="" // ThaiBulkSMS secret
   THAIBULKSMS_KEY ="" // ThaiBulkSMS key
   ```

3. Create a container for the database in the `qrthrough-docker` repository. You can find the repository [here](https://github.com/QRThrough/qrthrough-docker). Follow the instructions in the repository to create the database container.

4. Conect to database by PGAdmin or any database management tool and create a database with the name `qrthrough`.

5. Migration the database by running the following command:

   ```bash
    make dbmigrate
   ```

6. Create a container for message broker in the `qrthrough-docker` repository. You can find the repository [here](https://github.com/QRThrough/qrthrough-docker). Follow the instructions in the repository to create the message broker container.

7. Run the server by running the following command:

   ```bash
    go run main.go
   ```

## Project Architecture Overview

The project is built with the `Hexagonal Architecture` pattern. The `Hexagonal Architecture` is a software architecture pattern that allows the application to be designed in a way that it is independent of the external services, such as the database, the message broker, and the web framework.

<img src="https://miro.medium.com/v2/resize:fit:2000/1*mGLO5IfhJv4o0NYOAZI60A.png" alt="Hexagonal Architecture" style="width:800px;height:auto;">

The `Hexagonal Architecture` pattern consists of the following components:

- **Application Core**: The application core contains the business logic of the application. It consists of the following components:

  - **Domain**: The domain contains the business rules and the business entities of the application.
  - **Service**: The service contains the business logic of the application.
  - **Port**: The port contains the interfaces that define the interactions between the application core and the external services.

- **Adapter**: The adapter contains the implementation of the interfaces defined in the port. It consists of the following components:
  - **Handler**: The handler contains the implementation of the API endpoints.
  - **Repo**: The repo contains the implementation of the repository interfaces.

## Project Structure

The project structure implements the `Hexagonal Architecture`. The project is divided into the following directories:

```
qrthrough-api
├── api // Contains the API Endpoints definition
├── cmd // Contains go runable command files
├── config // Contains the configuration files and configuration code
├── db // Contains the database scripts (SQL)
├── infrastructure // Contains the infrastructure code
├── internal // Contains the internal code
│   ├── adapter // Contains the adapter code
│   │   ├── handler // Contains the handler code
│   │   ├── repo // Contains the repository code
│   ├── core // Contains the core code
│   │   ├── domain // Contains the interface of driving port (usecase)
│   │   ├── dto // Contains the data transfer object code
│   │   ├── model // Contains the model or entity code
│   │   ├── port // Contains the interface of driven port (repository)
│   │   ├── service // Contains the service or business logic code
├── pkg // Contains the package code
│   ├── errors // Contains the error handling code
│   ├── middleware // Contains the middleware code
│   ├── rest // Contains the REST handling code
│   ├── utils // Contains the utility code
├── .dockerfile
├── main.go
├── makefile
├── README.md
```

## API Endpoints

The server provides the following API endpoints:

| Method | Endpoint                         | Description                                                  |
| ------ | -------------------------------- | ------------------------------------------------------------ |
| GET    | /                                | Check the health of the server                               |
| GET    | /api/v1/dashboard/signin         | Sign in to the dashboard application                         |
| GET    | /api/v1/dashboard/configurations | Get the configurations for the dashboard application         |
| POST   | /api/v1/dashboard/configurations | Update the configurations for the dashboard application      |
| GET    | /api/v1/dashboard/moderators     | Get the moderators for the dashboard application             |
| PUT    | /api/v1/dashboard/moderators/:id | Update the moderators for the dashboard application by id    |
| DELETE | /api/v1/dashboard/moderators/:id | Delete the moderators for the dashboard application by id    |
| GET    | /api/v1/dashboard/logs           | Get the usage transaction logs for the dashboard application |
| GET    | /api/v1/dashboard/users          | Get the users for the dashboard application                  |
| PUT    | /api/v1/dashboard/users/:id      | Update the users for the dashboard application by id         |
| DELETE | /api/v1/dashboard/users/:id      | Delete the users for the dashboard application by id         |
| GET    | /api/v1/device/scanner/:token    | Verify the token for qr code scanner                         |
| GET    | /api/v1/liff/alumni/:id          | Get the alumni information by id                             |
| POST   | /api/v1/liff/user                | Create the user for the liff application                     |
| POST   | /api/v1/liff/otp/request         | Request the OTP for the liff application                     |
| PUT    | /api/v1/liff/otp/verify          | Verify the OTP for the liff application                      |
| POST   | /api/v1/line/webhook             | Handle the webhook from the LINE Messaging API               |

## Deployment

To deploy the server, you need to have a server with Docker installed. You can follow the instructions in the installation section to create a container for the database and the message broker.

1. Create a `.env.prod` file in the config directory. The file should contain the following environment variables:

```

# Application Config

APP_NAME="qr-through-api" // Application name
PORT="9000" // Port number for the server

# Line Config

LINE_BOT_CHANNEL_SECRET="" // Line bot channel secret
LINE_BOT_CHANNEL_TOKEN="" // Line bot channel token
LINE_CHANNEL_ID="" // Line channel id for user application
LINE_DASHBOARD_CHANNEL_ID="" // Line channel id for dashboard application

# Postgres Config

POSTGRES_HOST="" // Postgres host
POSTGRES_NAME="" // Postgres database name
POSTGRES_USER="" // Postgres user
POSTGRES_PORT="" // Postgres port
POSTGRES_PASSWORD="" // Postgres password

#ThaiBulksms
THAIBULKSMS_SECRET ="" // ThaiBulkSMS secret
THAIBULKSMS_KEY ="" // ThaiBulkSMS key

```

2. Configure the `makefile` file to deploy the server to the server. You can find the `makefile` file in the root directory of the project. You need to change the following variables:

- First, change the `MODE` variable to `prod`.
- Second, change the `REMOTE_HOST` variable to the server's IP address or hostname.

3. Run the following command to deploy the server to the server with two situations:

- Initial deployment:

  ```bash
  make init-deploy
  ```

- Update deployment:

  ```bash
  make build-deploy
  ```

4. During the deployment, If you not have the `ssh` key to connect to the server, you need to enter the password to connect to the server.

## Authors

- [JM jirapat](https://github.com/JMjirapat) - Developer
