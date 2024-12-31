<template>
    <fault-rules-list-custom-list :custom-operate="customOperate"/>
</template>
<script setup lang="ts">
import FaultRulesListCustomList from 'hooks/custom/custom.vue';
import customOperate from './custom';

</script>
