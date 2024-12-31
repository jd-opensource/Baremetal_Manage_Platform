interface DataType {
    importDeviceDiaLog: Ref<boolean>;
    repeatData: string[];
};

const deviceImport = (reset: any) => {
    const data: DataType = {
        // 导入设备弹窗组件状态
        importDeviceDiaLog: ref<boolean>(false),
        repeatData: ['']
    };

    /**
     * 导入设备点击事件
     * @return {boolean} importDeviceDiaLog.value 导入设备组件弹窗状态
    */
    const importDeviceClick = (): boolean => {
        return data.importDeviceDiaLog.value = true;
    };

    /**
     * 导入设备组件取消事件
     * @param {boolean} type 弹窗状态
     * @return {boolean} importDeviceDiaLog.value 导入设备组件弹窗状态
    */
    const importDeviceCancel = (type: boolean): boolean => {
        return data.importDeviceDiaLog.value = type;
    };

    /**
     * 导入设备成功的回调操作
    */
    const importDeviceSure = () => {
        reset.reset();
    };

    return {
        ...data,
        importDeviceClick,
        importDeviceCancel,
        importDeviceSure
    };
};

export default deviceImport;
