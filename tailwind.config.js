/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{go,html,templ}",
    "./components/*.{html, templ, go}",
  ],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ["dark", "lofi"],
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
