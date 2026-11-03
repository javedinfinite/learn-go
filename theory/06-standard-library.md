# 6. Useful Go library examples

Go comes with a **standard library**: packages of ready-made code for common jobs. These examples show a few of them. Open the linked file and look for the function named in the notes.

- [`random.go`](../random.go) uses `math/rand/v2` to generate random numbers. It also converts text to numbers and numbers back to text with `strconv`.
- [`time.go`](../time.go) gets the current time, compares times, and adds a duration. It also parses a date string and formats a time. Go uses a special example date as its layout: `Mon Jan 2 15:04:05 MST 2006`. The numbers in that example stand for parts of the date and time.
- [`regularexp.go`](../regularexp.go) searches text using regular expressions, which are patterns such as `\d+` for one or more digits.
- [`json.go`](../json.go) changes Go values into JSON and reads JSON into Go values. Struct tags such as ``json:"full_name"`` tell Go which JSON key matches a field.
- [`logging.go`](../logging.go) prints log messages with dates, times, and other useful details.

Some of these operations can fail. For example, a date may not match the expected layout, or JSON may be invalid. Check the returned error before using the result so the program can respond to the problem.
