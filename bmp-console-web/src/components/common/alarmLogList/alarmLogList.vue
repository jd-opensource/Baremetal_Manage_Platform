<template>
    <div class="alarm-log-list">
        <div class="operation-header">
            <div class="setting-position mr0">
                <!-- 刷新 -->
                <p
                    class="operate-refresh"
                    @click="refreshAlarmLoglistData"
                >
                    <span class="my-update-gray"></span>
                    <!-- 刷新 -->
                    <span class="refresh-title">
                        {{$t('list.update')}}
                    </span>
                </p>
                <p
                    class="operate-set-up"
                    @click="customVisible"
                >
                    <span class="my-settings-gray"></span>
                    <!-- 设置 -->
                    <span class="set-up-title">
                        {{$t('list.settings')}}
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
        <div class="table-content ml0 mr0">
            <el-config-provider :locale="locale">
                <el-table 
                    border
                    :data="reactiveArr.tableData" 
                    v-loading="alarmLogListLoading"
                    style="width: 100%;"
                    :max-height="reactiveArr.tableData.length ? tableMaxHeight : 350" 
                >
                    <!-- 序号 -->
                    <el-table-column 
                        prop="logid" 
                        v-if="reactiveArr?.hasCustomInfo['logid']?.selected"
                        :label="$t('instance.detail.number')" 
                        width = 60
                    >
                        <template v-slot="scope">
                            {{scope.row.logid || $filter.emptyFilter()}}
                        </template>
                    </el-table-column>
                    <!-- 故障类型 -->
                    <el-table-column 
                        prop="alert_type" 
                        v-if="reactiveArr?.hasCustomInfo['alertType']?.selected"
                        :label="$t('instance.detail.faultType')" 
                    >
                        <template v-slot="scope">
                            {{scope.row.alert_type || $filter.emptyFilter()}}
                        </template>
                    </el-table-column>
                    <!-- 故障描述 -->
                    <el-table-column 
                        prop="alert_desc" 
                        v-if="reactiveArr?.hasCustomInfo['alertDesc']?.selected"
                        :label="$t('instance.detail.faultDesc')"
                    >
                        <template v-slot="scope">
                            {{scope.row.alert_desc || $filter.emptyFilter()}}
                        </template>
                    </el-table-column>
                    <!-- 故障报警时间 -->
                    <el-table-column 
                        prop="event_time" 
                        v-if="reactiveArr?.hasCustomInfo['eventTime']?.selected"
                        :label="$t('instance.detail.faultAlarmTime')" 
                    >
                        <template v-slot="scope">
                            {{scope.row.event_time || $filter.emptyFilter()}}
                        </template>
                    </el-table-column>
                    <!-- 更新时间 -->
                    <el-table-column 
                        prop="update_time" 
                        v-if="reactiveArr?.hasCustomInfo['updateTime']?.selected"
                        :label="$t('instance.detail.updatedTime')" 
                    >
                        <template v-slot="scope">
                            {{scope.row.update_time || $filter.emptyFilter()}}
                        </template>
                    </el-table-column>
                    <!-- 报警次数 -->
                    <el-table-column 
                        prop="count" 
                        v-if="reactiveArr?.hasCustomInfo['count']?.selected"
                        :label="$t('instance.detail.alarmNumbers')" 
                    >
                        <template v-slot="scope">
                            {{scope.row.count || $filter.emptyFilter()}}
                        </template>
                    </el-table-column>
                    <!-- 状态 -->
                    <el-table-column 
                        prop="status" 
                        v-if="reactiveArr?.hasCustomInfo['status']?.selected"
                        :label="$t('instance.detail.status')" 
                    >
                        <template v-slot="scope">
                            <!-- 0代表未处理 1代表已恢复 2代表已忽略 -->
                            <span :class="scope.row.status === 0 ? 'status-notprocessed' : 'status-processed'">
                                {{alarmStatusConstant[scope.row.status] || $filter.emptyFilter() }}
                            </span>	
                        </template>
                    </el-table-column>
                    <template #empty>
                        <div class="noData">                        
                        </div>
                    </template>
                </el-table>
            </el-config-provider>   
        </div>
        <div class="footer-count">
            <my-pagination
                v-if="reactiveArr.tableData.length > 0 && !alarmLogListLoading"
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
        <div v-if="openCustom">
            <custom
                :page-key="'device_alert_log_list'"
                :open-visible="openCustom"
                :check-list-arr="reactiveArr.checkListArr"
                :has-custom-info="reactiveArr.hasCustomInfo"
                sign="alarm"
                @close="openCustomCancel"
                @determine="publicStore.alarmLogCustomList('device_alert_log_list', reactiveArr)"
            >
            </custom>
        </div>        
    </div>
</template>

<script setup lang="ts">
import { 
    ref,
    Ref,
    reactive,
    nextTick,
    onMounted,
    onUnmounted,
    onBeforeUnmount,
    watch,
    onActivated
} from 'vue';
import {ElMessage} from 'element-plus';
import {
    getDate, // 获取日期
    languageSwitch, // 语言切换
    generateRandomNum, // 生成随机数
    debounce //防抖函数
} from 'utils/index.ts';
import publicIndexStore from 'store/index.ts'; // 公共store
import custom from 'components/custom/custom.vue';
import myPagination from 'components/Pagination/MyPagination.vue';
import {
    alarmLogCustomIntro, // 列表自定义初始数据
    alarmLogCustomInfo,
    alarmStatusConstant // 报警日志状态
} from 'utils/constants.ts';
import {
    alarmLogListExportAPI, 
    alarmLogListAPI
} from 'api/request.ts';

// 计算表格的最大高度
const tableMaxHeight: Ref<number> = ref<number>(window.innerHeight - 500)
// 监听窗口大小变化，以便动态调整表格高度
const updateSize = () => {
    // 触发响应式更新
    tableMaxHeight.value = window.innerHeight - 500
};

onMounted(() => {
    window.addEventListener('resize', updateSize);
});

onUnmounted(() => {
    window.removeEventListener('resize', updateSize);
});

/**
 * 语言
*/
const locale: any = languageSwitch();

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    instanceInfo: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    instanceInfo: {},
});


/**
 * 事件回传
 */
 const emitData = defineEmits(['refresh', 'goBack']);

 /**
  * 列表刷新
  */
const refreshData : () => void = () => {
    emitData('refresh');
}

/**
 * 报警日志列表item类型
 */
 type AlarmLoglistType = {
    logid: string,
    alert_type: string,
    alert_desc: string,
    event_time: string,
    update_time: string,
    count: string,
    status: string,
};

/**
 * 列表类型
 */
 interface alarmLog {
    tableData: AlarmLoglistType[]
    hasCustomInfo: {},
    checkListArr: {}
}

/**
 * 列表数据
 */
 const reactiveArr: alarmLog = reactive<alarmLog>({
    tableData: [],
    hasCustomInfo: alarmLogCustomIntro,
    checkListArr: alarmLogCustomInfo
})

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();

/**
 * 自定义列表操作打开标志
 */
 const openCustom: Ref<boolean> = ref<boolean>(false)

/**
 * 自定义列表操作
 * @param value 
 */
const customVisible: () => void = () => {
    openCustom.value = !openCustom.value;
} 
/**
 * 自定义列表弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} openCustom.value 弹窗关闭
 */
const openCustomCancel = (type: boolean): boolean => {
    return openCustom.value = type;
};

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
        alarmLogListExportAPI(
            {
                exportType: '1',
                isAll: '1',
                sn: props.instanceInfo.sn,
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
            link.download = `alarm_log_list_${getDate()}_${generateRandomNum(6)}`;
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
 * 报警日志列表loadig态
*/
const alarmLogListLoading: Ref<boolean> = ref<boolean>(false);
let timer: null | number = null;
onMounted(() => clearTimeout((timer as number)));
/**
 * 页面销毁时-触发，清空延时器
*/
onUnmounted(() => clearTimeout((timer as number)));
onBeforeUnmount(() => clearTimeout((timer as number)));
/**
 * 请求报警日志列表接口
*/
const requestAlarmLoglistData = (params: any): void => {
    alarmLogListLoading.value = true
    alarmLogListAPI({
        ...params,
    }).then(({data} : {data: any}) => {
        reactiveArr.tableData = data.result.detail;
        totalCount.value = data.result.total_count;
    }).catch(({message} : {message: string;}) => {      
        if (message === '找不到对象' || message === 'Not found') { 
            emitData('goBack');
            return;
        }
        if (message.indexOf('undefined') > -1) return;
        ElMessage.error( message );
    }).finally(() => {    
        alarmLogListLoading.value = false     
    });
};
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
 * 每页展示条数切换
 * @param {number} size 每页展示条数
*/
const pageSizesChange = (size: number) => {
    pageNumber.value = 1;
    pageSize.value = size;
    processingParameter();
};

/**
 * 页数切换
 * @param {number} num 页数切换
*/
const changePage = (num: number): void => {
    pageNumber.value = num;
    processingParameter();
};

/**
 * 处理参数
*/
const processingParameter = () => {
    const params = {
        page_size: pageSize.value,
        page_num: pageNumber.value,
        sn: props.instanceInfo.sn
    };
    publicStore.deleteEmtpyData(params);
    requestAlarmLoglistData(params);
    publicStore.alarmLogCustomList('device_alert_log_list', reactiveArr); //自定义列表接口
};

/**
 * 刷新报警日志列表接口
*/
const refreshAlarmLoglistData = (): void => {
    pageNumber.value = 1;
    pageSize.value = 20;
    processingParameter();
}

/**
 * 暴露方法
 */
defineExpose({ refreshAlarmLoglistData })
</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import 'assets/css/icon';
@import 'assets/css/list';
@import './alarmLogList.scss';
</style>