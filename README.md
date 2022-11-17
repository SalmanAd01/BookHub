# BookHub
Link : https://bookkhub.herokuapp.com/
#### Get Started
copy the .env.exmaple to .env

```
  cp .\.env.example .\.env
```

To Install dependencies

```
    go mod download
```

To run 

```
    air
```

#### Rus using docker

copy the .env.exmaple to .env

```
  cp .\.env.example .\.env
```

Build and run the image
```
   docker compose -f .\docker-compose-dev.yml up
```

### To Run Lint
```
  golangci-lint run
```
