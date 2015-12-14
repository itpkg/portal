var path = require("path");
var webpack = require("webpack");
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = function (options) {
    var entry = {
        vendor: [
            'jquery',
            'bootstrap',
            'history',
            'js-base64',
            'highlight.js',
            'markdown',
            'i18next/lib',
            'i18next-browser-languagedetector/lib'
        ],
        react: [
            'react',
            'react-dom',
            'react-bootstrap',
            'react-router'
        ],
        main: "./app/main"
    };


    var loaders = [
        {test: /\.jsx?$/, exclude: /node_modules/, loader: 'babel', query: {presets: ['react', 'es2015']}},
        {test: /\.css$/, loader: "style!css"},
        {test: /\.(png|jpg|jpeg|gif|svg|ttf|woff|woff2|eot)$/, loader: "file-loader"}
    ];

    var plugins = [
        new webpack.ProvidePlugin({
            //fix 'jQuery is not defined' bug
            $: "jquery",
            jQuery: "jquery",
            //fix 'ReactIntl is not defined' bug
            ReactIntl: 'react-intl'
        })
    ];

    var htmlOptions = {
        title: 'IT-PACKAGE Portal',
        inject: true,
        template: './app/index.html'
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

        plugins.push(new webpack.DefinePlugin({
            "process.env": {
                NODE_ENV: JSON.stringify("production")
            }
        }));
        plugins.push(new webpack.NoErrorsPlugin());

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