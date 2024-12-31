import {defineStore} from 'pinia';
import STORE_NAMES from 'store/storeName.ts';

interface StateType {
    loginUserName: string;
}
interface UserType {
    userId: string;
}

// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
export const loginInfoStore = defineStore(STORE_NAMES.LOGIN_INFO, {
    state: (): StateType => {
        return {
            loginUserName: ''
        }
    }
});

export const projectListStore = defineStore(STORE_NAMES.PROJECT_LIST, {
    state: (): UserType => {
        return {
            userId: ''
        }
    }
});
