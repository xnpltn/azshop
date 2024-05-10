/** @type {import('tailwindcss').Config} */
export const content = [
  './templates/**/*.templ',
  './handler/**/*.go',
];
export const theme = {
  extend: {
    fontFamily: {
      'inter': ['Inter', 'sans-serif'],
    },
    colors: {
      'light-gray': '#D8DEE9',
      'dark-gray': '#2E3440',
      'light-b': '#ECEFF4'
    }
  }
};

