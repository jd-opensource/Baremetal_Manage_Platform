<template>
   <p
        v-if="!hasVisibility ? status : true"
        :style="status ? {visibility: 'inherit'} : {visibility: 'hidden'}"
        class="search-filter-tip"
    >
        <span v-if="total > 0">
            {{$t('searchTip.title1', [total])}} 
        </span>
        <span v-else>
            {{$t('searchTip.title2', [total])}}
        </span>
        <span
            class="go-back-list"
            @click="backClick"
        >
            {{$t('searchTip.title3')}}
        </span>
    </p>
</template>

<script lang="ts" setup>

interface PropsType {
    status: boolean;
    total: number;
    hasVisibility?: boolean;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    hasVisibility: false
});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['back-click']);
const backClick = () => {
    emitValue('back-click');
};


</script>

<style lang="scss" scoped>
.search-filter-tip {
    font-size: 12px;
    color: #333;
    font-weight: 400;
    margin-bottom: 10px;

    .go-back-list {
        font-size: 12px;
        color: #0F8FE9;
        font-weight: 400;
        cursor: pointer;
    }
}
</style>