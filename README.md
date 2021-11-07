# go-url-shortener
This project describes how to create url shortner of long url.

This project shows how to write a url shortener in Go programming language
using Redis as store mechanism for super fast data retrieval in the implementation.

prerequisite
- Golang installed on machine
- Redis Installed on machine
- Before running main.go file, please configure store_service.go file according to the redis
  setup done in your own machine.

You can just clone the project and run main.go file to see in effect.

After successful running of project, you can check functionality using below curl call.

curl --location --request POST 'http://localhost:9808/create-short-url' \
--header 'Content-Type: application/json' \
--data-raw '{
    "long_url": "https://www.google.com"
}'
