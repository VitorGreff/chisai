## Chisai 小さい
Chisai is a minimalistic url shortner.

### Stack
- NextJs for the client side
    - Typescript
    - Tailwindcss
- Echo for the server side
    - Golang
    - Sqlc (go code generator) 
    - Postgres

### How to run the project
First things first, you'll need to create a file named '.env', at root directory, with the following format:
```.env
POSTGRES_USER=<your_user>
POSTGRES_PASSWORD=<your_password>
POSTGRES_PORT=<db_port>
POSTGRES_DBNAME=<db_name>
```
Then, you should be able to run the backend by going to the server directory and running:

```bash
go mod tidy # install go dependencies
go run main.go
```
After that's done, run the client at the root directory by typing:
```bash
npm install # install the local node dependencies
npm run dev
```
![chisai](https://github.com/VitorGreff/chisai/assets/73392743/08adbef0-f3a2-4f01-94ab-8368fa08f99a)
