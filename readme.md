# Saints Quotes API

A RESTful API that provides inspirational quotes from Catholic and Orthodox saints, along with their biographical details, spiritual attributes, and patronages.

## Description

This API serves inspirational quotes from Catholic and Orthodox saints, enriched with biographical information, spiritual attributes, patronages, and categorized quote themes (e.g., faith, hope, prayer).

## Features

- Inspirational quotes from Catholic and Orthodox saints.
- Biographies of saints.
- Patronages associated with each saint.
- Categorized quotes.

## Motivation

To inspire reflection, spiritual growth, and inner peace by preserving and sharing the timeless wisdom of saints.

## Installation Guide

1. Clone the repository:

    ```bash
    git clone https://github.com/stfuxbm/quote-saints.git
    ```

2. Navigate to the project directory:

    ```bash
    cd quote-saints
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

**Order Endpoints**

- `GET /api/v1/search?order={query}`: Retrieve order information by name.
- `POST /api/v1/ordos`: Add ordos.

**Quote Endpoints**

- `POST /api/v1/quotes`: Add quotes.
- `GET /api/v1/random-quotes`: Retrieve a random quote.
- `GET /api/v1/quotes/search?name={query}`: Retrieve quotes by author name.
- `GET /api/v1/quotes/category?category={query}`: Retrieve quotes by category.

## Contribution Guidelines

Contributions are welcome. Fork the repository, create a branch, and submit a pull request.

## License

The content provided in this application, including quotes from saints and holy figures, is available for free use under the MIT License. You are free to use, modify, and share the content as long as appropriate credit is given to the original sources.
