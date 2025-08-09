# urlLengthener
A Go program that artificially lengthens URLs by adding random tracking parameters and query strings. This is primarily a demonstration of URL manipulation in Go.

## Features

- Adds common tracking parameters (utm_source, utm_medium, etc.)
- Inserts random query parameters
- Optionally adds fragment identifiers
- Interactive command-line interface
- Basic URL validation

## Installation

1. Ensure you have [Go installed](https://golang.org/doc/install) (version 1.16+ recommended)
2. Clone this repository or download the source files
3. Build the program:

```bash
go build
```
## User Input:

Now prompts the user to enter a URL

Uses bufio.NewReader to read the input properly

Trims whitespace from the input

## Basic Validation:

Checks if the input starts with "http" (though this is very basic validation)

Shows a warning if it doesn't look like a complete URL

## User Experience:

Clear prompts and output formatting

Waits for user to press enter before exiting (useful when running from explorer)

## Additional Parameters:

Added fbclid (Facebook click ID) as another common tracking parameter

## How to Use:
Run the program

Enter the URL you want to lengthen when prompted

View the lengthened URL

Press enter to exit

### Example Usage:
```text
URL Lengthener
--------------
Enter the URL to lengthen: https://example.com/test

Original URL: https://example.com/test
Lengthened URL: https://example.com/test?utm_source=4jFg7K2m&utm_medium=9Hq3z&utm_campaign=pR8sW2nL5v&ref=7kT9mN2sP4qY&fbclid=wX5yK8nJ2qR7tP4&abc=12345&xyz=9876abcd&longparam=verylongvalue#3kf7mLp2

Press enter to exit...
```
