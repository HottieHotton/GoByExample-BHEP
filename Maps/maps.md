Maps are Go's version of dicts or hashes(hasmaps). Utilizing the make() command creates the empty map.

And you can declare the key value by doing this:

```go
var m map[string]int
m = make(map[string]int)
m["test"] = 42
fmt.Printf("%v", m)
```

You can also delete a key/value via the `delete()` command, and clear the entire thing by doing `clear()`.

When reading, I wanted to better understand this instruction:
> The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.

Going to Gemini, I learned this: https://share.gemini.google/ujrMQkqZzkO5

This helps explain why this is helpful.