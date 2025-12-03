# aoc-2025
Solutions for https://adventofcode.com/2025 using Go. Each day lives in its own module so you can copy yesterday's code forward without shared dependencies.

## Layout
- `days/day01/` – starter module with `main.go`, `main_test.go`, `go.mod`, `Makefile`, `input.txt`, and `test.txt`.
- Each new day should be a full copy of the previous day's folder so it stays self-contained. Update the module name in `go.mod` to match the new day.

## Starting a new day
1) Copy yesterday's folder, e.g. `cp -R days/day01 days/day02`.
2) Update the module name inside the new day's `go.mod`, e.g. change `day01` to `day02`.
3) Fill `test.txt` with the sample input and set expected outputs in `main_test.go`.
4) Drop your puzzle input into `input.txt`.
5) Run from within the day's folder:
   - `make test` to validate against the sample.
   - `make run` to solve the puzzle input (reads `input.txt` by default; you can also pass a path: `go run . other-input.txt`).
