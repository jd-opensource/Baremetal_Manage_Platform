import {resourcesAPI} from 'api/public/request.ts';
import {language} from 'utils/publicClass.ts';
import {CallbackType} from '@utils/publicType';
import RegularContent from 'utils/regular.ts';

// type ParamsType = Partial<{name: string; deviceType: string;}>;

class RegExpCheck {
    memSizeFlag: Ref<boolean> = ref<boolean>(false);
    memFrequencyFlag: Ref<boolean> = ref<boolean>(false);
    memTypeFlag: Ref<boolean> = ref<boolean>(false);
    heightFlag: Ref<boolean> = ref<boolean>(false);
    interfaceModeFlag: Ref<boolean> = ref<boolean>(false);
    nicRateFlag: Ref<boolean> = ref<boolean>(false);
    cpuManufacturerFlag: Ref<boolean> = ref<boolean>(false);
    memInfoFlag: Ref<boolean> = ref<boolean>(false);
    cpuInfoFlag: Ref<boolean> = ref<boolean>(false);
    cpuGHzFlag: Ref<boolean> = ref<boolean>(false);
    cpuModelFlag: Ref<boolean> = ref<boolean>(false);
    nameFlag: Ref<boolean> = ref<boolean>(false);
    deviceTypeFlag: Ref<boolean> = ref<boolean>(false);
    nameLoading: Ref<boolean> = ref<boolean>(false);
    deviceTypeLoading: Ref<boolean> = ref<boolean>(false);
    ruleForm: {
        minValUnit: string;
    } = {
        minValUnit: 'GB',
    };

    constructor (minValUnit: string = 'GB') {
        this.ruleForm.minValUnit = minValUnit;
    };

    /**
     * 机型名称校验
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    nameCheck: unknown = (_: unknown, value: string, callback: CallbackType) => {
        const nameArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.nameFlag.value = true;
                    callback(new Error(language.t('modelForm.emptyTip.modelName')))
                }
            ],
            [
                (value: string) => value.length > 64,
                () => {
                    this.nameFlag.value = true;
                    callback(new Error(language.t('modelForm.errorTip.currency2')))
                }
            ],
            [
                (value: string) => (!RegularContent.modelNameReg.test(value)),
                () => {
                    this.nameFlag.value = true;
                    callback(new Error(language.t('modelForm.errorTip.currency')))
                }
            ],
            [
                (value: string) => value,
                () => {
                    const name: string = sessionStorage.getItem('propsName')??'';
                    const status: string = sessionStorage.getItem('hasTemplate')??'';
                    this.nameLoading.value = true;
                    if (name === value && status !== 'success') {
                        this.nameFlag.value = false;
                        this.nameLoading.value = false;
                        callback();
                    }
                    else {
                        resourcesAPI(
                            {
                                name: value
                            }
                        )
                        .then(({data} : {data: {result: {success: boolean;}}}) => {
                            if (data?.result?.success) {
                                this.nameFlag.value = true;
                                callback(new Error(language.t('modelForm.errorTip.name')));
                                return;
                            }
                            return Promise.reject();
                        })
                        .catch(() => {
                            this.nameFlag.value = false;
                            callback();
                        })
                        .finally(() => {
                            this.nameLoading.value = false;
                        })
                    }
                }
            ]
        ];
        for (const key of nameArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    /**
     * 机型规格校验
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    deviceTypeCheck: unknown = (_: unknown, value: string, callback: CallbackType) => {
        const deviceTypeArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.deviceTypeFlag.value = true;
                    callback(new Error(language.t('modelForm.emptyTip.modelSpecifications')))
                }
            ],
            [
                (value: string) => value.length > 64,
                () => {
                    this.deviceTypeFlag.value = true;
                    callback(new Error(language.t('modelForm.errorTip.currency2')))
                }
            ],
            [
                (value: string) => (value && sessionStorage?.getItem('repeat')),
                () => {
                    this.deviceTypeFlag.value = true;
                    callback(new Error(language.t('modelForm.errorTip.modelSpecifications')))
                }
            ],
            [
                (value: string) => value,
                () => {
                    const deviceType = sessionStorage?.getItem('propsDeviceType');
                    const status = sessionStorage.getItem('hasTemplate');
                    this.deviceTypeLoading.value = true;
                    if (deviceType === value && status !== 'success') {
                        this.deviceTypeFlag.value = false;
                        this.deviceTypeLoading.value = false;
                        callback();
                    }
                    else {
                        resourcesAPI(
                            {
                                deviceType: value
                            }
                        )
                        .then(({data} : {data: {result: {success: boolean;}}}) => {
                            if (data?.result?.success) {
                                this.deviceTypeFlag.value = true;
                                callback(new Error(language.t('modelForm.errorTip.specifications')));
                                return;
                            }
                            return Promise.reject();
                        })
                        .catch(() => {
                            this.deviceTypeFlag.value = false;
                            callback();
                        })
                        .finally(() => {
                            this.deviceTypeLoading.value = false;
                        })
                    }
                }
            ]
        ];
        for (const key of deviceTypeArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    /**
     * cpu型号
     * @param {string} str 提示项
     * @param {string} trigger 判断时机-change/blur
    */
    cpuModelCheck = (_: {}, value: string, callback: (arg0?: Error) => void) => {
        this.cpuModelFlag.value = true;
        const cpuModelArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('modelForm.emptyTip.modelChoose')))
            ],
            [
                (value: string) => value.length > 64,
                () => callback(new Error(language.t('modelForm.errorTip.currency2')))
            ],
            [
                (value: string) => value,
                () => {
                    callback();
                    this.cpuModelFlag.value = false;
                }
            ]
        ];
        for (const key of cpuModelArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    /**
     * CPU-主频
    */
    cpuFrequencyCheck = (_: unknown, value: string, callback: (arg0?: Error) => void) => {
        this.cpuGHzFlag.value = true;
        const cpuFrequencyArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('modelForm.emptyTip.dominantFrequency')))
            ],
            [
                (value: string) => (!RegularContent.capacityReg.test(value)),
                () => callback(new Error(language.t('modelForm.errorTip.number')))
            ],
            [
                (value: string) => value,
                () => {
                    callback();
                    this.cpuGHzFlag.value = false;
                }
            ]
        ];
        for (const key of cpuFrequencyArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };
};

const regExpCheck: RegExpCheck = new RegExpCheck();

export {
    RegExpCheck,
    regExpCheck
};
