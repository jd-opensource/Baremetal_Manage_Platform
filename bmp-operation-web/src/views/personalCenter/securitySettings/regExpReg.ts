import RegularContent from 'utils/regular.ts';
import {language} from 'utils/publicClass.ts';

class RegExpCheck {
    oldPasswordFlag: Ref<boolean> = ref<boolean>(false);
    confirmPasswordFlag: Ref<boolean> = ref<boolean>(false);
    errorTip: Ref<string> = ref<string>(language.t('userCenter.errorTip.password'));
    colorStatus: Ref<boolean> = ref<boolean>(false);
    newPassword: Ref<string> = ref<string>('');

    /**
     * 密码正则校验
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    oldPasswordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.oldPasswordFlag.value = true;
        const passwordArr = [
            [
                (value: string) => !value?.length,
                () => {
                    callback(new Error(language.t('userCenter.emptyTip.currectPassword')));
                }
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => {
                    callback(new Error(language.t('userCenter.errorTip.password')));
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.oldPasswordFlag.value = false;
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

    newPasswordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.colorStatus.value = true;
        const newPasswordArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.errorTip.value = language.t('userCenter.emptyTip.newPassword');
                    callback(new Error()); 
                }
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => {
                    this.errorTip.value = language.t('userCenter.errorTip.password');
                    callback(new Error());
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.errorTip.value = language.t('userCenter.errorTip.password');
                    this.colorStatus.value = false;
                    callback();
                }
            ]
        ];
        for (const key of newPasswordArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    confirmPasswordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.confirmPasswordFlag.value = true;
        const confirmPasswordArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('userCenter.emptyTip.confirmPassword')))
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => callback(new Error(language.t('userCenter.errorTip.password')))
            ],
            [
                (value: string) => (value !== this.newPassword.value),
                () => callback(new Error(language.t('userCenter.errorTip.password2')))
            ],
            [
                (value: string) => value,
                () => {
                    this.confirmPasswordFlag.value = false;
                    callback();
                }
            ]
        ];
        for (const key of confirmPasswordArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };
};

export default RegExpCheck;
