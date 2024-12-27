# Crankset

This is a simple web application built using the Gin framework in Go. It demonstrates how to set up a basic RESTful API with routing, controllers, and models.

## Project Structure

```
crankset
├── src
│   ├── main.go          # Entry point of the application
│   ├── controllers      # Contains the controller logic
│   │   └── controller.go
│   ├── routes           # Defines the application routes
│   │   └── routes.go
│   └── models           # Contains data structures and methods
│       └── model.go
├── go.mod               # Module definition and dependencies
└── README.md            # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone https://github.com/wiscotrashpanda/crankset.git
   cd crankset
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run src/main.go
   ```

## Usage

Once the application is running, you can access the API endpoints:

- **GET /**: Returns a welcome message.
- **POST /data**: Accepts data and processes it.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.