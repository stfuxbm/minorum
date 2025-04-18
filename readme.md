# Minorum

Minorum is a RESTful API providing access to inspirational quotes and biographical information for Catholic and Orthodox saints.

## Motivation

The motivation behind **Minorum** is to make the profound wisdom and inspiring lives of Catholic and Orthodox saints readily accessible to a wider audience. In a world often filled with noise and distraction, the timeless teachings of these holy figures offer guidance, comfort, and a path towards spiritual growth and inner peace.

## Features

- Inspirational quotes from Catholic and Orthodox saints.
- Biographies of saints.
- Patronages associated with each saint.
- Categorized quotes (e.g., faith, hope, prayer).
- Information about Catholic and Orthodox religious orders.

## Installation Guide

1. Clone the repository:

    ```bash
    git clone https://github.com/stfuxbm/minorum.git
    ```

2. Navigate to the project directory:

    ```bash
    cd minorum
    ```

3. Install Go dependencies:

    ```bash
    go mod download
    ```

4. Run the application locally with:

    ```bash
    make dev
    ```

    This will start the server and allow you to interact with the API locally.

## API Endpoints

### Quote

- `POST /api/v1/quotes`  
  Add a new quote

- `GET /api/v1/random-quotes`  
  Retrieve a random quote.

- `GET /api/v1/quotes/search?name={query}`  
  Retrieve quotes by saint name or category.

### Ordo

- `POST /api/v1/ordos`  
  Add a new ordo

- `GET /api/v1/ordos/search?order={query}`  
  Retrieve ordo information by order name or nickname.

### Saint

- `POST /api/v1/saints`  
  Add a new saint

- `GET /api/v1/saints/search?order={query}`  
  Retrieve saint information by order name or nickname.

## Contribution Guidelines

Contributions are welcome. Fork the repository, create a branch, and submit a pull request.

## License

The content provided in this application, including quotes from saints and holy figures, is available for free use under the MIT License. You are free to use, modify, and share the content as long as appropriate credit is given to the original sources.
