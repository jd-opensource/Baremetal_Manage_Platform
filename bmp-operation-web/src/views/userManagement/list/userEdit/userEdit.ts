import {StatusParamsType} from '@utils/publicType';
import {EditUserType} from '../typeManagement';

interface DataType {
    listData: {[x: string]: {}};
    prefix: Ref<string>;
    editingUserDiaLog: Ref<boolean>;
}

const editOperate = <T extends {getUserList(): void}>(userList: T) => {
    const data: DataType = {
        listData: reactive<{[x: string]: {}}>({}),
        prefix: ref<string>(''), // 手机号类型
        // 编辑用户组件弹窗状态
        editingUserDiaLog: ref<boolean>(false)
    };

    /**
     * 点击编辑用户操作
     * @param {Object} item 当前点击的这一项
     * @return {boolean} addUserDialog.value 编辑用户组件弹窗状态
    */
    const editUserClick = async (item: EditUserType, phoneType: string): Promise<boolean> => {
        const {userName, roleName, description, phoneNumber, roleId, userId, phonePrefix, email, defaultProjectId, createdTime}: EditUserType = item;
        data.prefix.value = phoneType;
        await nextTick(() => {
            data.listData.value = {
                userName,
                roleName,
                description,
                phoneNumber,
                roleId,
                userId,
                phonePrefix,
                email,
                defaultProjectId,
                createdTime
            };
        });
        return data.editingUserDiaLog.value = !data.editingUserDiaLog.value;
    };

    /**
     * 编辑用户取消事件
     * @param {boolean} status 状态
     * @return {boolean} xxx 蒙层状态
    */
    const editUserCancel: StatusParamsType <(status: boolean) => boolean> = (status: boolean): boolean => {
        return data.editingUserDiaLog.value = status;
    };

    /**
     * 编辑用户操作后的回调-请求列表接口
    */
    const editUserSure = () => {
        userList.getUserList();
    };

    return {
        ...data,
        editUserClick,
        editUserCancel,
        editUserSure
    }
};

export default editOperate;