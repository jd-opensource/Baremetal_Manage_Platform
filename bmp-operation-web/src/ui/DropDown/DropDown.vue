<template>
    <el-dropdown>
        <slot name="dropdownTitle"></slot>
        <template #dropdown>
            <el-dropdown-menu>
                <slot></slot>
            </el-dropdown-menu>
        </template>
    </el-dropdown>
</template>
  
<script setup lang="ts">

</script>
<style lang="scss">
@import './dropDown.scss';
</style>
