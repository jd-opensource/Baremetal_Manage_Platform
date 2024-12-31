const components = require('unplugin-vue-components/webpack');
const autoImport = require('unplugin-auto-import/webpack');

module.exports = config => {
    config.plugins.push(
        components({
            dirs: ['./src/components'],
            extensions: ['vue'],
            dts: './src/components.d.ts'
        }),
        autoImport({
            imports: ['vue'],
            dts: './src/auto-imports.d.ts'
        })
    );
    return config;
};
