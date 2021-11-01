# Advent of Code 2019
Just my solutions for [Advent of Code 2019](https://adventofcode.com/2020).

Uses [Cobra](https://github.com/spf13/cobra) as CLI framework.

Requirements installation:
~~~~
> go get -u github.com/spf13/cobra
~~~~

(not working) How to run specific puzzle (input files included in /input-files):
~~~~
> go run day <day:1-25> <part:1/2> <input-file>
~~~~

(2021-11) How to run specific puzzle (input files included in /input-files):
~~~~
> go run main.go day*.go day <day:1-25> <part:1/2> <input-file>
~~~~


Because of TDD approach, tests are also included:
~~~~
> go test
~~~~
