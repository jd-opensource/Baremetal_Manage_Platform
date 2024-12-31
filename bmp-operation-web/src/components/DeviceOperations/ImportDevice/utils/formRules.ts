import {RuleFormType, RuleFormRefType2, RulesType} from '../typeManagement'
import {language} from 'utils/publicClass.ts';

interface DataType {
    ruleFormRef: Ref<RuleFormRefType2 | undefined>;
    ruleForm: RuleFormType;
    rules: RulesType;
};

const formRules = () => {
    const data: DataType = {
        // 表单ref
        ruleFormRef: ref<RuleFormRefType2>(),
        ruleForm: reactive<RuleFormType>({
            fileParams: {},
            idcName: '', // 机房名称
            // modelName: '', // 机型名称
            deviceInfo: [] // 设备信息
        }),
        // 表单提交规则
        rules: reactive<RulesType>({
            idcName: [ // 机房名称
                {
                    required: true,
                    message: language.t('importDevice.errTip.computerRoomName'),
                    trigger: 'change'
                }
            ],
            // modelName: [ // 机型名称
            //     {
            //         required: true,
            //         message: language.t('importDevice.errTip.modelName'),
            //         trigger: 'change'
            //     }
            // ],
            deviceInfo: [ // 设备信息
                {
                    required: true,
                    trigger: 'change'
                }
            ]
        })
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    const getFormRef = (formEl: {value: RuleFormRefType2}): void => {
        data.ruleFormRef.value = formEl.value;
    };

    return {
        ...data,
        getFormRef
    };
};

export default formRules;
