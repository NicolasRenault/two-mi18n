![Static Badge](https://img.shields.io/badge/Mirrored_from-Forgejo-%23FB923C?style=for-the-badge&logo=forgejo)

# two-mi18n

> Minimalist Javascript library for internationalization in only two methods. Too minimalist.

`Two Mi18n` is a Javascript library for translating websites directly on the client.
The goal of this library is to provide the easiest way to translate any website.

-   Minimalist and easy to use: Only two methods
-   Lightweight: ~1.6 KB minified
-   Client-side oriented

## Inspiration

I was searching for a simple way to translate my website. I have found some libraries but they were too complicated for my needs. So I decided to build my own.

I had the idea of a client-side i18n library in mind, then found [simple-translator](https://github.com/andreasremdt/simple-translator) that had the same idea but 3 years before. But I wanted something more minimalist.
So I take inspiration from [Apline.js](https://alpinejs.dev/) for its minimalist approach.

## Getting started

### Installation

Since it is Client-side oriented, the best way to use it is to call the CDN with this script tag:

```html
<script
	defer
	src="https://unpkg.com/two-mi18n@latest/dist/TwoMi18n.umd.js"
></script>
```

### Usage

#### The translation object

The `translation object` is a simple Javascript object that contains the translations. Its keys are the language keys and its values are objects that contain the translations with keys and values.

**The `translation object` must have a `default` key that contains the default language. The default language value must exist in the `translation object`**

```js
const translations = {
	default: "en",
	en: {
		hello: "Hello",
		world: "World",
	},
	fr: {
		hello: "Bonjour",
		world: "Monde",
	},
};
```

#### The `TwoMi18n` object

The `TwoMi18n` object is the main object of the library. It contains the two methods of the library.

**The translation object will be validated by the `TwoMi18n` constructor. If the `translation object` is not valid, an error will be thrown.**

```js
const twoMi18n = new TwoMi18n(translations);
```

#### Translate in Javascript

The `translate` method is the method that will translate a single string. It takes two arguments:

-   `key`: The translation key.
-   `lang`: The language to translate to from the `translation object`.

```js
twoMi18n.translate("hello", "fr"); // Bonjour
```

You can use any type of language code, are even create new ones.

**If you pass a language that is not in the `translation object`, it will try the first 2 letters.**
This behavior is useful if you want to specify language variations depending on the location.
For example `en-GB`.

```js
const translations = {
	default: "en",
	en: {
		hello: "Hello",
		color: "Color",
	},
	en-GB: {
		hello: "Hello",
		color: "Colour",
	},
};

twoMi18n.translate("color", "en"); // Color
twoMi18n.translate("colour", "en-GB"); // Colour
```

Now if you want to get a translation for a country not defined in the `translation object`, like `en-US`. It will try to match the two first letters.

```js
twoMi18n.translate("Color", "en-US"); // Color
//Fallback to the first two letters. Here "en".
```

If a not-defined variable that doesn't fall back with its two first letters is passed in the param, it will take the value from the `default`.

```js
twoMi18n.translate("color", "it"); // Color
```

#### Translate in HTML

Add the `data-twomi18n` attribute to the elements that need to be translated in the HTML. The value of the attribute is the translation key.

```html
<h1 data-twomi18n="hello"></h1>
```

The `translateHTML` method is the method that will translate all the elements with the `data-twomi18n` attribute.

It takes one argument:

-   `lang`: The language to translate to

```js
twoMi18n.translateHTML("fr");
```

##### Translate HTML attributes

The `translateHTML` method can also translate HTML attributes. Add the attribute name after the translation key between brackets to specify the attribute to translate.

```html
<input type="text" data-twomi18n="hello[placeholder]"></input>
```

You can translate multiple attributes by separating them with a space.

```html
<input
    type="text"
    data-twomi18n="world hello[placeholder] world[title]"
></input>
```

**Note: When the attribute is not specified, the `innerText` attribute is used.**

#### Example

```html
<header>
	<h1 data-twomi18n="hello"></h1>
    <input
        type="text"
        data-twomi18n="world hello[placeholder] world[title]"
    ></input>
</header>

<script
	defer
	src="https://unpkg.com/two-mi18n@latest/dist/TwoMi18n.umd.js"
></script>

<script defer>
	const translations = {
		default: "en",
		en: {
			hello: "Hello",
			world: "World",
		},
		fr: {
			hello: "Bonjour",
			world: "Monde",
		},
	};

	const twoMi18n = new TwoMi18n(translations);

	twoMi18n.translateHTML("fr");
</script>
```
See a full working example:

[![Two Mi18n example](https://codesandbox.io/static/img/play-codesandbox.svg)](https://codesandbox.io/p/sandbox/bold-snowflake-tt6chr)

You can find real world projects examples:

- [My personal portfolio](https://nicolasrenault.com) - [See code](https://github.com/NicolasRenault/portfolio)
- [QrCode Prevention](https://qrcode.nicolasrenault.com) - [See code](https://github.com/NicolasRenault/qrcode-prevention)

See the full documentation on [twomi18n.nicolasrenault.com](https://twomi18n.nicolasrenault.com/).

## Helper Script

Automatically extract translation keys from your project to generate your JSON translation object.

Using the script is as minimalist as the library itself. Simply pass the path of your project folder to the command. It will go recursively through all the files from there.

```bash frame="none"
curl -fsSL https://two-mi18n-helper-script.nicolasrenault.com/install.sh | sh
```

```bash frame="none"
twomi18n-helper-script ./src
```

Running the script will generate the language array for the `translation object` for you:

```json
{
    "homepage": "Homepage",
    "contact-me": "Contact me",
    "contact-button": ""
}
```

Learn more about the `helper-script` [here](https://twomi18n.nicolasrenault.com/go-further/helper-script).

## Contribute

If you want to contribute to the project, you can open an [issue](https://github.com/NicolasRenault/two-mi18n/issues/new) or a [pull request](https://github.com/NicolasRenault/two-mi18n/compare).

## License

[MIT](https://github.com/NicolasRenault/two-mi18n/blob/main/LICENSE)
