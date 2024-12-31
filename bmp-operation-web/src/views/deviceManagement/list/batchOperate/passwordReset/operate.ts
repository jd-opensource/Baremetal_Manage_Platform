const batchPasswordReset = (deviceList: any) => {
    const batchResetPasswordDiaLog: Ref<boolean> = ref<boolean>(false);

    const batchResetPasswordClick = (status: boolean) => {
        if (status) return;
        batchResetPasswordDiaLog.value = !batchResetPasswordDiaLog.value;
    };

    const batchResetPasswordCancel = (type: boolean) => {
        batchResetPasswordDiaLog.value = type;
    };

    const batchResetPasswordSure = () => {
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        batchResetPasswordDiaLog,
        batchResetPasswordClick,
        batchResetPasswordCancel,
        batchResetPasswordSure
    }
};

export default batchPasswordReset;
