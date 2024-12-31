<template>
    <fault-list-custom-list :custom-operate="customOperate"/>
</template>
<script setup lang="ts">
import FaultListCustomList from 'hooks/custom/custom.vue';
import customOperate from './custom';

</script>
