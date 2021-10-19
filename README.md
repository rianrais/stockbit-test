# stockbit-test
This is a public repository for the test question from stockbit.

Some of the module is using go.mod so ensure that your Golang version is able to support this.

## Usage

First, ensure that you are in the correct folder (which is separated by which question it denotes to) and then simply run the go program.

Example:

```bash
cd question3
go run answer.go
```

For the REST API question, do the same as above by entering folder `/question2` but instead run:

```bash
go get
go run main.go
```

To get the dependency. And then use API platform (such as Postman) to hit the endpoint with the required query parameters.

The endpoint would be: `localhost:3070/omdb` with GET method.

## Important Notes

Since the OMDB API key given to me via email can not be used any further, I've taken steps to get my own free API key for OMDB API.

## License
[MIT](https://choosealicense.com/licenses/mit/)
