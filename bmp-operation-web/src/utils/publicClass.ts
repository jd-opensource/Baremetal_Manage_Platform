/**
 * @file
 * @author
*/

import store from 'store/index.ts';
import i18n from 'lib/i18n/index.ts';
import VueCookies from 'vue-cookies';
import {useRouter} from 'vue-router';
import {Router, CurrencyType, CurrencyType2, CurrencyType3, CurrencyType4, CustomInfoType, CheckListArrType, RequestParamsType} from '@utils/publicType';
import RegularContent from 'utils/regular.ts';
import {msgTip, customTip, getDate, encrypt, decrypt, tableScroll, languageSwitch, generateRandomNum} from 'utils/index.ts';
import {AxiosError} from 'axios';

// 语言-使用国际化
class Language {
    t = i18n.global.t
};
const language = new Language();

class TimerOperate {
    timer: null | number = null;
};
new TimerOperate();

// element-plus 国际化
class UiLocale {
    // 国际化
    locale: CurrencyType3 = languageSwitch();
};
// 实例化
const uiLocale: UiLocale = new UiLocale();

class CurrentInstance {
    proxy = getCurrentInstance()!.proxy;
    // 使用mitt传值
    instanceMitt = getCurrentInstance();
};

/**
 * 国际化
*/
class LocationItem {
    // cookie ts规范校验
    cookie: {
        [x: string]: unknown;
        get?: Function;
        set?: Function;
    } = VueCookies;
    getLocationItem = (this.cookie?.get && this.cookie.get('X-Jdcloud-Language'))?? 'zh_CN';
};
const locationItem: LocationItem = new LocationItem();

/**
 * 检查userPurview
*/
class InspectuserPurview {
    userPurview: {userPurview: string;} = store.userPurviewInfo();

    /**
     * 检查userPurview 状态
    */
    inspectuserPurview = () => {
        // 'role-admin-uuid-01'
        const {proxy}: any = getCurrentInstance();
        const hasStauts = localStorage?.getItem('userPurview')?? '';
        const status: string = (decrypt(hasStauts));
        if (!status || status !== proxy.$defInfo.purview('admin')) {
            return Promise.reject();
        }
        return Promise.resolve();
    };
};

const inspectuserPurview: InspectuserPurview = new InspectuserPurview();

/**
 * 路由操作
*/
class RouterOperate {
    // 路由
    router: Router = useRouter();
    // 路由地址
    routerPath?: string = '';

    proxy = new CurrentInstance().proxy;

    /**
     * 构造器-接收实例化传递过来的值
     * @param {string} path 路由地址 默认机房列表
    */
    constructor (path?: string) {
        this.routerPath = path?? this.proxy!.$defInfo.routerPath('idcDetail')!;
    };

    /**
     * 跳转路由
     * @param {Object} query 路由参数
    */
    jumpRouter = (query: CurrencyType = {}): void => {
        this.router.push({
            path: this.routerPath,
            query: query?.status ? {} : query
        });
        const encryInfo: string[] = this.proxy!.$defInfo.encryptDecrypt(1)!;
        if (query?.status) sessionStorage.setItem(query?.type, encrypt(query.params, encryInfo[0], encryInfo[1]));
    };
};

/**
 * 分页操作
*/
class PaginationOperate {
    // 总条数
    total: Ref<number> = ref<number>(0);
    // 当前页面页数条数
    pageSize: Ref<number> = ref<number>(20);
    // 当前页面页数
    pageNumber: Ref<number> = ref<number>(1);
    // 路由改变
    routerChange: Ref<boolean> = ref<boolean>(false);

    /**
     * 分页操作
     * @param {number} count 每页展示条数、页数
     * @param {string} type size、num类型
    */
    paginationChange = (count: number, type: string) => {
        this.routerChange.value = false;
        const paginationData = [
            [
                (type: string) => !type.localeCompare('num'),
                () => this.pageNumber.value = count
            ],
            [
                (type: string) => !type.localeCompare('size'),
                () => {
                    this.pageSize.value = count;
                    this.pageNumber.value = 1;
                }
            ]
        ];
        for (const key of paginationData) {
            if (key[0](type)) {
                key[1](type);
                break;
            }
        }
    };
};
// 实例化
const paginationOperate: PaginationOperate = new PaginationOperate();

class UserCheck {
    phoneNumberType: string = '';
    type: string = '';

    constructor (phonePrefix: string = '', type: string = 'editUser') {
        this.phoneNumberType = phonePrefix;
        this.type = type;
    };

    /**
     * 手机号正则
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    phoneChcek: unknown = (_: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        const phoneArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t(`${this.type}.emptyTip.cellPhone`)))
            ],
            [
                (value: string) => (!/^[0-9]*$/.test(value)),
                () => callback(new Error(language.t(`errorTip.title`)))
            ],
            [
                (value: string) => (!RegularContent.phoneReg.test(value) && Object.is(this.phoneNumberType, language.t('addUser.phoneData.China'))),
                () => callback(new Error(language.t(`${this.type}.errorTip.cellPhone`)))
            ],
            [
                (value: string) => (!RegularContent.hongKongPhoneReg.test(value) && this.phoneNumberType.indexOf('852') > -1),
                () => callback(new Error(language.t(`${this.type}.errorTip.HongKong`)))
            ],
            [
                (value: string) => (!RegularContent.aOmenPhoneReg.test(value) && this.phoneNumberType.indexOf('853') > -1),
                () => callback(new Error(language.t(`${this.type}.errorTip.MacaMacao`)))
            ],
            [
                (value: string) => (!RegularContent.taiWanPhoneReg.test(value) && this.phoneNumberType.indexOf('886') > -1),
                () => callback(new Error(language.t(`${this.type}.errorTip.Taiwan`)))
            ],
            [
                (value: string) => value,
                () => callback()
            ]
        ];
        for (const key of phoneArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    /**
     * 邮箱正则
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    emailChcek: unknown = (_: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        const emailArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t(`${this.type}.emptyTip.email`)))
            ],
            [
                (value: string) => (!RegularContent.emailReg.test(value)),
                () => callback(new Error(language.t(`${this.type}.errorTip.email`)))
            ],
            [
                (value: string) => value,
                () => callback()   
            ]
        ];
        for (const key of emailArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };
};

/**
 * 自定义操作
*/
class CustomOperate {
    pageKey: string = ''; // pageKey
    customListDiaLog: Ref<boolean> = ref<boolean>(false); // 自定义列表蒙层
    customStore: { // store-操作自定义列表
        customList(arg0: string, arg1: CustomInfoType): unknown;
    } = store.customInfo;
    reactiveArr: { // 自定义列表数据
        checkListArr: CheckListArrType;
        hasCustomInfo: CurrencyType4;
    } = reactive<{
        checkListArr: CheckListArrType;
        hasCustomInfo: CurrencyType4;
    }>({
        hasCustomInfo: {}, // 列表显示状态
        checkListArr: [] // 角色列表自定义数据
    });

    /**
     * 构造器-接收实例化传递过来的值
     * @param {string} pageKey 自定义列表 pageKey
     * @param {Object} hasCustomInfo 自定义信息
     * @param {Array} checkListArr 自定义列表数据
    */
    constructor (
        pageKey: string,
        hasCustomInfo: CurrencyType4,
        checkListArr: CheckListArrType
    ) {
        this.pageKey = pageKey;
        this.reactiveArr.hasCustomInfo = hasCustomInfo;
        this.reactiveArr.checkListArr = checkListArr;
        this.getCustomList()
    };

    /**
     * 获取自定义列表
    */
    getCustomList = () => {
        this.customStore.customList(this.pageKey, this.reactiveArr);
    };

    /**
     * 自定义取消操作
     * @return {boolean} xxx 自定义列表组件状态
    */
    customCancelOperate = (): boolean => {
        return this.customListDiaLog.value = false;
    };

    /**
     * 自定义点击操作-蒙层显示、隐藏
     * @return {boolean} xxx 蒙层状态
    */
    customClickOperate = (): boolean => {
        return this.customListDiaLog.value = !this.customListDiaLog.value;
    };
};

/**
 * 导出数据操作
*/
class ExportDataOperate {
    exportDataDiaLog: Ref<boolean> = ref<boolean>(false);
    downloadName: string = '';
    // hasExportData: Ref<boolean> = ref<boolean>(true);
    requestUrl: Function = () => {};
    requestParams: RequestParamsType = {exportType: '1', isAll: '1'};
    templateParams: {
        type: string;
    } = {
        type: ''
    };

    /**
     * 构造器-接收实例化传递的参数
     * @param {Function} url 请求url地址
     * @param {Object} params 请求参数
    */
    constructor (
        url:(arg0: RequestParamsType) => Promise<{data: ResponseType}>,
        params: CurrencyType2 = {},
        name: string = 'idc_list'
    ) {
        this.requestUrl = url;
        if (params.type === 'template') {
            this.templateParams = {
                type: params.type
            };
            return;
        };
        this.requestParams = {...this.requestParams, ...params};
        this.downloadName = name;
    };

    diaLogClick = (): boolean => {
        return this.exportDataDiaLog.value = !this.exportDataDiaLog.value;
    };

    exportCancel = (status: boolean): boolean => {
        return this.exportDataDiaLog.value = status;
    };

    /**
     * 导出数据
    */
    exportData = async (status: {cancelClick(): void; isLoading: {value: boolean}}) => {
        const params = this.templateParams?.type ? {} : this.requestParams;
        store.filterEmpty.deleteEmtpyData(params);
        try {
            const exportApi = await this.requestUrl(params);
            msgTip({
                message: language.t('operate.success'),
                type: 'success',
                customClass: 'custom-message',
                duration: 1000,
                onClose: () => {
                    status.cancelClick();
                    const blob: Blob = new Blob([exportApi.data]);
                    // 创建a标签
                    let link = document.createElement('a');
                    // 创建下载的链接
                    link.href = window.URL.createObjectURL(blob);
                    // 下载后的文件名
                    link.download = this.templateParams?.type ? 'device_import' : `${this.downloadName}_${getDate()}_${generateRandomNum(6)}.xlsx`;
                    // 点击下载
                    link.click();
                    // 释放这个url
                    window.URL.revokeObjectURL(link.href);
                    // store.progressInfo.progressNum = 100;
                }
            })
        }
        catch (e) {
            if (Object.is(e, void 0)) {
                throw new Error('下载失败');
            }
            const err = e as AxiosError;
            customTip('error', err.message, 1000, () =>  status.isLoading.value = false);
        }
    };
};


class ScrollOperate {

    constructor(countHeight: number, tableMaxHeight: {value: number}) {
        tableScroll(countHeight, tableMaxHeight);
        this.onResize(countHeight, tableMaxHeight);
    };

    onResize = (countHeight: number, tableMaxHeight: {value: number}) => {
        window.onresize = () => {
            tableScroll(countHeight, tableMaxHeight);
        };
    };
};

export {
    language,
    uiLocale,
    locationItem,
    paginationOperate,
    inspectuserPurview,
    CurrentInstance,
    UserCheck,
    TimerOperate,
    RouterOperate,
    CustomOperate,
    ScrollOperate,
    ExportDataOperate
};
