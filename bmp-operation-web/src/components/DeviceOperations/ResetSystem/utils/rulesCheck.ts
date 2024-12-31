import {CurrencyType3} from '@utils/publicType';
import RegularContent from 'utils/regular.ts';
import {language} from 'utils/publicClass.ts';

class RulesCheck {
    hostNameFlag: Ref<boolean> = ref<boolean>(false);
    confirmPasswordFlag: Ref<boolean> = ref<boolean>(false);
    errorTip: Ref<string> = ref<string>(language.t('resetSystem.regCheckTip.password'));
    colorStatus: Ref<boolean> = ref<boolean>(false);
    newPassword: Ref<string> = ref<string>('');
    hasKeyFlag: Ref<boolean> = ref<boolean>(false);

    hostNameCheck = (_: CurrencyType3, value: string, callback: (arg0?: Error | string) => void) => {
        this.hostNameFlag.value = true;
        if (!value?.length) {
            this.hostNameFlag.value = false;
            callback();
        }
        else if (!RegularContent.hostNameReg.test(value)) {
            callback(new Error(language.t('resetSystem.regCheckTip.hostName')));
        }
        else {
            this.hostNameFlag.value = false;
            callback();
        }
    };

    passwordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.newPassword.value = value;
        this.colorStatus.value = true;
        const newPasswordArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.errorTip.value = language.t('resetSystem.emptyCheck.password');
                    callback(new Error()); 
                }
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => {
                    this.errorTip.value = language.t('resetSystem.regCheckTip.password');
                    callback(new Error());
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.errorTip.value = language.t('resetSystem.regCheckTip.password');
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
                () => callback(new Error(language.t('resetSystem.emptyCheck.confirmPassword')))
            ],
            [
                (value: string) => (!RegularContent.passwordReg.test(value)),
                () => callback(new Error(language.t('resetSystem.regCheckTip.password')))
            ],
            [
                (value: string) => (value !== this.newPassword.value),
                () => callback(new Error(language.t('resetSystem.regCheckTip.password2')))
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
