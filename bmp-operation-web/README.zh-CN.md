# bmp私有化运营平台项目

[English](README.md) | 简体中文

## 简介
* 本工程基于@vue/cli 4搭建而成(vue3 + ts + pinia)

## 目录及文件功能简介
* config目录存放webpack相关配置
* public目录由@vue/cli默认创建生成，存放一些不需要经过webpack打包处理的静态资源
* public/robots.txt-是一个纯文本文件,可以声明该网站中不想被搜索引擎访问的部分，或者指定搜索引擎只收录指定的内容
* src目录存放项目开发文件
* App.vue根组件
* api目录存放根据路由划分接口相关配置及调用方法
* utils目录存放公共方法、正则表达式
* components目录存放页面通用组件
* lib目录存放除vue全家桶以外的第三方类库及插件引用，当前包含element-plus、filters、vue-i18n、mitt、pinia的引用
* reuqest目录存放axios二次封装
* router目录存放路由配置文件
* store目录存放状态管理库pinia
* main.ts工程入口ts
* assets公共资源关联：图片、样式、字体
* views目录存储根据路由划分的页面组件
* .env文件配置全局环境变量
* .eslintrsc文件配置eslint检测代码相关规则
* .gitignore文件 指定忽略规则
* shims-vue.d.ts 该文件是为了 typescript 做的适配定义文件
* package.json 配置文件
* tsconfig.json ts文件的一些配置
* babel.config.js配置babel解析转换js代码相关规则、解析es6可选链式语法
* jest.config.js jest单元测试相关
* tsconfig.json ts配置相关
* vue.config.js配置webpack、跨域相关
* node版本-依赖-建议10.16.0，启动-建议-16.19.0
* 或者npm i —-legacy-peer-deps 避免各依赖版本问题，npm run dev[用这个命令node版本指定为16.19.0即可，不用切换]

## 本地开发
* npm run dev, 接口走测试环境的接口，具体url见dev.server.js中的代理配置