<template>
    <!-- 用户管理-用户列表 -->
    <div class="operate-management">
        <!-- 头部标题内容 -->
        <header-info @header-ref="userCount.tableScrollOperate.getHeaderRef"/>
        <!-- 主体内容 -->
        <main class="operate-management-content">
            <div class="search-tip-operate2">
                <!-- 操作 -->
                <refresh-custom-export
                    :status="userCount.userList.searchTip.value"
                    :search-operate="userCount.search"
                    :place-holder="[
                        $t('userList.search.placeholder.userName')
                    ]"
                    @operate-ref="userCount.tableScrollOperate.getOperateRef"
                    @refresh="userCount.refreshReset.refresh"
                    @custom="customOperate.default?.customClickOperate"
                    @export="userCount.exportDataOperate.diaLogClick"
                    @btn-ref="userCount.tableScrollOperate.getBtnRef"
                    @btn-operate="userCount.addUserOperate.addUserClick"
                />
            </div>
            <!-- 表格数据-数据列表 -->
            <article class="operate-management-count no-top user-list">
                <no-data-tip
                    :status="userCount.userList.searchTip.value"
                    :total="paginationOperate.total.value"
                    @back-click="userCount.refreshReset.reset"
                />
                <table-list
                    :user-list="userCount.userList"
                    :table-max-height="userCount.tableScrollOperate.tableMaxHeight.value"
                    :filter-operates="userCount.filter"
                    @edit-user="userCount.editUser.editUserClick"
                    @delete-user="userCount.deleteUser.deleteUserOperate"
                    @empty-click="userCount.addUserOperate.addUserClick"
                />
            </article>
        </main>
        <!-- 分页组件 -->
        <custom-pagination
            :user-list="userCount.userList"
            :search="userCount.search"
            :filter="userCount.filter"
        />
        <!-- 编辑用户组件 -->
        <user-edit :edit-operate="userCount.editUser"/>
        <!-- 添加用户组件 -->
        <user-add :add-user-operate="userCount.addUserOperate"/>
        <!-- 删除用户组件 -->
        <user-delete :delete-operate="userCount.deleteUser"/>
        <!-- 自定义列表组件 -->
        <custom/>
        <!-- 导出数据组件 -->
        <export-data
            :export-filter-data="userCount.exportFilterData"
            :export-data-operate="userCount.exportDataOperate"
        />
    </div>
</template>

<script setup lang="ts">
import pluginComp from './module';
import methods from './methods';
import {methodsTotal} from 'utils/index.ts';
import {CurrentInstance, paginationOperate} from 'utils/publicClass.ts';

const [Custom, ExportData, CustomPagination, TableList, UserAdd, UserDelete, UserEdit] = pluginComp.map((item) => item.default);
const [customOperate, exportInfo, SearchOperate, refreshResetOperate, FilterOperate, TableScrollOperate, unmountOperate, UserList, addUserInfo, deleteOperate, editOperate] = methods;

const user = () => {

    unmountOperate.default();

    const userList = new UserList.default();

    const proxy = new CurrentInstance().proxy;

    const editUser = editOperate.default(userList);

    const search = new SearchOperate.default(userList);

    const filter = new FilterOperate.default(userList);

    const exportDataOperate = new exportInfo.ExportDataOperate(proxy.$userApi.userListExportAPI);

    const exportFilterData = new exportInfo.ExportFilterData(search.reactiveArr, filter.reactiveArr);

    const refreshReset = refreshResetOperate.default(userList, filter, search);

    const addUserOperate = addUserInfo.default(refreshReset);

    const deleteUser = deleteOperate.default(userList, refreshReset);

    const tableScrollOperate = new TableScrollOperate.default(filter, search, userList.searchTip);

    const useListenMsg = methodsTotal.listenMsg((info: {type: string; content: {type: string; userId: string;} & string;}) => {
        if (info.type === 'update-user-info') {
            const val = userList.reactiveArr.tableData;
            const id = val.findIndex((item: {userId: string}) => item.userId === info.content.userId);
            if (id >= 0) {
                val[id] = info.content;
                userList.setList(val, 'edit');
            }
        }
        // else if (info.type === 'delete-user-info') {
        //     const val = userList.reactiveArr.tableData;
        //     const id = val.findIndex((item: {userId: string}) => item.userId === info.content);
        //     if (id >= 0) {
        //         // val[id] = info.content;
        //         userList.reactiveArr.tableData.splice(info.content, 1);
        //         userList.setList( userList.reactiveArr.tableData, 'delete');
        //     }
        // }
    })

    onUnmounted(useListenMsg);

    return {
        userList,
        editUser,
        search,
        filter,
        exportDataOperate,
        exportFilterData,
        refreshReset,
        addUserOperate,
        deleteUser,
        tableScrollOperate
    }
};

const userCount = user();
</script>

<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
</style>
