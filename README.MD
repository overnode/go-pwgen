# Go-Pwgen
A CLI password generator made in Go [https://golang.org]. Fast generation of passwords with upper-/lowercase, numbers, symbols. Also possible to write results to file.

## Compile
- Run `go build`

## Running
- running `./go-pwgen`
generates one password:
```
Aiq573tV
---
Duration 103.233µs
```

### Options
```
Usage of ./go-pwgen:
  -f string
    	Write to specified filename
  -l int
    	Length of password (default 8)
  -lc
    	Lowercase letters (default true)
  -n int
    	Number of generated passwords (default 1)
  -num
    	Numbers (default true)
  -symbol
    	Symbols
  -uc
    	Uppercase letters (default true)
```