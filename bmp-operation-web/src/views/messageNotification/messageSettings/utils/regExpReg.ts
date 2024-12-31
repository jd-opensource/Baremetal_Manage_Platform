// import RegularContent from 'utils/regular.ts';
import {language} from 'utils/publicClass.ts';
import RegularContent from 'utils/regular.ts';

class RegExpCheck {
    emailPasswordFlag: Ref<boolean> = ref<boolean>(false);
    emailAddressFlag: Ref<boolean> = ref<boolean>(false);
    portFlag: Ref<boolean> = ref<boolean>(false);
    ipAddressFlag: Ref<boolean> = ref<boolean>(false);
    testEmailFlag: Ref<boolean> = ref<boolean>(false);

    ipAddressCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.ipAddressFlag.value = true;
        const ipAddressArr = [
            [
                (value: string) => !value?.length,
                () => {
                    callback(new Error(language.t('messageSettings.empty.ipAddress')));
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.ipAddressFlag.value = false;
                    callback();
                }
            ]
        ];
        for (const key of ipAddressArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    portCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.portFlag.value = true;
        const portArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('messageSettings.empty.port')))
            ],
            [
                (value: string) => (!RegularContent.NumReg.test(value) || Number(value) >= 65536),
                () => callback(new Error(language.t('messageSettings.errorTip.port')))
            ],
            [
                (value: string) => value,
                () => {
                    this.portFlag.value = false;
                    callback();
                }
            ]
        ];
        for (const key of portArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    emailAddressCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.emailAddressFlag.value = true;
        const emailAddressArr = [
            [
                (value: string) => !value?.length,
                () => {
                    callback(new Error(language.t('messageSettings.empty.emailAddress')));
                }
            ],
            [
                (value: string) => !RegularContent.emailReg.test(value),
                () => callback(new Error(language.t('messageSettings.errorTip.emailAddress')))
            ],
            [
                (value: string) => value,
                () => {
                    this.emailAddressFlag.value = false;
                    callback();
                }
            ]
        ];
        for (const key of emailAddressArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    testEmailCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.testEmailFlag.value = true;
        const testEmailArr = [
            [
                (value: string) => !value?.length,
                () => {
                    callback(new Error(language.t('messageSettings.empty.testEmail')));
                }
            ],
            [
                (value: string) => !(RegularContent.emailReg.test(value)),
                () => callback(new Error(language.t('messageSettings.errorTip.testEmail')))
            ],
            // [
            //     (value: string) => !(RegularContent.moreEmailReg.test(value)),
            //     () => callback(new Error(language.t('messageSettings.errorTip.testEmail')))
            // ],
            [
                (value: string) => value,
                () => {
                    this.testEmailFlag.value = false;
                    callback();
                }
            ]
        ];
        for (const key of testEmailArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    emailPasswordCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        this.emailPasswordFlag.value = true;
        const emailPasswordArr = [
            [
                (value: string) => !value?.length,
                () => {
                    callback(new Error(language.t('messageSettings.empty.emailPassword')));
                }
            ],
            // [
            //     (value: string) => !RegularContent.passwordReg.test(value),
            //     () => callback(new Error(language.t('messageSettings.errorTip.emailPassword')))
            // ],
            [
                (value: string) => value,
                () => {
                    this.emailPasswordFlag.value = false;
                    callback();
                }
            ]
        ];
        for (const key of emailPasswordArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    resetFlag = () => {
        const {emailPasswordFlag, emailAddressFlag, portFlag, ipAddressFlag, testEmailFlag} = this;
        emailPasswordFlag.value = false;
        emailAddressFlag.value = false;
        portFlag.value = false;
        ipAddressFlag.value = false;
        testEmailFlag.value = false;
    }
};

export default RegExpCheck;
