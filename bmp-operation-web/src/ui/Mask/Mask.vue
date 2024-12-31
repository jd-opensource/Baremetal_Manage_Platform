<template>
    <!-- 自定义蒙层 -->
    <div
        class="custom-mask"
        :style="style"
    >
    </div>
</template>
  
<script setup lang="ts">

/**
 * style类
*/
type StyleType = {
    background: string;
};

/**
 * props 类
*/
interface PropsType {
    style?: StyleType;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {
    style: () => {
        return {
            background: 'rgba(212, 217, 221, .6)'
        }
    }
});

</script>
<style lang="scss" scoped>
@import './mask.scss';
</style>
