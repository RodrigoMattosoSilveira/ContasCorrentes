# Abstract
This is a SPA, with four areas:
- `ui-nav` - The header, at the top of the page;
- `ui-rail` - A rail, on the left of the page;
- `ui-footer`- The footer, at the bottom of the page;
- `ui-body` - The main content area, which will affect how to handle the other three;

I want the ability to combine versions `ui-nav`, `ui-rail`, `ui-footer` with `ui-body`, based on the `ui-body` type.

# Approach
I'll use the Go Template {{ define }} and {{ block }} actions.
## Create the templates
Here, I define the target template, with th4 {{ block }} tags, where  I'll insert in the {{ define }} sources.

### Blocck
``` html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ block "title" . }}Default Title{{ end }}</title>
</head>
<body>
    {{ block "content" . }}{{ end }}
</body>
</html>
```
<sub> layout.html<sub>

### Define
Here, I define the source templates that I'll insert in the {{ block }} targets.
```html
{{ define "title" }}Home Page{{ end }}
{{ define "content" }}
<h1>Welcome to the Home Page!</h1>
<p>This is a simple example of using templates in Go Fiber.</p>
{{ end }}
```

## Load and Parse the templates
```go
package main

import (
	"html/template"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
)

// TemplateEngine wraps the Go html/template package
type TemplateEngine struct {
	templates *template.Template
}

// Render renders a template with data
func (t *TemplateEngine) Render(w io.Writer, name string, data interface{}, c *fiber.Ctx) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Parse templates
	templates, err := template.ParseFiles("layout.html", "home.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		Views: &TemplateEngine{templates: templates},
	})

	// Define routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("layout.html", fiber.Map{
			"Title": "Home Page",
		})
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
```

## Notes
1. `layout.html`: Acts as the base template with placeholders for title and content.
1. `home.html`: Defines specific content for the title and content blocks.
1. `Fiber's Render Method`: Renders the layout.html template, which dynamically includes the blocks defined in home.html.
This approach allows you to create modular and reusable templates in your Go Fiber application. Let me know if you need further clarification

# My implementation
## The UiBody Structute

```go
type UiBody struct {
	uiBody template.Template
}
```