interface DataType {
    deleteModelDiaLog: Ref<boolean>;
    modelInfo: Ref<string>;
}
interface BaseInfoType {
    reactiveArr: {
        detail: {
            name: string;
        }
    }
}
const deleteModelOperate = (baseInfo: BaseInfoType) => {
    const data: DataType = {
        // 删除机型弹窗状态-默认false
        deleteModelDiaLog: ref<boolean>(false),
        // 机型信息，删除机型时传递的信息
        modelInfo: ref<string>('')
    }

    /**
     * 删除机型点击事件
     * @return {boolean} deleteModelDiaLog.value 删除机型组件弹窗状态
    */
    const deleteModelClick = (): boolean => {
        data.modelInfo.value = baseInfo.reactiveArr.detail.name;
        return data.deleteModelDiaLog.value = !data.deleteModelDiaLog.value;
    };

    /**
     * 删除机型弹窗取消事件
     * @param {boolean} type false 弹窗关闭
     * @return {boolean} data.deleteModelDiaLog.value 删除机型弹窗关闭
    */
    const deleteModelCancel = (type: boolean): boolean => {
        return data.deleteModelDiaLog.value = type;
    };

    const deleteModelError = () => {
        (baseInfo as any).getModelDetail();
        // location.reload();
    };

    return {
        ...data,
        deleteModelClick,
        deleteModelCancel,
        deleteModelError
    }
};

export default deleteModelOperate;
