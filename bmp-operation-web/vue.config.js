// 页面路径
const setAlias = require('./config/setAlias');
// 跨域配置
const devServer = require('./config/devServer');

const setOutPut = require('./config/setOutPut');

// 避免文件的频繁变更导致浏览器缓存失效，更好的利用缓存。提升用户体验
const runtimeChunk = require('./config/runtimeChunk');
// 可以看到启动时的消耗模块，方便后续做优化
const speedMeasure = require('./config/speedMeasure');
// prod环境移除console
// const terserPlugin = require('./config/terserPlugin');
// 减少不必要的解析
const noParse  = require('./config/noParse');

const publicConfig = require('./config/plugin');

const devTool = require('./config/devTool');

// devtool: 'inline-source-map'
module.exports = {
    publicPath: '/',
    productionSourceMap: false, // 是否在构建生产包时生成 sourceMap 文件，false将提高构建速度
    outputDir: 'dist',
    // 放置生成的静态资源 (js、css、img、fonts) 的 (相对于 outputDir 的) 目录
    assetsDir: 'static',
    // 指定生成的 index.html 的输出路径 (相对于 outputDir),也可以是一个绝对路径
    indexPath: 'index.html',
    // 生成的静态资源在它们的文件名中包含了 hash 以便更好的控制缓存
    filenameHashing: true,
    // 修改默认的webpack的配置
    configureWebpack: config => {
        setOutPut(config);
        runtimeChunk(config);
        speedMeasure(config);
        noParse(config);
        publicConfig(config);
        devTool(config);
        
    },
    // 调整内部的webpack配置
    chainWebpack: config => {
        config.plugin('html').tap(args => {
            args[0].unknownContextCritical = false;
            return args;
        });
        config.resolve.symlinks(true); // 修复热更新失效
        // 配置alias别名
        setAlias(config);
        // 删除懒加载模块的prefetch，降低带宽压力
        config.plugins.delete('prefetch');
        config.module
        .rule('images')
            .set('parser', {
                dataUrlCondition: {
                    maxSize: 30 * 1024
                }
            }
        );
    },
    ...devServer,
    css: {
        loaderOptions: {
            scss: {
                prependData: '@import "~@/assets/css/communal.scss";'
            }
        }
    }
};
