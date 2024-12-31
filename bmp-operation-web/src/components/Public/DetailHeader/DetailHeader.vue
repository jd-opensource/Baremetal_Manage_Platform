<template>
    <header
        ref="headerTitle"
        class="operate-management-detail-header"
    >
        <!-- 标题 & 操作 -->
        <div class="operate-management-detail-header-count">
            <!-- 标题信息 -->
            <div class="operate-management-detail-header-left">
                <!-- 返回箭头 -->
                <div
                    class="detail-arrow-left"
                    @click="emitValue('header')"
                >
                    <div class="header-arrow-left"></div>
                </div>
                <!-- 上一页标识 -->
                <span class="header-list-title">
                    {{$t(`${title}.header.title2`)}}
                </span>
                <p v-if="name !== 'no'">
                    <!-- 分隔符 -->
                    <span class="header-left-icon">/   </span>
                    <!-- 当前详情标题 -->
                    <!-- <span class="header-title">
                        {{$filter.emptyFilter(name)}}
                    </span> -->
                    <span class="header-title" v-html="name"></span>
                </p>
            </div>
            <!-- 操作 -->
            <div
                v-if="optFlag"
                :class="[selectStatus ? 'hover-select' : 'header-select']"
            >
                <ui-dropdown @visible-change="selectHover">
                    <template #dropdownTitle>
                        <!-- 操作 -->
                        <div class="drop-down-operate">
                            <p class="drop-down-operate-count">
                                <span :style="selectStatus ? {color: '#108ee9'} : {color: '#333'}">
                                    {{$t('operate.name')}}
                                </span>
                                <!-- 下拉箭头 -->
                                <p :class="[selectStatus ? 'arrow-img' : 'arrow-bottom-img']"></p>
                            </p>
                        </div>
                    </template>
                    <slot></slot>
                </ui-dropdown>
            </div>
        </div>
    </header>
</template>

<script setup lang="ts">
import {HeightType} from '@utils/publicType';

interface PropsType {
    name: string;
    title?: string;
    optFlag?: boolean;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    name: '',
    title: 'idcDetail',
    optFlag: true
});
// defineEmits API 来替代 emits
const emitValue = defineEmits(['header-ref', 'header']);

const selectStatus: Ref<boolean> = ref<boolean>(false);
const selectHover = (val: boolean) => {
    selectStatus.value = val;
};
const headerTitle: Ref<HeightType> = ref<HeightType>({offsetHeight: 0, offsetTop: 0});
onMounted(() => {
    emitValue('header-ref', headerTitle);
});
</script>

<style lang="scss">
@import './detailHeader.scss';
</style>