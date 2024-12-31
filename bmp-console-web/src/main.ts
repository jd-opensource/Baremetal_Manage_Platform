import {createApp} from 'vue'; // 返回一个提供应用上下文的应用实例。应用实例挂载的整个组件树共享同一个上下文
import App from '@/App.vue'; // 根组件
import router from 'router/index.ts'; // 路由
import i18n from 'lib/i18n/index.ts'; // 国际化
import ElementPlus from 'lib/element/index.ts'; // elementui
import pinia from 'lib/pinia/pinia.ts'; // pinia-store状态库
import filter from 'lib/filters/filter.ts'; // 过滤方法
import mittCommunication from 'lib/mitt/mitt.ts'; // mitt-跨组件通信
import * as ElementPlusIconsVue from '@element-plus/icons-vue'; // Element-plus-icon使用
import VueCookies from 'vue-cookies';

const app = createApp(App);
// 全局挂载
app.config.globalProperties.$cookies = VueCookies;
filter(app);
mittCommunication(app);
app.use(router).use(pinia).use(ElementPlus).use(i18n).mount('#app');
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
};