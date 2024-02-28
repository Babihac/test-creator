/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{go,html,templ}",
    "./components/*.{html, templ, go}",
    "./pages/**/*.{go,html,templ}",
    "./pages/*.{html, templ, go}",
    "./handlers/*/*.{go,html,templ}",
    "./js/*.ts",
    "./js/**/*.ts",
  ],
  theme: {
    extend: {
      backgroundImage: (theme) => ({
        sun: "url('/icons/sun.svg')",
        moon: "url('/icons/moon.svg')",
        delete: "url('/icons/delete.svg')",
      }),
    },
  },
  daisyui: {
    themes: [
      {
        light: {
          ...require("daisyui/src/theming/themes")["light"],
          ".form-shadow": {
            "box-shadow":
              "0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.3);",
          },
        },
        dark: {
          ...require("daisyui/src/theming/themes")["dark"],
          ".form-shadow": {
            "box-shadow":
              "0 10px 15px -3px rgb(147 197 253 / 0.5), 0 4px 6px -4px rgb(147 197 253 / 0.5);",
          },
        },
      },
    ],
  },
  plugins: [require("@tailwindcss/typography"), require("daisyui")],
};
