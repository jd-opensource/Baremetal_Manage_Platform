const batchRecoveryInstance = (deviceList: any) => {
    const batchRecoveryInstanceDiaLog: Ref<boolean> = ref<boolean>(false);

    const batchRecoveryInstanceClick = () => {
        batchRecoveryInstanceDiaLog.value = !batchRecoveryInstanceDiaLog.value;
    };

    const batchRecoveryInstanceCancel = (type: boolean) => {
        batchRecoveryInstanceDiaLog.value = type;
    };

    const batchRecoveryInstanceSure = () => {
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        batchRecoveryInstanceDiaLog,
        batchRecoveryInstanceClick,
        batchRecoveryInstanceCancel,
        batchRecoveryInstanceSure
    }
};

export default batchRecoveryInstance;
