# 3. Functions, strings, and structs

### Functions

A function is a named piece of code that can receive values, do some work, and return a result. Go functions can return more than one value. A common example is returning a result together with an error. A variadic function can receive any number of values of the same type.

Open [`functions.go`](../functions.go). The `closureFunc` remembers the variable `x` from the surrounding function, so each call can increase it. The factorial function calls itself; this is called **recursion**.

### Strings and characters

A Go string stores text as UTF-8 bytes. For simple English text, one character is usually one byte. Some characters use more than one byte, so `len(text)` counts bytes and `text[0]` gives one byte. See [`strings_rune.go`](../strings_rune.go) for examples.

A **rune** represents one Unicode code point. Converting text to `[]rune` is useful when you want to work with characters instead of individual bytes. The same file also shows common string functions and ways to format values with `fmt.Printf`.

### Structs and methods

A struct groups related values together. For example, a `Person` can have a name, age, and height. [`struct_method.go`](../struct_method.go) creates `Person` values and adds a method that calculates using the person's fields.

A method is a function attached to a type. The method in this example uses a pointer receiver (`*Person`), so it can update the original person's age. Field names that begin with a capital letter can be used from other packages; lowercase fields are private to their package.
