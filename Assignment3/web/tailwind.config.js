/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        'primary': '#0B131E',
        'secondary': '#202B3B',
      },
      screens: {
        xsm: '500px',
      },
    },
  },
  plugins: [],
}

