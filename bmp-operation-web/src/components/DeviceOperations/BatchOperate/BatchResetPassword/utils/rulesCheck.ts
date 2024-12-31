import RegularContent from 'utils/regular.ts';
import {language} from 'utils/publicClass.ts';

class RulesCheck {
    confirmPasswordFlag: Ref<boolean> = ref<boolean>(false);
    errorTip: Ref<string> = ref<string>(language.t('batchOperate.resetPassword.ipt.errorTip.newPassword'));
    colorStatus: Ref<boolean> = ref<boolean>(false);
    newPassword: Ref<string> = ref<string>('');
    newPasswordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.newPassword.value = value;
        this.colorStatus.value = true;
        const passwordArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.errorTip.value = language.t('batchOperate.resetPassword.ipt.emptyTip.newPassword');
                    callback(new Error()); 
                }
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => {
                    this.errorTip.value = language.t('batchOperate.resetPassword.ipt.errorTip.newPassword');
                    callback(new Error());
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.errorTip.value = language.t('batchOperate.resetPassword.ipt.errorTip.newPassword');
                    this.colorStatus.value = false;
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

    confirmPasswordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.confirmPasswordFlag.value = true;
        const confirmPasswordArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('batchOperate.resetPassword.ipt.emptyTip.confirmPassword')))
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => callback(new Error(language.t('batchOperate.resetPassword.ipt.errorTip.newPassword')))
            ],
            [
                (value: string) => (value !== this.newPassword.value),
                () => callback(new Error(language.t('batchOperate.resetPassword.ipt.errorTip.confirmPassword')))
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

export default RulesCheck;