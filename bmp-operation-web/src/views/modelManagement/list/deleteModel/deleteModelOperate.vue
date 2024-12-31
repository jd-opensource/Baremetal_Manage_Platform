<template>
    <div v-if="deleteModelOperate.deleteModelDiaLog.value">
        <delete-model
            :dia-log="deleteModelOperate.deleteModelDiaLog.value"
            :model-info="deleteModelOperate.modelName.value"
            :device-type-id="deleteModelOperate.deviceTypeId.value"
            @dia-log-close="deleteModelOperate.deleteModelCancel"
            @determine-click="deleteModelOperate.deleteModelSure"
            @error-click="deleteModelOperate.deleteModelError"
        />
    </div>
</template>
<script lang="ts" setup>
interface PropsType {
    deleteModelOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>