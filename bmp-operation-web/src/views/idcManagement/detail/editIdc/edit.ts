const editIdcOperate = (idcDetail: any) => {
    // 编辑机房弹窗状态
    const editIdcDiaLog: Ref<boolean> = ref<boolean>(false);

    /**
     * 编辑机房弹窗取消事件，弹窗关闭
     * @param {Array} type 控制弹窗显示隐藏 true 显示 false 隐藏
     * @return {boolean} editIdcDiaLog.value true-显示 false-隐藏
    */
    const editIdcCancel = (type: boolean): boolean => {
        return editIdcDiaLog.value = type;
    };

    const editIdcSure = () => {
        idcDetail.isLoading.value = true;
        idcDetail.getDetailData();
    };

    /**
     * 是否可以点击编辑操作
     * @param {boolean} type false 说明不可以在验证完安全校验后进行编辑操作
    */
    const hasEditOperateClick = (type: boolean): void => {
        if (!type) {
            // 遍历数据
            idcDetail.getDetailData();
        }
    };

    return {
        editIdcDiaLog,
        editIdcCancel,
        editIdcSure,
        hasEditOperateClick
    };
};

export default editIdcOperate;
