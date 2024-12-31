
import {CurrentInstance} from 'utils/publicClass.ts';
import {QueryType} from '../type';

const propsOperate = (fn: Function) => {
    const instanceMitt = new CurrentInstance();
    const query: Pick<QueryType, 'deviceTypeId' | 'idcId'> = {
        idcId: '',
        deviceTypeId: '',
    };

    const getTableFn = (item: QueryType) => {
        query.deviceTypeId = item.deviceTypeId;
        query.idcId = item.idcId;
        fn(item, query);
    };

    // 监听mitt-device-table，获取表格数据
    instanceMitt?.proxy?.$Bus?.on('reset-system', getTableFn);

    onUnmounted(() => {
        instanceMitt?.proxy?.$Bus?.off('reset-system', getTableFn);
    })
}

export default propsOperate;
