import {defineStore} from 'pinia'; // 定义容器名
import {encrypt} from 'utils/index.ts';
import {userInfoAPI} from 'api/userCenter/request.ts';
import storName from 'store/storeName.ts'; // 容器名
import {CurrencyType} from '@utils/publicType';
import {locationItem} from 'utils/publicClass.ts';
// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
const userPurviewStore = defineStore(storName.userInfo, {
    state: (): {userPurview: string;} => {
        return {
            userPurview: '',
        }
    },

    actions: {

        /**
         * 获取用户权限
        */
        getUserPurview (api: {userPurviewAPI: Function}) {
            return new Promise(async (resolve, reject) => {
                try {
                    const userPurviewApi = await api.userPurviewAPI();
                    if (userPurviewApi?.data?.result?.role?.roleId) {
                        this.userPurview = userPurviewApi.data.result.role.roleId;
                        localStorage.setItem('userPurview', encrypt(this.userPurview))
                        return resolve(userPurviewApi.data.result.role.roleId);
                    }
                    throw new Error('');
                }
                catch {
                    this.userPurview = '';
                    localStorage.setItem('userPurview', encrypt(this.userPurview))
                    return reject();
                }
                finally {
                    // this.userInfo();
                }
            })
        },

        userInfo() {
            userInfoAPI({})
            .then(({data} : {data: {result: CurrencyType}}) => {
                if (data?.result && Object.keys(data.result).length) {
                    const email: string = data?.result?.email?? '';
                    localStorage.setItem('email', window.btoa(email));
                    const userName = window.btoa(data.result.userName);
                    localStorage.setItem('userName', userName);
                    locationItem.cookie.set('X-Jdcloud-Language', data.result.language);
                }
            })
            .catch(() => {
                throw new Error('异常');
            })
        },

        promiseUserInfo() {
            return new Promise((resolve, reject) => {
                return userInfoAPI({})
                    .then(({data}: {data: {result: CurrencyType}}) => {
                        return resolve(data.result);
                    })
                    .catch(() => {
                        return reject('异常');
                    })
            })
        }
    }
});

export default userPurviewStore;
