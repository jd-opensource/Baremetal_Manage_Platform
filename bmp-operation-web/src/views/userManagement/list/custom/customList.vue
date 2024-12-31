<template>
    <user-custom-list :custom-operate="customOperate"/>
</template>
<script setup lang="ts">
import UserCustomList from 'hooks/custom/custom.vue';
import customOperate from './custom';

</script>
