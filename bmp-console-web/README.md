# bmp-console-web-私有化物理机管理平台项目

English | [简体中文](README.zh-CN.md) 
## Introduction
* This project is built based on @vue/cli 4 (vue3 + ts + pinia)

## Introduction to directory and file functions
* The config directory stores webpack related configurations
* The public directory is created by @vue/cli by default and stores some static resources that do not need to be packaged by webpack
* public/robots.txt- is a plain text file that can declare the parts of the website that you do not want to be accessed by search engines, or specify that search engines only include specified content
* The src directory stores project development files
* App.vue root component
* The api directory stores interface-related configurations and calling methods based on routing
* The common directory stores public methods and regular expressions
* The components directory stores common components for pages
* The src/config directory stores third-party class libraries and plug-in references other than the vue family bucket, currently including element-plus, filters, vue-i18n, mitt, and pinia references
* The reuqest directory stores axios secondary encapsulation
* The router directory stores routing configuration files
* The store directory stores the state management library pinia
* main.ts project entry ts
* assets public resource association: pictures, styles, fonts
* views directory stores page components divided by routes
* .env file configures global environment variables
* .eslintrsc file configures eslint detection code related rules
* .gitignore file specifies ignore rules
* shims-vue.d.ts This file is an adaptation definition file for typescript
* package.json configuration file
* tsconfig.json ts file configuration
* babel.config.js configures babel to parse and convert js code related rules, parse es6 optional chain syntax
* vue.config.js configures webpack, cross-domain related

## Local development
* npm run dev, the interface goes to the test environment interface, see the proxy configuration in dev.server.js for the specific url
* Or npm i —-legacy-peer-deps to avoid dependency version issues, npm run dev [use this command to specify the node version as 16.19.0, no need to switch]