type ImageLiatType = {
    getDataList(): void;
    reactiveArr: {
        tableData: {
            length: number;
        };
    }
};

type ResetOperateType = {
    reset(): void;
};

interface ParamsType {
    deviceList: ImageLiatType;
    resetOperate: ResetOperateType;
};

const refreshOperate = (deviceList: ParamsType['deviceList'], resetOperate: ParamsType['resetOperate']) => {
    const refresh = () => {
        if (!deviceList?.reactiveArr?.tableData?.length) {
            resetOperate.reset();
            return;
        }
        deviceList.getDataList();
    };

    return {
        refresh
    };
};

export default refreshOperate;
