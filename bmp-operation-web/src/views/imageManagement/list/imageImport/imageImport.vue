<template>
    <div v-if="importImageOperate?.importImageDiaLog?.value">
        <import-image
            :dia-log="importImageOperate?.importImageDiaLog?.value"
            @dia-log-close="importImageOperate?.importImageCancel"
            @determine-click="importImageOperate?.importImageSure"
        />
    </div>
</template>
<script setup lang="ts">
interface PropsType {
    importImageOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>