import {defineStore} from 'pinia'; // 定义容器名
import storeName from 'store/storeName.ts'; // 容器名
import { language } from 'utils/publicClass.ts';
// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
const navigationBarStore = defineStore(storeName.navigationBar, {
    state: (): {
        hasNavigationBar: boolean;
    } => {
        return {
            hasNavigationBar: true
        }
    },

    actions: {

        /**
         * 侧边栏展开收起点击事件
         * @return {boolean} this.store.hasNavigationBar 状态取反
        */
        hasNavigationBarClick(): boolean {
            return this.hasNavigationBar = !this.hasNavigationBar;
        },

        tootlipDisabled(locationItem: {getLocationItem: string}, title: string) {
            if (!this.hasNavigationBar) return true;
            const tooltipData = [
                language.t('navigationBar.children.hardwareAlarmStatus'),
                language.t('navigationBar.children.myMessages'),
                language.t('navigationBar.children.messageSet'),
                language.t('navigationBar.children.resourceEcharts'),
                language.t('navigationBar.children.historyAlarmInfo'),
                language.t('navigationBar.children.allAlarmRules'),
                language.t('securitySettings.header.securitySettings'),
                language.t('apiKey.header.apiKey')
            ]
            return !(locationItem.getLocationItem === 'en_US' && tooltipData.indexOf(title) > -1)
        }
    }
});

export default navigationBarStore;
