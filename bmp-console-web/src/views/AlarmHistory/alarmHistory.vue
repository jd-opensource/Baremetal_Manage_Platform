<template>
    <div class="page-position alarm-history-list">
        <div class="page-content pos-fixed">
            <el-config-provider :locale="locale">
                <div class="operation-header">
                    <!-- 日期搜索 -->
                    <div class="operation-search h30">
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
                    <div class="setting-position mt-40"> 
                        <!-- 刷新 -->
                    <p
                        class="operate-refresh"
                        @click="processingParameter"
                    >
                        <span class="my-update-gray"></span>
                        <!-- 刷新 -->
                        <span class="refresh-title">
                            {{$t('list.update')}}
                        </span>
                    </p>
                    <p
                        :class="[exportDataOperate.hasExportData.value ? 'operate-export' : 'no-export']"
                        @click="exportDataOperate.exportData"
                    >
                        <span class="my-download-gray"></span>
                        <!-- 导出 -->
                        <span class="export-title">
                            {{$t('list.download')}}
                        </span>
                    </p> 
                    </div>

                </div>
                <div class="table-content"
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
                        v-loading="isLoading"
                        style="width: 100%"
                        :max-height="alarmHistoryArr.tableData.length ? tableMaxHeight : 420" 
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
                        <!-- 通知对象 -->
                        <el-table-column 
                            prop="noticeObject" 
                            :label="$t('alarm.alarmRule.noticeObject')" 
                            min-width=150
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['noticeObject'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['noticeObject'].showTooltip"
                                    :content= scope.row.userName
                                >
                                    <div
                                        class="long-instance-name"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['noticeObject'], scope.row.userName, 1.17, 'list')"
                                    >
                                        <span  @click.stop="jumpToPersonCentre()"
                                        >{{scope.row.userName || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <template #empty>
                            <div class="noData">                        
                            </div>
                            <p 
                                v-if="!alarmHistoryArr.tableData.length && !isLoading"
                                class="no-data-jump"
                            >
                            <span class="c333 f12">{{$t('instance.create.noData')}}</span>
                            </p>
                        </template>
                    </el-table>                
                </div>
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
            </el-config-provider>
        </div>
    </div>
</template>

<script setup lang="ts">
import { 
    Ref, 
    ref,
    watch,
    reactive,
    nextTick,
    onMounted,
    getCurrentInstance,
    onUnmounted
 } from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {
    convertStringsToTimestampObject, //转换时间为时间戳
    getDateMinutes,
    hasShowTooltip, // 是否显示提示
    getDate, // 获取日期
    generateRandomNum, // 生成随机数
    languageSwitch, // 语言切换
    deepCopy, //深拷贝
} from 'utils/index.ts';
import myPagination from 'components/Pagination/MyPagination.vue';
import  {ElMessage } from 'element-plus';
import { useRouter, useRoute } from 'vue-router';
import { useI18n } from 'vue-i18n'; // 国际化
import { alarmHistoryAPI, alarmHistoryListExportAPI } from 'api/request.ts'; // 报警历史接口
import publicIndexStore from 'store/index.ts'; // 公共store

/**
 * 国际化
*/
const {t} = useI18n();

const router = useRouter();
/**
 * 路由带值
 */
 const route = useRoute();

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();

// 计算表格的最大高度
const tableMaxHeight: Ref<number> = ref<number>(window.innerHeight - 300)
// 监听窗口大小变化，以便动态调整表格高度
const updateSize = () => {
  // 触发响应式更新
  tableMaxHeight.value = window.innerHeight - 300
};

onMounted(() => {
  window.addEventListener('resize', updateSize);
});

onUnmounted(() => {
  window.removeEventListener('resize', updateSize);
});

/**
 * 禁用的日期
 */
 const disabledDate = (time: Date) => {
    return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 180;
}

/**
 * 导出数据操作
*/
class ExportDataOperate {
    hasExportData: Ref<boolean> = ref<boolean>(true);
    /**
     * 导出数据
    */
    exportData = (): void => {
        if (!this.hasExportData.value) return;
        this.hasExportData.value = false;
        alarmHistoryListExportAPI(
            {
                exportType: '1',
                isAll: '1',
                projectId: route.query.projectId
            }
        )
        .then(({data} : {data: string;}) => {
            const blob: Blob = new Blob([data], {
                // blob类型
                type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet;charset=utf-8',
            });
            // 创建a标签
            let link = document.createElement('a');
            // 创建下载的链接
            link.href = window.URL.createObjectURL(blob);
            // 下载后的文件名
            link.download = `alarm_rule_list_${getDate()}_${generateRandomNum(6)}`;
            // 点击下载
            link.click();
            // 释放这个url
            window.URL.revokeObjectURL(link.href);
            nextTick(() => {
                this.hasExportData.value = true;
            });
        })
        .catch(() => {
            this.hasExportData.value = true;
        });
    };
};
// 实例化
const exportDataOperate: ExportDataOperate = new ExportDataOperate();

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
    userName: string
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
 * 设置类
*/
type SetType<T> = T extends {} ? any : T;

/**
 * 使用mitt传值
*/
const instanceMitt: Exclude<Required<SetType<{}>> | null, never> = getCurrentInstance();

/**
 * 请求报警历史列表接口
*/
const requestAlarmHistory = (params: any): void => {
    alarmHistoryAPI({
        ...params,
        projectId: route.query.projectId
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
                        'noticeObject': {
                            showTooltip: false
                        }
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
        (instanceMitt as any).proxy.$Bus.emit('project-name', route.query.projectName); 
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
        alarmHistoryArr.operationTime = []
        searchResultMark.value = false
    }
    if(alarmHistoryArr.operationTime.length) {
        timeObject = convertStringsToTimestampObject(alarmHistoryArr.operationTime, 'startTime', 'endTime'); 
        searchResultMark.value = true
    }
    const params = {
        pageSize: pageSize.value,
        pageNumber: pageNumber.value,
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
 * 搜索结果
*/
const searchOperationData = (): void => {
    pageNumber.value = 1;
    pageSize.value = 20;
    processingParameter();
};

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



const jumpToPersonCentre = () =>{
    router.push({
        path: `/user`,
        query: {
            type: 'account',   
        }
    });
}

</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import 'assets/css/addButton';
@import 'assets/css/icon';
@import 'assets/css/list';
@import './alarmHistory.scss';

</style>