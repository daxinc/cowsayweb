package cowsay

import "bytes"

const (
	cowASCII = `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`
	// IndexHTML defines the HTML of the page
	IndexHTML = `<!doctype html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Cow Say</title>
</head>

<body>
<form method="get" action=".">
<input type="text" name="quote" size="100" value="%s"><input type="submit" value="Say">
</form>

<pre id="Quote">%s</pre>
</body>
</html>
`
)

// Say says the quote with a cow
func Say(quote string) string {
	var text string
	if quote == "" {
		text = "Hello there!"
	} else {
		text = quote
	}

	line := make([]byte, len(text)+3)
	copy(line[0:], " ")
	for i := 1; i < cap(line); i++ {
		copy(line[i:], "-")
	}

	var result bytes.Buffer
	result.WriteString(string(line))
	result.WriteString("\n")

	result.WriteString("< " + text + " >\n")
	result.WriteString(string(line))

	result.WriteString(cowASCII)

	return result.String()
}
