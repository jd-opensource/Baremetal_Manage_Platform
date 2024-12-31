import {paginationOperate} from 'utils/publicClass.ts';

interface DataType {
    deleteModelDiaLog: Ref<boolean>;
    modelName: Ref<string>;
    deviceTypeId: Ref<string>;
};

const deleteModelOperate = (modelList: { reactiveArr: { tableData: string | unknown[]; }; getModelList: () => void; }, reset: { reset: () => void; }) => {
    const data: DataType = {
        // 删除机型弹窗状态
        deleteModelDiaLog: ref<boolean>(false),
        // 当前点击删除的机型信息
        modelName: ref<string>(''),
        // 当前点击删除的机型deviceTypeId
        deviceTypeId: ref<string>('')
    };

    /**
     * 删除机型点击事件
     * @param {string} name 当前点击删除的机型名称
     * @param {string} id 当前点击的deviceTypeId
    */
    const deleteModelClick = (item: any): void => {
        if (item.instanceCount > 0 || item.deviceCount > 0) return;
        data.deleteModelDiaLog.value = !data.deleteModelDiaLog.value;
        data.modelName.value = item.name;
        data.deviceTypeId.value = item.deviceTypeId;
    };

    /**
     * 删除机型弹窗取消事件
     * @param {boolean} type false 弹窗关闭
     * @return {boolean} deleteModelDiaLog.value 删除机型弹窗关闭
    */
    const deleteModelCancel = (type: boolean): boolean => {
        return data.deleteModelDiaLog.value = type;
    };

    const deleteModelSure = () => {
        if (paginationOperate.pageNumber.value > 1 && modelList.reactiveArr.tableData.length === 1) {
            paginationOperate.pageNumber.value = paginationOperate.pageNumber.value - 1;
            return;
        }
        if (modelList.reactiveArr.tableData.length === 1) {
            reset.reset();
        }
        else {
            modelList.getModelList();
        }
    };

    const deleteModelError = () => {
        modelList.getModelList();
    };

    return {
        ...data,
        deleteModelClick,
        deleteModelCancel,
        deleteModelSure,
        deleteModelError
    };
};

export default deleteModelOperate;
