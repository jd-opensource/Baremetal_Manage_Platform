<template>
    <!-- 机房详情 -->
    <div class="operate-management-detail">
        <!-- 头部信息 -->
        <detail-header
            :name="idcDetail.reactiveArr.detailInfo.name"
            @header="router.jumpRouter"
        >
            <!-- 编辑 -->
            <ui-dropdown-item>
                <span
                    class="drop-down-operate-content"
                    @click="validate.operateClick"
                >
                    {{$t('idcDetail.header.operate.edit')}}
                </span>
            </ui-dropdown-item>
        </detail-header>
        <!-- 主体内容信息 -->
        <main class="operate-management-detail-info">
            <!-- 基本信息 -->
            <pluginComp.BasicInfo :idc-detail="idcDetail"/>
            <!-- 分割线 -->
            <div class="split-line"></div>
            <!-- 全局管理头部信息 -->
            <pluginComp.TableData :show-password="showPassword" :idc-detail="idcDetail"/>
        </main>
        <!-- 编辑机房-安全验证 -->
        <pluginComp.Verification :operate="validate" :edit-idc-operate="editIdcOpt"/>
        <!-- 编辑机房 -->
        <pluginComp.EditIdcData :idc-detail="idcDetail" :edit-idc-operate="editIdcOpt"/>
    </div>
</template>
<script setup lang="ts">
import * as all from './methods';
import pluginComp from './module';
import {RouterType} from './typeManagement';
import {RouterOperate, CurrentInstance} from 'utils/publicClass.ts';

const instanceProxy = new CurrentInstance().proxy;

const path: string = instanceProxy.$defInfo.routerPath('idcList');

// 路由
const router: RouterType = new RouterOperate(path);

const idcDetail = new all.IdcDetailData(router);

const editIdcOpt = all.editIdcOperate(idcDetail);

const validate = all.verificationOperate(editIdcOpt);

const showPassword = all.hasShowPassword(validate, idcDetail);

</script>
<style lang="scss">
@import 'assets/css/detail.scss';
</style>
