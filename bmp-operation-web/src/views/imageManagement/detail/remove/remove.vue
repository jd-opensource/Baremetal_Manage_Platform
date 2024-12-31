<template>
    <div v-if="remove.removeOperateDiaLog.value">
        <remove-model
            :dia-log="remove.removeOperateDiaLog.value"
            :image-name="query?.imageName"
            :model-name="remove.name.value"
            :imageId="query?.imageId"
            :deviceType="remove.deviceType.value"
            :device-type-id="remove.deviceTypeId.value"
            @dia-log-close="remove.removeOperateCancel"
            @determine-click="remove.removeOperateSure"
        />
    </div>
</template>
<script lang="ts" setup>
// import * as all from '../all';
interface PropsType {
    remove: any;
    query: any;
}

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {});

</script>