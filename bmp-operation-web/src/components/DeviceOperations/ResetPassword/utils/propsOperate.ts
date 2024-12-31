import {CurrentInstance} from 'utils/publicClass.ts';

const propsOperate = (fn: any) => {
    const instanceMitt: any = new CurrentInstance();

    const getTableFn = (item: any) => {
        fn(item);
    };

    // 监听mitt-device-table，获取表格数据
    instanceMitt?.proxy?.$Bus?.on('reset-password', getTableFn);

    onUnmounted(() => {
        instanceMitt?.proxy?.$Bus?.off('reset-password', getTableFn);
    })
}

export default propsOperate;
