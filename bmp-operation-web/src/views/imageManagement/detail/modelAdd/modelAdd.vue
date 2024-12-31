<template>
    <div v-if="modelAdd.addModelDiaLog.value">
        <image-add-model
            :dia-log="modelAdd.addModelDiaLog.value"
            :image-id="query?.imageId"
            :architecture="query?.architecture"
            :image-name="query?.imageName"
            @dia-log-close="modelAdd.addModelCancel"
            @determine-click="modelAdd.addModelSure"
        />
    </div>
</template>
<script lang="ts" setup>
// import * as all from '../all';
interface PropsType {
    modelAdd: any;
    query: any;
}

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {});

</script>