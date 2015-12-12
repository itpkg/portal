var path = require("path");
var webpack = require("webpack");
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = function (options) {
    var entry = {
        vendor: [
            'jquery',
            'history',
            'js-base64',
            'highlight.js',
            'markdown'
        ],
        react: [
            'react',
            'react-bootstrap',
            'react-router'
        ],
        main: options.prerender ? "./config/mainPrerender" : "./config/mainApp"
    };


    var loaders = [
        {test: /\.jsx?$/, exclude: /node_modules/, loader: 'babel'},
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

        plugins.push(new webpack.optimize.UglifyJsPlugin({
            output: {
                comments: false
            }
        }));

    }
    plugins.push(new HtmlWebpackPlugin(htmlOptions));

    var output = {
        path: path.join(__dirname, options.prerender ? 'assets' : 'public'),
        filename: options.prerender ? "[id]-[chunkhash].js" : '[name].js'
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
