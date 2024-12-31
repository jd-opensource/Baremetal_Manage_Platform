<template>
    <section
        ref="deviceHeader"
        :class="setDiffClass('en-model-info', 'zh-model-info')">
        <!-- 基本信息、关联机型 -->
        <ui-tabs
            v-model="tabs.activeName.value"
            :class="setDiffClass('en-info', 'default-info')"
            :active-name="tabs.activeName.value"
            :data="tabs.setData(routerOperate.status, routerOperate.monitorStatus)"
            @on-change="tabs.tabsChange"
        />
    </section>
</template>
<script lang="ts" setup>
import {RouterOperate} from 'utils/publicClass.ts';
import {setDiffClass} from 'utils/index.ts';
import {HeightType} from '@utils/publicType';

const routerOperate = new RouterOperate().router.currentRoute.value.query;

interface PropsType {
    tabs: {
        activeName: {
            value: string;
        };
        setData(arg0: string, arg1: string): void;
        tabsChange(arg0: string): void;
    };
}

// emit 事件
const emitValue = defineEmits(['header-ref']);
const deviceHeader: Ref<HeightType> = ref<HeightType>({
    offsetHeight: 0,
    offsetTop: 0
});

nextTick(() => {
    emitValue('header-ref', deviceHeader)
})
   
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

</script>