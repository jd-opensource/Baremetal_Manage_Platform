import {StatusParamsType} from '@utils/publicType';
import {paginationOperate} from 'utils/publicClass.ts';
import {DeleteUserType, EditUserType} from '../typeManagement';

interface DataType {
    listData: {[x: string]: {}};
    deleteUserDiaLog: Ref<boolean>;
};

const deleteOperate = <T extends {reactiveArr: {tableData: EditUserType[]}, getUserList(): void}, K extends {reset(): void}>(userList: T, refreshReset: K) => {
    const data: DataType = {
        listData: reactive<{[x: string]: {}}>({}),
        // 删除用户组件弹窗状态
        deleteUserDiaLog: ref<boolean>(false)
    };

    /**
     * 点击删除用户操作
     * @param {Object} item 当前点击的这一项
     * @return {boolean} addUserDialog.value 删除用户组件弹窗状态
    */
    const deleteUserOperate = async (item: DeleteUserType): Promise<boolean> => {
        const {userName, roleName, phoneNumber, userId, email}: DeleteUserType = item;
        await nextTick(() => {
            data.listData.value = {
                userName,
                roleName,
                phoneNumber,
                userId,
                email
            };
        });
        return data.deleteUserDiaLog.value = !data.deleteUserDiaLog.value;
    };

    /**
     * 删除操作取消事件
    */
    const deleteCancel: StatusParamsType <(status: boolean) => boolean> = (status: boolean): boolean => {
        return data.deleteUserDiaLog.value = status;
    };

    /**
     * 删除成功的回调
    */
    const deleteSure = () => {
        if (paginationOperate.pageNumber.value > 1 && !userList.reactiveArr.tableData.length) {
            paginationOperate.pageNumber.value = paginationOperate.pageNumber.value - 1;
            return;
        }
        const data = [
            [
                (len: number) => len === 1,
                () => refreshReset.reset()
            ],
            [
                (len: number) => len !== 1,
                () => userList.getUserList()
            ]
        ];
        for (const key of data) {
            if (key[0](userList.reactiveArr.tableData.length)) {
                key[1](userList.reactiveArr.tableData.length);
                break;
            }
        }
    };

    return {
        ...data,
        deleteUserOperate,
        deleteCancel,
        deleteSure,
    }
};

export default deleteOperate;
