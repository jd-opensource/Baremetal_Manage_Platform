<template>
    <div v-if="delImage.deleteImageDiaLog.value">
        <delete-image
            :dia-log="delImage.deleteImageDiaLog.value"
            :image-info="delImage.imageInfo.value"
            :image-id="delImage.imageId.value"
            :device-type-id="baseInfo.reactiveArr.detail.deviceTypeId"
            @dia-log-close="delImage.deleteImageCancel"
            @determine-click="delImage.deleteImageSure"
        />
    </div>
</template>
<script lang="ts" setup>
interface PropsType {
    delImage: any;
    baseInfo: any
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>