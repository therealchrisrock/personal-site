const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/**/*.{go,js,templ,html}",
    ],
    theme: {
        extend: {
            fontFamily: {
                mono: ['"Fira Code"', ...defaultTheme.fontFamily.mono]
            }
        }
    },
    plugins: [],
  }