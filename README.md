# Igbo API Go Client

This is a Go client for the Igbo API, which allows you to search for and retrieve information about Igbo words.

## Installation
To use this client, you need to import `igboapi-go` package in your Go project.
```
import "github.com/bontusss/igboapi-go"
```

## Usage

First, you need to create a new client with your API key:

```
key := "your-api-key"
client := igboapi.NewClient(key)
```

### Searching for a Term

To search for a term, use the `Search` method of the `TermService`:

```
t, err := client.TermService.Search(context.Background(), &igboapi.TermParams{Keyword: "bia"})
if err != nil {
	log.Fatal(err)
	return
}

fmt.Println(len(t))
```
This will return a list of terms that match the keyword "bia"

### Finding a Specific Term

To find a specific term by its ID, use the `Find` method of the `TermService`:

```
t, err := client.TermService.Find(context.Background(), "5f90c35e49f7e863e92b7269")
if err != nil {
	log.Fatal(err)
	return
}

fmt.Println(t.Word)
```
This will return the Term with the ID "5f90c35e49f7e863e92b7269"

## Error Handling

If an error occurs during a request, the error will be returned by the method. You should always check for errors after making a request.

## Contributing

Contributions are welcome. Please submit a pull request or create an issue to discuss the changes you want to make.

## License

This project is licensed under the MIT License.
