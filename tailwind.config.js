/** @type {import('tailwindcss').Config} */

const defaultTheme = require("tailwindcss/defaultTheme");
module.exports = {
  content: ["./view/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["OpenSans", ...defaultTheme.fontFamily.sans],
        serif: ["Antipasto", ...defaultTheme.fontFamily.serif],
      },
    },
  },
  plugins: [
    // ...
    require("@tailwindcss/forms"),
    require("@tailwindcss/typography"),
    require("@tailwindcss/aspect-ratio"),
  ],
};
