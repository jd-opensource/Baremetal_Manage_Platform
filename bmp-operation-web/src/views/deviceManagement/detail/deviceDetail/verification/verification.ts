const verificationOperate = (fn: any) => {
    // 安全验证弹窗状态
    const securityVerificationDiaLog: Ref<boolean> = ref<boolean>(false);

    const verificationCancel = (type: boolean) => {
        securityVerificationDiaLog.value = type;
    };

    /**
     * 是否可以点击编辑操作
     * @param {boolean} type false 说明不可以在验证完安全校验后进行编辑操作
    */
    const hasEditOperateClick = (type: boolean): void => {
        if (!type) {
            fn()
        }
    };

    return {
        securityVerificationDiaLog,
        verificationCancel,
        hasEditOperateClick
    };
};

export default verificationOperate;
