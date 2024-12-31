<template>
    <div v-if="operate.upDownFrameDiaLog.value">
        <up-down-frame
            :dia-log="operate.upDownFrameDiaLog.value"
            :device-id="routerOperate.router?.currentRoute?.value?.query?.deviceId"
            :title="operate.upDownFrameTitle.value"
            @dia-log-close="operate.upDownDeleteCancel"
            @determine-click="operate.operateClick"
        />
    </div>
</template>
<script lang="ts" setup>
import {RouterOperate} from 'utils/publicClass.ts';
// import OpenCloseRestart from './operate';
// defineEmits API 来替代 emits
interface PropsType {
    operate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
// const operate = new OpenCloseRestart(props.deviceDetail);

const routerOperate = new RouterOperate();
</script>