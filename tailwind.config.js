/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["**/*.templ"],
  theme: {
    fontFamily: {
      sans: ["Inter", "sans-serif"],
      display: ["'IBM Plex Sans'", "sans-serif"],
    },
    extend: {
      colors: {
        primary: "rgb(29 78 216)",
      },
    },
  },
  plugins: [],
};
