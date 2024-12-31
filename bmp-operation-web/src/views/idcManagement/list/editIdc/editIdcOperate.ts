/**
 * @file
 * @author
*/

import {StatusParamsType} from '@utils/publicType';

interface IdcListType {
    getDataList(): void;
};

const editIdcOperate = (idcList: IdcListType) => {

    // 编辑机房弹窗状态
    const editingComputerRoomDiaLog: Ref<boolean> = ref<boolean>(false);

    /**
     * 编辑机房弹窗取消事件，弹窗关闭
     * @param {Array} type 控制弹窗显示隐藏 true 显示 false 隐藏
     * @return {boolean} editingComputerRoomDiaLog.value true-显示 false-隐藏
    */
    const editIdcCancel: StatusParamsType <(type: boolean) => boolean> = (type: boolean): boolean => {
        return editingComputerRoomDiaLog.value = type;
    };

    /**
     * 编辑机房成功的操作
    */
    const editIdcSure = () => {
        idcList.getDataList();
    };

    return {
        editingComputerRoomDiaLog,
        editIdcCancel,
        editIdcSure
    }
};

export default editIdcOperate;
