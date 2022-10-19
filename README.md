
# are-you-bombed

A Golang application which can generate flat and nested zip bombs and serve them using Gin HTTP web framework.

A zip bomb, also known as a decompression bomb or zip of death, is a malicious archive file designed to crash or render useless the program or system reading it. It is often employed to disable antivirus software, in order to create an opening for more traditional malware. -**Wikipedia**

## Run Locally

Clone the project

```bash
  git clone https://github.com/SANTHOSH17-DOT/are-you-bombed.git
```

Go to the project directory

```bash
  cd are-you-bombed
```

## Usage/Examples

### Generate flat zip bomb

```bash
  go run main.go generate flat <number of files>
```

#### Example

```
$ go run main.go generate flat 100

bomb file: bomb/flat/bomb-flat.zip 104 KB
Decompression size 100 MB
Total time elapsed: 508.3543ms milliseconds
```


### Generate nested zip bomb

```bash
  go run main.go generate nested <depth>
```

#### Example

```
$ go run main.go generate nested 10

bomb file: bomb/nested/bomb-nested.zip 29 KB
Decompression size 1000000000 MB
Total time elapsed: 299.697ms milliseconds
```

### Run the server

```bash
  go run main.go host
```

### The zip bombs can be downloaded from

http://localhost:8080/static/flat/bomb-flat.zip  
http://localhost:8080/static/nested/bomb-nested.zip

## Contributing

Contributions are always welcome!  
Please open an issue if you find any improvement.

## License

[MIT](https://choosealicense.com/licenses/mit/)


**FOR EDUCATIONAL PURPOSE ONLY**
