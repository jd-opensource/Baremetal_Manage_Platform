module.exports = {
    // ts编译之后，在进行一个babel编译
    'preset': '@vue/cli-plugin-unit-jest/presets/typescript-and-babel',
    'transform': {
        '^.+\\.vue$': 'vue-jest'
    }
};