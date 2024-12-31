<template>
	<el-dropdown @command="selectOps" size="small" class="batch-operate-dropdown">
		<el-button>
			{{ $t('list.operate.more') }}
			<el-icon><CaretBottom /></el-icon>
		</el-button>
		<template #dropdown>
			<el-dropdown-menu >
				<el-dropdown-item
					v-for="(item, index) in operations"
					:key="index"
					:command="item"
					:disabled="item.disabled"
					class="select-form-item"
				>
					<el-tooltip
						placement="right"
						:disabled="!item.disabled"
						:content="item.tip"
					>
						<label>
							<div>{{item.operateName}}</div>
						</label>
					</el-tooltip>
				</el-dropdown-item>
			</el-dropdown-menu>
		</template>
	</el-dropdown>
</template>

<script setup lang="ts">

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    operations: Array<any>;
    instanceInfo: any;
};

/**
 * withDefaults API 可以指定默认值，defineProps API 来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    instanceInfo: {},
});

const emitValue = defineEmits(['clickSelectItem']);

const selectOps = (item: any) => {
    if (!item.disabled) {
        emitValue('clickSelectItem', item.value,props.instanceInfo)
    }
}


</script>

<style lang="scss" scoped>

.batch-operate-dropdown {
	.el-button {
		width: 50px;
		height: 24px;
		font-size: 12px;
		--el-button-border-color: #333 !important;
		--el-button-hover-text-color: #333 !important;
		&:focus,
		&:hover {
			--el-button-hover-bg-color: #fff !important;
			--el-button--hover-border-color: #333 !important;
			--el-button-hover-text-color: #333 !important;
			border-color: #333 !important;
		}
		&:active {
			--el-button-active-text-color: #0f81d4 !important;
			--el-button-border-color: #0f81d4 !important;
			border-color: #0f81d4 !important;
		}
	}
}

.form-tip {
	display: none;
	position: absolute;
	background: #fff;
	border: 1px solid #d9d9d9;
	top: 0;
	line-height: 25PX;
	height: 25px;
	padding: 0 5px;
	white-space:nowrap;
	left: 0;
	margin: 0;
	transform: translate(-105%,20%);
}

.select-form-item {
	position: relative;
	min-width: 60px;

}

.operate-select {
  	width: 102px;
}

.select-form-item.is-disabled {
	pointer-events: initial;
	&:hover {
		.form-tip {
			display: block;
		}
	}
}

:deep( .el-dropdown-menu__item ) {	 
	font-size: 12px;
	line-height: 18px;
	color: #333333;
	&:hover {
		color:#108EE9 !important
	}
	&.is-disabled{
		color:#999999;      
		&:hover {
			color:#999999 !important
		}
	}
		
}

.fc333 {
  color: #333333;
}

</style>