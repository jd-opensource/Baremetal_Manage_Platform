const batchOpenCloseRestart = (deviceList: any) => {
    const types: Ref<string> = ref<string>('');
    const batchOpenCloseRestartDiaLog: Ref<boolean> = ref<boolean>(false);

    const batchOpenCloseRestartClick = (type: string) => {
        types.value = type;
        batchOpenCloseRestartDiaLog.value = !batchOpenCloseRestartDiaLog.value;
       
    };

    const batchOpenCloseRestartCancel = (type: boolean) => {
        batchOpenCloseRestartDiaLog.value = type;
    };

    const batchOpenCloseRestartSure = () => {
        deviceList.filter.tableRef?.value?.clearSelection()
        deviceList.getDataList();
    };

    return {
        types,
        batchOpenCloseRestartDiaLog,
        batchOpenCloseRestartClick,
        batchOpenCloseRestartCancel,
        batchOpenCloseRestartSure
    };
};

export default batchOpenCloseRestart;
