<template>
    <el-table-column v-if='!hasUserTemplate'/>
    <el-table-column v-else>
        <template #default="scope">
            <slot :scope="scope"></slot>
        </template>
    </el-table-column>
</template>

<script setup lang="ts">

/**
 * props 类
*/
interface PropType {
    hasUserTemplate?: boolean;
};

/**
 * withDefaults API 用来书写默认值 defineProps API 来替代 props
*/
withDefaults(defineProps<PropType>(), {
    hasUserTemplate: false // 是否使用template
});
</script>

<style lang="scss">
@import './tableColumn.scss';
</style>
