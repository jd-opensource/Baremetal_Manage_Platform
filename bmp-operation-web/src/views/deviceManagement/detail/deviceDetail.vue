<template>
    <!-- 设备详情 -->
    <div class="operate-management-detail device-detail">
        <detail-header
            :name="$filter.emptyFilter($route.query.sn as string)"
            :title="'deviceDetail'"
            @header="deviceDetail.routerJump"
            @header-ref="tableS?.getHeaderRef"
        >
            <drop-down-operate
                :detail="deviceDetail?.reactiveArr.detail"
                :open-close-restart="openCloseRestart"
                :up-down-delete-operate="upDownDeleteOperate"
                :remove-recovery-operate="removeRecoveryOperate"
                :reset-instance-operate="resetInstanceOperate"
                :lock-deblocking="lockDeblocking"
                :reset-password-opt="passwordResetOpt"
                :reset-system-opt="resetSystemConfirmOpt"
            />
        </detail-header>
         <tabs-data
            :tabs="activeOperate"
            @header-ref="tableS.getRef"
        />
       <custom-device-detail
            v-if="activeOperate.activeName.value === 'basicInfo'"
            :device-detail="deviceDetail"
            :notes-edit="notesEdit"
            :collect="collectConfirmOpt"
        />
        <div v-else-if="activeOperate.activeName.value === 'diskDetail'">
            <disk-table
                :detail="(deviceDetail.reactiveArr.detail as any)"
                :disk-table-opt="diskTableOpt"
                @add-models="associatedModelOpt.associatedModelClick"
            />
        </div>
        <!-- 操作日志 -->
        <div
            class="opt-logs-count"
            v-else-if="activeOperate.activeName.value === 'operatLog'"
        >
            <operate-log
                :sn="($route.query.sn as unknown as string)"
                :device-detail="deviceDetail"
            />
        </div>
        <div v-else-if="activeOperate.activeName.value === 'performanceMonitoring'">
            <performance-monitoring :device-detail="deviceDetail"/>
        </div>
        <div
            class="operate-management-count hardware-monitoring"
            v-else
        >
           <hardware-list
                :query="query"
                :table-s="tableS"
                :device-detail="deviceDetail"
                :activeOperate="activeOperate"
            />
        </div>
        <custom-verification :operate="verification"/>
        <custom-up-down-delete :operate="upDownDeleteOperate"/>
        <custom-notes-edit
            :operate="notesEdit"
            :description="deviceDetail.reactiveArr.detail.description"
        />
        <custom-open-close-restart :operate="openCloseRestart"/>
        <recovery-remove :operate="removeRecoveryOperate"/>
        <instance-reset :operate="resetInstanceOperate"/>
        <lock-un-lock :operate="lockDeblocking"/>
        <reset-system-confirm :operate="resetSystemConfirmOpt"/>
        <system-reset :operate="resetSystemOpt"/>
        <custom-password-reset :operate="passwordResetOpt"/>
        <collect-confirm
            :operate="collectConfirmOpt"
            :device-detail="deviceDetail"
        />
        <clear-raid
            :device-detail="deviceDetail"
            :operate="raidClearOpt"
        />
        <models-associated
            :device-id="($route.query.deviceId as string)"
            :operate="associatedModelOpt"
            :disk-table-opt="diskTableOpt"
        />
    </div>
</template>
<script setup lang="ts">
import methods from './methods';
import pluginComp from './module';
import {RouterOperate} from 'utils/publicClass.ts';
const [DropDownOperate, InstanceReset, LockUnLock, CustomOpenCloseRestart, CustomPasswordReset, RecoveryRemove, ResetSystemConfirm, SystemReset, CustomUpDownDelete, TabsData, DiskTable, PerformanceMonitoring, ClearRaid, CollectConfirm, CustomDeviceDetail, CustomNotesEdit, CustomVerification, ModelsAssociated, OperateLog, HardwareList] = pluginComp;

const [resetInstanceOpt, lockDeblockingOpt, openCloseRestartOpt, passwordReset, removeRecoveryOpt, resetSystemConfirmOperate, resetSystemOperate, upDownDeleteOpt, tabsOperate, DiskTableOpt, RaidClear, CollectOpt, DeviceDetail, notesEditOperate, verificationOperate, AssociatedModelOpt, TableScrollOperate] = methods;
// 实例化-路由操作
const routerOperate = new RouterOperate().router;

const query = routerOperate.currentRoute.value.query;

const verification = verificationOperate(() => deviceDetail.getDeviceDetail());

const deviceDetail = new DeviceDetail(verification);

const diskTableOpt = new DiskTableOpt(query);

const activeOperate = tabsOperate(deviceDetail, query, diskTableOpt);

const notesEdit = notesEditOperate(deviceDetail);

const lockDeblocking = lockDeblockingOpt(deviceDetail);

const openCloseRestart = openCloseRestartOpt(deviceDetail);

const upDownDeleteOperate = upDownDeleteOpt(deviceDetail);

const resetInstanceOperate = resetInstanceOpt(deviceDetail);

const removeRecoveryOperate = removeRecoveryOpt(deviceDetail);

const collectConfirmOpt = new CollectOpt(deviceDetail, (type: boolean) => raidClearOpt.clearInfoClick(type));

const raidClearOpt = new RaidClear(collectConfirmOpt, deviceDetail)

const resetSystemOpt = resetSystemOperate(deviceDetail);

const resetSystemConfirmOpt = resetSystemConfirmOperate(resetSystemOpt);

// new all.PaginationOperate(faultLogTable);

const passwordResetOpt = passwordReset(deviceDetail);

// surveillanceOpt.isLoading
const tableS = new TableScrollOperate();

const associatedModelOpt = new AssociatedModelOpt(deviceDetail);

</script>
<style lang="scss">
@import './deviceDetail.scss';

</style>