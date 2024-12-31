<template>
    <p class="search-filter-tip">
        <span v-if="customPagination.total.value > 0">
            {{$t('searchTip.title1', [customPagination.total.value])}} 
        </span>
        <span v-else>
            {{$t('searchTip.title2', [customPagination.total.value])}}
        </span>
        <span
            class="go-back-list"
            @click="resetOperate.reset"
        >
            {{$t('searchTip.title3')}}
        </span>
    </p>
</template>
<script lang="ts" setup>

/**
 * props 类
*/
interface PropsType {
    resetOperate: any;
    customPagination: any;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {});
</script>