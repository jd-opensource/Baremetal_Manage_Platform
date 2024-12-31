<template>
    <div v-if="createKey.createKeyDiaLog.value">
        <create-api-key
            :dia-log="createKey.createKeyDiaLog.value"
            :key-name-data="apiKey.reactiveArr.keyNameData"
            @dia-log-close="createKey.createKeyCancel"
            @determine-click="createKey.createKeySure"
        />
    </div> 
</template>
<script lang="ts" setup>

interface PropsType {
    apiKey: any;
    createKey: any;
    // tableMaxHeight: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
