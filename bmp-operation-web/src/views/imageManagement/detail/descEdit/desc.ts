interface BaseInfoType {
    getImageDetail(): void;
};

const descEdit = (baseInfo: BaseInfoType) => {
    const editDescDiaLog: Ref<boolean> = ref<boolean>(false);

    const editDescClick = () => {
        editDescDiaLog.value = !editDescDiaLog.value;
    };
    
    
    const editDescCacnel = (type: boolean): boolean => {
        return editDescDiaLog.value = type;
    };
    
    const editDescSure = () => {
        baseInfo.getImageDetail();
    };

    return {
        editDescDiaLog,
        editDescClick,
        editDescCacnel,
        editDescSure
    };
};

export default descEdit;
