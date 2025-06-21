# Gistbox

A `Go` web application which is used to store code snippets similar to GitHub Gists and Pastebin.

## Table of Contents

- [About the Project](#about-the-project)
- [Application Routes](#application-routes)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation & Setup](#installation--setup)
- [Setting up HTTPS for Local Development](#setting-up-https-for-local-development)
- [Building a Standalone Binary](#building-a-standalone-binary)
- [Testing](#testing)
- [License](#license)

## About the Project

"Gistbox" is a simple web application that allows users to create, view, and share text-based gists. The primary goal of this project is to learn and apply best practices for building **clean, secure, and idiomatic web applications in Go**.

Key features and concepts covered include:

- **Project Structure:** Organizing code in a clean, scalable, and maintainable way.
- **Advanced Routing:** Implementing advanced routing patterns with the standard `net/http` library and third-party routers like `httprouter`.
- **Custom Middleware:** Creating and using custom middleware for logging, error recovery, and rate-limiting.
- **Form Handling and Security:** Managing form submissions, implementing robust validation, and preventing common vulnerabilities like Cross-Site Scripting (XSS) and Cross-Site Request Forgery (CSRF).
- **Database Integration:** Interacting with a MySQL database to store and retrieve data.
- **Real-Time Communication:** Using WebSocket examples for real-time communication.
- **Graceful Operations:** Implementing graceful shutdown and using context for request handling.
- **Comprehensive Testing:** Writing unit and integration tests for handlers and other application components.
- **Session Management:** Managing user sessions and authentication.
- **HTML Templating:** Rendering dynamic HTML templates.

## Application Routes

The application provides the following endpoints:

| HTTP Method | URI Path         | Description                                       | Authentication |
| ----------- | ---------------- | ------------------------------------------------- | -------------- |
| `GET`       | `/`              | Display the home page with a list of recent gists | No             |
| `GET`       | `/gist/view/:id` | Display a specific gist                           | No             |
| `GET`       | `/gist/create`   | Display the form for creating a new gist          | Yes            |
| `POST`      | `/gist/create`   | Create a new gist                                 | Yes            |
| `GET`       | `/user/signup`   | Display the user signup form                      | No             |
| `POST`      | `/user/signup`   | Create a new user account                         | No             |
| `GET`       | `/user/login`    | Display the user login form                       | No             |
| `POST`      | `/user/login`    | Authenticate and log in a user                    | No             |
| `POST`      | `/user/logout`   | Log a user out                                    | Yes            |
| `GET`       | `/static/`       | Serve a specific static file                      | No             |

## Project Structure

The project follows the structure below:

```
.
├── cmd
│   └── web
│       ├── handlers.go
│       ├── helpers.go
│       ├── main.go
│       ├── middleware.go
│       ├── routes.go
│       └── templates.go
├── internal
│   ├── models
│   │   ├── errors.go
│   │   └── gists.go
│   └── validator
│       └── validator.go
├── ui
│   ├── html
│   │   ├── pages
│   │   │   ├── create.tmpl
│   │   │   ├── home.tmpl
│   │   │   └── view.tmpl
│   │   ├── partials
│   │   │   └── nav.tmpl
│   │   └── base.tmpl
│   └── static
│       ├── css
│       │   └── main.css
│       ├── img
│       │   ├── favicon.ico
│       │   └── logo.png
│       └── js
│           └── main.js
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## Configuration

The application can be configured through command-line flags.

- `-addr`: The network address to run the server on (e.g., `:4000`).
- `-dsn`: The Data Source Name for the MySQL database connection.

<!-- For convenience and security, you can store sensitive information like the DSN in a `.env` file in the project root. The application will read this file if it exists. -->

<!-- **Example `.env` file:**
```
SNIPPETBOX_DB_DSN="user:password@/snippetbox?parseTime=true"
``` -->

The application gives precedence to the command-line flag if both a flag and an environment variable are set.

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

- Go (version 1.24 or later recommended)
- MySQL

### Installation & Setup

1.  **Clone the repository:**

    ```sh
    git clone https://github.com/bkandh30/gistbox.git
    cd gistbox
    ```

2.  **Download Dependencies:** Download the modules required for the project.

    ```sh
    go mod download
    ```

3.  **Set up the Database:**

    - Create a MySQL database for the project.
    - Configure your database connection string as described in the [Configuration](#configuration) section.
    <!-- - Run the database migrations provided in the book to set up the necessary tables. -->

4.  **Run the application:**
    ```sh
    go run ./cmd/web
    ```
    The application will be accessible at `http://localhost:4000`.

## Setting up HTTPS for Local Development

To run the web server over HTTPS locally, you'll need to generate a self-signed TLS certificate.

1.  **Ensure you have a `tls` directory:** The project is structured to look for the certificate and key in a `/tls` directory at the root of the project. If it doesn't exist, create it:

    ```sh
    mkdir tls
    cd tls
    ```

2.  **Generate the certificate and private key:** Using the command-line, generate `cert.pem` and `key.pem` file inside the `tls` directory.

    ```sh
    go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
    ```

3.  **Run the application:** The application is configured to automatically look for `tls/cert.pem` and `tls/key.pem`. If found, it will serve traffic over HTTPS on the configured address.

    **Note:** Since the certificate is self-signed and not issued by a Trusted Certificate Authority, the browser will show a security warning. This is expected. You can safely proceed past the warning for local development purposes.

After creating the certificates, you are going to get "Warning: Potential Security Risk Ahead" screen. Click on "Advanced" and then click on "Accept the Risk and Continue" on Firefox based browsers or "Proceed to localhost (unsafe)" on Chromium based browsers.

## Building a Standalone Binary

You can build a standalone executable binary and run it in a clean directory. This is useful for deploying the application.

```sh
# Build the executable, placing it in the /tmp directory.
go build -o /tmp/web ./cmd/web/

# Copy the UI assets (HTML templates and static files) to the target directory.
# The binary needs these to render pages correctly.
cp -r ./ui /tmp/

# Copy the TLS certificates to the target directory.
cp -r ./tls /tmp/

# Change into the target directory.
cd /tmp/

# Run the binary.
# You can use flags like -addr and -dsn as needed.
./web
```

## Testing

To run the full suite of unit and integration tests, use the following command from the root of the project:

```sh
go test ./...
```

To run the tests in verbose mode and see the results for each individual test, use the `-v` flag:

```sh
go test -v ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
