interface DataType {
    title: Ref<string>;
    instanceId: Ref<string>;
    instanceInfo: Ref<string>;
    openShutDownRestartDiaLog: Ref<boolean>;
};

const openCloseRestart = (deviceDetail: any) => {
    const data: DataType = {
        // 开机、关机、重启名称
        title: ref<string>(''),
        instanceId: ref<string>(''),
        // 开机、关机、重启实例信息
        instanceInfo: ref<string>(''),
        // 开机、关机、重启组件状态
        openShutDownRestartDiaLog: ref<boolean>(false)
    };

    /**
     * 开机、关机、重启点击事件
     * @param {string} item.instanceName 当前点击这一项实例
     * @param {string} type 当前点击这一项状态
     * @return {string} title.value 当前点击这一项的title标题
    */
    const openShutDownRestartClick = (instanceName: string, id: string, type: string, status: boolean): string | void => {
        if (!status) return;
        data.openShutDownRestartDiaLog.value = !data.openShutDownRestartDiaLog.value;
        data.instanceInfo.value = instanceName;
        data.instanceId.value = id;
        switch(type) {
            case 'open':
                return data.title.value = '开机';
            case 'restart':
                return data.title.value = '重启';
            default:
                return data.title.value = '关机';
        };
    };
    /**
     * 开机、关机、重启组件取消事件
     * @param {boolean} type 弹窗状态
     * @return {boolean} openShutDownRestartDiaLog.value 开机、关机、重启组件弹窗状态
    */
    const openShutDownRestartCancel = (type: boolean): boolean => {
        return data.openShutDownRestartDiaLog.value = type;
    };

    const openShutDownRestartSure = () => {
        deviceDetail.initData();
    };

    return {
        ...data,
        openShutDownRestartClick,
        openShutDownRestartCancel,
        openShutDownRestartSure
    };
};

export default openCloseRestart;
