/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.templ"],
  theme: {
    extend: {
      fontFamily: {
        sans: ['Graphik', 'sans-serif'],
      },
      colors: {
        primary: {
          DEFAULT: "#1DA1F2",
          dark: "#1A8CD8",
          white: "#FFFFFF",
        },
        secondary: "#14171A",
        neutral: {
          dark: "#657786",
          light: "#AAB8C2",
          xlight: "#E1E8ED",
          xxlight: "#F5F8FA",
        }
      },
    },
  },
  plugins: [],
}
