Notes from the Go By Example Website(Formmating provided by Gemini):
# Go String Formatting

Go offers excellent support for string formatting in the `printf` tradition. Here are some examples of common string formatting tasks.

Go offers several printing “verbs” designed to format general Go values:

* **General Struct Formatting (`%v`)**: Prints an instance of a struct.
* **Struct Field Names (`%+v`)**: If the value is a struct, this variant will include the struct’s field names.
* **Go Syntax Representation (`%#v`)**: Prints a Go syntax representation of the value, i.e., the source code snippet that would produce that value.
* **Type Information (`%T`)**: Prints the type of a given value.

---

## Formatting Specific Types

* **Booleans (`%t`)**: Straightforward boolean formatting.
* **Integers (`%d`)**: Standard, base-10 formatting.
* **Binary (`%b`)**: Prints a binary representation of an integer.
* **Character (`%c`)**: Prints the character corresponding to the given integer.
* **Hex Encoding (`%x`)**: Provides hex encoding for integers.
* **Floats (`%f`)**: Basic decimal formatting for floating-point numbers.
* **Scientific Notation (`%e` / `%E`)**: Formats the float in (slightly different versions of) scientific notation.
* **Basic Strings (`%s`)**: Standard string printing.
* **Double-Quoted Strings (`%q`)**: Double-quotes strings as in Go source.
* **String Hex Encoding (`%x`)**: Renders the string in base-16, with two output characters per byte of input.
* **Pointers (`%p`)**: Prints a representation of a pointer address.

---

## Width and Precision Control

When formatting numbers, you will often want to control the width and precision of the resulting figure:

* **Integer Width**: Specify the width of an integer using a number after the `%` in the verb. By default, the result will be right-justified and padded with spaces.
* **Float Width & Precision**: Specify width and restrict decimal precision using the `width.precision` syntax.
* **Left-Justify Numbers**: To left-justify numbers, use the `-` flag inside the width specifier.
* **String Width (Right-Justified)**: Control string width to ensure alignment in table-like output using standard right-justified width syntax.
* **String Width (Left-Justified)**: To left-justify strings, use the `-` flag as with numbers.

---

## Formatting Variants

So far we’ve seen `Printf`, which prints the formatted string to `os.Stdout`:

* **`Sprintf`**: Formats and returns a string without printing it anywhere.
* **`Fprintf`**: Allows you to format and print directly to `io.Writer` instances other than `os.Stdout`.