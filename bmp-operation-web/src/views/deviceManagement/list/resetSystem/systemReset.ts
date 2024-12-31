const resetSystemOperate = (deviceList: any) => {
    const resetSystemDiaLog: Ref<boolean> = ref<boolean>(false);

    const resetSystemClick = () => {
        resetSystemDiaLog.value = !resetSystemDiaLog.value;
    };

    const resetSystemCancel = (type: boolean) => {
        resetSystemDiaLog.value = type;
    };

    const resetSystemSure = () => {
        deviceList.getDataList();
    };

    return {
        resetSystemClick,
        resetSystemDiaLog,
        resetSystemCancel,
        resetSystemSure
    };
};

export default resetSystemOperate;
