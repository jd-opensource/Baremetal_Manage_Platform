import VueCookies from 'vue-cookies'; // cookie
import { ref, onUnmounted } from 'vue';
import en from 'element-plus/lib/locale/lang/en'; // 英文
import zhCn from 'element-plus/lib/locale/lang/zh-cn'; // 中文
const CryptoJS = require('crypto-js');

/**
 * AES加密
*/
export const encrypt = (str: string, keyStr: string, ivStr: string) => {
    let key: string = CryptoJS.enc.Utf8.parse(keyStr);
    let iv: string = CryptoJS.enc.Utf8.parse(ivStr);
    let srcs: string = CryptoJS.enc.Utf8.parse(str);
    let encrypt: {ciphertext: string} = CryptoJS.AES.encrypt(srcs, key, {
        iv,
        mode: CryptoJS.mode.CBC, // AES加密模式
        padding: CryptoJS.pad.Pkcs7
    });
    return CryptoJS.enc.Base64.stringify(encrypt.ciphertext);
}

/**
 * AES解密
*/
export const decrypt = (str: string, keyStr: string, ivStr: string) => {
    let key: string = CryptoJS.enc.Utf8.parse(keyStr);
    let iv: string = CryptoJS.enc.Utf8.parse(ivStr);
    let base64: string = CryptoJS.enc.Base64.parse(str);
    let src: string = CryptoJS.enc.Base64.stringify(base64);
    let decrypt = CryptoJS.AES.decrypt(src, key, {
        iv,
        mode: CryptoJS.mode.CBC, // AES解密模式
        padding: CryptoJS.pad.Pkcs7
    });
    return decrypt.toString(CryptoJS.enc.Utf8).toString();
};

/**
 * 深拷贝
 * @param {Object | Array | number | string | null} target 需要拷贝的数据
 * @return {number | string | null | undefined} data 深拷贝的数据
 * @return {Object | Array} result 深拷贝的数据
*/
export function deepCopy(data: {[x: string]: any; hasOwnProperty: (arg0: string) => any} | null) {
    if (typeof data !== 'object' || data === null) {
        return data;
    }
    // 判断是数组还是对象
    const result: [] | {} = data instanceof Array ? [] : {};
    for (const key in data) {
        // 只拷贝对象自身的属性
        if (data.hasOwnProperty(key)) {
            (result as any)[key] = deepCopy(data[key]);
        }
    }
    return result;
}

/**
 * 获取cookie里面的语言
*/
const vueCookiesInfo = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * 语言切换
 * @return {string} zhCn | en 中文|英文
*/
export const languageSwitch = (): string | unknown => {
    return vueCookiesInfo === 'zh_CN' ? zhCn : en;
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

export const hasShowTooltip = (event: Event, row: {showTooltip: boolean}, text: string, num: number = 0, type: string) =>{
    // event为鼠标移入时的事件对象
    // textWidth 为文本在页面中所占的宽度，创建标签，加入到页面，获取textWidth ,最后在移除
    const createSpan: HTMLSpanElement = document.createElement('span');
    // 文本内容
    createSpan.innerText = text;
    // 文本class类
    createSpan.className = 'getTextWidth';
    // body里添加
    document?.querySelector('body')?.appendChild(createSpan);
    // 类型断言，获取文本offsetWidth
    const textWidth: number = (document?.querySelector('.getTextWidth') as HTMLElement)?.offsetWidth;
    // 移除
    document?.querySelector('.getTextWidth')?.remove();
    // tableWidth为表格容器的宽度、类型断言
    const tableWidth: number = (event?.target as HTMLElement)?.offsetWidth;
    // 列表页中tableWidth固定，详情页不固定
    if(type === 'detail') {
        tableWidth <= num ? row.showTooltip = false : row.showTooltip = true;
    }else{
        // 当文本宽度小于等于容器宽度两倍时，这块会有点小误差，根据业务来换算，代表文本显示未超过两行
        textWidth <= Math.ceil(num * tableWidth) ? row.showTooltip = false : row.showTooltip = true;
    }    
};

/**
 * 过滤重复数据
 * @param {Array} data 需要过滤的数据
 * @return {Array} newArr 过滤完毕的数据
*/
export const filterData = (data: any[], filterType: string) => {
    const newArr: string[] = []
    const map: Map<unknown, unknown> = new Map();
    for(let i: number = 0; i < data.length; i++){
        // 如果 map里面不包含，就设置进去
        if (!map.has(data[i][filterType])) {
            map.set(data[i][filterType], true);
            newArr.push(
                {
                    [filterType]: data[i][filterType],
                    ...data[i]
                }
            );
        }
    };
    return newArr;
};

/**
 * 防抖函数
 * @param {Any} fn 需要加上防抖的函数
 * @param {Any} delay 延迟的时间
*/
export function debounce(fn:any, delay:any) {
    let timer:any = null; // 形成闭包
        return function () {
            if (timer) {
                clearTimeout(timer); // 防抖
            }
            timer = setTimeout(() => {
                fn(); // 执行传入的函数
            }, delay);
        };
}

/**
 * 合并两个对象，如果两个对象中有相同的键，且第二个对象中对应的值为非空，
 * 则使用第二个对象中的值覆盖第一个对象中的空值。
 * 
 * @param {T} obj1 第一个对象
 * @param {U} obj2 第二个对象，其值可能会覆盖第一个对象中的空值
 * @returns {T & U} 返回一个新的对象，包含合并后的键值对
 */
 export const mergeObjectsWithNonEmptyOverride = <T extends object, U extends object>(obj1: T, obj2: U): T & U => {
    // 创建一个新对象，初始包含obj1的所有键值对
    const result: Partial<T & U> = { ...obj1 };
  
    // 遍历obj2的键值对
    Object.entries(obj2).forEach(([key, value]) => {
      // 检查obj2中的值是否为非空（不是空字符串、null或undefined）
      const isValueNonEmpty = value !== "" && value !== null && value !== undefined;
  
      // 获取result中对应键的值
      const originalValue = result[key as keyof typeof result];
  
      // 检查原对象中的值是否为空（空字符串、null或undefined）
      // 首先确保originalValue是一个字符串，然后检查是否为空字符串
      const isOriginalValueEmpty = typeof originalValue === 'string' && originalValue === "" || originalValue === null || originalValue === undefined;
  
      // 如果obj2中的值为非空，或者obj1中对应的键的值为空，则使用obj2的值
      if (isValueNonEmpty || isOriginalValueEmpty) {
        result[key as keyof typeof result] = value;
      }
    });
  
    // 返回合并后的新对象
    return result as T & U;
};

type TimeStrings = [string, string];
interface TimeObject {
  [key: string]: number;
}

/**
 * 将包含两个日期字符串的数组转换为一个对象，对象包含两个键值对，
 * 键由参数指定，值为日期字符串对应的时间戳。
 * 
 * @param operationTimeStrings - 包含两个日期字符串的数组。
 * @param startKey - 转换后对象的起始时间键名。
 * @param endKey - 转换后对象的结束时间键名。
 * @returns 返回一个对象，包含起始和结束时间的时间戳。
 */
 export const convertStringsToTimestampObject = (operationTimeStrings: TimeStrings, startKey: string, endKey: string): TimeObject => {
    if (operationTimeStrings.length !== 2) {
        throw new Error('operationTimeStrings数组必须恰好包含两个元素。');
    }

    const [startString, endString] = operationTimeStrings;
    const startTime = new Date(startString);
    const endTime = new Date(endString);

    if (isNaN(startTime.getTime()) || isNaN(endTime.getTime())) {
        throw new Error('提供的日期字符串无效。');
    }

    return {
        [startKey]: startTime.getTime() /1000,
        [endKey]: endTime.getTime()/1000,
    };
};

/**
 * 创建一个延时动作，该动作在指定延时后执行，并且如果在延时期间再次被触发，则重新计时。
 * 
 * @param action - 要执行的动作，一个无参数的函数。
 * @param delay - 延时时间，单位为毫秒。
 * @returns 一个触发延时动作的函数。
 */
export const useDebouncedAction = (action: () => void, delay: number) => {
    // 存储当前的延时器 ID
    const timeoutId = ref<number | undefined>(undefined);

    /**
     * 清除当前的延时器。
     */
    const clearTimer = () => {
        if (timeoutId.value) {
            clearTimeout(timeoutId.value);
            timeoutId.value = undefined;
        }
    };

    /**
     * 触发延时动作。如果在延时期间再次被触发，则重新计时。
     */
    const trigger = () => {
        clearTimer();
        timeoutId.value = window.setTimeout(() => {
            action();
            clearTimer();
        }, delay);
    };

    // 组件卸载时清除延时器
    onUnmounted(clearTimer);

    return trigger;
};

/**
 * 将日期字符串数组转换为一个对象，该对象包含 'start' 和 'end' 属性。
 * 这些属性代表日期字符串中的时间部分（HH:MM:SS）。
 *
 * @param dateStrings - 需要转换的日期字符串数组。
 * @returns 一个对象，包含从输入日期中提取的 'start' 和 'end' 时间部分。
 */
export const createTimeObjectFromDateStringArray = (dateStrings: string[]): { start: string, end: string } => {
    // 如果dateStrings为null或undefined，返回start和end为空
    if (!dateStrings || dateStrings.length === 0) {
        return { start: '', end: '' };
    }
    /**
     * 将日期字符串解析为 Date 对象。
     *
     * @param dateString - 需要解析的日期字符串。
     * @returns 从输入字符串解析得到的 Date 对象。
     */
    const parseDate = (dateString: string): Date => new Date(dateString);
  
    /**
     * 从 Date 对象中提取时间部分。
     *
     * @param date - 需要从中提取时间的 Date 对象。
     * @returns 代表 Date 对象时间部分（HH:MM:SS）的字符串。
     */
    const getTimeFromDate = (date: Date): string => date.toTimeString().split(' ')[0];
  
    // 将输入的日期字符串转换为 Date 对象数组
    const dates = dateStrings.map(parseDate);
  
    // 返回一个包含日期 'start' 和 'end' 时间部分的对象
    return {
      start: getTimeFromDate(dates[0]),
      end: getTimeFromDate(dates[1])
    };
  };


  
type Item = {
// 假设这是你的对象的类型，可以根据实际情况添加其他属性
[key: string]: any;
};

  /**
 * 给对象数组中的每个对象添加一个独一无二的ID。
 * ID生成策略是结合当前时间戳和数组索引。
 * 
 * @param items - 待处理的对象数组。
 * @returns 返回一个新的数组，每个对象都包含一个唯一的ID。
 */
export const addUniqueId = (items: Item[]): Item[] => {
    return items.map((item, index) => ({
      ...item,
      id: `unique_${Date.now()}_${index}` // 使用当前时间戳和索引生成一个简单的唯一ID
    }));
};

// 假设的Item类型定义，这里我们假设每个对象都可能有一个id属性
type IdItem = {
    id?: string;
    [key: string]: any;
};

/**
 * 移除对象数组中每个对象的id属性。
 * 
 * @param items - 包含id属性的对象数组。
 * @returns 返回一个新的数组，每个对象都不包含id属性。
 */
export const removeIdFromItems = (items: IdItem[]): Item[] => {
    return items.map(({ id, ...rest }) => rest);
};
  
/**
 * 将包含方括号的时间字符串转换为 Date 对象数组。
 * 
 * @param startTime - 开始时间字符串，格式为 'HH:MM:SS'。
 * @param endTime - 结束时间字符串，格式为 'HH:MM:SS'。
 * @returns 返回一个包含两个 Date 对象的数组，分别对应开始时间和结束时间。
 */
export const convertBracketedTimeStringsToDateArray = (startTime: string, endTime: string): Date[] => {
    // 移除时间字符串中的方括号并解析时间
    const parseTime = (time: string) => {
        const [hours, minutes, seconds] = time.split(':').map(Number);
        return new Date(0, 0, 0, hours, minutes, seconds);
    };
  
    // 使用 parseTime 函数转换开始和结束时间
    const startDate = parseTime(startTime);
    const endDate = parseTime(endTime);
  
    // 返回包含两个日期对象的数组
    return [startDate, endDate];
};

/**
 * 将原始数组转换为对象数组，每个对象包含指定的键和值。
 * 
 * @param array - 原始数组，其元素将被用作对象数组中的值。
 * @param keys - 需要的键数组，可以是字符串或者是一个包含键名和转换函数的对象。
 * @returns 返回一个对象数组，每个对象包含指定的键和值。
 * 
 * 示例
 * const originalArray = ['Apple', 'Banana', 'Cherry'];
 * const keys = [
 *     'original', // 直接使用数组元素作为值
 *     { key: 'lowercase', value: (item: string) => item.toLowerCase() }, // 使用函数转换值
 *     { key: 'length', value: (item: string) => item.length } // 使用函数转换值
 * ];
 */
export const convertArrayToObjectArray = <T>(
    array: T[],
    keys: (string | { key: string, value: (item: T) => any })[]
  ): Object[] => {
    return array.map(item => {
        const obj: { [key: string]: any } = {};
        keys.forEach(key => {
                if (typeof key === 'string') {
                // 如果 key 是字符串，直接使用数组元素作为值
                obj[key] = item;
            } else {
                // 如果 key 是对象，使用提供的函数来计算值
                obj[key.key] = key.value(item);
            }
        });
        return obj;
    });
};
  

/**
 * 从一个对象中移除指定的键
 * @param {Object} obj - 需要操作的对象
 * @param {string} key - 需要移除的键
 * @returns {Object} - 移除指定键后的新对象
 */
export const removeObjectKey = <T extends object, K extends keyof T>(
        obj: T,
        key: K
    ): Omit<T, K> => {
    const { [key]: _, ...rest } = obj;
    return rest;
};


/**
 * 根据给定的键和值在对象数组中查找，返回另一个指定键的值。
 * 
 * @param items - 对象数组
 * @param keyToFind - 要匹配的键
 * @param valueToFind - 要匹配的键的值
 * @param keyToReturn - 找到对象后要返回的键的值
 * @returns 找到的对象中指定键的值，如果没有找到则返回 undefined
 */
 export const findValueByKey = <T, K1 extends keyof T, K2 extends keyof T>(
    items: T[],
    keyToFind: K1,
    valueToFind: T[K1],
    keyToReturn: K2
): T[K2] | undefined => items.find(item => item[keyToFind] === valueToFind)?.[keyToReturn];

/**
 * 根据给定的键和值数组在对象数组中查找，返回另一个指定键的值数组。
 * 
 * @param items - 对象数组
 * @param keyToFind - 要匹配的键
 * @param valuesToFind - 要匹配的键的值数组
 * @param keyToReturn - 找到对象后要返回的键的值
 * @returns 找到的对象中指定键的值数组，如果某个值没有找到则跳过
 */
 export const findValuesByKey = <T, K1 extends keyof T, K2 extends keyof T>(
    items: T[],
    keyToFind: K1,
    valuesToFind: T[K1][],
    keyToReturn: K2
  ): T[K2][] => {
    return valuesToFind.reduce((acc: T[K2][], valueToFind) => {
        const foundItem = items.find(item => item[keyToFind] === valueToFind);
        if (foundItem) {
            acc.push(foundItem[keyToReturn]);
        }
        return acc;
    }, []);
  };

/** 
 * 渲染数组内容，每个元素变成唯一对象
 */
export const renderItemsWithBreaks = (items: any) => {
    return items.map((item: any) => ({ item, key: hashCode(item) }));
}; 

/**
 * 一个简单的哈希函数，用于生成字符串的哈希值。
 * 它将给定的值转换为字符串，然后计算该字符串的哈希值。
 * 
 * @param value - 任何可以被JSON.stringify的值。
 * @returns 哈希值的字符串表示。
 */
 export const hashCode = (value: any): string => {
    let hash = 0;
    const str = JSON.stringify(value); // 将对象转换为字符串
    if (str.length === 0) return hash.toString();
    for (let i = 0; i < str.length; i++) {
        const char = str.charCodeAt(i);
        hash = (hash << 5) - hash + char;
        hash |= 0; // 将hash转换为32位整数
    }
    return hash.toString();
};
  
  
/**
 * 查找并返回数组中最长的字符串。
 * 如果数组为空，返回`undefined`。
 * 如果有多个最长的字符串，返回第一个遇到的。
 * 
 * @param arr - 一个字符串数组。
 * @returns 返回数组中最长的字符串，或者在数组为空时返回`undefined`。
 */
export const findLongestString = (arr: string[]): string | undefined => {
    if (arr.length === 0) return undefined;
    return arr.reduce((longest, current) => {
        return current.length > longest.length ? current : longest;
    });
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
export function arrayToStringWithPlus(arr: any[]): string {
    return arr.join(' + ');
}

export function joinObjectKeyValues<T>(array: T[] | undefined, key: keyof T): string {
    // 如果数组是undefined或空数组，返回空字符串
    if (!array || array.length === 0) {
      return '';
    }
    return array.map(item => String(item[key])).join('+');
  }