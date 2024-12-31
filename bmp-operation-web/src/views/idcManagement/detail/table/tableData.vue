<template>
    <!-- 全局管理头部信息 -->
    <div class="info-header idc-administration">
        <span class="info-line"></span>
        <!-- 全局管理 -->
        <span class="info-title">
            {{$t('idcDetail.allAdministration.title')}}
        </span>
    </div>
    <!-- 全局管理表格信息 -->
    <div class="operate-management-count idc-detail">
        <ui-table
            style="width: 100%;"
            v-loading="idcDetail.isLoading.value"
            border
            :class="tableClass(idcDetail.reactiveArr.tableData, idcDetail.isLoading.value)"
            :data="idcDetail.reactiveArr.tableData"
        >
            <!-- 名称 -->
            <ui-table-column
                prop="name"
                align="center"
                :label="$t('idcDetail.allAdministration.name')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.name)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 登录账号 -->
            <ui-table-column
                prop="user"
                align="center"
                :label="$t('idcDetail.allAdministration.loginUserName')"
                :has-user-template="true"
            >
                <template #default="{scope}">
                    <span>
                        {{$filter.emptyFilter(scope.row.user)}}
                    </span>
                </template>
            </ui-table-column>
            <!-- 登录密码 -->
            <ui-table-column
                prop="password"
                align="center"
                :has-user-template="true"
                :label="$t('idcDetail.allAdministration.loginPassWord')"
            >
                <template #default="{scope}">
                    <!-- 密码-明文-加密-eyeicon -->
                    <p class="login-password">
                        <!-- 通过传入的状态，判断密码显示状态 -->
                        <span
                            :style="scope.row.password !== '--' ? '' : {color: '#333', cursor: 'text'}"
                            :class="[scope.row.hasPassword ? 'set-size-small' : 'set-size']"
                        >
                            {{$filter.defaultPassword(scope.row.password)}}
                        </span>
                        <!-- open-eye -->
                        <img
                            class="login-password-img"
                            alt=""
                            v-if="!scope.row.hasPassword"
                            :src="($defInfo.imagePath('closeEye') as unknown as string)"
                            @click="showPassword.hasPasswordClick(scope.row, 'open')"
                        />
                        <!-- close-eye -->
                        <img
                            class="login-password-img"
                            alt=""
                            v-else
                            :src="($defInfo.imagePath('openEye') as unknown as string)"
                            @click="showPassword.hasPasswordClick(scope.row, 'close')"
                        />
                    </p>
                </template>
            </ui-table-column>
        </ui-table>
    </div>
</template>

<script lang="ts" setup>
import {tableClass} from 'utils/index.ts'; // 详情表格class类
import SetEmpty from './setEmpty';

interface PropsType {
    idcDetail: any;
    showPassword: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});
new SetEmpty(props);
</script>