/**
 * @file
 * @author
*/

import {userInfoAPI} from 'api/userCenter/request.ts';
import {encrypt, customTip} from 'utils/index.ts';
import {CurrencyType, RuleFormRefType} from '@utils/publicType';
import store from 'store/index.ts';
import {AxiosError} from 'axios';
import {locationItem, CurrentInstance} from 'utils/publicClass.ts';
import {CheckType, UserPurviewType} from '../typeManagement';

interface CheckType1 extends CheckType {
    ruleFormRef: Ref<RuleFormRefType | undefined>;
};

// 登录操作
class LoginOperate {
    userPurview: UserPurviewType = store.userPurviewInfo();
    errMessage: Ref<string> = ref<string>('');
    // 加载态-用于登录请求时，按钮的加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    ruleFormCheck: CheckType1;
    loginInfo: {email: string} = store.loginInfo;
    router: any;
    instanceProxy = new CurrentInstance().proxy;

    // 构造器
    constructor (ruleFormCheck: CheckType1, router: any) {
        this.router = router;
        this.ruleFormCheck = ruleFormCheck;
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.loginOperate();
        };
        watchEffect(() => {
            this.loadingKeyDownStatus();
        })
    };

    /**
     * 登录操作，校验输入项
    */
    loginOperate = async (): Promise<void> => {
        await nextTick(() => {
            this.ruleFormCheck.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入项符合条件
                this.errMessage.value = '';
                if (valid) {
                    this.isLoading.value = true;
                    this.#requestLogin();
                }
            });
        });
    };

    /**
     * 请求登录接口
    */
    #requestLogin = async () => {
        try {
            const params = {
                username: this.ruleFormCheck.ruleForm.loginUserName,
                password: encrypt(this.ruleFormCheck.ruleForm.loginPassword)
            };
            const loginApi = await this.instanceProxy.$loginUserApi.loginAPI(params);
            const status: boolean = (loginApi?.data?.result?.success)?? false;
            status && this.#loginDealWithResponse();
        }
        catch (e) {
            this.#loginDealWithCatch(e as AxiosError);
        }
    };

    #loginDealWithResponse = () => {
        const defaultPath: string = this.#routerUrl('idcList');
        const pathUrl: string = locationItem.cookie.get('path_url')?? '';
        this.getUserInfo(pathUrl || window.btoa(defaultPath));
    };

    #loginDealWithCatch = (e: AxiosError) => {
        const err = e as AxiosError;
        if ((err.code as unknown | number) === 400) {
            this.loadingKeyDownStatus();
            this.errMessage.value = err.message;
            return;
        }
        customTip('error', err.message, 1000, () => this.loadingKeyDownStatus());
    };

    /**
     * 获取用户信息
     * @param {string} pathUrl 页面地址，登出前最后一次停留的地址
     * @param {string} loginUserName 登录用户名
    */
    getUserInfo = (pathUrl: string) => {
        userInfoAPI({})
        .then(({data}: {data: {result: CurrencyType}}) => {
            if (data?.result && Object.keys(data.result).length) {
                // cookie获取存储的语言
                locationItem.cookie.set('X-Jdcloud-Language', data.result.language);
                this.loginInfo.email = window.btoa(data.result.email);
                const userName: string = data.result.userName?? '';
                const email: string = data.result.email?? '';
                // const userId = data.result.userId?? '';
                localStorage.setItem('userName', window.btoa(userName));
                localStorage.setItem('email', window.btoa(email));
                // 页面地址加密
                pathUrl = window.atob(pathUrl);
                const loginPath: string = this.#routerUrl('login');
                const idcListPath: string = this.#routerUrl('idcList');
                pathUrl === loginPath ? idcListPath : pathUrl;
                // const licenseExist = data.result.licenseExist;
                this.userPurview.getUserPurview(this.instanceProxy.$loginUserApi)
                .finally(() => {
                    this.router.push(pathUrl);
                    // loading false
                    // this.loadingKeyDownStatus();
                    // if (licenseExist) {
                    //     // 跳转页面
                    //     // window.open(pathUrl, '_self');
                    //     this.router.push(pathUrl);
                    //     return;
                    // }
                    // const cardPath = this.#routerUrl('license');
                    // // 跳转页面
                    // // window.open(cardPath, '_self');
                    // this.router.push(cardPath);
                })
            }
        })
        .catch(({message}: {message: string}) => {
            customTip('error', message, 1000, () => this.loadingKeyDownStatus());
        })
    };

    #routerUrl = (type: string): string => {
        return this.instanceProxy!.$defInfo.routerPath(type) as unknown as string;
    };

    /**
     * 登录失败-异常后的操作状态
    */
    loadingKeyDownStatus = () => {
        this.isLoading.value = false;
    };
};

export default LoginOperate;
