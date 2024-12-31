import {language} from 'utils/publicClass.ts';
import RegularContent from 'utils/regular.ts';
import {resourcesAPI} from 'api/public/request.ts';

class RegExpCheck {
    userNameTip: Ref<string> = ref<string>('');
    heightFlag: Ref<boolean> = ref<boolean>(false);
    passwordFlag: Ref<boolean> = ref<boolean>(false);
    isShowLoading: Ref<boolean> = ref<boolean>(false);

    /**
     * 用户名正则校验
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    userNameCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        const userNameArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.userNameTip.value = language.t('addUser.emptyTip.userName');
                    callback(new Error());
                }
            ],
            [
                (value: string) => (!RegularContent.name1Reg.test(value)),
                () => {
                    this.userNameTip.value = language.t('addUser.errorTip.userName');
                    callback(new Error());
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.userNameTip.value = '';
                    this.isShowLoading.value = true;
                    this.getData(value).then(() => {
                        this.userNameTip.value = language.t('addUser.errorTip.repeatUserName');
                        callback(new Error());
                    })
                    .catch(() => {
                        this.userNameTip.value = '';
                        callback();
                    })
                    .finally(() => {
                        this.isShowLoading.value = false;
                    });
                }
            ]
        ];
        for (const key of userNameArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    getData = (userName: string): Promise<void> => {
        return resourcesAPI({userName})
        .then(({data} : {data: {result: {success: boolean}}}) => {
            if (data?.result?.success) {
                return Promise.resolve();
            }
            return Promise.reject();
        })
        .catch(() => {
            return Promise.reject();
        });
    };

    /**
     * 密码正则校验
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    passwordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        const passwordArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.heightFlag.value = false;
                    this.passwordFlag.value = true;
                    callback(new Error(language.t('addUser.emptyTip.password')));
                }
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => {
                    this.heightFlag.value = true;
                    this.passwordFlag.value = false;
                    callback(new Error(language.t('addUser.errorTip.password')));
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.heightFlag.value = false;
                    this.passwordFlag.value = false;
                    callback();
                }
            ]
        ];
        for (const key of passwordArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };
};

export default RegExpCheck;
