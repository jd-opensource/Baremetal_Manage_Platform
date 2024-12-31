<template>
    <!-- 设备管理-设备列表 -->
    <div class="operate-management">
        <!-- header宽度 -->
        <header-info
            :type="'deviceList'"
            @header-ref="tableScroll.getHeaderRef"
        />
        <!-- 设备主体内容 -->
        <main class="operate-management-content">
            <div class="search-tip-operate2">
                <!-- 操作 -->
                <refresh-custom-export
                    :type="'deviceList.operate.exportDevice'"
                    :status="deviceList.searchTip.value"
                    :search-operate="search"
                    :place-holder="[
                        $t('deviceList.operate.placeholder.sn'),
                        $t('deviceList.operate.placeholder.instanceOwner')
                    ]"
                    @operate-ref="tableScroll.getOperateRef"
                    @refresh="refresh.refresh"
                    @custom="all.custom.customClickOperate"
                    @export="exportDataOperate.diaLogClick"
                    @btn-operate="importDevice.importDeviceClick"
                    @btn-ref="tableScroll.getBtnRef"
                />
            </div>
            <!-- 列表内容 -->
            <article class="operate-management-count device-list">
                <no-data-tip
                    :status="deviceList.searchTip.value"
                    :total="all.paginationOperate.total.value"
                    @back-click="reset.reset"
                />
                <pluginComp.TableData
                    :device-list="deviceList"
                    :table-max-height="tableScroll.tableMaxHeight.value"
                    :filter-operate="filter"
                    :search="search"
                    @up-down-delete="upDownDel.upDownDeleteClick"
                    @open-close-restart="openCloseRestartOperate.openShutDownRestartClick"
                    @reset-instance="instanceReset.resetInstanceClick"
                    @remove-recovery="removeRecoveryOperate.removeRecoveryClick"
                    @lock-nolock="lockUnLock.lock"
                    @import-device="importDevice.importDeviceClick"
                    @empty-click="importDevice.importDeviceClick"
                    @reset-system="resetSystemConfirmOpt.resetSystemConfirmClick"
                    @reset-password="resetPasswordOpt.resetPasswordClick"
                />
            </article>
        </main>
        <pluginComp.Pagination
            :device-list="deviceList"
            :filter="filter"
            @open-close-restart="batchOperate1.batchOpenCloseRestartClick"
            @recovery-instance="batchOperate2.batchRecoveryInstanceClick"
            @instance-name="batchOperate3.batchEditInstanceNameClick"
            @reset-password="batchOperate4.batchResetPasswordClick"
        />
        <pluginComp.Custom/>
        <pluginComp.OpenCloseRestart :operate="openCloseRestartOperate"/>
        <pluginComp.UpDownDelete :up-down-delete="upDownDel"/>
        <pluginComp.DeviceImport :operate="importDevice"/>
        <pluginComp.RecoveryRemove :operate="removeRecoveryOperate"/>
        <pluginComp.InstanceResets :instance-reset="instanceReset"/>
        <pluginComp.ExportData
            :export-filter-data="exportFilterData"
            :export-data-operate="exportDataOperate"
        />
        <pluginComp.LockOperate :operate="lockUnLock"/>
        <pluginComp.BatchOpenCloseRestartOperate
            :device-list="deviceList"
            :operate="batchOperate1"
        />
        <pluginComp.BatchRecoveryInstanceOperate
            :device-list="deviceList"
            :operate="batchOperate2"
        />
        <pluginComp.BatchEditInstanceNameOperate
            :device-list="deviceList"
            :operate="batchOperate3"
        />
        <pluginComp.BatchPasswordResetOperate
            :device-list="deviceList"
            :operate="batchOperate4"
        />
        <pluginComp.ResetSystemConfirm
            :operate="resetSystemConfirmOpt"
        />
        <pluginComp.SystemReset
            :operate="resetSystemOpt"
        />
        <pluginComp.PasswordReset :operate="resetPasswordOpt"/>
    </div>
</template>

<script setup lang="ts">
import * as all from './all';
import pluginComp from './module';

const search = new all.SearchOperate(() => deviceList.getDataList());

const filter = new all.FilterOperate(() => deviceList.getDataList());

const deviceList = new all.DeviceList(search, filter);

const batchOperate1 = all.batchOpenCloseRestart(deviceList);

const batchOperate2 = all.batchRecoveryInstance(deviceList);

const batchOperate3 = all.batchEditInstanceName(deviceList);

const batchOperate4 = all.batchPasswordReset(deviceList);

const reset = all.resetOperate(filter, search);

const importDevice = all.deviceImport(reset);

const lockUnLock = all.lockDeblocking(deviceList);

const instanceReset = all.resetInstance(deviceList);

// all.customPagination(deviceList)

const refresh = all.refreshOperate(deviceList, reset);

const upDownDel = all.upDownDelete(deviceList, reset);

const removeRecoveryOperate = all.removeRecovery(deviceList);

const openCloseRestartOperate = all.openCloseRestart(deviceList);

const tableScroll = all.tableScrollOperate(filter.reactiveArr, search.reactiveArr, deviceList.searchTip);

const exportDataOperate = new all.ExportDataOperate(all.devicesListExportAPI);
const exportFilterData = new all.ExportFilterData(search.reactiveArr, filter.reactiveArr);

const resetSystemOpt = all.resetSystemOperate(deviceList);
const resetSystemConfirmOpt = all.resetSystemConfirmOperate(resetSystemOpt);

const resetPasswordOpt = all.passwordReset(deviceList);

</script>

<style lang="scss" scoped>
@import 'assets/css/listSearch.scss';
</style>
