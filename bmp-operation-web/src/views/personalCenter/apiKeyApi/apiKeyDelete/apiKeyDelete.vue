<template>
    <div v-if="deleteKey.deleteKeyDiaLog.value">
        <delete-api-key
            :dia-log="deleteKey.deleteKeyDiaLog.value"
            :name="deleteKey.name.value"
            :apikey-id="deleteKey.apikeyId.value"
            @dia-log-close="deleteKey.deleteKeyCancel"
            @determine-click="deleteKey.deleteKeySure"
        />
    </div>
</template>
<script lang="ts" setup>

interface PropsType {
    apiKey: any;
    deleteKey: any;
    // tableMaxHeight: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
