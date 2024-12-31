import RegularContent from 'utils/regular.ts';
import {language} from 'utils/publicClass.ts';
import {VerifyRulesType} from '@utils/publicType';

interface DataType {
    keyNameTip: Ref<string>;
    hasErrorStatus: Ref<boolean>;
    rules: {
        name: VerifyRulesType[];
        readOnly: VerifyRulesType[];
    };
};

const regExpCheck = (props: any) => {
    const data: DataType = {
        keyNameTip: ref<string>(language.t('createKey.errorTip.keyName')),
        hasErrorStatus: ref<boolean>(false),

        rules: reactive<DataType['rules']>({
            name: [
                {
                    required: true,
                    trigger: 'blur',
                    validator: ''
                }
            ],
            readOnly: [
                {
                    required: true,
                    trigger: 'change',
                    message: language.t('createKey.empty.permissions')
                }
            ]
        })
    };

    /**
     * 密钥名称正则校验
     * @param {Object} _ 占位符
     * @param {string} value 输入项
     * @param {Function} callback 回调函数，返回对应状态
    */
    const keyNameCheck: unknown = (_: {[x: string]: unknown;}, value: string, callback: (arg0?: Error | string) => void) => {
        data.hasErrorStatus.value = true;
        const keyNameArr = [
            [
                (value: string) => !value?.length,
                () => {
                    data.keyNameTip.value = language.t('createKey.empty.keyName');
                    callback(new Error());
                }
            ],
            [
                (value: string) => (!RegularContent.nameReg.test(value)),
                () => {
                    data.keyNameTip.value = language.t('createKey.errorTip.keyName');
                    callback(new Error());
                }
            ],
            [
                (value: string) => (props.keyNameData.includes(value)),
                () => {
                    data.keyNameTip.value = language.t('createKey.errorTip.repeat');
                    callback(new Error());
                }
            ],
            [
                (value: string) => value,
                () => {
                    data.hasErrorStatus.value = false;
                    data.keyNameTip.value = language.t('createKey.errorTip.keyName');
                    callback();
                }
            ]
        ];
        for (const key of keyNameArr) {
            if (key[0](value)) {
                key[1](value);
                break;
            }
        }
    };

    data.rules.name[0].validator = keyNameCheck;

    return {
        ...data,
        keyNameCheck
    };
};

export default regExpCheck;
