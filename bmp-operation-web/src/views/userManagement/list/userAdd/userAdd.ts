/**
 * @file
 * @author
*/

import {ClassParamsType} from '../typeManagement';
import {StatusParamsType, StatusType} from '@utils/publicType';

interface DataType {
    addUserDialog: Ref<boolean>;
    refreshReset: ClassParamsType['refreshOperate'];
};

const addUserInfo = (refreshReset: ClassParamsType['refreshOperate']) => {
    const data: DataType = {
        // 添加用户组件弹窗状态
        addUserDialog: ref<boolean>(false),
        refreshReset
    };

    /**
     * 点击添加用户按钮
     * @return {boolean} xxx 添加用户组件蒙层状态
    */
    const addUserClick: StatusType<() => boolean> = (): boolean =>{
        return data.addUserDialog.value = true;
    };

    /**
     * 添加用户取消事件
     * @param {boolean} status 添加用户弹窗状态
     * @return {boolean} xxx 点击操作后的弹窗状态
    */
    const addUserCancel: StatusParamsType<(status: boolean) => boolean> = (status: boolean): boolean => {
        return data.addUserDialog.value = status;
    };

    /**
     * 添加用户操作后成功的返回
    */
    const addUserSure = () => {
       // 清空搜索值，确保可以看到最新的值
       data.refreshReset.reset();
   };

    return {
        ...data,
        addUserSure,
        addUserClick,
        addUserCancel
    }
};

export default addUserInfo;
