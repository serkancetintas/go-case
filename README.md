# go-case

Yemeksepeti case

## Up

- go to directory **/cmd**
- go run main.go <br/>
  locate => http://localhost:8080/swagger

## Api

**Set key**
```
curl --location --request POST 'localhost:8080/api/set' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "delivery",
    "value": "hero"
}'
```

**Get key**
```
curl --location --request GET 'localhost:8080/api/get?key=delivery'
```

**Flush**
```
curl --location --request GET 'localhost:8080/api/flush'
```

## Docker

- docker build -t go-case -f docker/Dockerfile .
- docker run -it --rm -p 8080:8080 go-case

