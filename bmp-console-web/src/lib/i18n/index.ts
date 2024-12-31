import {createI18n, I18n} from 'vue-i18n';
import zh from 'lib/i18n/zh-CN.ts';
import en from 'lib/i18n/en-US.ts';
import zhLocale from 'element-plus/lib/locale/lang/zh-cn';
import enLocale from 'element-plus/lib/locale/lang/en';
import VueCookies from 'vue-cookies';

/**
 * 配置对象-国际化
*/
const messages = {
    en_US: {
        ...en,
        ...enLocale
    },
    zh_CN: {
        ...zh,
        ...zhLocale
    }
};

/**
 * 获取语言类型
 * @return {string} key 语言种类
 * @return {string} isCookie 语言种类
 * @return {string} zh_CN 语言种类
*/
const getLocale = () => {
    const isCookie = (VueCookies as any).get('X-Jdcloud-Language');
    // 初次进入页面，cookie里面还未获取到值
    if (isCookie === null) {
        // 浏览器语言
        const language: string = navigator.language.toLocaleLowerCase();
        // message语言
        const locale: string[] = Object.keys(messages);
        for (const key of locale) {
            // 如果浏览器语言类型包括message里面的类型
            if (language.indexOf(key) > -1) {
                // 返回
                return key;
            }
        }
    }
    // 有cookie，优先用cookie里面的
    else if (isCookie) {
        return isCookie;
    }
    // 默认语言
    return 'zh_CN';
};

/**
 * 首先从缓存里拿，没有的话就用浏览器语言
*/
const i18n: I18n<any, unknown, unknown, false> = createI18n({
    legacy: false, // 使用 vue3 组合式API 时必须设置为 false
    locale: getLocale(), // 本地语言
    globalInjection: true, // 全局注入 $t 函数
    fallbackLocale: 'zh_CN', // 设置备用语言
    messages // 配置对象
});

export default i18n;
