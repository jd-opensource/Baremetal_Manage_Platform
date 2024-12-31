<template>
    <div class="info-header" >
        <span class="info-line"></span>
        <!-- 设备信息 -->
        <span
            style="fontWeight: 600"
            class="info-title"
        >
            {{title}}
        </span>
    </div>
</template>
<script lang="ts" setup>

interface PropsType {
    title: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>