<template>
    <p class="currency-count-name-title">
        {{$filter.withClon($t('addImage.header.label.modelName'))}}
        {{$filter.emptyFilter(modelName)}}
    </p>
    <p class="currency-count-title">
        <span>
            {{$filter.withClon($t('addImage.header.label.choose')) + ' ' + chooseNum }}
        </span>
    </p>
</template>
<script lang="ts" setup>

/**
 * props 类
*/
interface PropsType {
    modelName: any;
    imageList: any;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props = withDefaults(defineProps<PropsType>(), {});

// 选择的镜像数
const chooseNum = computed((): number => {
    return props.imageList.reactiveArr.arr.filter((item: {isBind: boolean;}) => !item.isBind).length;
});
</script>