<template>
    <div v-if="operate.batchEditInstanceNameDiaLog.value">
        <batch-edit-instance-name
            :dia-log="operate.batchEditInstanceNameDiaLog.value"
            :select-arr="deviceList.reactiveArr.defTableData"
            @dia-log-close="operate.batchEditInstanceNameCancel"
            @determine-click="operate.batchEditInstanceNameSure"
        />
    </div>
</template>
<script setup lang="ts">
interface PropsType {
    operate: any;
    deviceList: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
