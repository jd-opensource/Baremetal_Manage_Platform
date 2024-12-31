<template>
    <div v-if="imageDelete.deleteImageDiaLog.value">
        <delete-image
            :dia-log="imageDelete.deleteImageDiaLog.value"
            :image-info="query?.imageName"
            :has-relation-images="true"
            :image-id="query?.imageId"
            @dia-log-close="imageDelete.deleteImageCancel"
            @determine-click="imageDelete.deleteImageSure"
            @error-click="imageDelete.imageDeleteError"
        />
    </div>
</template>
<script lang="ts" setup>
// import * as all from '../all';
interface PropsType {
    imageDelete: any;
    query: any;
}

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {});

</script>