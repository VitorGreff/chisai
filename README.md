## Chisai
Chisai is a minimalist url shortner.

### Stack
- NextJs for the client side
    - Typescript
    - Tailwindcss
- Echo for the server side
    - Golang
    - Sqlc (go code generator) 
    - Postgres

### How to run the project
First, you'll need to run the backend by navigating to the server directory, then run the commands below:
```bash
go mod tidy # install go dependencies
go run main.go
```
After that's done, run the client at the root directory by typing:
```bash
npm instal # install the local node dependencies
npm run dev
```
