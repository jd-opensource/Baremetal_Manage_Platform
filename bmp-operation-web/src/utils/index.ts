import VueCookies from 'vue-cookies'; // cookie
import i18n from 'lib/i18n/index.ts'; // 国际化
import en from 'element-plus/lib/locale/lang/en'; // 英文
import zhCn from 'element-plus/lib/locale/lang/zh-cn'; // 中文
import {ElMessage} from 'element-plus';
import {CurrencyType} from './publicType';

const CryptoJS = require('crypto-js');

/**
 * 国际化1
*/
const {global} = i18n;

class MethodsTotal {
    channel = new BroadcastChannel('');

    sendMsg = (type: string, content: CurrencyType[]) => {
        this.channel.postMessage({
            type,
            content
        })
    }

    listenMsg = (fn: Function) => {
        const handle = (e: {data: CurrencyType[];}) => {
            fn && fn(e.data);
        }

        this.channel.addEventListener('message', (handle));
        return () => {
            this.channel.removeEventListener('message', handle);
        }
    }

    static deepCopy = (data: any[]) => {
        if (typeof data !== 'object' || !data) {
            return data;
        }
        // 判断是数组还是对象
        const result: any = data instanceof Array ? [] : {};
        for (const key in data) {
            // 只拷贝对象自身的属性
            if (data.hasOwnProperty(key)) {
                result[key] = this.deepCopy(data[key]);
            }
        }
        return result;
    };

    toLine(params: any) {
        const conversionData = [
            [
                (params: string) => typeof params === 'string',
                () => this.humpConversion(params)
            ],
            [
                (params: {[x: string]: string | number}) => params instanceof Object,
                () => {
                    const obj = [];
                    for (const index of Object.keys(params)) {
                        obj.push(
                            {
                                [this.humpConversion(index)]: params[index]
                            }
                        )
                    }
                    return Object.assign({}, ...obj);
                }
            ]
        ];

        for (const key of conversionData) {
            if (key[0](params)) {
                return key[1](params);
            }
        }
    }

    humpConversion = (params: string) => {
        return params.replace(/([A-Z])/g, '_$1').toLowerCase()
    };

    humpSplit = (params: string) => {
        return params.replace(/([A-Z])/g, '-$1').toLowerCase()
    };

    lineConverting (data: any) {
        if (typeof data !== 'object' || !data) return data;
        if (Array.isArray(data)) {
            return data.map((item): any => this.lineConverting(item))
        }
        const newData: any = {}
        for (let key in data) {
          const newKey = key.replace(/_([a-z])/g, (_, m) => m.toUpperCase())
          newData[newKey] = this.lineConverting(data[key])
        }
        return newData;
    };

    initScrollLeft = (value: any) => {
        value.$refs.bodyWrapper.getElementsByClassName('el-scrollbar__wrap')[0].scrollLeft = 0;
    };

    /**
     * 是否明文显示密码
     * @param {boolean} type 密码当前状态 false 加密 true 明文
     * @param {string} len 密码长度
     * @return {string} encryption.join('') 数组转字符串，','隔开 len 密码长度
    */
    hasShowPassword = (type: boolean, len: string): string => {
        if (!type) {
            if (!len) {
                return '*********';
            }
            const encryption: string[] = [];
            for (let i: number = 1; i <= len.length; i++) {
                encryption.push('*');
            }
            return encryption.join('');
        }
        return len;
    };

    /**
     * 16进制颜色转rgba
     * @param {string} color 十六进制颜色
     * @param {string} opacity 透明度
     * @return {string} xxx rgba颜色
    */
    conversionRgba = (color: string, opacity: string): string | null => {
        // exec() 方法用于检索字符串中的正则表达式的匹配。
        // 如果字符串中有匹配的值返回该匹配值，否则返回 null。
        let result: RegExpExecArray | null = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(color);
        if (!result) {
            return null;
        }
        return `rgba(${parseInt(result[1], 16)}, ${parseInt(result[2], 16)},${parseInt(result[3], 16)}, ${opacity})`;
    };

    /**
     * 复制内容
     * @param {string} text 需要复制的内容
    */
    copyInfo = (text: string) => {
        if (navigator?.clipboard) {
            // this.copyInfo = () => {
                navigator.clipboard?.writeText(text).then(() => {
                    msgTip.success(global.t('operate.copy'));
                });
            // }
            return;
        }
        // this.copyInfo = () => {
            const iptCopy: HTMLInputElement = document.createElement('input');
            iptCopy.value = text;
            document.body.appendChild(iptCopy);
            iptCopy.select();
            msgTip.success(global.t('operate.copy'));
            document.execCommand('Copy');
            iptCopy.remove();
        // }
        // this.copyInfo(text);
    };

    // 搜索，短时间内连续点击搜索按钮只执行一次搜索函数(节流)
    throttleFun = (fn: Function, wait = 500) => {
        const that = this;
        let last: number,
        now: number;
        return function () {
            now = Date.now()
            if (last && now - last < wait) {
                last = now
            } else {
                last = now
                fn.call(that, ...arguments)
            }
        }
    }

    // 监听搜索框内容的变化，等输入完成才会执行搜索函数(防抖)
    static debounceFun = (fn: Function, wait = 500) =>{
        let timer: number | null;
        const that = this;
        return function () {
            let context = that;
            let args = arguments;
            if (timer) clearTimeout(timer);
            timer = setTimeout(() => {
                fn.apply(context, args)
            }, wait)
        }
    }

    /**
     * 数组去重
     * @param {Array} arr 传入的数组
     * 
    */
    uniaueArray = (arr: any) => {
        if (!Array.isArray(arr)) {
            throw new Error('传入的不是数组');
        }
        const result = [];
        // 表示它现在代表的某个循环
        outer: for (const item of arr) {
            for (const r of result) {
                if (this.#equal(item, r)) {
                    // 结束outer的循环
                    continue outer;
                }
            }
            result.push(item);
        }
        return result;
    }

    isPrimitive = (val: unknown) => {
        return val === null || (typeof val !== 'object' && typeof val !== 'function');
    }

    #equal = <T extends ArrayLike<unknown>, K extends ArrayLike<unknown> | { [s: string]: unknown; }>(a: T, b: K) => {
        if (this.isPrimitive(a) || this.isPrimitive(b)) {
            return Object.is(a, b);
        }
        const ent1 = Object.entries(a);
        const ent2 = Object.entries(b);
        if (ent1.length !== ent2.length) {
            return false;
        }
        for (const [key, val] of ent1) {
            // @ts-ignore
            if (!this.equal(val, b[key])) {
                return false
            }
        }
        return true;
    }

    getLastFiveDaysTime = () => {
        const result = [];
        for (let i = 0; i < 5; i++) {
            // 计算一天毫秒数
            const timestamp = Date.now() - (1 * 24 * 60 * 60 * 1000) * i;
            result.unshift(timestamp);
        }
        return result;
    }

    beautifyLog = () => {
        const isPre = ['pre', 'prod'].includes(process.env.VUE_APP_ENV);
        const isEmpty = (value: string | null | undefined) => {
            return ['', null, void 0].includes(value);
        };
        const befautifyPrint = (title: string, text: string, color: string) => {
            if (isPre) return;
            if (typeof text !== 'string') {
                console.log(title, text);
                return;
            }
            console.log(
                `%c ${title} %c ${text} %c`,
                `background: ${color}; border: 1px solid ${color}; padding: 1px; border-radius: 2px 0 0 2px; color: #fff;`,
                `border: 1px solid ${color}; padding: 1px; border-radius: 0 2px 2px 0; color: ${color};`,
                `background: transparent`
            );
        };

        const info = (textInfo: string, content: string = '') => {
            setPublicInfo(textInfo, content, 'Info', '#909399');
        };

        const error = (textError: string, content: string = '') => {
            setPublicInfo(textError, content, 'Error', '#f56c6c');
        };

        const warning = (textWarning: string, content: string = '') => {
            setPublicInfo(textWarning, content, 'Warning', '#e6a23c');
        };

        const table = <T>(data: T) => {
            return console.table(data);
        }

        const success = (textSuccess: string, content: string = '') => {
            setPublicInfo(textSuccess, content, 'Success', '#67c23a');
        };

        const setPublicInfo = (typeVal: string, content: string = '', type: string, color: string) => {
            const title = isEmpty(content) ? type : typeVal;
            const text = isEmpty(content) ? typeVal : content;
            befautifyPrint(title, text, color);
        }

        return {
            info,
            table,
            error,
            warning,
            success
        }
    }
};

const methodsTotal = new MethodsTotal();

export {
    methodsTotal
}
export default MethodsTotal;


/**
 * 深拷贝
 * @param {Object | Array | number | string | null} target 需要拷贝的数据
 * @return {number | string | null | undefined} data 深拷贝的数据
 * @return {Object | Array} result 深拷贝的数据
*/
export function deepCopy(data: any[]) {
    if (typeof data !== 'object' || !data) {
        return data;
    }
    // 判断是数组还是对象
    const result: any = data instanceof Array ? [] : {};
    for (const key in data) {
        // 只拷贝对象自身的属性
        if (data.hasOwnProperty(key)) {
            result[key] = deepCopy(data[key]);
        }
    }
    return result;
};

/**
 * 获取cookie里面的语言
*/
const vueCookiesInfo = (VueCookies as any).get('X-Jdcloud-Language')?? 'zh_CN';

/**
 * 语言切换
 * @return {string} zhCn | en 中文|英文
*/
export const languageSwitch = (): string | unknown => {
    return vueCookiesInfo === 'zh_CN' ? zhCn : en;
};


/**
 * 实例状态文案颜色
 * @param {string} status 当前状态
 * @return {string} xxx.xxx 当前状态class类
*/
export const instanceStatusColor = (status: string) => {
    const i18n: boolean = vueCookiesInfo === 'zh_CN';
    const defaultClass: string = 'def-status';
    const enDefaultClass: string = 'en-def-status';
    const statusColorData = [
        [
            (status: string) => !status?.length,
            () => {return 'set-empty'}
        ],
        [
            (status: string) => ['created', 'Running', 'running'].indexOf(status) > -1,
            () => {return i18n ? `${defaultClass} set-status-create` : `${enDefaultClass} set-status-create-en`}
        ],
        [
            (status: string) => ['in', 'putaway', 'removed'].indexOf(status) > -1,
            () => {return i18n ? `${defaultClass} set-status-warehousing` : `${enDefaultClass} set-status-warehousing-en`}
        ],
        [
            (status: string) => ['error', 'Error', 'Mount failed', 'Creation failed', 'Startup failed', 'Shutdown failed', 'Reboot failed', 'Delete failed', 'Reinstall failed', 'Resetpasswd failed'].indexOf(status) > -1,
            () => {return i18n ? `${defaultClass} set-status-error` : `${enDefaultClass} set-status-error-en`}
        ],
        [
            (status: string) => ['Shutting Down', 'destroyed', 'Destroyed', 'stopped'].indexOf(status) > -1,
            () => {return i18n ? `${defaultClass} set-status-default` : `${enDefaultClass} set-status-default-en`}   
        ],
        [
            (status: string) => status,
            () => {return i18n ? `${defaultClass} set-status-on-down` : `${enDefaultClass} set-status-on-down-en`}
        ]
    ];
    for (const key of statusColorData) {
        if (key[0](status)) {
            return key[1](status);
        }
    }
};

/**
 * 管理状态文案颜色
 * @param {string} status 当前状态
 * @return {string} xxx.xxx 当前状态class类
*/
export const managementStatusColor = (status: string) => {
    const defaultClass: string = 'def-management-status';
    const statusColorData = [
        [
            (status: string) => !status?.length,
            () => {return 'set-empty'}
        ],
        [
            (status: string) => ['created'].indexOf(status) > -1,
            () => {return `${defaultClass} set-status-create`}
        ],
        [
            (status: string) => ['in', 'putaway', 'removed'].indexOf(status) > -1,
            () => {return `${defaultClass} set-status-warehousing`}
        ],
        [
            (status: string) => ['putawayfail'].indexOf(status) > -1,
            () => {return `${defaultClass} set-status-error`}
        ],
        [
            (status: string) => ['putawaying'].indexOf(status) > -1,
            () => {return `${defaultClass} set-status-on-down`}
        ]
    ];
    for (const key of statusColorData) {
        if (key[0](status)) {
            return key[1](status);
        }
    }
};

/**
 * 实例状态
 * @param {string} status 实例状态
 * @return {string} xxx 实例状态名称
*/
export const instanceStatus = (status: string) => {
    const instanceStatusColorData = [
        [
            (status: string) => ['creating', 'Creating'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.creating')
        ],
        [
            (status: string) => ['starting', 'Starting up'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.starting')
        ],
        [
            (status: string) => ['running', 'Running'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.running')
        ],
        [
            (status: string) => ['stopping', 'Shutting Down'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.stopping')
        ],
        [
            (status: string) => ['stopped', 'Stopped'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.stopped')
        ],
        [
            (status: string) => ['restarting', 'Restarting', 'Rebooting'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.restarting')
        ],
        [
            (status: string) => ['resetting_password', 'Reset Password', 'Resetting password', 'resettingPassword'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.resetting_password')
        ],
        [
            (status: string) => ['destroying', 'Destructing', 'Destroying'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.destroying')
        ],
        [
            (status: string) => ['destroyed', 'Destroyed'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.destroyed')
        ],
        [
            (status: string) => ['error', 'Error', 'Creation failed', 'Startup failed', 'Shutdown failed', 'Reboot failed', 'Delete failed'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.error')
        ],
        [
            (status: string) => ['upgrading', 'Upgrading'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.upgrading')
        ],
        [
            (status: string) => ['reinstalling', 'Reinstalling System'].indexOf(status) > -1,
            () => global.t('deviceList.instanceStatus.reinstalling')
        ]
    ];

    for(const key of instanceStatusColorData) {
        if (key[0](status)) {
            return key[1](status);
        }
    }
}

export const tableClass = computed(() => (table: {[x: string]: string}[], isLoading: boolean): string => {
    const defaultClass: string = 'operate-management-count-short-table';
    if (table?.length) {
        if (isLoading) {
            return `${defaultClass} img-show`;
        }
        return 'operate-management-count-table';
    }
    else {
        if (isLoading) {
            return `${defaultClass} img-show`;
        }
        return `${defaultClass} img-empty`;
    }
});

/**
* 设置文字class类
* @param {boolean} type 当前的class是否高亮 true 是 false 否
* @param {string} hasRight right class类
* @return {string} xxx 返回对应的class类
*/
export const setTextClass: ComputedRef<{}> = computed(() => (type: boolean, hasRight = ''): string => {
    if (type) {
        return `currency-style ${hasRight}`;
    }
    return 'default-style';
});

/**
 * 设置宽度
 * @param {string} width1 宽度1
 * @param {string} width2 快读2
 * @param {Array} data 表格数据
 * @return {string} width1 width2 宽度1、宽度2
*/
export const setWidth: ComputedRef<{}> = computed(() => (width1: string, width2: string, data: unknown[]): string => {
    if (data?.length) {
        return width1;
    }
    return width2;
});

/**
 * 过滤重复数据
 * @param {Array} data 需要过滤的数据
 * @return {Array} newArr 过滤完毕的数据
*/
export const filterData = (data: CurrencyType[], filterType: string) => {
    const newArr: CurrencyType[] = []
    const map: Map<unknown, unknown> = new Map();
    for (const item of data) {
        if (!map.has(item[filterType])) {
            map.set(item[filterType], true);
            newArr.push(
                {
                    [filterType]: item[filterType],
                    ...item
                }
            );
        }
    }
    return newArr;
};

/**
 * 设置不同的样式
 * @param {string} className1 样式1
 * @param {string} className2 样式2
 * @return {string} className1 || className2 根据判断，返回className1 || className2
*/
export const setDiffClass: ComputedRef<{}> = computed(() => (className1: string, className2: string): string => {
    if (vueCookiesInfo === 'en_US') {
        return className1;
    }
    return className2;
});

export const hasShowTooltip = (event: Event, row: {showTooltip: boolean}, text: string, num: number = 2.3) =>{
    // event为鼠标移入时的事件对象
    // textWidth 为文本在页面中所占的宽度，创建标签，加入到页面，获取textWidth ,最后在移除
    const createSpan: HTMLSpanElement = document.createElement('span');
    // 文本内容
    createSpan.innerText = text;
    // 文本class类
    createSpan.className = 'getTextWidth'
    // body里添加
    document?.querySelector('body')?.appendChild(createSpan);
    // 类型断言，获取文本offsetWidth
    const textWidth: number = (document?.querySelector('.getTextWidth') as HTMLElement)?.offsetWidth;
    // 移除
    document?.querySelector('.getTextWidth')?.remove();
    // tableWidth为表格容器的宽度、类型断言
    const tableWidth: number = (event?.target as HTMLElement)?.offsetWidth;
    // 当文本宽度小于等于容器宽度两倍时，这块会有点小误差，根据业务来换算，代表文本显示未超过两行
    textWidth <= Math.ceil(num * tableWidth) ? row.showTooltip = false : row.showTooltip = true;
};

/**
 * 获取日期
 * @return {string} xxx 年-月-日
*/
export const getDate = (): string => {
    const year: number = new Date().getFullYear();
    const month: number = new Date().getMonth() + 1;
    const day: number = new Date().getDate();
    return `${year}-${month}-${day}`
};

/**
 * 获取日期
 * @return {string} xxx 年-月-日
*/
export const getDateParams = (params: string): string => {
    const year: number = new Date(params).getFullYear();
    const month: number = new Date(params).getMonth() + 1;
    const day: number = new Date(params).getDate();
    return `${year}-${month}-${day}`
};

/**
 * 获取日期
 * @return {string} xxx 年-月-日
*/
export const getDateParamsNum = (params: number): string => {
    const newParams: number = params * 1000;
    const year: number = new Date(newParams).getFullYear();
    const month: number = new Date(newParams).getMonth() + 1;
    const day: number = new Date(newParams).getDate();
    return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`
};

/**
 * 计算两个时间之间的时间差 多少天时分秒
 * @param {number} startTime 开始时间
 * @param {number} endTime 结束时间
 * @return {string} days hours 计算出相差天数 计算出小时数
*/
export const intervalTime = (startTime: number, endTime: number) => {
    // 开始时间
    const date1 = startTime / 1000;
    // 结束时间
    const date2 = endTime / 1000;
    // 时间差的毫秒数
    const date3 =  (date2- date1) * 1000;
    // 计算出相差天数
    const days = Math.floor(date3 / (24 * 3600 * 1000));
    // 计算出小时数
    // 计算天数后剩余的毫秒数
    const leave1 = date3 % (24 * 3600 * 1000);
    const hours = Math.floor(leave1 / (3600 * 1000));
    return {
        days,
        hours
    }
}


/**
 * 计算两个时间之间相差多少天
 * @param {number} preTime 开始时间戳
 * @param {number} endTime 结束时间戳
*/
export const returnInt = (startTime: number, endTime: number) => {
    // 开始日期
    const startDate = new Date(startTime);
    // 结束日期
    const endDate = new Date(endTime);
    // if (endDate.getTime() > startDate.getTime()) {
    //     return '';
    // }
    // Math.abs 返回数的绝对值，防止出错
    const diffDate = Math.abs(endDate.getTime() - startDate.getTime());
    // 1000 毫秒换算为1s
    return Math.ceil(diffDate / (1000 * 60 * 60 * 24));
}
   
/**
 * 获取日期时分秒
 * @return {string} xxx 年-月-日-时分秒
*/
export const getDateMinutes = <T extends number | undefined>(params: T): string => {
    if ([undefined, null].includes((params as undefined))) return '--';
    const newParams: number = (params as number) * 1000;
    const year: number = new Date(newParams).getFullYear();
    const month: number = new Date(newParams).getMonth() + 1;
    const day: number = new Date(newParams).getDate();
    const hours: number = new Date(newParams).getHours();
    const minutes: number = new Date(newParams).getMinutes();
    const mill: number = new Date(newParams).getSeconds();
    return `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day} ${hours < 10 ? '0' + hours : hours}:${minutes < 10 ? '0' + minutes : minutes}:${mill < 10 ? '0' + mill : mill}`
};

/**
 * 获取日期时分
 * @return {string} xxx 月-日-时分
*/
export const getMonthHoursMinutes = <T extends number | undefined>(params: T): string => {
    if ([undefined, null].includes((params as undefined))) return '--';
    const newParams: number = (params as number) * 1000;
    // const year: number = new Date(newParams).getFullYear();
    const month: number = new Date(newParams).getMonth() + 1;
    const day: number = new Date(newParams).getDate();
    const hours: number = new Date(newParams).getHours();
    const minutes: number = new Date(newParams).getMinutes();
    // const mill: number = new Date(newParams).getSeconds();
    return `${hours < 10 ? '0' + hours : hours}:${minutes < 10 ? '0' + minutes : minutes} ${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`
};

/**
 * 生成随机数字
 * @param {number} num 随机数位数
 * @return {string} xxx 生成的随机数
*/
export const generateRandomNum = (num: number): string => {
    let randomNum: string = ''
    for(let i: number = 0; i < num; i++) {
        randomNum += String(Math.floor(Math.random() * 10));
    }
    return randomNum;
};

/**
 * AES加密
*/
export const encrypt = (str: string, keyStr = 'ABCDEFGHIJKLMNOP', ivStr = '0123456789ABCDEF') => {
    const key: string = CryptoJS.enc.Utf8.parse(keyStr);
    const iv: string = CryptoJS.enc.Utf8.parse(ivStr);
    const srcs: string = CryptoJS.enc.Utf8.parse(str);
    const encrypt: {ciphertext: string} = CryptoJS.AES.encrypt(srcs, key, {
        iv,
        mode: CryptoJS.mode.CBC, // AES加密模式
        padding: CryptoJS.pad.Pkcs7
    });
    return CryptoJS.enc.Base64.stringify(encrypt.ciphertext);
};

/**
 * AES解密
*/
export const decrypt = (str: string, keyStr = 'ABCDEFGHIJKLMNOP', ivStr = '0123456789ABCDEF') => {
    const key: string = CryptoJS.enc.Utf8.parse(keyStr);
    const iv: string = CryptoJS.enc.Utf8.parse(ivStr);
    const base64: string = CryptoJS.enc.Base64.parse(str);
    const src: string = CryptoJS.enc.Base64.stringify(base64);
    const decrypt = CryptoJS.AES.decrypt(src, key, {
        iv,
        mode: CryptoJS.mode.CBC, // AES解密模式
        padding: CryptoJS.pad.Pkcs7
    });
    return decrypt.toString(CryptoJS.enc.Utf8).toString();
};

export const tableScroll = (arg0: number, tableMaxHeight: {value: number}) => {
    const titleNav: HTMLElement | null = document!.querySelector('.operate-header');
    const windowHeight: number = document.documentElement.clientHeight;
    nextTick(() => {
        tableMaxHeight.value = windowHeight - arg0 - titleNav!.offsetHeight;
    });
};

export const msgTip = ElMessage;

export const customTip = (type: never, message: string, duration: number, fn: Function) => {
    ElMessage({
        type,
        message,
        duration,
        onClose: () => fn()
    });
};
