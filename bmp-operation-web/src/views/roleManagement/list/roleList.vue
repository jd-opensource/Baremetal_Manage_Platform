<template>
    <!-- 角色管理-角色列表 -->
    <div class="operate-management">
        <!-- 头部标题内容 -->
        <header-info
            :type="'roleList'"
            @header-ref="tableScrollOperate.getHeaderRef"
        />
        <!-- 操作内容 -->
        <main class="operate-management-content">
            <div class="search-tip-operate2">
                <!-- 操作 -->
                <refresh-custom-export
                    :has-btn-operate="false"
                    :has-export="false"
                    :other-class="'set-flex-end'"
                    @operate-ref="tableScrollOperate.getOperateRef"
                    @refresh="roleList.getRoleList"
                    @custom="all.customOperate.customClickOperate"
                />
            </div>
            <!-- 主体内容 -->
            <role-list-table
                :role-list="roleList"
                :table-scroll-operate="tableScrollOperate"
            />
        </main>
        <!-- 分页组件 -->
        <role-list-pagination :pagination-opt="roleList"/>
        <!-- 自定义组件 -->
        <custom/>
    </div>
</template>

<script setup lang="ts">
import * as all from './all';
import Custom from './custom/customList.vue';
import RoleListTable from './table/roleList.vue';
import usePagination from 'hooks/pagination/usePagination.ts';
import RoleListPagination from 'hooks/pagination/paginationSelect.vue';

const roleList = new all.RoleList();

const tableScrollOperate = new all.TableScrollOperate();

usePagination(roleList.getRoleList);

</script>

<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
</style>
