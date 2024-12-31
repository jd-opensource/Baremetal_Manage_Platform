<template>
    <!-- 主体内容 -->
    <article class="operate-management-count no-top role-list">
        <ui-table
            style="width: 100%;"
            v-loading="roleList.isLoading.value"
            border
            :class="tableClass(roleList.reactiveArr.tableData, roleList.isLoading.value)"
            :max-height="tableScrollOperate.tableMaxHeight.value"
            :data="roleList.reactiveArr.tableData"
        >
            <!-- 角色 -->
            <ui-table-column
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['roleName']?.selected"
                :label="$t('roleList.label.role')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span :class="setTextClass(false)">
                        {{$filter.emptyFilter(scope.row.roleName)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 关联用户 -->
            <ui-table-column
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['userName']?.selected"
                :has-user-template="true"
                :label="$t('roleList.label.relationUser')"
            >
                <template #default="{scope}">
                    <!-- 平台拥有者默认 -->
                    <span
                        v-if="scope.row.roleId === $defInfo.purview('admin')"
                        :class="setTextClass(false)"
                    >
                        admin
                    </span>
                    <!-- 跳转用户列表 -->
                    <span
                        v-else-if="scope.row.userCount > 0"
                        :class="setTextClass(true)"
                        @click="routerOperate.jumpRouter({params: scope.row.roleId, status: true, type: 'roleId'})"
                    >
                        {{$filter.emptyFilter(scope.row.userCount)}}
                    </span>
                    <span
                        v-else
                        :class="setTextClass(false)"
                    >
                        {{scope.row.userCount}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 权限 -->
            <ui-table-column
                align="center"
                v-if="customOperate?.reactiveArr?.hasCustomInfo['operate']?.selected"
                :has-user-template="true"
                :min-width="setDiffClass('280', '130')" 
                :label="$t('roleList.label.jurisdiction')"
            >
                <template #default="{scope}">
                    {{$filter.emptyFilter(scope.row.permission)}}
                </template>
            </ui-table-column>
        </ui-table>
    </article>
</template>

<script setup lang="ts">
import SetEmpty from './setEmpty';
import customOperate from '../custom/custom';
import {RouterOperate, CurrentInstance} from 'utils/publicClass.ts';
import {tableClass, setTextClass, setDiffClass} from 'utils/index.ts'; // 表格class类

interface PropsType {
    roleList: any;
    tableScrollOperate: any;
};

const instanceProxy = new CurrentInstance().proxy;

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

// 实例化-路由操作
const routerOperate = new RouterOperate(instanceProxy.$defInfo.routerPath('user'));

new SetEmpty(props);

</script>
