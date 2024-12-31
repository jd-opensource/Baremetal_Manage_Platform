<template>
    <ui-dropdown-item>
        <!-- 上架提示 -->
        <ui-tooltip
            placement="left"
            v-if="!dropDownOperate.up(detail?.manageStatus)"
            :disabled="dropDownOperate.up(detail?.manageStatus)"
            :content="$t('deviceDetail.tooltip.up')"
        >
            <!-- 上架 -->
            <span
                :class="dropDownOperate.up(detail?.manageStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="upDownDeleteOperate.upDownFrameClick('up', dropDownOperate.up(detail?.manageStatus))"
            >
                {{$t('deviceDetail.operate.up')}}
            </span>
        </ui-tooltip>
        <ui-tooltip
            placement="left"
            v-else
            :disabled="!!detail.deviceTypeId"
            :content="$t('deviceDetail.tooltip.model')"
        >
            <!-- 上架 -->
            <span
                :class="detail.deviceTypeId?.length
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="upDownDeleteOperate.upDownFrameClick('up', detail.deviceTypeId.length)"
            >
                {{$t('deviceDetail.operate.up')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <ui-dropdown-item>
        <!-- 下架提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.down(detail?.manageStatus)"
            :content="$t('deviceDetail.tooltip.down')"
        >
            <!-- 下架 -->
            <span
                :class="dropDownOperate.down(detail?.manageStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="upDownDeleteOperate.upDownFrameClick('down', dropDownOperate.down(detail?.manageStatus))"
            >
                {{$t('deviceDetail.operate.down')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <!-- 开机 -->
    <ui-dropdown-item v-if="dropDownOperate.stoppedStatus(detail?.instanceStatus)">
        <span
            class="drop-down-operate-content"
            @click="openCloseRestart.openShutDownRestartClick(detail?.instanceName, detail?.instanceId, 'open', dropDownOperate.stoppedStatus(detail?.instanceStatus))"
        >
            {{$t('deviceDetail.operate.open')}}
        </span>
    </ui-dropdown-item>
    <!-- 关机 -->
    <ui-dropdown-item v-else>
        <!-- 关机提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.running(detail?.instanceStatus)"
            :content="$t('deviceDetail.tooltip.turnOff')"
        >
            <!-- 关机 -->
            <span 
                :class="dropDownOperate.running(detail?.instanceStatus)
                ? 'drop-down-operate-content'
                : 'drop-down-disabled'"
                @click="openCloseRestart.openShutDownRestartClick(detail?.instanceName, detail?.instanceId, 'close', dropDownOperate.running(detail?.instanceStatus))"
            >
                {{$t('deviceDetail.operate.close')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <!-- 重启 -->
    <ui-dropdown-item>
        <!-- 重启提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.running(detail?.instanceStatus)"
            :content="$t('deviceDetail.tooltip.restart')"
        >
            <!-- 重启 -->
            <span
                :class="dropDownOperate.running(detail?.instanceStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="openCloseRestart.openShutDownRestartClick(detail?.instanceName, detail?.instanceId, 'restart', dropDownOperate.running(detail?.instanceStatus))"
            >
                {{$t('deviceDetail.operate.restart')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <!-- 重置实例状态 -->
    <ui-dropdown-item>
        <!-- 重置实例状态提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.resetInstance(detail?.instanceStatus)"
            :content="$t('deviceDetail.tooltip.reset')"
        >
            <!-- 重置实例状态 -->
            <span
                :class="dropDownOperate.resetInstance(detail?.instanceStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="resetInstanceOperate.resetInstanceClick(detail?.instanceId, detail?.instanceName, dropDownOperate.resetInstance(detail?.instanceStatus))"
            >
                {{$t('deviceDetail.operate.reset')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <!-- 设备移除 -->
    <ui-dropdown-item>
        <!-- 设备移除状态提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.deviceRemove(detail?.instanceStatus)"
            :content="$t('deviceDetail.tooltip.remove')"
        >
            <!-- 设备移除状态 -->
            <span
                :class="dropDownOperate.deviceRemove(detail?.instanceStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="removeRecoveryOperate.removeRecoveryClick(detail, 'remove', dropDownOperate.deviceRemove(detail?.instanceStatus))"
            >
                {{$t('deviceDetail.operate.remove')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <!-- 回收实例状态 -->
    <ui-dropdown-item>
        <!-- 回收实例状态提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.recoveryInstance(detail?.instanceStatus, detail?.locked)"
            :content="detail?.locked === 'locked' ? $t('deviceDetail.tooltip.lock') : $t('deviceDetail.tooltip.recovery')"
        >
            <!-- 回收实例状态 -->
            <span
                :class="dropDownOperate.recoveryInstance(detail?.instanceStatus, detail?.locked)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="removeRecoveryOperate.removeRecoveryClick(detail, 'recovery', dropDownOperate.recoveryInstance(detail?.instanceStatus, detail?.locked))"
            >
                {{$t('deviceDetail.operate.recovery')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <ui-dropdown-item>
        <!-- 锁定 -->
        <span
            class="drop-down-operate-content"
            v-if="detail?.locked === 'unlocked' && !DeviceStaticData.lockUnlocked.includes(detail?.instanceStatus)"
            @click="lockDeblocking.lock(detail?.instanceName, detail?.instanceId, 'lock')"
        >
            {{$t('deviceDetail.operate.lock')}}
        </span>
        <ui-tooltip
            placement="left"
            v-else-if="detail?.locked === 'unlocked'"
            :content=" $t('deviceDetail.tooltip.lock1')"
        >
            <!-- 锁定 -->
            <span class="drop-down-disabled">
                {{$t('deviceDetail.operate.lock')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <ui-dropdown-item>
        <!-- 解锁 -->
        <span
            class="drop-down-operate-content"
            v-if="detail?.locked === 'locked' && !DeviceStaticData.lockUnlocked.includes(detail?.instanceStatus)"
            @click="lockDeblocking.lock(detail?.instanceName, detail?.instanceId, 'unlock')"
        >
            {{$t('deviceDetail.operate.unlock')}}
        </span>
        <ui-tooltip
            placement="left"
            v-else-if="detail?.locked === 'locked'"
            :content=" $t('deviceDetail.tooltip.lock2')"
        >
            <!-- 解锁 -->
            <span class="drop-down-disabled">
                {{$t('deviceDetail.operate.unlock')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <ui-dropdown-item>
        <!-- 重置密码提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.resetPasswordStatus(detail?.instanceStatus)"
            :content="$t('deviceDetail.tooltip.resetPassword')"
        >
            <!-- 重置密码 -->
            <span
                :class="dropDownOperate.resetPasswordStatus(detail?.instanceStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="resetPasswordOpt.resetPasswordClick(detail)"
            >
                {{$t('deviceDetail.operate.resetPassword')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <ui-dropdown-item>
        <!-- 重装系统提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.resetSystemStatus(detail?.instanceStatus)"
            :content="$t('deviceDetail.tooltip.resetSystem')"
        >
            <!-- 重装系统 -->
            <span
                :class="dropDownOperate.resetSystemStatus(detail?.instanceStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="resetSystemOpt.resetSystemConfirmClick(detail)"
            >
                {{$t('deviceDetail.operate.resetSystem')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <ui-dropdown-item>
        <!-- 删除提示 -->
        <ui-tooltip
            placement="left"
            :disabled="dropDownOperate.deviceDelete(detail?.manageStatus)"
            :content="$t('deviceDetail.tooltip.delete')"
        >
            <!-- 删除 -->
            <span
                :class="dropDownOperate.deviceDelete(detail?.manageStatus)
                    ? 'drop-down-operate-content'
                    : 'drop-down-disabled'"
                @click="upDownDeleteOperate.upDownFrameClick('delete', dropDownOperate.deviceDelete(detail?.manageStatus))"
            >
                {{$t('deviceDetail.operate.delete')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
</template>

<script lang="ts" setup>
import DropDownOpt from './dropDownOpt';
import DeviceStaticData from 'staticData/device/index.ts';

// defineEmits API 来替代 emits
interface PropsType {
    detail: any;
    openCloseRestart: any;
    upDownDeleteOperate: any;
    removeRecoveryOperate: any;
    resetInstanceOperate: any;
    lockDeblocking: any;
    resetPasswordOpt: any;
    resetSystemOpt: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

const dropDownOperate = new DropDownOpt();
</script>
