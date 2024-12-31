<template>
    <!-- 编辑 -->
    <ui-dropdown-item>
        <!-- 编辑机型 -->
        <ui-tooltip
            placement="left"
            :disabled="!(reactiveArr.detail.deviceCount > 0)"
            :content="$t('modelList.edit')"
        >
            <span
                :class="[
                    'currency-style-right',
                    (reactiveArr.detail.deviceCount > 0) ? 'drop-down-disabled' : 'drop-down-operate-content'
                ]"
                @click="emitValue('edit', reactiveArr.detail)"
            >
                {{$t('modelDetail.header.operate.edit')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
    <!-- 添加模板 -->
    <ui-dropdown-item>
        <!-- 添加相同机型模板 -->
        <span
            class="drop-down-operate-content"
            @click="emitValue('template-add', reactiveArr.detail)"
        >
            {{$t('modelDetail.header.operate.addModel')}}
        </span>
    </ui-dropdown-item>
    <!-- 删除 -->
    <ui-dropdown-item>
        <!-- 删除提示 -->
        <ui-tooltip
            placement="left"
            :disabled="!(reactiveArr && (reactiveArr.type1 > 0 || reactiveArr.type2 > 0))"
            :content="$t('modelDetail.header.tip')"
        >
            <span
                :class="[
                    (reactiveArr && (reactiveArr.type1 > 0 || reactiveArr.type2 > 0)) ? 'drop-down-disabled' : 'drop-down-operate-content'
                ]"
                @click="deleteModel"
            >
                {{$t('modelDetail.header.operate.delete')}}
            </span>
        </ui-tooltip>
    </ui-dropdown-item>
</template>

<script lang="ts" setup>
// defineEmits API 来替代 emits
const emitValue = defineEmits(['edit', 'delete', 'template-add']);
interface PropsType {
    reactiveArr: any;
    // imgSrc?: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {});

const deleteModel = () => {
    if (!(props.reactiveArr && (props.reactiveArr.type1 > 0 || props.reactiveArr.type2 > 0))) {
        emitValue('delete')
    };
}

</script>
