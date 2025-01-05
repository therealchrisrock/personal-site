/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/**/*.{go,js,templ,html}",
        "./node_modules/flowbite/**/*.js"
    ],
    theme: {
      extend: {},
    },
    plugins: [
        require('flowbite/plugin')
    ],
  }