'use strict'

const { VueLoaderPlugin } = require('vue-loader')
const { VuetifyLoaderPlugin } = require('vuetify-loader')

module.exports = {
    mode: 'development',
    entry: [
        './src/app/app.js'
    ],
    watch: true,
    devtool: "inline-cheap-source-map",
    watchOptions: {
        ignored: ["node_modules/**"],
    },
    
    module: {
        rules: [
            // SASS has different line endings than SCSS
            // and cannot use semicolons in the markup
            {
                test: /\.sass$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                    {
                        loader: 'sass-loader',
                        // Requires sass-loader@^7.0.0
                        options: {
                            // This is the path to your variables
                            data: "@import './src/sass/variables.scss'"
                        },
                        // Requires sass-loader@^8.0.0
                        options: {
                            // This is the path to your variables
                            prependData: "@import './src/sass/variables.scss'"
                        },
                        // Requires >= sass-loader@^9.0.0
                        options: {
                            // This is the path to your variables
                            additionalData: "@import './src/sass/variables.scss'"
                        },
                    },
                ],
            },
            // SCSS has different line endings than SASS
            // and needs a semicolon after the import.
            {
                test: /\.scss$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                    {
                        loader: 'sass-loader',
                        // Requires sass-loader@^7.0.0
                        options: {
                            // This is the path to your variables
                            data: "@import './src/sass/variables.scss';"
                        },
                        // Requires sass-loader@^8.0.0
                        options: {
                            // This is the path to your variables
                            prependData: "@import './src/sass/variables.scss';"
                        },
                        // Requires sass-loader@^9.0.0
                        options: {
                            // This is the path to your variables
                            additionalData: "@import './src/sass/variables.scss';"
                        },
                    },
                    
                ],
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader',
            },
            {
                test: /\.js$/,
                loader: 'babel-loader',
                exclude: /node_modules/
            },

        ],
    },
    resolve: {
        extensions: ['*', '.js', '.vue', '.json'],
        alias: {
            'vue$': 'vue/dist/vue.esm.js' // Use the full build
        }
    },
    plugins: [
        new VueLoaderPlugin(),
        new VuetifyLoaderPlugin()
    ]

}

if(process.env.NODE_ENV === 'production') {
    module.exports.watch = false
}