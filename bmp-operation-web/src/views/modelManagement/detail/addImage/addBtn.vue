<template>
    <!-- 添加镜像 -->
    <ui-tooltip
        placement="right"
        :disabled="!tableDetail.hasDisabled.value"
    >
        <template #content>
            <div>
                {{$t('modelDetail.tabs.relationImage.header.tooltip')}}
            </div>
        </template>
        <ui-button
            class="add-image-btn"
            type="primary"
            :disabled="tableDetail.hasDisabled.value"
            @click="addImage.addImageClick"
        >
            +  {{$t('modelDetail.tabs.relationImage.header.addImage')}}
        </ui-button>
    </ui-tooltip>
</template>
<script lang="ts" setup>
interface PropsType {
    tableDetail: any;
    addImage: any
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>