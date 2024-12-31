class FormRulesOpt {
    ruleForm: {instanceName: string} = reactive<{instanceName: string}>({
        instanceName: ''
    });
    ruleFormRef: Ref<any | undefined> = ref<any>();


    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: any}): void => {
        this.ruleFormRef.value = formEl.value;
    };
   
};

export default FormRulesOpt;
