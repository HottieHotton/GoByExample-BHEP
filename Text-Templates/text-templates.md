When breaking down the reason why Text Templates are helpful, I turned to Gemini to get a better grasp between the concatenation (% type commands) and templates:
>Text templates in Go (via the standard library's text/template package) are data-driven engines used to generate formatted text automatically.
>
>When you're starting out with Go, it's easy to just use string concatenation `("Hello " + name)` or `fmt.Sprintf("Hello %s", name)`. However, as soon as your output gets larger or more dynamic, those approaches break down.
>
>1. Separation of Logic and Presentation
>Mixing Go code with formatted text makes code hard to read and maintain.
>
>Without templates: You end up chaining multiple fmt.Sprintf calls, handling loops manually, and building massive multi-line string buffers in code.
>
>With templates: You define your text structure in a clean layout or external file and simply inject your Go structs or maps into it.
>```
>Hello {{.Name}},
>
>Your recent order #{{.OrderID}} has shipped.
>Items:
>{{range .Items}}
>  - {{.Name}}: ${{.Price}}
>{{end}}
>```
>2. Built-in Control Flow (Loops & Conditionals)
>Templates have their own simple mini-language. You can iterate over slices, maps, or structs using {{range}}, and conditionally output sections using {{if}}, {{else}}, and {{with}} without pollution from Go logic.
>
>3. Safety and Security
>When generating web pages or raw text output, untrusted inputs can introduce security vulnerabilities (like Cross-Site Scripting or HTML injection).
>
>Go offers html/template alongside text/template. Both share the exact same syntax, but html/template automatically escapes dangerous characters (like `<script>`), preventing security bugs without extra work.
>
>4. Custom Function Pipelines (Funcs)
>Templates let you pipe output through custom functions using the Unix-style | operator. This keeps formatting logic out of your main Go code:
>
>```
>Joined: {{.JoinedDate | formatDate}}
>User: {{.Username | toUpper}}
>```
>5. Reusability and Modularity
>You can define smaller templates and nest them inside larger ones using `{{template "header" .}}`. This makes it easy to construct complex documents, configuration files, or emails out of reusable parts.
>
>Common Beginner Use Cases
>Dynamic Emails & Reports: Inserting user data into welcome emails, invoices, or alerts.
>
>Code Generation: Writing Go tools that generate other .go code files based on schemas.
>
>Config File Generation: Dynamically building yaml, json, or dockerfiles based on environment variables.
>
>CLI Tools: Formatting terminal outputs and help menus cleanly.