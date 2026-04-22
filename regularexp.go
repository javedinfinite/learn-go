package main
import ("fmt"
		"regexp")

func PrintRegularExp () {
	fmt.Println("*******************Printing regular expression************************")

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("matched expe : ", match)

	// Instead of calling MatchString repeatedly, compile once for more efficient and cleaner use
	re1 := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("matched compiled re : ",re1.MatchString("peach"))

	// Here are core building blocks:

	// | Pattern        | Meaning                         | Example                          |
	// |----------------|----------------------------------|----------------------------------|
	// | `.`            | any character                    | `g.` → "go", "ga"                |
	// | `\d`           | digit (0–9)                      | `\d+` → "123"                    |
	// | `\w`           | word char (a-zA-Z0-9_)           | `\w+` → "abc123"                 |
	// | `\s`           | whitespace                       | `\s+` → spaces/tabs              |

	// | `[abc]`        | a OR b OR c                      | `[ae]t` → "at", "et"             |
	// | `[a-z]`        | lowercase letters                | `[a-z]+` → "hello"               |
	// | `[A-Z]`        | uppercase letters                | `[A-Z]+` → "GO"                  |
	// | `[0-9]`        | digits                           | `[0-9]+` → "42"                  |
	// | `[^a-z]`       | NOT lowercase                    | `[^a-z]+` → "123!"               |

	// | `*`            | 0 or more                        | `go*` → "g", "goooo"             |
	// | `+`            | 1 or more                        | `go+` → "go", "goooo"            |
	// | `?`            | 0 or 1 (optional)                | `colou?r` → "color", "colour"    |
	// | `{n}`          | exactly n times                  | `\d{3}` → "123"                  |
	// | `{n,m}`        | between n and m times            | `\d{2,4}` → "12", "1234"         |

	// | `^`            | start of string                  | `^go` → "golang"                 |
	// | `$`            | end of string                    | `go$` → "lego"                   |

	// | `( ... )`      | capture group                    | `(\d+)` → extract "123"          |
	// | `|`            | OR                               | `cat|dog` → "cat", "dog"         |

	// | `\.`           | literal dot                      | `\.com` → ".com"                 |
	// | `\+` `\*` etc. | escape special chars             | `\+` → "+"                       |

	// Find first match
	re := regexp.MustCompile(`\d+`) // extracting digit from given string
	fmt.Println("first digit: ", re.FindString("Age: 25 34 45 67")) // just extract first digit
	fmt.Println("first 3 digit: ",re.FindAllString("Age: 25 34 45 67", 3)) // extract first 3 digits found
	fmt.Println("extract all digit: ",re.FindAllString("Age: 25 34 45 67", -1)) // extract all digits found

	// Groups capturing
	re2 := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
	matches := re2.FindStringSubmatch("test@gmail.com")

	fmt.Println(matches) // index 0 is full match value, rest index are as per group matching
	// ["test@gmail.com", "test", "gmail", "com"]

	// replace using regx
	re3 := regexp.MustCompile(`\d+`)
	result := re3.ReplaceAllString("I have 2 apples", "X")
	fmt.Println("replaced : ",result) // "I have X apples"

	// Email matching
	re4 := regexp.MustCompile(`^[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}$`)
	fmt.Println("email matching : ",re4.MatchString("test@gmail.com")) // true


	// Need to do (pending) :
	// Practice 3–4 problems once:
		// 	Extract all numbers from a string
		// 	Validate a simple email
		// 	Replace extra whitespace
		// 	Parse key-value text


}