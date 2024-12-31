<template>
    <el-menu ref="menuRef">
        <slot></slot>
    </el-menu>
</template>
  
<script setup lang="ts">
const menuRef = ref();
// defineEmits API 来替代 emits
const emitRefVal = defineEmits(['get-menu-ref']);
onMounted(() => {
    emitRefVal('get-menu-ref', menuRef)
})

</script>
<style lang="scss">
@import './menu.scss';
</style>
