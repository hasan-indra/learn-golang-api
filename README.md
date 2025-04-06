## GO API
This project is for learn or create backend api using golang. I used go version 1.21.0

## What I do
```bash
# init golang project & docker-compose
go mod init github.com/hasan-indra/learn-golang-api

# dependencies
go get github.com/gin-gonic/gin
go get github.com/joho/godotenv

# create structure project
go-api
--cmd
----api 
------main.go   # main files
--internal
----config      # configuration file
----handlers    # request handlers
----middleware  # custom middleware
----route       # routes definition
--pkg           # reusable package
```

