interface DataType {
    editNotesDiaLog: Ref<boolean>;
    descInfo: string;
};

const editNotesOperate = (deviceDetail: any) => {
    const data: DataType = {
        editNotesDiaLog: ref<boolean>(false),
        descInfo: String('')
    };

    const editNotesClick = () => {
        data.descInfo = deviceDetail?.reactiveArr.detail.description;
        data.editNotesDiaLog.value = !data.editNotesDiaLog.value;
    };

    const editNotesCacnel = (type: boolean) => {
        data.editNotesDiaLog.value = type;
    };

    const editNotesSure = () => {
        document.onkeyup = () => {return;}
        deviceDetail.initData();
    };

    return {
        ...data,
        editNotesClick,
        editNotesCacnel,
        editNotesSure
    };
};

export default editNotesOperate;
