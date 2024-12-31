// 这个插件可以看到启动时候的耗时模块,方便后续做优化使用
const SpeedMeasurePlugin = require('speed-measure-webpack-plugin');

module.exports = (config) => {
    config.plugins.push(
        new SpeedMeasurePlugin()
    );
    return config;
}
