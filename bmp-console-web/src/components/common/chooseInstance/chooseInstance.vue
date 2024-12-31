<template>
    <div class="choose-instance alarm-log-list">
        <div class="search-position">
            <el-input 
                v-model="searchInstance" 
                :placeholder="$t('instance.placeholder.instanceNameOrId')" 
                :disabled="props.allChooseTip"
            >
                <template #append>
                <!-- 搜索按钮 -->
                    <img
                        class="search-icon"
                        src="@/assets/img/search-icon.png"
                        alt=""
                    />
                </template>
            </el-input>
            <img
                class="input-search-clear"
                src="@/assets/img/clearInput.png"
                alt=""
                v-if="searchInstance.length"
                @click="clearInput"
            />
        </div>
        <div class="table-content table-search m0 table-change">
            <p class="mb10 f12">{{$t('list.content.allInstance', [tableData?.length])}}</p>           
            <!-- 全部实例 -->
            <el-table
                ref="tableRef"
                :data="tableData"   
                :row-key="getKey" 
                :header-cell-class-name="cellClass"
                :cell-class-name="cellStyle"
                :row-style="{height: '28px'}"
                @selection-change="handleSelectionChange"
                @row-click="rowClick"
                style="width: 100%"
                height="280"
                max-height="280"
            >  
                <!-- 选择多行 -->
                <el-table-column 
                    v-if="!props.allChooseTip"
                    type="selection" 
                    width="30"      
                    :reserve-selection="true"
                />
                <el-table-column 
                    v-else
                    width="30"
                >
                    <el-checkbox
                        v-model="props.allChooseTip"
                        disabled                    
                    >
                    </el-checkbox>
                </el-table-column>
                <!-- 实例名称 -->
                <el-table-column     
                    prop="instanceName" 
                    :label="$t('instance.detail.instanceName')"
                    width="115"
                    show-overflow-tooltip
                >
                    <template v-slot="scope">
                        <div
                            class="long-row"
                        >
                            <span>{{scope.row.instanceName || $filter.emptyFilter()}}</span>
                        </div>
                    </template>
                </el-table-column>
                <!-- 实例ID -->
                <el-table-column   
                    prop="instanceId" 
                    :label="$t('instance.detail.instanceId')"
                    width="250"
                    show-overflow-tooltip
                >
                    <template v-slot="scope">                   
                        <div
                            class="long-row"
                        >
                            <span>{{scope.row.instanceId || $filter.emptyFilter()}}</span>
                        </div>
                    </template>
                </el-table-column>
                <template #empty>
                    <div>{{$t('instance.create.noData')}}</div>
                </template>
            </el-table> 
        </div> 
        <!-- 穿梭图标 -->
        <img
            alt=""
            class="switch-img"
            src="@/assets/img/switch.png"
        /> 
        <!-- 清空 -->
        <div class="clear-position">
            <operate-unit
                v-if="!props.allChooseTip"
                :operateName="$t('list.clear')"
                :instanceInfo="[]"
                operateTip=''
                :disabled="props.allChooseTip"
                @handleClick="clearSelect"
            >
            </operate-unit>
        </div>
        <div class="table-content m0 table-change">
            <p class="mb10 f12">{{$t('list.content.selectInstance', [multipleSelection?.length])}}</p>  
           <!-- 已选实例  -->
            <el-table
                ref="multiTableRef"
                :data="multipleSelection" 
                style="width: 100%"
                :row-style="{height: '28px'}"
                height="280"
                max-height="280"
            >
                <!-- 实例名称 -->
                <el-table-column     
                    prop="instanceName" 
                    :label="$t('instance.detail.instanceName')"
                    width="112"
                    show-overflow-tooltip
                >
                    <template v-slot="scope">    
                            <div
                                class="long-row"
                            >
                                <span>{{scope.row.instanceName || $filter.emptyFilter()}}</span>
                            </div>
                    </template>
                </el-table-column>
                <!-- 实例ID -->
                <el-table-column     
                    prop="instanceId" 
                    :label="$t('instance.detail.instanceId')"
                    :width="props.allChooseTip ? 300 : 254"
                    show-overflow-tooltip
                >
                    <template v-slot="scope">
                        <div
                            class="long-row"
                        >
                            <span>{{scope.row.instanceId || $filter.emptyFilter()}}</span>
                        </div>
                    </template>
                </el-table-column>
                <!-- 已选实例回退 -->
                <el-table-column
                    width="36"
                    v-if="!props.allChooseTip"
                >
                    <template v-slot="scope">
                        <el-icon
                            class="clear-img"
                            :size="16" 
                            @click="deleteSelect(scope.row)"
                            ><CircleClose />
                        </el-icon>
                    </template>
                </el-table-column>
                <template #empty>
                    <div>{{$t('instance.create.noData')}}</div>
                </template>
            </el-table> 
        </div>       
    </div>
</template>

<script setup lang="ts">
import {
    ref,
    Ref,
    reactive,
    onMounted,
    computed,
    watch
} from 'vue';
import { instanceListAPI } from 'api/request.ts'; // 实例列表数据接口
import { ElTable } from 'element-plus'
import operateUnit from 'components/OperateUnit/operateUnit.vue';
import {deepCopy} from 'utils/index.ts';
import {
    CircleClose,
} from '@element-plus/icons-vue'
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    allChooseTip: boolean;
    instanceData: any;
    nextStepTip: boolean;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    allChooseTip: false,
    instanceData: [],
    nextStepTip: false,
});

/**
 * 实例item类型
 */
 type InstancesType = {
    instanceName: string,
    instanceId: string,
}

/**
 * 实例类型
 */
 interface instance {
    tableData: InstancesType[],
}

/**
 * 全部实例列表数据
 */
 const reactiveArr: instance = reactive<instance>({
    tableData: [],
})

/**
 * 回传对话框数据
 */
 const emitData = defineEmits(['getData', 'selectedData']);

/**
 * 点击下一步获取全部实例
 */
watch(() => props.nextStepTip, (): void => {
    if(props.nextStepTip) {
        emitData('getData')        
    }
});

/**
 * 筛选实例关键字
 */
const searchInstance: Ref<string> = ref<string>('');

/**
 * 点击x号，清空当前输入项
*/
const clearInput = (): void => {
    searchInstance.value = '';
};

/**
 * 全部转移或无实例时，全选复选框隐藏
 */
const cellClass = () => {
    if(props.allChooseTip || !props.instanceData.length) {
        return 'selectAllbtnDis'
    }
}

/**
 * 全部实例复选框样式修改
 */
 const cellStyle = ({row, column}: any) => {
    if(column.no === 0) {
        return 'select-column'
    }
}

/**
 * 筛选后实例数据
 */
const tableData = computed(() => {  
    if(searchInstance.value){
        return props.instanceData.filter(function(dataNews:any){
            return Object.keys(dataNews).some(function(key){
              return String(dataNews[key]).toLowerCase().indexOf(searchInstance.value) > -1
            })
        })
    }
    return props.instanceData
});

/**
 * 得到每行数据的key
 * @param row 每行的数据
 */
const getKey = (row:any) => {
    return row.instanceId
}

/**
* 列表已选实例
*/
const multipleSelection = ref<InstancesType[]>([])

/**
 * 回传已选实例
 */
 watch(() => multipleSelection.value, (): void => {
    emitData('selectedData', multipleSelection.value)   
});

/**
 * 表格ref
*/
const tableRef = ref<InstanceType<typeof ElTable>>()

/**
* 全部实例列表全选 
*/
const handleCheckAllChange = () => {
    tableRef.value!.toggleAllSelection()
}

/**
 * 全部实例列表数据勾选
 * @param row 每行的数据
 */
const selectChange = (row: any) => {
    tableRef.value!.toggleRowSelection(row, true)
}

/**
 * 表格点击某一行触发选中事件
 * @param {Object} row 当前选中的这一项
 */
 const rowClick = (row: any) => {
    row.rowChange = !row.rowChange;
    tableRef.value!.toggleRowSelection(row, row.rowChange);
};

/**
* 已选实例列表选项清空 
*/
const clearSelect = () => {
    tableRef.value!.clearSelection()
}

/**
 * 点击上一步清空筛选和实例勾选
 */
const handlePreStep = () => {
    clearSelect();
    searchInstance.value = ''

}

/**
 * 该行已选实例撤销
 * @param row 每行的数据
 */
const deleteSelect = (row:any) => {
    tableRef.value!.toggleRowSelection(row, false)
}

/**
* 列表选中项 
* @param val 已选中的实例数据
*/
const handleSelectionChange = (val: InstancesType[]) => {
    // 选中项
    multipleSelection.value = val 

}

/**
 * 暴露方法
 */
 defineExpose({ handlePreStep, handleCheckAllChange, selectChange })
</script>
<style lang="scss">
@import 'assets/css/table';
@import './chooseInstance.scss';
</style>