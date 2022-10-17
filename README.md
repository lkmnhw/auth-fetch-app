# auth-fetch-app

A simple app for authentication and fetching data

## Run Locally

Clone the project

```bash
  git clone https://github.com/lkmnhw/auth-fetch-app.git
```

Run with docker

- build docker compose
```bash
  docker compose build
```
- run auth-app
```bash
  docker compose run -p 3000:3000 auth-app
```
- run fetch-app
```bash
  docker compose run -p 4000:4000 fetch-app
```

Run without docker

- go to each directory
```bash
  cd auth-app
  cd fetch-app
```
- run go
```bash
  go mod tidy
  go run main.go
```

## Author
- [@lkmnhw](https://github.com/lkmnhw)