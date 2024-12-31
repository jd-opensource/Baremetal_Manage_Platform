const batchEditInstanceName = (deviceList: any) => {
    const batchEditInstanceNameDiaLog: Ref<boolean> = ref<boolean>(false);

    const batchEditInstanceNameClick = (status: boolean) => {
        if (status) return;
        batchEditInstanceNameDiaLog.value = !batchEditInstanceNameDiaLog.value;
    };

    const batchEditInstanceNameCancel = (type: boolean) => {
        batchEditInstanceNameDiaLog.value = type;
    };

    const batchEditInstanceNameSure = () => {
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        batchEditInstanceNameDiaLog,
        batchEditInstanceNameClick,
        batchEditInstanceNameCancel,
        batchEditInstanceNameSure
    }
};

export default batchEditInstanceName;
