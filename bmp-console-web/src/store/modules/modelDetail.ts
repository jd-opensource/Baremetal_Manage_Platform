import {defineStore} from 'pinia'; // 定义容器名
import {
    manageMentStatus // 管理状态
} from 'utils/constants.ts';
import STORE_NAMES from 'store/storeName.ts'; // 容器名

interface StateType {
    status: any;
    deviceSeries: any;
};

// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
const modelDetailStore = defineStore(STORE_NAMES.MODEL_DETAIL, {
    state: (): StateType => {
        return {
            status: manageMentStatus,
            deviceSeries: []
        }
    },
});

export default modelDetailStore;
