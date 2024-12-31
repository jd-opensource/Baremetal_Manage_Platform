<template>
    <div v-if="descEdit.editDescDiaLog.value">
        <image-desc
            :dia-log="descEdit.editDescDiaLog.value"
            :image-id="query?.imageId"
            :description="description"
            @dia-log-close="descEdit.editDescCacnel"
            @determine-click="descEdit.editDescSure"
        />
    </div>
</template>
<script setup lang="ts">
// import * as all from '../all';
interface PropsType {
    descEdit: any;
    description: string;
    query: any;
}

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {
    description: ''
});

</script>