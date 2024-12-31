<template>
    <!-- 删除 -->
    <ui-dropdown-item>
        <!-- 删除提示 -->
        <ui-tooltip
            placement="left"
            v-if="reactiveArr.deviceTypeNum > 0"
            :content="$t('imageDetail.tooltip.title')"
        >
            <span class="drop-down-disabled">
                {{$t('imageDetail.header.operate.delete')}}
            </span>
        </ui-tooltip>
        <!-- 预置镜像不支持删除 -->
        <ui-tooltip
            placement="bottom"
            v-else-if="Object.is(reactiveArr.detail.source, 'common')"
            :content="$t('imageList.tooltip.image')"
        >
            <span class="drop-down-disabled">
                {{$t('imageList.content.label.operate.delete')}}
            </span>
        </ui-tooltip>
        <span
            class="drop-down-operate-content image-delete"
            v-else
            @click="emitValue('delete')"
        >
            {{$t('imageDetail.header.operate.delete')}}
        </span>
    </ui-dropdown-item>
</template>

<script lang="ts" setup>
// defineEmits API 来替代 emits
const emitValue = defineEmits(['delete']);
interface PropsType {
    reactiveArr: any;
    // imgSrc?: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
