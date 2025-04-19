# Minorum

Minorum is a RESTful API designed to provide easy access to inspirational quotes and biographical information about Catholic and Orthodox saints. It offers a spiritual resource that encourages personal growth, reflection, and inspiration.

## Motivation

The primary motivation behind **Minorum** is to make the timeless wisdom and inspiring lives of Catholic and Orthodox saints more accessible to a global audience. In a world that often feels chaotic and uncertain, the teachings of these saints provide comfort, direction, and a pathway toward spiritual growth and inner peace.

## Features

- A vast collection of inspirational quotes from Catholic and Orthodox saints.
- Detailed biographies of saints.
- Patronages associated with each saint.
- Categorized quotes based on themes (e.g., faith, hope, prayer, love).
- Information on Catholic and Orthodox religious orders.

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

4. Run the application locally:

    ```bash
    make dev
    ```

    This command will start the development server, and you can begin interacting with the API locally.

## Host

 ```bash
   [https://minorum.onrender.com](https://minorum.onrender.com)
  ```

## API Endpoints

### Quotes

- `POST /api/v1/quotes`  
  Add a new quote.

- `GET /api/v1/random-quotes`  
  Retrieve a random quote.

- `GET /api/v1/quotes/search?name={query}`  
  Search for quotes by saint name or category.

### Ordos (Religious Orders)

- `POST /api/v1/ordos`  
  Add a new religious order.

- `GET /api/v1/ordos/search?order={query}`  
  Retrieve information about religious orders by name or nickname.

### Saints

- `POST /api/v1/saints`  
  Add a new saint.

- `GET /api/v1/saints/search?order={query}`  
  Retrieve saint information by order name or nickname.

## Contribution Guidelines

Contributions are welcome. Fork the repository, create a branch, and submit a pull request.

## License

The content provided in this application, including quotes and biographical information, is made freely available under the MIT License. You are free to use, modify, and distribute the content, provided that appropriate credit is given to the original sources.
