<template>
    <export
        v-if="exportDataOperate.exportDataDiaLog.value"
        :dia-log="exportDataOperate.exportDataDiaLog.value"
        :event="exportData"
        @dia-log-close="exportDataOperate.exportCancel"
    />
</template>

<script setup lang="ts">
import {ExportData, exportDataOperate} from './export';
const exportData = new ExportData();
</script>
