<template>
    <p class="currency-count-name-title">
        {{$filter.withClon($t('imageAddModel.header.imageName'))}}
        {{$filter.emptyFilter(imageName)}}
    </p>
    <p class="currency-count-title">
        <span>{{$filter.withClon($t('imageAddModel.header.choose')) + ' ' + chooseNum}}</span>
    </p>
</template>
<script lang="ts" setup>

/**
 * props 类
*/
interface PropsType {
    imageName: any;
    imageDeviceTypes: any;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props = withDefaults(defineProps<PropsType>(), {});

// 选择的镜像数
const chooseNum = computed((): number => {
    return props.imageDeviceTypes.reactiveArr.arr.filter((item: {isBind: boolean;}) => !item.isBind).length;
});
</script>