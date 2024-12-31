import {CurrencyType} from '@utils/publicType';
import {paginationOperate} from 'utils/publicClass.ts';
interface DataType {
    removeOperateDiaLog: Ref<boolean>;
    name: Ref<string>;
    deviceType: Ref<string>;
    deviceTypeId: Ref<string>;
};

interface AssociatedModelType {
    reactiveArr: any;
    getTable(): Function;
};

const remove = (associatedModel: AssociatedModelType) => {
    const data: DataType = {
        removeOperateDiaLog: ref<boolean>(false),
        name: ref<string>(''),
        deviceType: ref<string>(''),
        deviceTypeId: ref<string>('')
    };

    const removeClick = (item: CurrencyType) => {
        data.removeOperateDiaLog.value = !data.removeOperateDiaLog.value;
        data.name.value = item.name;
        data.deviceType.value = item.deviceType;
        data.deviceTypeId.value = item.deviceTypeId;
    
    };
    const removeOperateCancel = (type: boolean): boolean => {
        return data.removeOperateDiaLog.value = type;
    };
    
    const removeOperateSure = () => {
        if (paginationOperate.pageNumber.value > 1 && associatedModel.reactiveArr.tableData.length === 1) {
            paginationOperate.pageNumber.value = paginationOperate.pageNumber.value - 1;
            return;
        }
        associatedModel.getTable();
    };

    return {
        ...data,
        removeClick,
        removeOperateCancel,
        removeOperateSure
    }
};

export default remove;
