<template>
    <el-dialog>
        <template #title>
            <slot name="title"></slot>
        </template>
        <slot></slot>
        <template #footer>
            <slot name="footer"></slot>
        </template>
    </el-dialog>
</template>
  
<script lang="ts" setup>

</script>

<style lang="scss">
@import './diaLog.scss';
</style>
