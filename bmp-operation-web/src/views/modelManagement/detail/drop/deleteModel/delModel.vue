<template>
    <div v-if="delModel.deleteModelDiaLog.value">
        <delete-model
            :dia-log="delModel.deleteModelDiaLog.value"
            :model-info="delModel.modelInfo.value"
            :device-type-id="baseInfo.reactiveArr.detail.deviceTypeId"
            @dia-log-close="delModel.deleteModelCancel"
            @determine-click="router.push($defInfo.routerPath('modelList'))"
            @error-click="delModel.deleteModelError"
        />
    </div>
</template>
<script lang="ts" setup>
import {RouterOperate} from 'utils/publicClass.ts';
interface PropsType {
    delModel: any;
    baseInfo: any
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

// 路由-使用路由
const router = new RouterOperate().router;
</script>