<template>
    <fault-log-list-custom-list :custom-operate="customOperate"/>
</template>
<script setup lang="ts">
import FaultLogListCustomList from 'hooks/custom/custom.vue';
import customOperate from './custom';

</script>
