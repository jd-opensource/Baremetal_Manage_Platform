import {RuleFormRefType, CurrencyType5} from '@utils/publicType';
import {language} from 'utils/publicClass';

interface DataType {
    ruleFormRef: Ref<RuleFormRefType | undefined>;
    ruleForm: {
        name: string;
        readOnly: number;
        type: string;
    };
    reactiveArr: {
        readOnlyData: CurrencyType5[];
    };
};

const formRules = () => {
    const data: DataType = {
        // ref 操作
        ruleFormRef: ref<RuleFormRefType>(),
        // 表单输入项
        ruleForm: reactive<DataType['ruleForm']>({
            name: '',
            readOnly: 0,
            type: 'user'
        }),
        reactiveArr: reactive<DataType['reactiveArr']>({
            readOnlyData: [
                {
                    value: 0,
                    label: language.t('createKey.select.title')
                }
            ]
        })
    };

     /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    const getFormRef = (formEl: {value: RuleFormRefType}): void => {
        data.ruleFormRef.value = formEl.value;
    };

    return {
        ...data,
        getFormRef
    };
};

export default formRules;
