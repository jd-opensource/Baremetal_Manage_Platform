<template>
    <p
        class="empty-text"
        v-if="show"
    >
        <span class="empty-text-title">{{emptyTitle}}</span>
        <span class="empty-text-title" v-if="hasEmptyCount">，</span>
        <span v-if="hasEmptyCount">
            <span class="empty-text-title">
                {{$t('table.tip')}}
            </span>
            <span
                class="empty-text-operate"
                @click="emitValue('emit-click')"
            >
                {{emptyCount}}
            </span>
        </span>
    </p>
</template>

<script setup lang="ts">

/**
 * props 类
*/
interface PropsType {
    hasEmptyCount?: boolean;
    emptyTitle?: string;
    emptyCount?: string;
    show: boolean;
};

/**
 * withDefaults API 用来定义默认值，defineProps API 用来代替props
*/
withDefaults(defineProps<PropsType>(), {
    hasEmptyCount: true,
    emptyTitle: '',
    emptyCount: '',
    show: false
});

/**
 * defineEmits API 来替代 emits
*/
const emitValue= defineEmits(['emit-click']);

</script>
<style lang="scss" scoped>
@import './text.scss';
</style>
