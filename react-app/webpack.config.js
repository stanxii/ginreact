const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");


module.exports = {
    entry: "./src/index.js",
    output: {
        path: path.join(__dirname, "/dist"),
        filename: "bundle.js"
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ["babel-loader"]
            },
            {
                test: /\.css$/,
                use: [
                    "style-loader", // load and concat all .css to single file
                    "css-loader", // adds all styles into <style>
                ]
            }
        ]
    },
    plugins: [
        /**
         * Generates .html then put in script
         */
        new HtmlWebpackPlugin({
            template: "./src/index.html"
        })
    ]
};