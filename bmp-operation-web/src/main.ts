import App from '@/App.vue'; // 根组件
import router from 'router/index.ts'; // 路由
import i18n from 'lib/i18n/index.ts'; // 国际化
import ElementPlus from 'lib/element/index.ts'; // elementui
import pinia from 'lib/pinia/pinia.ts'; // pinia-store状态库
import defInfo from 'lib/default/defInfo.ts'; // 默认
import filter from 'lib/filters/filter.ts'; // 过滤方法
import mittCommunication from 'lib/mitt/mitt.ts'; // mitt-跨组件通信
import * as ElementPlusIconsVue from '@element-plus/icons-vue'; // Element-plus-icon使用
import customUi from '@/ui/utils.ts'; // 自定义封装的ui
import plugin from 'plugin/index.ts'; // 插件
import request from 'lib/request/request.ts';
// import './upDate';

// 创建app实例
// createApp-返回一个提供应用上下文的应用实例。应用实例挂载的整个组件树共享同一个上下文
const app = createApp(App);

defInfo(app);

request(app);

// 仿过滤器
filter(app);

// 跨组件通信
mittCommunication(app);

// ui挂载全局
customUi.install(app);

// Element-plus-icon
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    // 挂载，key-名称，component-组件
    app.component(key, component)
};

// 创建实例后，通过链式调用方法
app.use(plugin).use(router).use(i18n).use(pinia).use(ElementPlus).mount('#app');
