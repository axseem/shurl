![logo](./assets/logo.png)

## What is shurl?

"shurl" is an elegant backend solution for a customizable URL shortener application. It allows users to generate shortened URLs while providing the option to customize the endpoint, specify its length, and set an expiry time. The application is written in GO and utilizes Redis as a database.

## Showcase

Make a shurl to "exampl.com"
```js
// POST "localhost:3000"
{ 
    "url": "example.com" 
}

// Response
{
    "url": "http://example.com",
    "shurl": "localhost:3000/xQ9DC",
    "length": 5,
    "expiry": 0
}
```
Set custom endpoint
```js
// POST "localhost:3000"
{
    "url": "example.com",
    "shurl": "example"
}

// Response
{
    "url": "http://example.com",
    "shurl": "localhost:3000/example",
    "length": 7,
    "expiry": 0
}
```
Specifie length of endpoint and set expiry time (in hours)
```js
// POST "localhost:3000"
{
    "url": "example.com",
    "length": 8,
    "expiry": 24
}

// Response
{
    "url": "http://example.com",
    "shurl": "localhost:3000/FpQpLdFN",
    "length": 8,
    "expiry": 24
}

```
Visiting `localhost:3000/xQ9DC`, `localhost:3000/example` or `localhost:3000/FpQpLdFN` will redirect to [exmaple.com](http://example.com)

## Installation

To build "shurl" and launch server, follow these steps:

1. Clone and build "shurl":
```
git clone github.com/axseem/shurl
cd shurl
go build
```
2. Set environment variables in ".env" file:
```sh
DB_ADDR="localhost:6379"
DB_PASS=""
PORT=":3000"
DOMAIN="localhost${PORT}"
```
3. Execute the built application to launch the server:
```sh
./shurl
```

## Usage

Here are some enpoints end its methods:
- `/` `POST <shurl>` - Create a shurl
- `/:id` `GET` - Redirect to full URL

## License

This project is licensed under the [MIT License](LICENSE.md).