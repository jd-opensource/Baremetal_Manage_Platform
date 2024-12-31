import {RuleFormRefType} from '@utils/publicType';

class FormRulesOperate {
    reactiveArr: any = reactive({
        tableData: []
    });
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    ruleForm: any = reactive({
        newPassword: '',
        confirmPassword: ''
    });

    setData = (item: any) => {
        this.reactiveArr.tableData.push(item);
    };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };
};

export default FormRulesOperate;
