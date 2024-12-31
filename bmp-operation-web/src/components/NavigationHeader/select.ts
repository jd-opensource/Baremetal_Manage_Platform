/**
 * @file
 * @author
*/

import {useI18n} from 'vue-i18n'; // 国际化
import {Router, useRouter} from 'vue-router';
import {locationItem, CurrentInstance} from 'utils/publicClass.ts';
import store from 'store/index.ts';

class NavBarHeader {
    email: Ref<string> = ref<string>('');
    router: Router = useRouter();
    locale: Ref<string> = ref<string>(useI18n().locale.value);
    // 退出登录弹窗状态
    loginOutDiaLog: Ref<boolean> = ref<boolean>(false);
    // 登录用户名
    loginUserName: Ref<string> = ref<string>('');
    instanceProxy = new CurrentInstance().proxy;
    hasNoRead: Ref<boolean> = ref<boolean>(false);

    // 构造器
    constructor () {
        this.#setInfo();
        this.watchRouter();
        this.#getHasNoRead();
    };

    /**
     * 监听路由变化
    */
    watchRouter = () => {
        
        /**
         * 监听路由path
         * @param {string} newValue 每次更新的路由的值
        */
        watch(() => this.router.currentRoute.value.path, (newValue: string) => {    
            if (newValue) {
                this.#setInfo();
                this.#getHasNoRead();
            }
        }, {deep: true});
    };

    #getHasNoRead = async () => {
        const res = await this.instanceProxy.$messageApi.hasUnreadMessageAPI({});
        this.hasNoRead.value = res?.data?.result?.hasUnread;
    }

    #setInfo = async () => {
        store.userPurviewInfo().promiseUserInfo()
        .then((res: {email: string; userName: string;}) => {
            if (res?.email?.length) {
                this.email.value = res.email;
                this.loginUserName.value = res.userName;
                return;
            }
            return Promise.reject();
        })
        .catch(() => {
            const email: string = localStorage?.getItem('email')?? '';
            const userName: string = localStorage?.getItem('userName')?? '';
            this.email.value = email ? window.atob(email) : '';
            this.loginUserName.value = userName ? window.atob(userName) : '';
        })
    };

    /**
     * 改变语言类型
     * @param {string} lang 语言类型
    */
    changeLang = (lang: string): void => {
        this.locale.value = lang;
        // if (lang === locationItem.cookie.get('X-Jdcloud-Language')) return;
        // setUserInfoAPI(
        //     {
        //         language:  this.locale.value
        //     }
        // )
        // .then(({data} : {data: any}) => {
        //     if (data?.result && Object.keys(data.result)?.length) {
        locationItem.cookie.set('X-Jdcloud-Language', this.locale.value);
        location.reload();
    };

    /**
     * 退出登录
     * @return {boolean} xxx 返回对应的弹窗状态
    */
    loginOut = (): boolean => {
        return this.loginOutDiaLog.value = !this.loginOutDiaLog.value;
    };

    /**
     * 退出登录取消操作
     * @param {boolean} type 弹窗状态
     * @return {boolean} xxx 返回对应的弹窗状态
    */
    loginOutCancel = (type: boolean): boolean => {
        return this.loginOutDiaLog.value = type;
    };

    userClick = (type: string) => {
        this.router.push(this.instanceProxy!.$defInfo.routerPath(type));
    };

    messageListClick = () => {
        this.router.push(this.instanceProxy!.$defInfo.routerPath('myMessage'));
    }

    jumpDefaultList = () => {
        const path: string = this.instanceProxy!.$defInfo.routerPath('idcList');
        this.router.push(path);
    };
};

export default NavBarHeader;
