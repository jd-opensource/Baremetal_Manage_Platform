<template>
    <div class="middle-dialog big-dialog" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="list dialog-icon"></div>
                    <span class="my-title">{{$t('alarm.operate.history.name')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <el-config-provider :locale="locale">
                <!-- 日期搜索 -->
                <div class="operation-search">
                    <el-date-picker
                        v-model="alarmHistoryArr.operationTime"
                        type="datetimerange"
                        unlink-panels
                        range-separator="-"
                        :disabled-date="disabledDate"
                        :start-placeholder="$t('instance.placeholder.startTime')"
                        :end-placeholder="$t('instance.placeholder.endTime')"
                    />
                    <div class="alarm-history-search-button">
                        <el-button 
                            type="primary" 
                            :loading="isLoading" 
                            @click="searchOperationData"
                        >
                            {{$t('list.search')}}
                        </el-button>
                    </div>   
                </div>
                <div class="batch-table-content m0 more-height"
                    :class="[
                            searchResultMark ? 'mt16' : 'mt30'
                        ]"
                    >
                    <!-- 搜索结果显示 -->
                    <div v-if="searchResultMark" class="search-result">
                        <span>{{totalCount ? $t('instance.list.searchHasResult', [totalCount]) : $t('instance.list.searchNoResult', [totalCount])}}</span>
                        <a 
                            class="mouse-point"
                            @click="backToList()">{{$t('instance.list.backList')}}
                        </a>
                    </div>
                    <el-table 
                        border
                        ref="tableRef"
                        :data="alarmHistoryArr.tableData" 
                        style="width: 100%"
                        max-height="255"
                    >
                        <!-- 报警时间 -->
                        <el-table-column 
                            prop="alarmTime" 
                            :label="$t('alarm.alarmRule.alarmTime')" 
                            min-width=100
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['alarmTime'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['alarmTime'].showTooltip"
                                    :content= scope.row.alertTime
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['alarmTime'], scope.row.alertTime, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.alertTime || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 报警资源 -->
                        <el-table-column 
                            prop="alarmResource" 
                            :label="$t('alarm.alarmRule.alarmResourceId')" 
                            min-width=100
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['alarmResource'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['alarmResource'].showTooltip"
                                    :content= scope.row.resourceId
                                    :show-after="500"
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['alarmResource'], scope.row.resourceId, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.resourceId || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 触发条件 -->
                        <el-table-column 
                            prop="triggeringCondition" 
                            :label="$t('alarm.alarmRule.triggeringCondition')" 
                            min-width=200
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['triggeringCondition'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['triggeringCondition'].showTooltip"
                                    :content= scope.row.triggerDescription
                                    :show-after="500"
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['triggeringCondition'], scope.row.triggerDescription, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.triggerDescription || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 报警值 -->
                        <el-table-column 
                            prop="alarmValue" 
                            :label="$t('alarm.alarmRule.alarmValue')" 
                            min-width=60
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['alarmValue'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['alarmValue'].showTooltip"
                                    :content= scope.row.alertValue
                                    :show-after="500"
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['alarmValue'], scope.row.alertValue, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.alertValue || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 报警级别 -->
                        <el-table-column 
                            prop="alarmLevel" 
                            :label="$t('alarm.alarmRule.alarmLevel')" 
                            min-width=70
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['alarmLevel'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['alarmLevel'].showTooltip"
                                    :content= scope.row.alertLevelDescription
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['alarmLevel'], scope.row.alertLevelDescription, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.alertLevelDescription || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 持续时间 -->
                        <el-table-column 
                            prop="continueTime" 
                            :label="$t('alarm.alarmRule.continueTime')" 
                            min-width=70
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['continueTime'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['continueTime'].showTooltip"
                                    :content= scope.row.alertPeriod
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['continueTime'], scope.row.alertPeriod, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.alertPeriod || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <template #empty>
                            <p 
                                v-if="!alarmHistoryArr.tableData.length && !isLoading"
                                class="no-data-jump"
                            >
                            {{$t('instance.create.noData')}}
                            </p>
                        </template>
                    </el-table>                
                </div>
            </el-config-provider>
            <div class="footer-count">
                <my-pagination
                    v-if="alarmHistoryArr.tableData.length > 0 && !isLoading"
                    :hasUseDefault="false"
                    :page-sizes="[20, 50, 100]"
                    :total="totalCount" 
                    :page-size="pageSize" 
                    :page-number="pageNumber" 
                    @change-page="changePage"
                    @page-sizes-change="pageSizesChange" 
                >
                </my-pagination>
            </div> 
        </el-dialog>
    </div>
</template>
<script setup lang="ts">
import { Ref, ref,watch, reactive, onUnmounted, onMounted, nextTick } from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {
    hasShowTooltip, // 是否显示提示
    languageSwitch, // 语言切换
    convertStringsToTimestampObject,
    getDateMinutes
} from 'utils/index.ts';
import myPagination from 'components/Pagination/MyPagination.vue';
import  {ElMessage, ElTable } from 'element-plus';
import { useI18n } from 'vue-i18n'; // 国际化
import { alarmHistoryAPI } from 'api/request.ts'; // 报警历史接口
import publicIndexStore from 'store/index.ts'; // 公共store

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();
    
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    projectId: string;
    instanceInfo: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    projectId: '',
    instanceInfo: []
});

/**
 * 回传对话框关闭
 */
const emitClose = defineEmits(['close']);

/**
 * 监听蒙层开关
 * @param {boolean} newValue 蒙层显示隐藏 true false
*/
watch(() => props.openVisible, (newValue: Required<boolean>): Readonly<void> => {  
    emitClose('close', newValue);
});

/**
 * 语言
*/
const locale: any = languageSwitch();

/**
 * 实例item类型
 */
 type InstancesType = {
    alertTime: string,
    resourceName: string,
    triggerDescription: string,
    alertValue: string,
    alertLevelDescription: string,
    alertPeriod: string
}

/**
 * 报警规则类型
 */
interface instance {
    tableData: InstancesType[],
    operationTime: any

}

/**
 * 项目列表数据
 */
const alarmHistoryArr: instance = reactive<instance>({
    tableData: [],
    operationTime: []
})

/**
 * loadig态
*/
const isLoading: Ref<boolean> = ref<boolean>(true);

/**
* 当前页面页数条数
*/
const pageSize: Ref<number> = ref<number>(20);

/**
* 当前页面页数
*/
const pageNumber: Ref<number> = ref<number>(1);

/**
* 总条数
*/
const totalCount: Ref<number> = ref<number>(0);

/**
 * 禁用的日期
 */
 const disabledDate = (time: Date) => {
    return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 180;
}

/**
 * 请求报警历史列表接口
*/
const requestAlarmHistory = (params: any): void => {
    alarmHistoryAPI({
        ...params,
        projectId: props.projectId,
    }).then(({data} : {data: any}) => {
        if (data?.result?.instances?.length) {
            alarmHistoryArr.tableData = data.result.instances.map((item:any) => {
                item.alertTime = getDateMinutes(item.alertTime);
                item.alertValue = item.alertValue;
                item.alertPeriod = t('alarm.alarmRule.minute', [item.alertPeriod])
                return {
                    ...item,
                    tooltipStatus: {
                        'alarmTime': {
                            showTooltip: false
                        },
                        'alarmResource': {
                            showTooltip: false
                        },
                        'triggeringCondition': {
                            showTooltip: false
                        },
                        'alarmValue': {
                            showTooltip: false
                        },
                        'alarmLevel': {
                            showTooltip: false
                        },
                        'continueTime': {
                            showTooltip: false
                        },
                    }
                }
            });
            totalCount.value = data.result.totalCount;
            return;
        }
        alarmHistoryArr.tableData = [];
        totalCount.value = 0;      
    }).catch(({message} : {message: string;}) => {
        if (message.indexOf('undefined') > -1) return;
    }).finally(() => {
        isLoading.value = false;     
    });
}

/**
 * 搜索结果提示
 */
 const searchResultMark: Ref<boolean> = ref<boolean>(false)


/**
 * 处理参数
*/
const processingParameter = () => {
    isLoading.value = true;
    let timeObject
    if(alarmHistoryArr.operationTime === null) {
        alarmHistoryArr.operationTime = [];
        searchResultMark.value = false
    }
    if(alarmHistoryArr.operationTime.length) {
        timeObject = convertStringsToTimestampObject(alarmHistoryArr.operationTime, 'startTime', 'endTime'); 
        searchResultMark.value = true
    }
    const params = {
        pageSize: pageSize.value,
        pageNumber: pageNumber.value,
        ruleId: props.instanceInfo.ruleId,
        ...timeObject,
    };
    publicStore.deleteEmtpyData(params);
    requestAlarmHistory(params);
};
processingParameter()

/**
 * 返回列表
 */
 const backToList = () =>{
    alarmHistoryArr.operationTime = null;
    searchOperationData()
}

/**
 * 每页展示条数切换
 * @param {number} size 每页展示条数
*/
const pageSizesChange = (size: number) => {
    pageNumber.value = 1;
    pageSize.value = size;
    isLoading.value = true;
    processingParameter();
};

/**
 * 页数切换
 * @param {number} num 页数切换
*/
const changePage = (num: number): void => {
    pageNumber.value = num;
    isLoading.value = true;
    processingParameter();
};

/**
 * 搜索结果
*/
const searchOperationData = (): void => {
    pageNumber.value = 1;
    pageSize.value = 20;
    processingParameter();
};


/**
 * 点击取消
 */
const closeDialog:  () => void = () => {
    emitClose('close', false)
}

const handleTableScroll = () => {
    // 遍历表格数据，将所有 tooltip 的 v-model 设置为 false 来隐藏它们
    alarmHistoryArr.tableData.forEach((row: any) => {
        Object.keys(row.tooltipStatus).forEach((key) => {
            row.tooltipStatus[key].showTooltip = false;
        });
    });
};

// 在组件的 onMounted 生命周期钩子中设置监听器
onMounted(() => {
  window.addEventListener('wheel', handleWheelEvent);
});

// 处理鼠标滚轮滚动事件
const handleWheelEvent = (event: WheelEvent) => {
  // event.deltaY 会告诉你滚轮的方向和距离
  if (event.deltaY < 0) {
    handleTableScroll()
  } else if (event.deltaY > 0) {
    handleTableScroll()
  }
};

// 在组件的 onUnmounted 生命周期钩子中清除监听器
onUnmounted(() => {
  window.removeEventListener('wheel', handleWheelEvent);
});

</script>
<style lang="scss">
@import 'assets/css/middleDialog';
@import 'assets/css/batch-table';
@import 'assets/css/icon';
.big-dialog {
    .el-dialog {
        --el-dialog-width: 840px !important; 
        .el-dialog__body {
            padding: 10px 20px 20px 20px;
        }    
    } 
    .footer-count {
        margin-top: 10px;
        .el-select {
            width: 70px !important;
        }
        .el-input {
            width: 70px !important;
            .el-input__inner {
                background-color: #FAFAFB;
            }
            
        }
    }
    .more-height {
        .el-table {
            max-height: 300px !important;
        }
    }

    .operation-search {
        height: 50px;
        .el-date-editor {
            width: 260px;
            .el-range-input {
                font-size: 12px;
            }
        }
        .alarm-history-search-button {
            margin-left: 20px;
            display: inline-block;
            vertical-align: bottom;
            .el-button {
                font-size: 12px;
            }
        }
    }

    .search-result {
        font-size: 12px;
        margin-bottom: 10px;
    }
        
}
</style>