<template>
    <!-- 搜索筛选框 -->
    <p
        class="operate-ipt"
    >
        <ui-input
            :class="[setClass]"
            :search-input="true"
            :select-option="searchOperate.reactiveArr.selectOption"
            :place-holder="placeHolder"
            :has-clear="searchOperate.hasClear.value"
            @ipt-val="searchOperate.iptValue"
            @enter-value-ipt="searchOperate.enterValueIpt"
            @change-select="searchOperate.changeSelect"
            @select="searchOperate.selectChange"
            @clear-operate="searchOperate.clearClick"
        />
    </p>
</template>

<script lang="ts" setup>
import {locationItem} from 'utils/publicClass.ts';

interface PropsType {
    type?: string;
    placeHolder?: string[];
    searchOperate: any;
}
/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {
    type: 'userList'
});

const setClass = computed(() => {
    const defaultClass1: string = 'operate-management-input';
    const defaultClass2: string = 'small-input';
    // const trendsClass: string = locationItem.getLocationItem === 'zh_CN' ? 'small-input' : 'en-select';
    // return `${defaultClass} ${trendsClass}`
    const trendsClass = locationItem.getLocationItem === 'zh_CN' ? defaultClass1 : defaultClass2;
    return trendsClass;
    // 'operate-management-input',
                // locationItem.getLocationItem !== 'zh_CN' ? 'small-input' : 'en-select'
})
</script>

<style lang="scss" scoped>
.operate-ipt {
    display: flex;
    align-items: center;
    margin-right: 30px;
}
</style>