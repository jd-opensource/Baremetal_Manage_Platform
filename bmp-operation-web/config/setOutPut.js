module.exports = (config) => {
    config.output.filename = '[name].[hash].js';
    config.output.chunkFilename = '[name].[hash].js';
    return config;
}