<template>
    <div v-if="addImageOperate.addImageDialog.value">
        <image-add
            :dia-log="addImageOperate.addImageDialog.value"
            :model-name="basicInformation.reactiveArr.detail.name"
            :architecture="architectureName"
            :device-type-id="basicInformation.reactiveArr.detail.deviceTypeId"
            @dia-log-close="addImageOperate.addImageCancel"
            @determine-click="addImageOperate.addImageSure"
        />
    </div>
</template>
<script lang="ts" setup>
// import * as all from '../all';
interface PropsType {
    addImageOperate: any;
    basicInformation: any;
    architectureName: any;
}

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {});
</script>