# Helper Script
> Automatically extract translation keys from your project to generate your JSON translation object.

After using the library myself, I noticed that manually keeping track of all the `data-twomi18n` keys scattered across my project was a nightmare. 

That's why I created the **helper-script** to solve this. It crawls your project folder, finds every translation key, and automatically generates the clean JSON object needed for your `translations` configuration.

## Installation

### Install the command

Following the concept of the library, I tried my best to make the helper-script minimal and lightweight. You can install the pre-compiled binary for macOS, Linux, or Windows using a simple one-liner in your terminal:

```bash frame="none"
curl -fsSL https://two-mi18n.helper-script.nicolasrenault.com/install.sh | sh
```

> The script will automatically detect your OS and architecture, and install the `twomi18n-helper-script` command to your terminal.

### Manual Installation

If you don't want to install the command to your machine, you can download the script from the 

## Usage

Using the script is as minimalist as the library itself. Simply pass the path of your project folder to the command. It will go recursively through all the files from there.

```bash frame="none"
twomi18n-helper-script ./src
```

By default, the script outputs the generated JSON directly in your terminal. To save this output to a file, simply pipe it into a `.json` file:

```bash frame="none"
twomi18n-helper-script ./src > locales/en.json
```

## How it works

The script scans common web files looking for the `data-twomi18n` attribute. See [Supported File Extensions](#supported-file-extensions).

It handles two specific behaviors to make your life easier:

1. **Default Values:** It captures the inner text of the HTML tag and cleans up any messy whitespace or newlines, providing a ready-to-use default translation.
2. **Other Attribute:** If you use bracket notation for other attributes (e.g., `[data-copied]`), the script isolates the key and initializes it with an empty string, ready for manual input.

### Example

Given the following HTML file in your project:

```html
<nav>
    <a href="/" data-twomi18n="homepage">
        Homepage
    </a>
    <button data-twomi18n="contact-me contact-button[title]" title="Contact Button">
        Contact me
    </button>
</nav>

```

Running the script will generate this exact JSON structure:

```json
{
    "homepage": "Homepage",
    "contact-me": "Contact me",
    "contact-button": ""
}

```

You can now directly copy this output into the `translations` object when initializing the `TwoMi18n` class in your Javascript!
> Pro tip: Search for all the `""` in your string to quickly find the **Other Attribute** that don't have a translation linked to them. 

## Supported File Extensions

The script is designed to work seamlessly with almost any web framework, static site generator, or templating engine. 

Out of the box, it automatically scans the following file types for `data-twomi18n` attributes:

```go
var htmlExtensions = map[string]bool{
    // --- Standard & Legacy ---
    ".html": true, ".htm": true,
    ".xhtml": true, ".xml": true,
    ".asp": true, ".aspx": true, // Classic ASP / WebForms
    ".cfm": true, ".cfml": true, // ColdFusion

    // --- Modern Frontend (JavaScript/TypeScript) ---
    ".jsx": true, ".tsx": true, // React
    ".vue":     true,                   // Vue
    ".svelte":  true,                   // Svelte
    ".astro":   true,                   // Astro
    ".html.js": true, ".html.ts": true, // Lit / HTML-in-JS

    // --- PHP & Content Management ---
    ".php": true, ".phtml": true,
    ".twig":      true, // Symfony
    ".blade.php": true, // Laravel
    ".tpl":       true, // Smarty / Dwoo

    // --- Python & Ruby ---
    ".jinja": true, ".jinja2": true, // Flask / Django
    ".pyhtml": true,
    ".erb":    true, // Ruby on Rails
    ".rhtml":  true,

    // --- Go (Golang) Specific ---
    ".gohtml": true, // Standard html/template
    ".tmpl":   true, // Common generic template

    // --- Java & .NET ---
    ".jsp": true, ".jspx": true, // Java Server Pages
    ".tag":    true,
    ".cshtml": true, ".razor": true, // ASP.NET Core / Blazor

    // --- Static Site Generators & Logicless ---
    ".liquid":   true, // Shopify / Jekyll
    ".mustache": true,
    ".hbs":      true, ".handlebars": true, // Handlebars
    ".ejs": true, // Embedded JS
    ".njk": true, // Nunjucks
    ".mdx": true, // Markdown + JSX

    // --- Abstraction Layers (HTML-like syntax) ---
    ".pug": true, ".jade": true, // Pug
    ".haml": true, // Haml
    ".slim": true, // Slim
}
```