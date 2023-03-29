'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  GOOGLE_API_KEY: "AIzaSyAPhJXKB6lW-phonNc98yzH-x1pJ6aOQso",
  API_HOST: "http://localhost:3000/",
  MARKER_SVG_PATH: "M13.9304 17.869C13.6084 18.7988 12.2931 18.798 11.9721 17.8678L10.3524 13.1739L5.78287 11.2307C4.87733 10.8456 4.96832 9.53344 5.91837 9.27705L16.1497 6.51591C16.9526 6.29922 17.6707 7.0693 17.3986 7.85518L13.9304 17.869Z"
})
