import {CurrentInstance} from 'utils/publicClass.ts';

interface ParamsType {
    deviceList: {
        reactiveArr: {
            tableData: string | unknown[];
        };
        getDataList: () => void;
    };
    reset: {
        reset: () => void;
    };
};

interface DataType {
    deviceId: Ref<string>;
    upDownDeleteDiaLog: Ref<boolean>;
    upDownDeleteTitle: Ref<string>;
};

const upDownDelete = (deviceList: ParamsType['deviceList'], reset: ParamsType['reset']) => {
    const data: DataType = {
        deviceId: ref<string>(''),
        // 上架、下架、删除设备组件状态
        upDownDeleteDiaLog: ref<boolean>(false),
        // 上架、下架、删除设备标题名称
        upDownDeleteTitle: ref<string>('')
    };

    const mitt = new CurrentInstance();

    /**
     * 上架、下架、删除设备点击事件
     * @param {Object} item 当前点击的这一项
     * @param {string} type 当前点击这一项状态
    */
    const upDownDeleteClick = async (item: {deviceId: string;}, type: string, status: boolean): Promise<void> => {
        if (!status) return;
        data.deviceId.value = item.deviceId;
        data.upDownDeleteDiaLog.value = !data.upDownDeleteDiaLog.value;
        data.upDownDeleteTitle.value = type;
        await nextTick(() => {
            mitt.instanceMitt?.proxy?.$Bus?.emit('device-table', {list: item});
        });
    };

    /**
     * 上架、下架、删除组件取消事件
     * @param {boolean} type 弹窗状态
     * @return {boolean} upDownDeleteDiaLog.value 上架、下架、删除设备组件弹窗状态
    */
    const upDownDeleteCancel = (type: boolean): boolean => {
        return data.upDownDeleteDiaLog.value = type;
    };

    const upDownDeleteSure = (type: string) => {
        if (type === 'delete') {
            if (deviceList.reactiveArr.tableData.length === 1) {
                reset.reset();
            }
            else {
                deviceList.getDataList();
            }
        }
        else {
            deviceList.getDataList();
        }
    };

    return {
        ...data,
        upDownDeleteClick,
        upDownDeleteCancel,
        upDownDeleteSure
    }
};

export default upDownDelete;
