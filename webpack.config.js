const path = require('path');

module.exports = {
  mode: 'development',

  entry: {
    bundle: './_dist/index.js'
  },

  output: {
    path: path.resolve(__dirname, '_dist'),
    filename: '[name].js'
  },
  devtool: 'inline-source-map'
};
