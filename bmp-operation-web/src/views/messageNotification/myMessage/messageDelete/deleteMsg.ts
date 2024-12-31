/**
 * @file
 * @author
*/

class DeleteMsgOpt {
    deleteMsgDiaLog: Ref<boolean> = ref<boolean>(false);
    myMessageListOpt;

    constructor(myMessageListOpt: {restoreDefaultStatus(): void}) {
        this.myMessageListOpt = myMessageListOpt;
    }

    deleteMsgClick = () => {
        this.deleteMsgDiaLog.value = !this.deleteMsgDiaLog.value;
    }

    deleteMsgCancel = (type: boolean): boolean => {
        return this.deleteMsgDiaLog.value = type;
    }

    deleteMsgSure = () => {
        this.myMessageListOpt.restoreDefaultStatus();
    }
}

export default DeleteMsgOpt;
