// import {CurrencyType1} from '@utils/publicType';
import {language} from 'utils/publicClass.ts';
import {resourcesAPI} from 'api/public/request.ts';
import {RulesType} from '../typeManagement';

class RulesVerification {
    isShowLoading: Ref<boolean> = ref<boolean>(false);
    imageNameFlag: Ref<boolean> = ref<boolean>(false);
    customOSType: Ref<boolean> = ref<boolean>(false);
    customOSVersion: Ref<boolean> = ref<boolean>(false);
    // 表单提交规则
    rules: RulesType = reactive<RulesType> ({
        imageName: [ // 镜像名称
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        architecture: [ // 体系架构
            {
                required: true,
                trigger: 'change',
                message: language.t('importImage.emptyTip.architecture')
            }
        ],
        osType: [ // 操作系统平台
            {
                required: true,
                trigger: 'change',
                message: language.t('importImage.emptyTip.operateSysPlatform')
            }
        ],
        customOperatePlatform: [ // 自定义操作系统平台
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        customVersion: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        imageFormat: [
            {
                required: true,
                trigger: 'change',
                message: language.t('importImage.emptyTip.imageFormat')
            }
        ],
        bootMode: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
                // message: language.t('importImage.emptyTip.bootMode')
            }
        ],
        version: [ // 操作系统版本
            {
                required: true,
                trigger: 'change',
                message: language.t('importImage.emptyTip.operateSysVersion')
            }
        ],
        systemPartition: [ // 系统盘分区模板
            {
                required: true
            }
        ],
        imageFile: [ // 选择镜像
            {
                required: true,
                trigger: 'change'
            }
        ]
    });
    fn: any;

    constructor(fn: any) {
        this.fn = fn;
        this.rules.imageName[0].validator = this.imageNameChcek;
        this.rules.customOperatePlatform[0].validator = this.customOperatePlatformCheck;
        this.rules.customVersion[0].validator = this.customVersionCheck;
        this.rules.bootMode[0].validator = this.bootModeCheck;
    }

    /**
     * 镜像名称校验
     * @param {Object} _ 对应规则的标识
     * @param {string} value 输入框输入的值
     * @param {Function} callback 回调函数，异常/正常时的返回
    */
    imageNameChcek: unknown = (_: unknown, value: string, callback: (arg0?: Error | string) => void) => {
        const imageNameArr = [
            [
                (value: string) => !value?.length,
                () => {
                    this.imageNameFlag.value = true;
                    callback(new Error(language.t('importImage.emptyTip.imageName')));
                }
            ],
            [
                (value: string) => (value.length > 64),
                () => {
                    this.imageNameFlag.value = true;
                    callback(new Error(language.t('importImage.errTip.imageName')));
                }
            ],
            [
                (value: string) => value,
                () => {
                    this.isShowLoading.value = true;
                    resourcesAPI(
                        {
                            imageName: value
                        }
                    )
                    .then(({data}: {data: {result: {success: boolean}}}) => {
                        if (data?.result?.success) {
                            this.imageNameFlag.value = true;
                            callback(new Error(language.t('importImage.errTip.imageRepeat')));
                            return;
                        }
                        // importImageOperate.repeatData = images.map((item) => item.imageName);
                        return Promise.reject();
                    })
                    .catch(() => {
                        this.imageNameFlag.value = false;
                        callback();
                        // importImageOperate.repeatData = [];
                    })
                    .finally(() => {
                        this.isShowLoading.value = false;
                    });
                }
            ]
        ];
        for (const key of imageNameArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    customOperatePlatformCheck: unknown = (_: unknown, value: string, callback: (arg0?: Error | string) => void) => {
        this.customOSType.value = true;
        const cusotmPlatFormArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('importImage.emptyTip.customOperatePlatform')))
            ],
            [
                (value: string) => value,
                () => {
                    this.customOSType.value = false;
                    callback();
                }
            ]
        ];
        for (const key of cusotmPlatFormArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    customVersionCheck: unknown = (_: unknown, value: string, callback: (arg0?: Error | string) => void) => {
        this.customOSVersion.value = true;
        const customVersionArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('importImage.emptyTip.customVersion')))
            ],
            [
                (value: string) => value,
                () => {
                    this.customOSVersion.value = false;
                    callback();
                }
            ]
        ];
        for (const key of customVersionArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    bootModeCheck: unknown = (_: unknown, value: string, callback: (arg0?: Error | string) => void) => {
        this.fn(true);
        const customVersionArr = [
            [
                (value: string) => !value?.length,
                () => callback(new Error(language.t('importImage.emptyTip.bootMode')))
            ],
            [
                (value: string) => value,
                () => {
                    this.fn(false);
                    callback();
                }
            ]
        ];
        for (const key of customVersionArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };
};

export default RulesVerification;
