The description from gobyexample:
> Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function. Here’s a function that will take an arbitrary number of ints as arguments. Within the function, the type of nums is equivalent to []int. We can call len(nums), iterate over it with range, etc. Variadic functions can be called in the usual way with individual arguments. If you already have multiple args in a slice, apply them to a variadic function using func(slice...) like this.

Even as fmt.Println() is a variadic function, and it can take both string and ints..I wonder if functions can allow it at the same time?

Answer: https://share.google/aimode/lBdXD5mdmjR9ooIlr

I see, this is where `any` comes into play. That is very neat, very strange as well. But still neat and helpful in terms of clean code.