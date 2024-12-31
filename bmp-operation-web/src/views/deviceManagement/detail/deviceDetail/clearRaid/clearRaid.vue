<template>
    <div v-if="operate.clearInfoDiaLog.value">
        <clear-raid-confirm
            :dia-log="operate.clearInfoDiaLog.value"
            :sn="deviceDetail.reactiveArr.detail?.sn"
            :checkbox-status="operate.checkboxStatus.value"
            @dia-log-close="operate.clearRaidCancel"
            @determine-click="operate.clearRaidSure"
        >
        </clear-raid-confirm>
    </div>
</template>
<script lang="ts" setup>
// import {RouterOperate} from 'utils/publicClass.ts';

interface PropsType {
    operate: any;
    deviceDetail: any
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
