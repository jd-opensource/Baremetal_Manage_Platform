<template>
    <!-- 携带下拉框搜索的输入框 -->
    <div>
        <el-input
            v-model.trim="iptValue"
            :placeholder="computedPlaceholder"
            @keydown.enter="inputChange"
        >
            <template #append>
                <!-- 搜索按钮 -->
                <img
                    class="search-icon"
                    src="@/assets/img/search-icon.png"
                    alt=""
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
                src="@/assets/img/clearInput.png"
                alt=""
                v-if="iptValue.length"
                @click="clearInput"
            />
        </div>
    </div>
</template>
  
<script setup lang="ts">
import {
    ref, // 实现响应式数据的方法
    Ref, // ref类
    watch, // 监听
    computed, // 计算属性
    ComputedRef // 计算属性类
} from 'vue';
import {ElInput, ElSelect, ElOption} from 'element-plus'; // input、select、option组件

/**
 * 筛选框类
*/
type SelectOptionType = {
    value?: number | any;
    label?: string;
};

/**
 * props类
*/
interface PropsType {
    hasClear?: boolean;
    placeHolder?: string[];
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
const emitValue = defineEmits(['ipt-value', 'change-select', 'select', 'clear-operate']);

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    hasClear: false,
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
const inputChange = (): void => {
    emitValue('ipt-value', iptValue.value, selectValue.value);
};

</script>
<style lang="scss">
@import './searchInput.scss';

</style>