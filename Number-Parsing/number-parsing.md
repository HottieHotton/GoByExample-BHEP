Notes regarding the functionality of `strconv` and the options you can use(Thanks Gemini for formatting):
# Go Number Parsing

* **`ParseFloat`**: The `64` parameter specifies how many bits of precision to parse.
* **`ParseInt` (Base Inference & Bit Size)**: Setting the base parameter to `0` means inferring the base directly from the string, while `64` requires that the result fit in 64 bits.
* **`ParseInt` (Hexadecimal Support)**: Automatically recognizes hex-formatted numbers (e.g., strings starting with `0x`).
* **`ParseUint`**: Available for parsing unsigned integers.
* **`Atoi`**: A convenience function for basic base-10 integer parsing.
* **Error Handling**: `Parse` functions return an error when encountering invalid input.