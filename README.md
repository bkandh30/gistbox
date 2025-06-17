# Gistbox

A `Go` web application which is used to store code snippets similar to GitHub Gists and Pastebin.

## Table of Contents

- [About the Project](#about-the-project)
- [Project Structure](#project-structure)
- [Configuration](#configuration)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation & Setup](#installation--setup)
- [Testing](#testing)
- [License](#license)

## About the Project

"Gistbox" is a simple web application that allows users to create, view, and share text-based snippets. The primary goal of this project is to learn and apply best practices for building **clean, secure, and idiomatic web applications in Go**.

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

## Project Structure

The project follows the structure recommended in the "Let's Go" book:

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
    - Run the database migrations provided in the book to set up the necessary tables.

4.  **Run the application:**
    ```sh
    go run ./cmd/web
    ```
    The application will be accessible at `http://localhost:4000`.

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

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
