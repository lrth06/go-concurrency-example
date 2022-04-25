# go-concurrency-example

## Purpose

You've likely heard that [Go](https://golang.org/doc/install) is a great language for writing **FAST** code... This example repository serves as an example of just how fast... 

The program scrapes the [jsonplaceholder](https://jsonplaceholder.typicode.com/) API in its entirety, and then writes the results to a series of csv and image files.


## Usage

### Prerequisites

* [Go](https://golang.org/doc/install)

Clone Repo:

```bash
git clone git@github.com/lrth06/go-concurrency-example.git
cd go-concurrency-example
```
Build Binary:

```bash
go build -o scraper
./scraper
```

From this point, the program will run and print the resulting times to the console, and all files will be written to the  `output` directory.