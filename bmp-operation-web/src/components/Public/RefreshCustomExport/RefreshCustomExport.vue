<template>

<!-- :class="[
            'operate-management-content-operate2',
            otherClass,
        ]" -->
    <aside
        ref="operate"
        :class="[
            defaultClass,
            otherClass
        ]"
        :style="!status ? {marginBottom: '30px'} : {marginBottom: '16px'}"
    >
        <!-- 左侧操作 -->
        <btn-operate
            v-if="hasBtnOperate"
            :type="type"
            @btn-operate="emitValue('btn-operate')"
            @btn-ref="btnRefClick"
        />
        <!-- 右侧操作 -->
        <div
            :class="`${defaultClass}-right`"
        >
            <!-- 搜索筛选框 -->
            <search
                v-if="searchOperate"
                :place-holder="placeHolder"
                :search-operate="searchOperate"
            />
            <!-- 刷新 设置 导出 -->
            <p
                v-if="refreshStatus"
                :class="[
                    disabled ? 'disabled-list-data-operate0' : 'list-data-operate0'
                ]"
                @click="emitValue('refresh')"
            >
                <span class="default-img-style"></span>
                <!-- 刷新 -->
                <span class="default-text-size">
                    {{$t('operate.refresh')}}
                </span>
            </p>
            <p
                v-if="customStatus"
                :class="[
                    disabled ? 'disabled-list-data-operate1' : 'list-data-operate1'
                ]"
                :style="!hasExport ? {marginRight: 0} : {}"       
                @click="emitValue('custom')"
            >
                <span class="default-img-style"></span>
                <span class="default-text-size">
                    {{$t('operate.setUp')}}
                </span>
            </p>
            <p
                style="marginRight: 0"
                v-if="hasExport"
                :class="[
                    disabled ? 'disabled-list-data-operate2' : 'list-data-operate2'
                ]"
                @click="emitValue('export')"
            >
                <span class="default-img-style"></span>
                <span class="default-text-size">
                    {{$t('operate.export')}}
                </span>
            </p>
        </div>
    </aside>
</template>

<script lang="ts" setup>
import {HeightType2, HeightType, HeightType3} from '@utils/publicType';

interface PropsType {
    status?: boolean;
    searchOperate?: any;
    hasBtnOperate?: boolean;
    placeHolder?: string[];
    hasExport?: boolean;
    refreshStatus?: boolean;
    customStatus?: boolean;
    type?: string;
    defaultClass?: string;
    otherClass?: string;
    disabled?: boolean;
}
/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
withDefaults(defineProps<PropsType>(), {
    hasBtnOperate: true,
    hasExport: true,
    refreshStatus: true,
    customStatus: true,
    disabled: false,
    defaultClass: 'operate-management-content-operate2',
    otherClass: ''
});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['operate-ref', 'btn-operate', 'refresh', 'custom', 'export', 'btn-ref']);

const operate: Ref<HeightType2 | HeightType | HeightType3> = ref<HeightType2 | HeightType | HeightType3>({offsetHeight: 0, offsetTop: 0, scrollHeight: 0});

onMounted(() => {
    emitValue('operate-ref', operate);
});

const btnRefClick = (val: HeightType) => {
    emitValue('btn-ref', val)
}

</script>

<style lang="scss">
@import './refreshCustomExport';
</style>
