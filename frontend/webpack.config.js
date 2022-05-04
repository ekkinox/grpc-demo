const path = require("path");
const Dotenv = require("dotenv-webpack");
const { CleanWebpackPlugin }  = require('clean-webpack-plugin')
const HtmlWebpackPlugin = require("html-webpack-plugin");
const isProduction = process.env.NODE_ENV === "production";
const env = new Dotenv({
  path: './.env', // load this now instead of the ones in '.env'
  safe: true, // load '.env.example' to verify the '.env' variables are all set. Can also be a string to a different file.
  allowEmptyValues: true, // allow empty variables (e.g. `FOO=`) (treat it as empty string, rather than missing)
  systemvars: true,
  silent: true, // hide any errors
  defaults: false // load '.env.defaults' as the default values if empty.
});
const config = {
  entry: {
    "gweb": path.resolve(__dirname, 'src/index.ts')
  },
  output: {
    path: path.resolve(__dirname, "dist"),
    chunkFilename: '[name].js',
    filename: '[name].js'
  },
  devtool:'source-map',
  devServer: {
    open: true,
    host: "localhost",
  },
  plugins: [
    env,
    new CleanWebpackPlugin(),
    new HtmlWebpackPlugin({
      filename: 'index.html',
      inject: true,
      template: path.resolve(__dirname, 'src', 'index.html'),
    }),
  ],
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/i,
        loader: "ts-loader",
        exclude: ["/node_modules/"],
      },
    ],
  },
  resolve: {
    extensions: [".tsx", ".ts", ".js"],
  },
};

module.exports = () => {
  if (isProduction) {
    config.mode = "production";
  } else {
    config.mode = "development";
  }
  return config;
};
