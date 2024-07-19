# Go Movie CRUD

This project implements a movie CRUD application using Go.

## Features

- **Create**: Add new movies.
- **Read**: List all movies.
- **Update**: Modify movie details.
- **Delete**: Remove movies.
- **Get Movie**: Retrieve details of a specific movie.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ChandanJnv/go-movie-crud.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go-movie-crud
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

## Usage

Run the application:
```bash
go run main.go
```

Access the API at `http://localhost:8080`.
Note: port can be chagned in main.go

## Endpoints

- **List all movies.**
    - `GET /movies`

- **Get a specific movie**
    - `GET /movies/{id}`

- **Add a new movie**
    - `POST /movies`
    - **Request Body:**
        ```json
        {
        "isbn": "564648",
        "title": "Movie Three",
        "director": {
            "firstname": "Joe",
            "lastname": "Doe"
            }
        }
        ```

- **Update a movie**
    - `PUT /movies/{id}`
    - **Request Body:**
        ```json
        {
        "isbn": "564648",
        "title": "Movie Three",
        "director": {
            "firstname": "Joe",
            "lastname": "Alpha"
            }
        }
        ```

- **Delete a movie**
    - `DELETE /movies/{id}`


## Dependencies

- Gorilla Mux: For routing.

## License

This project is licensed under the MIT License.
