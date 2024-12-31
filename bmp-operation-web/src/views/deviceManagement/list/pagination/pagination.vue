<template>
    <!-- 分页组件 -->
    <device-list-pagination
        :other-class="'flex-end-right'"
        :pagination-opt="deviceList"
    >
        <div class="check-all checkbox-text-ellipsis">
            <ui-checkbox
                slot="check-box"
                v-model="deviceList.hasCheckAll.value"
                @change="deviceList.checkAll"
            />
        </div>
        <div class="batch-operate">
            <!-- 关机 -->
            <ui-button
                type="primary"
                :disabled="deviceList.closeBtnDisabled.value"
                @click="batchOperate.openCloseRestartClick(!deviceList.closeBtnDisabled.value, 'close')"
            >
                <ui-tooltip
                    placement="bottom"
                    :content="batchOperate.closeTip(deviceList.reactiveArr.selectArr.length)"
                    :disabled="!deviceList.closeBtnDisabled.value"
                >
                    {{$t('deviceList.batchOperate.btn.close')}}
                </ui-tooltip>
            </ui-button>
            <!-- 开机 -->
            <ui-button
                type="primary"
                :disabled="deviceList.openBtnDisabled.value"
                @click="batchOperate.openCloseRestartClick(!deviceList.openBtnDisabled.value, 'open')"
            >
                <ui-tooltip
                    placement="bottom"
                    :content="batchOperate.openTip(deviceList.reactiveArr.selectArr.length)"
                    :disabled="!deviceList.openBtnDisabled.value"
                >
                    {{$t('deviceList.batchOperate.btn.open')}}
                </ui-tooltip>
            </ui-button>
            <!-- 下拉菜单 -->
            <ui-dropdown
                popper-class="custom-batch-operate"
                @visible-change="batchOperate.selectHover"
            >
                <!-- 更多 -->
                <template #dropdownTitle>
                    <p
                        class="batch-operate-device-list"
                        :style="batchOperate.selectStatus.value ? {borderColor: '#0f8fe9'} : {}"
                    >
                        <span :class="batchOperate.selectStatus.value ? 'text-color' : ''">
                            {{$t('deviceList.batchOperate.btn.more')}}
                        </span>
                        <span
                            class='arrow-bottom-img'
                            :style="{
                                backgroundImage: `url(${
                                    batchOperate.selectStatus.value ? $defInfo.imagePath('arrowBottom'): $defInfo.imagePath('defArrowBottom')
                                })`
                            }"
                        >
                        </span>
                    </p>
                </template>
                <!-- 重启 -->
                <ui-dropdown-item>
                    <!-- 重启提示 -->
                    <ui-tooltip
                        placement="right"
                        :disabled="!deviceList.restartBtnDisabled.value"
                        :content="batchOperate.restartTip(deviceList.reactiveArr.selectArr?.length)"
                    >
                        <!-- 重启 -->
                        <span
                            :class="batchOperate.setClass(deviceList.restartBtnDisabled.value)"
                            @click="batchOperate.openCloseRestartClick(!deviceList.restartBtnDisabled.value, 'restart')"
                        >
                            {{$t('deviceList.batchOperate.btn.restart')}}
                        </span>
                    </ui-tooltip>
                </ui-dropdown-item>
                <!-- 重置密码 -->
                <ui-dropdown-item>
                    <ui-tooltip
                        placement="right"
                        :disabled="!deviceList.resetPasswordBntDisabled.value"
                        :content="batchOperate.resetPasswordTip(deviceList)"
                    >
                        <span
                            :class="batchOperate.setClass(deviceList.resetPasswordBntDisabled.value)"
                            @click="emitValue('reset-password', deviceList.resetPasswordBntDisabled.value)"
                        >
                            {{$t('deviceList.batchOperate.btn.resetPassword')}}
                        </span>
                    </ui-tooltip>
                </ui-dropdown-item>
                <!-- 编辑实例名称 -->
                <ui-dropdown-item>
                    <ui-tooltip
                        placement="right"
                        :disabled="!deviceList.instanceNameBntDisabled.value"
                        :content="batchOperate.editInstanceNameTip(deviceList)"
                    >
                        <span
                            :class="batchOperate.setClass(deviceList.instanceNameBntDisabled.value)"
                            @click="emitValue('instance-name', deviceList.instanceNameBntDisabled.value)"
                        >
                            {{$t('deviceList.batchOperate.btn.edit')}}
                        </span>
                    </ui-tooltip>
                </ui-dropdown-item>
                <!-- 回收实例 -->
                <ui-dropdown-item>
                    <!-- 回收实例提示 -->
                    <ui-tooltip
                        placement="right"
                        :disabled="!deviceList.recoveryInstanceBtnDisabled.value"
                        :content="batchOperate.recoveryInstanceTip(deviceList.reactiveArr.selectArr)"
                    >
                        <!-- 回收实例 -->
                        <span
                            :class="batchOperate.setClass(deviceList.recoveryInstanceBtnDisabled.value)"
                            @click="batchOperate.recoveryInstanceClick(!deviceList.recoveryInstanceBtnDisabled.value)"
                        >
                            {{$t('deviceList.batchOperate.btn.recovery')}}
                        </span>
                    </ui-tooltip>
                </ui-dropdown-item>
            </ui-dropdown>
        </div>
    </device-list-pagination>
</template>

<script lang="ts" setup>
import BatchOperate from './batchOpt';
import usePagination from 'hooks/pagination/usePagination.ts';
import DeviceListPagination from 'hooks/pagination/paginationSelect.vue';

interface PropsType {
    deviceList: any;
    filter: any;
};


// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

usePagination(props.deviceList.getDataList, props.deviceList.filter.tableRef?.value);

const emitValue = defineEmits(['open-close-restart', 'recovery-instance', 'instance-name', 'reset-password']);

const batchOperate = new BatchOperate(emitValue);

</script>
<style lang="scss" scoped>
@import 'assets/css/batchOperatePagination.scss';
</style>