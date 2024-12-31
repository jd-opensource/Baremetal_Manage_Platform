import {RuleFormRefType} from '@utils/publicType';

class FormRulesOperate {
    // 表单ref
    ruleFormRef: Ref<RuleFormRefType | undefined> = ref<RuleFormRefType>();
    ruleForm: any = reactive({
        newPassword: '',
        confirmPassword: ''
    });
    rules = {
        newPassword: [ // 邮箱
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
        confirmPassword: [
            {
                required: true,
                trigger: 'blur',
                validator: ''
            }
        ],
    };
    rulesCheck: any;


    constructor (rulesCheck: any) {
        this.rulesCheck = rulesCheck;
        this.rules.newPassword[0].validator = rulesCheck.newPasswordCheck;
        this.rules.confirmPassword[0].validator = rulesCheck.confirmPasswordCheck;
    };

    // setData = (item: any) => {
    //     // this.reactiveArr.tableData.push(item);
    // };

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType}): void => {
        this.ruleFormRef.value = formEl.value;
    };
};

export default FormRulesOperate;
