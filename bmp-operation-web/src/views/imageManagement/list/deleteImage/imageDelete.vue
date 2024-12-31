<template>
    <div v-if="deleteOperate?.deleteImageDiaLog?.value">
        <delete-image
            :dia-log="deleteOperate?.deleteImageDiaLog?.value"
            :image-info="deleteOperate?.imageInfo?.value"
            :has-relation-images="true"
            :image-id="deleteOperate?.imageId?.value"
            @dia-log-close="deleteOperate?.deleteImageCancel"
            @determine-click="deleteOperate?.deleteImageSure"
            @error-click="deleteOperate.imageDeleteError"
        />
    </div>
</template>
<script lang="ts" setup>
interface PropsType {
    deleteOperate: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>