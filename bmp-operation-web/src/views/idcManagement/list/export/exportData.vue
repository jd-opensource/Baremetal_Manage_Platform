<template>
    <export
        v-if="exportDataOperate.exportDataDiaLog.value"
        :dia-log="exportDataOperate.exportDataDiaLog.value"
        :event="exportData"
        @dia-log-close="exportDataOperate.exportCancel"
    />
</template>

<script lang="ts" setup>
import {ExportData} from './export';
const exportData = new ExportData();

interface PropsType {
    exportDataOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>