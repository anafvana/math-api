# Simple Maths API

This API performs very simple maths operations between two numbers, namely:

- adding
- subtracting
- multiplying
- dividing

Send a POST request with a JSON body in the following format:

```
{
    "first_number": integer,
    "second_number": integer,
    "operation": string ("ADD", "SUBTRACT", "MULTIPLY" or "DIVIDE")
}
```

You will receive your result as:

```
{
    "result": integer
}
```

## Running

To run the application, you will need Go 1.18 or above.

```
$ go mod download
$ go run startup/mathapi/main.go
```

## Tests

To run unit tests (assuming you are in the project root folder):

```
$ cd api && go test
```
