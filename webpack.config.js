const path = require('path');

module.exports = (_, { mode }) => ({
  mode: mode === 'development' ? 'development' : 'production',
  devtool: mode === 'development' ? 'inline-source-map' : undefined,
  entry: path.resolve(__dirname, 'src', 'app.tsx'),
  output: {
    path: path.resolve(__dirname, 'public'),
    filename: 'app.js',
  },
  module: {
    rules: [
      {
        test: /\.[jt]sx?$/,
        exclude: /node_modules/,
        use: [
          {
            loader: 'babel-loader',
            options: {
              presets: [
                ['@babel/preset-env', {
                  targets: {
                    chrome: '58',
                  },
                }],
                '@babel/preset-react',
              ],
            },
          },
          {
            loader: 'ts-loader'
          },
        ],
      },
    ],
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
  },
  watchOptions: {
    ignored: /node_modules/,
  },
});

