var path = require("path");
var webpack = require("webpack");
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = function (options) {
    var entry = {
        main: options.prerender ? "./config/mainPrerender" : "./config/mainApp"
    };


    var loaders = [
        {test: /\.css$/, loader: "style!css"},
        {test: /\.jpg$/, loader: "file-loader"},
        {test: /\.png$/, loader: "url-loader?mimetype=image/png"}
    ];

    var plugins = [];

    var htmlOptions = {
        favicon: './favicon.ico',
        title: 'IT-PACKAGE Portal'
    };
    if (options.minimize) {
        htmlOptions.minify = {
            collapseWhitespace: true,
            removeComments: true
        };
    }
    plugins.push(new HtmlWebpackPlugin(htmlOptions));

    var output = {
        path: path.join(__dirname, options.prerender ? 'assets' : 'public'),
        filename: "[id]-[hash].js"
    };

    return {
        entry: entry,
        output: output,
        plugins: plugins,
        module: {
            loaders: loaders
        }
    }
};
