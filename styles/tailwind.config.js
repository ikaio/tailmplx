/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')

module.exports = {
  content: ["../components/**/*.templ"],
  theme: {
    extend: {
      boxShadow: {
        
      },
      fontFamily: {
        'sans': ['Jost', ...defaultTheme.fontFamily.sans],
      },
    },
  },
  plugins: [],
}

