class HeightStatus {
    hasClick: Ref<boolean> = ref<boolean>(false);

    constructor(formRulesOperate: any) {
        watch(() => formRulesOperate.ruleForm.modelName, (newValue) => {
            if (newValue) {
                this.hasClick.value = false;
            }
        })
    };
};

export default HeightStatus;
