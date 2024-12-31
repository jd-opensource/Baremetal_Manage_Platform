/**
 * @file
 * @author
*/

import VueCookies from 'vue-cookies';
import router from '@/router/index.ts'; // 路由

class Error402 {
    cookie: {
        [x: string]: unknown;
        get?: Function;
        set?: Function;
    } = VueCookies;

    constructor() {
        this.#init();
    }

    #init = () => {
        sessionStorage?.removeItem('roleId');
        localStorage.clear();
        // 默认地址
        const defaultPath: string = '/IdcManagement/idcList';
        const lastPath: string = router?.currentRoute?.value?.fullPath;
        const path: string = lastPath === '/Login/login' ? defaultPath : lastPath;
        // 跳转地址-用户最后一次访问的地址
        const pathUrl: string = window.btoa(path || defaultPath);
        // 设置地址-存储到cookie
        this.cookie?.set && this.cookie.set('path_url', pathUrl);
        window.open('/Login/login', '_self');
    }
}

export default Error402;
