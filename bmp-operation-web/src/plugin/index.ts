import CustomList from 'components/Public/CustomList/CustomList.vue'; // 自定义列表组件

type PluginType = {
    name: string;
    components: Function;
};

/**
 * 插件&插件名称
*/
const pluginName: PluginType[] = [
    {
        name: 'custom-list',
        components: CustomList
    }
];

export default {
    install: (app: {component: Function;}) => {
        for (const key of pluginName) {
            app?.component(key.name, key.components);
        }
    }
};