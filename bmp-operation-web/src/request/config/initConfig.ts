/**
 * @file
 * @author
*/
import VueCookies from 'vue-cookies';

class InitConfig {
    cookie: {
        [x: string]: unknown;
        get?: Function;
        set?: Function;
    } = VueCookies;

    config = {
        baseURL: '/operation-web/', // 所有实例添加请求前缀
        timeout: 10 * 1000, // 超时时间
        headers: { // 请求头
            'Content-Type': 'application/json; charset=utf-8',
            'X-Jdcloud-Language': this.cookie?.get && this.cookie.get('X-Jdcloud-Language') || 'zh_CN'
        },
        // 设置是否允许携带cookie等凭证信息
        withCredentials: true,
        // 响应编码
        responseEncoding: 'utf8',
        // 设置重定向的最大次数
        maxRedirects: 5,
    };
}

export default InitConfig;
