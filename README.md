# Socket Craft

## Development

This project includes a tool to manage building and deployment
to facilitate local testing

To setup your environment for development of this project run

```bash
$ npm install
$ go build -o tool ./_tool
```

Run the tool with `./tool` to recieve help information about available commands

For a quickstart building everything and standing up the full stack run

```bash
$ ./tool init
```

The site should now be available at [https://localhost](https://localhost) provided your port `443`
isn't already used. Port `3306` should also be available for local mysql connection.
