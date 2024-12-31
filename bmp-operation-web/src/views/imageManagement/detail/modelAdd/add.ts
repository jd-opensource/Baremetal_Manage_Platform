interface AssociatedModelType {
    getTable(): Function;
};

const modelAdd = (associatedModel: AssociatedModelType) => {
    const addModelDiaLog: Ref<boolean> = ref<boolean>(false);


    const addModelClick = () => {
        return addModelDiaLog.value = !addModelDiaLog.value;
    };

    const addModelCancel = (type: boolean): boolean => {
        return addModelDiaLog.value = type;
    };
    
    const addModelSure = () => {
        associatedModel.getTable();
    };

    return {
        addModelDiaLog,
        addModelClick,
        addModelCancel,
        addModelSure
    };
};

export default modelAdd;
