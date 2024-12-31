import {CurrentInstance} from 'utils/publicClass.ts';
import router from 'router/index.ts';
// import {DeviceDetail} from '../../methods';

interface DataType {
    upDownFrameTitle: Ref<string>;
    upDownFrameDiaLog: Ref<boolean>;
};

const upDownDeleteOperate = (deviceDetail: any) => {
    const data: DataType = {
        // 上架、下架、删除设备标题名称
        upDownFrameTitle: ref<string>(''),
        // 上架、下架、删除设备组件状态
        upDownFrameDiaLog: ref<boolean>(false)
    };

    const mitt = new CurrentInstance().instanceMitt;

    const instanceProxy = new CurrentInstance().proxy;

    /**
     * 上架、下架、删除设备点击事件
     * @param {string} type 当前点击这一项状态
    */
    const upDownFrameClick = async (type: string, status: boolean): Promise<void> => {
        if (!status) return;
        data.upDownFrameDiaLog.value = !data.upDownFrameDiaLog.value;
        data.upDownFrameTitle.value = type;
        await nextTick(() => {
            mitt?.proxy?.$Bus?.emit('device-table', {list: deviceDetail.reactiveArr.detail});
        });
    };

    const upDownDeleteCancel = (type: boolean) => {
        data.upDownFrameDiaLog.value = type;
    };

    const operateClick = (type: string) => {
        const path: string = instanceProxy.$defInfo.routerPath('deviceList');
        if (type === 'delete') {
            clearTimeout((deviceDetail.timer as number));
            router.push(path);
            return;
        }
        deviceDetail.initData();
    };

    return {
        ...data,
        upDownFrameClick,
        upDownDeleteCancel,
        operateClick
    };
};

export default upDownDeleteOperate;
