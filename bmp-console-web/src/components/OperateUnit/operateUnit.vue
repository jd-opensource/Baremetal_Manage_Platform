<template>
	<div class="one-line">
		<!-- 操作置灰 -->
		<el-tooltip
			v-if="disabled"
			placement="bottom"
			:content="operateTip"
		>
			<span
				:class= "disabled ? 'disabled-currency-style' : 'change-show-style' "
			>
				{{ operateName }}
			</span>
		</el-tooltip>
		<!-- 可操作 -->
		<a 		
			v-else
			:class= "disabled ? 'disabled-currency-style' : 'change-show-style' "
			@click.stop="click"
		>
			{{ operateName }}
		</a>
	</div>
</template>

<script setup lang="ts">
import {Ref, ref, computed} from 'vue';

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    operateName: string;
    operateTip: string;
    disabled: boolean;
    instanceInfo: object;
};

/**
 * withDefaults API 可以指定默认值，defineProps API 来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    operateName: '',
    operateTip: '',
    disabled: false,
});

/**
 * defineEmits API 来替代 emits，子传父，事件回传
*/
const emitValue = defineEmits(['handle-click']);

/**
 * 按键是否置灰
 */
const disabled = computed(()=>{
	return props.disabled
})
/**
 * 当前列的值
 */   
const instanceInfo: Ref<any> = ref(props.instanceInfo);


const click: () => void = () => {
    emitValue('handle-click', disabled.value, instanceInfo.value)
}

</script>

<style lang="scss">
.change-show-style {
    cursor: pointer;     
}
.one-line {
	display: inline-block;
	margin-left: 5px;
	margin-right: 5px;
}
.disabled-currency-style {
	cursor: no-drop;
	color: #aaa;
}
</style>