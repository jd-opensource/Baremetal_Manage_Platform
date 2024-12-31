<template>
    <div v-if="exportDataOperate.exportDataDiaLog.value">
        <export
            :dia-log="exportDataOperate.exportDataDiaLog.value"
            :loading="exportDataOperate"
            :event="exportFilterData"
            :search="exportFilterData"
            :has-search="true"
            @dia-log-close="exportDataOperate.exportCancel"
        />
    </div>
</template>

<script setup lang="ts">
interface PropsType {
    exportFilterData: any;
    exportDataOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
