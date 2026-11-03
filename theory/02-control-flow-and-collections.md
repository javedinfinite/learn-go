# 2. Control flow and collections

Go uses a few simple statements to control a program:

- `if` chooses whether to run some code.
- `for` repeats code. Go uses `for` for all kinds of loops.
- `switch` chooses one matching case from several options.

Look at [`loops_case.go`](../loops_case.go). It shows a condition loop, a counted loop, a loop that stops with `break`, a loop over a number of values, and two kinds of `switch`.

### Arrays, slices, and maps

Go has three common ways to keep several values together:

- An **array** has a fixed number of items. Its length is part of its type. See [`arrays.go`](../arrays.go).
- A **slice** is used when the number of items can change. See [`slices.go`](../slices.go).
- A **map** stores values under keys, like a dictionary. See [`maps.go`](../maps.go).

In `slices.go`, `append` adds items to a slice. A slice made from another slice can share the same underlying data, so changing one may change the other. Use `copy` or `slices.Clone` when you need a separate copy. `len` tells you how many items are currently in a slice; `cap` tells you how much room it has before Go needs to make more space.

In `maps.go`, a lookup for a key that is not present gives the value type's zero value. The two-result form, such as `value, found := myMap[key]`, also tells you whether that key existed. Map entries can appear in any order when you loop over a map.
