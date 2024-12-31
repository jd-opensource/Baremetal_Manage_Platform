<template>
    <!-- 携带下拉框搜索的输入框 -->
    <div v-if="searchInput">
        <el-input
            v-model.trim="iptValue"
            :placeholder="computedPlaceholder"
            @input="inputVal"
            @blur="iptBlur"
            @focus="iptFocus"
        >
            <template #append>
                <!-- 搜索按钮 -->
                <img
                    class="search-icon"
                    alt=""
                    :src="($defInfo.imagePath('searchIcon') as unknown as string)"
                    @click="inputChange"
                />
            </template>
        </el-input>
        <div class="input-select">
            <!-- 分割线 -->
            <div class="input-select-division-line"></div>
            <!-- 筛选框 -->
            <el-select
                class="input-select-count"
                v-model="selectValue"
                popper-class="custom-select"
            >
                <el-option
                    style="fontSize: 12px;"
                    v-for="item in selectOption"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                />
            </el-select>
            <!-- close清空icon -->
            <img
                class="input-search-clear"
                alt=""
                v-if="iptValue.length"
                :src="($defInfo.imagePath('clear') as unknown as string)"
                @click="clearInput"
            />
        </div>
    </div>
    <el-input v-else-if="hasSuffix">
        <template #suffix>
            <slot name="suffix"></slot>
        </template>
    </el-input>
    <el-input v-else/>
</template>
  
<script setup lang="ts">

/**
 * 筛选框类
*/
type SelectOptionType = {
    value?: any;
    label?: string;
};

/**
 * props类
*/
interface PropsType {
    hasSuffix?: boolean;
    hasClear?: boolean;
    placeHolder?: string[];
    searchInput?: boolean;
    selectOption?: SelectOptionType[];
};

/**
 * 输入框value值
*/
const iptValue: Ref<string> = ref<string>('');

/**
 * 筛选框value
*/
const selectValue: Ref<number> = ref<number>(0);

/**
 * defineEmits API 用来替代emits
*/
const emitValue = defineEmits(['ipt-val', 'enter-value-ipt', 'change-select', 'select', 'clear-operate']);

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    hasSuffix: false,
    hasClear: false,
    searchInput: false, // 是否是搜索&input框
    placeHolder: () => [], // placeholder提示
    selectOption: () => [] // select筛选框的值
});

/**
 * 监听input输入框、监听select筛选搜索框
 * @param {Array} newValue input、select更新的值
*/
watch(() => [iptValue.value, selectValue.value], (newValue: any) => {
    iptValue.value = newValue[0];
    emitValue('change-select', selectValue.value);
});

watch(() => selectValue.value, (newValue: any) => {
    selectValue.value = newValue;
    iptValue.value = ''
    emitValue('select');
});

watch(() => props.hasClear, (newValue) => {
    if (newValue) {
        iptValue.value = ''
    }
})
// @keydown.enter="inputChange"


const iptFocus = () => {
    document.onkeyup = (event: {keyCode: number;}) => {
        event.keyCode === 13 && inputChange();
    };
}

/**
 * 点击x号，清空当前输入项
*/
const clearInput = (): void => {
    iptValue.value = '';
    emitValue('clear-operate', iptValue.value);
};
/**
 * 计算placeholder提示
 * @return {string} props.placeHolder[selectValue.value] 输入框placeholder提示
*/
const computedPlaceholder: ComputedRef<string> = computed((): string => {
    return props.placeHolder![selectValue.value]; 
});

/**
 * 点击搜索按钮
*/
const inputChange = () => {
    emitValue('enter-value-ipt', iptValue.value, selectValue.value);
};

const inputVal = () => {
    emitValue('ipt-val', iptValue.value, selectValue.value);
}

const iptBlur = () => {
    document.onkeyup = () => {return;}
}
</script>
<style lang="scss">
@import './input.scss';

</style>
