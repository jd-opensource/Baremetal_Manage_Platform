<template>
    <div class="page-position alarm-rule-list">
        <div class="page-content pos-fixed">
            <div class="operation-header">
                <el-button 
                    type="primary" 
                    class="add-button-style" 
                    :class="getLocationItem === 'zh_CN' ? '' : 'w140'" 
                    @click="jumpToCreate"
                >
                     + &nbsp;&nbsp;{{$t('alarm.alarmRule.addRule') }}
                </el-button>
                <div class="setting-position">
                    <!-- 搜索筛选框 -->
                    <p class="operate-ipt">
                        <search-input
                            :class="[
                                getLocationItem === 'zh_CN' ? 'operate-management-input' : 'small-input'
                            ]"
                            :search-input="true"
                            :select-option="searchOperate.reactiveArr.selectOption"
                            :place-holder="[
                                $t('alarm.placeHolder.ruleName'),
                                $t('alarm.placeHolder.ruleId')

                            ]"
                            :has-clear="searchOperate.hasClear.value"
                            @ipt-value="searchOperate.iptValue"
                            @change-select="searchOperate.changeSelect"
                            @select="searchOperate.selectChange"
                            @clear-operate="searchOperate.clearClick"
                        />
                    </p>
                    <!-- 刷新 -->
                    <p
                        class="operate-refresh"
                        @click="refreshData"
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
            <div 
                class="table-content"
                :class="[
                    searchResultMark ? 'mt16' : 'mt30'
                ]"
            >
                <el-config-provider :locale="locale">
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
                        :data="reactiveArr.tableData" 
                        v-loading="isAlarmRuleLoading"
                        style="width: 100%"
                        :max-height="reactiveArr.tableData.length ? tableMaxHeight : 420" 
                        @filter-change="filterChange"
                    >
                        <!-- 规则名称 -->
                        <el-table-column 
                            prop="ruleName" 
                            fixed
                            v-if="reactiveArr?.hasCustomInfo['ruleName']?.selected"
                            :label="$t('alarm.alarmRule.ruleName')" 
                            min-width=150
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['ruleName'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['ruleName'].showTooltip"
                                    :content= scope.row.ruleName
                                >
                                    <div
                                        class="long-instance-name"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['ruleName'], scope.row.ruleName, 1.17, 'list')"
                                    >
                                        <span  @click.stop="jumpToDetail(scope.row)"
                                        >{{scope.row.ruleName || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 规则ID -->
                        <el-table-column 
                            prop="ruleId" 
                            v-if="reactiveArr?.hasCustomInfo['ruleId']?.selected"
                            :label="$t('alarm.alarmRule.ruleId')" 
                            min-width=130
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['ruleId'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['ruleId'].showTooltip"
                                    :show-after="500"
                                    placement="right"
                                    :content= scope.row.ruleId
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['ruleId'], scope.row.ruleId, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.ruleId || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 资源类型 -->
                        <el-table-column 
                            prop="resource" 
                            v-if="reactiveArr?.hasCustomInfo['resourceName']?.selected"
                            :label="$t('alarm.alarmRule.resourceType')" 
                            min-width=100
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['resource'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['resource'].showTooltip"
                                    :content= scope.row.resourceName
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['resource'], scope.row.resourceName, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.resourceName || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 关联实例数量 -->
                        <el-table-column 
                            prop="instanceCount" 
                            v-if="reactiveArr?.hasCustomInfo['instanceCount']?.selected"
                            :label="$t('alarm.alarmRule.instanceNumber')" 
                            width="120"
                        >
                            <template v-slot="scope">
                                <el-button 
                                    type="text"
                                    class="f12"
                                    @click="jumpToDetail(scope.row)"
                                >
                                    {{ scope.row.relatedResourceCount || $filter.numberEmptyFilter()}}
                                </el-button>
                            </template>
                        </el-table-column>
                        <!-- 触发条件 -->
                        <el-table-column 
                            prop="triggeringCondition" 
                            v-if="reactiveArr?.hasCustomInfo['triggerDescription']?.selected"
                            :label="$t('alarm.alarmRule.triggeringCondition')" 
                            min-width=320
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['triggeringCondition'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['triggeringCondition'].showTooltip"
                                    :show-after="500"
                                >
                                <template #content>
                                    <span v-for="{ item, key } in renderItemsWithBreaks(scope.row.triggerDescription)" :key="key">
                                        {{ item || $filter.emptyFilter() }}<br>
                                    </span>
                                </template>
                                    <div
                                        class="triggering-condition-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['triggeringCondition'], findLongestString(scope.row.triggerDescription), 1.17, 'list')"
                                    >                                       
                                        <span v-for="{ item, key } in renderItemsWithBreaks(scope.row.triggerDescription).slice(0, 1)" :key="key">
                                            {{ $t('alarm.alarmRule.total' ,[renderItemsWithBreaks(scope.row.triggerDescription).length]) + item || $filter.emptyFilter() }}<br>
                                        </span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 状态 -->
                        <el-table-column     
                            prop="status" 
                            key="status"
                            column-key="status"
                            v-if="reactiveArr?.hasCustomInfo['statusName']?.selected"
                            :label="$t('alarm.alarmRule.status')" 
                            filter-placement="bottom-end"
                            :filters="alarmRuleStatus"
                            :filter-method="alarmRuleStatusFilter"
                            min-width=100
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['status'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['status'].showTooltip"
                                    :content= scope.row.statusName
                                >
                                    <div
                                        class="long-row"
                                        :class="stateChange(scope.row.status)"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['status'], scope.row.statusName, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.statusName || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 操作 -->
                        <el-table-column 
                            :fixed="'right'"
                            prop="operation" 
                            :label="$t('personCentre.form.operation')" width="215">
                            <template v-slot="scope"
                        >  
                            <operate-unit
                               v-if="scope.row.status === 1 || scope.row.status === 3"
                                :operateName="$t('alarm.operate.disable.name')"
                                :operateTip="$t('alarm.operate.disable.tip')"
                                :disabled="false"
                                :instanceInfo="scope.row"
                                @handleClick="clickOperateDisable(scope.row)"
                            >
                            </operate-unit>
                            <operate-unit
                                v-else-if="scope.row.status === 2"
                                :operateName="$t('alarm.operate.enable.name')"
                                :operateTip="$t('alarm.operate.enable.tip')"
                                :disabled="false"
                                :instanceInfo="scope.row"
                                @handleClick="clickOperateEnable(scope.row)"
                            >
                            </operate-unit>
                            <operate-unit
                                :operateName="$t('alarm.operate.edit.name')"
                                :operateTip="$t('alarm.operate.edit.tip')"
                                :disabled="false"
                                @handleClick="clickOperateEdit(scope.row)"
                                :instanceInfo="scope.row"
                            >
                            </operate-unit> 
                            <div class="one-line-drop" @click.stop="">
                                <drop-down-operation
                                    class="batch-operation-button" 
                                    :operations="scope.row.operations"
                                    :instanceInfo="scope.row"
                                    @clickSelectItem="clickMoreOperate($event, scope.row)"
                                >
                                </drop-down-operation>     
                            </div>
                            </template>   
                        </el-table-column>
                        <template #empty>
                            <div class="noData">                        
                            </div>
                            <p 
                                v-if="!reactiveArr.tableData.length && !isAlarmRuleLoading"
                                class="no-data-jump"
                            >
                                <span class="c333 f12">{{$t('alarm.alarmRule.noData')}}</span>
                                <el-button type="text" class="f12 pt7" @click="jumpToCreate">{{$t('alarm.alarmRule.addRule') }}</el-button>
                            </p>
                        </template>
                    </el-table>    
                </el-config-provider>
                <div v-if="openCustom">
                    <custom
                        :page-key="'monitorRulesList'"
                        :open-visible="openCustom"
                        :check-list-arr="reactiveArr.checkListArr"
                        :has-custom-info="reactiveArr.hasCustomInfo"
                        @close="openCustomCancel"
                        @determine="publicStore.customList('monitorRulesList', reactiveArr)"
                    >
                    </custom>
                </div>
                <div v-if="disable">
                    <alarm-rule-disable
                        :openVisible="disable"
                        :instanceInfo="instanceInfo"
                        :project-id="route.query.projectId"
                        @close="disableCancel"
                        @refresh="refreshData"
                    >
                    </alarm-rule-disable> 
                </div>
                <div v-if="enable">
                    <alarm-rule-enable
                        :openVisible="enable"
                        :instanceInfo="instanceInfo"
                        :project-id="route.query.projectId"
                        @close="enableCancel"
                        @refresh="refreshData"
                    >
                    </alarm-rule-enable> 
                </div>
                <div v-if="history">
                    <alarm-rule-history
                        :openVisible="history"
                        :instanceInfo="instanceInfo"
                        :project-id="route.query.projectId"
                        @close="historyCancel"
                        @refresh="refreshData"
                    >
                    </alarm-rule-history>
                </div>
                <div v-if="deleteRule">
                    <alarm-rule-delete
                        :openVisible="deleteRule"
                        :instanceInfo="instanceInfo"
                        :project-id="route.query.projectId"
                        @close="deleteCancel"
                        @refresh="refreshData"
                    >
                    </alarm-rule-delete>
                </div>
            </div>
            <div class="footer-count">
                <my-pagination
                    v-if="reactiveArr.tableData.length > 0 && !isAlarmRuleLoading"
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
        </div>
    </div>
</template>

<script setup lang="ts">
import {
    Ref, 
    ref,
    reactive,
    watch,
    computed,
    onMounted,
    onUnmounted,
    onBeforeUnmount,
    getCurrentInstance,
    nextTick
} from 'vue';
import VueCookies from 'vue-cookies'; // cookie
import { useRouter, useRoute } from 'vue-router';
import i18n from 'lib/i18n/index.ts'; // 国际化
import searchInput from 'components/SearchInput/searchInput.vue';
import allProjectStore from '@/store/modules/headerDetail.ts';
import { ElMessage, ElTable } from 'element-plus';
import {
    alarmRuleCustomInfo, // 报警规则列表自定义初始数据
    alarmRuleCustomIntro, // 报警规则列表自定义信息
    alarmRuleStatus,
} from 'utils/constants.ts';
import {
    getDate, // 获取日期
    generateRandomNum, // 生成随机数
    filterData, // 过滤重复数据
    languageSwitch, // 语言切换
    deepCopy, //深拷贝
    hasShowTooltip, // 是否显示提示
    renderItemsWithBreaks,
    findLongestString
} from 'utils/index.ts';
import publicIndexStore from 'store/index.ts'; // 公共store
import {projectListStore} from 'store/loginInfo.ts'; // store库存储的用户名
import { alarmRuleListAPI, alarmRuleListExportAPI } from 'api/request.ts'; // 列表数据接口
import custom from 'components/custom/custom.vue';
import myPagination from 'components/Pagination/MyPagination.vue';
import operateUnit from 'components/OperateUnit/operateUnit.vue';
import dropDownOperation from 'components/DropdownOperation/DropdownOperation.vue';
import alarmRuleDisable from 'components/AlarmRuleOperate/alarmRuleDisable.vue';
import alarmRuleEnable from 'components/AlarmRuleOperate/alarmRuleEnable.vue';
import alarmRuleHistory from 'components/AlarmRuleOperate/alarmRuleHistory.vue';
import alarmRuleDelete from 'components/AlarmRuleOperate/alarmRuleDelete.vue';

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: any = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * 国际化
*/
const {global} = i18n;

const router = useRouter();
/**
 * 路由带值
 */
 const route = useRoute();

 /**
 * 设置类
*/
type SetType<T> = T extends {} ? any : T;

/**
 * 使用mitt传值
*/
const instanceMitt: Exclude<Required<SetType<{}>> | null, never> = getCurrentInstance();

/**
 * 跳入创建页
 */
 const jumpToCreate = () => {
    router.push({
        path: `/instance/rules/create`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName
        }
    });
}

/**
 * 搜索结果提示
 */
 const searchResultMark: Ref<boolean> = ref<boolean>(false)

interface SearchArrType {
    searchParams: {
        ruleName?: string;
        ruleId?: string;
    };
    selectOption: {
        value: number;
        label: string;
    }[];
};

class SearchOperate {
    hasClear: Ref<boolean> = ref<boolean>(false);
    // 搜索框值
    selectVal: Ref<number> = ref<number>(0);
    reactiveArr: SearchArrType = reactive<SearchArrType>({
        searchParams: {},
        selectOption: [
            {
                value: 0,
                label: global.t('alarm.alarmRule.ruleName')
            },
            {
                value: 1,
                label: global.t('alarm.alarmRule.ruleId')
            }
        ]
    });

    /**
     * input 输入框 点击搜索按钮或者回车触发
     * @param {string} iptValue 输入框输入的值
     * @param {number} selectValue 筛选框筛选的值
    */
    iptValue = (iptValue: string, selectValue: number): void => {
        switch(selectValue) {
            case 0 :
                this.reactiveArr.searchParams = {ruleName: iptValue}
                break; 
            case 1 :
                this.reactiveArr.searchParams = {ruleId: iptValue}
                break; 
            default:
                break;

        }
        searchResultMark.value = true
        this.request();
    };

    /**
     * 搜索框筛选
     * @param {number} val 搜索框切换的搜索key
     * @return {number} selectVal 对应的key
    */
    changeSelect = (val: number): number => {
        return this.selectVal.value = val;
    };

    clearClick = (val: string) => {
        searchResultMark.value = false
        if (!val) {
            this.selectChange();
        }
    };

    selectChange = () => {
        searchResultMark.value = false
        const {
            ruleId, 
            ruleName
        } : {
            ruleId?: string; 
            ruleName?: string;
        } = this.reactiveArr.searchParams;
        if (ruleId || ruleName) {
            this.reactiveArr.searchParams = {};
            this.request();
        }
    };

    request = () => {
        if (pageNumber.value > 1) {
            pageNumber.value = 1;
            //return;
        }
        refreshData();
    };

};

const searchOperate: SearchOperate = new SearchOperate();

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
        alarmRuleListExportAPI(
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
 * 表格ref
*/
const tableRef = ref<InstanceType<typeof ElTable>>()

/**
 * 返回列表
 */
 const backToList = () =>{
    searchOperate.hasClear.value = true;
    searchOperate.reactiveArr.searchParams = {};
    reactiveArr.filterParams['status'] = '';
    // @ts-ignore
    tableRef.value!.clearFilter();
    searchOperate.request();
    searchResultMark.value = false
}


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
 * filter参数类
*/
interface FilterParamsType {
    status: string,
};

const filterChange = (filter: {[x: string]: any;}) => {
    const filterParams: FilterParamsType = {
        status: 'status',
    };
    const tag: string[] = [];
    for (const key in filter) {
        if (key === filterParams[key]) {
            for (const item of filter[key]) {
                tag.push(item);    
            }
            reactiveArr.filterParams[filterParams[key]] = tag.join(',');
        }
    }
    refreshData();
    if(reactiveArr.filterParams['status']) {
        searchResultMark.value = true 
    } else if(!reactiveArr.filterParams['status'] && !searchOperate.reactiveArr.searchParams.ruleName && !searchOperate.reactiveArr.searchParams.ruleId) {
        searchResultMark.value = false
    }
        
    
};

/**
 * 状态筛选
*/
const alarmRuleStatusFilter = (value: string, row: any): boolean => {
    return value === row.status;
};

const jumpToDetail = (row: any) => {  
    router.push({
        path: `/instance/rules/detail`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName,
            ruleId: row.ruleId
        }
    });
    
}

/**
 * store库存储的项目列表
*/
const store = allProjectStore();

/**
 * 实例item类型
 */
type InstancesType = {
    ruleName: string,
    ruleId: string,
    resource: string,
    instanceIds: [],
    triggeringCondition: string,
    status: string,
    operations: any
}

/**
 * 报警规则类型
 */
interface instance {
    tableData: InstancesType[],
    filterParams: any,
    hasCustomInfo: {},
    checkListArr: {},
}

/**
 * 项目列表数据
 */
const reactiveArr: instance = reactive<instance>({
    tableData: [],
    filterParams: {
        status: ''
    },
    hasCustomInfo: alarmRuleCustomIntro,
    checkListArr: alarmRuleCustomInfo,
})

/**
 * loadig态
*/
const isAlarmRuleLoading: Ref<boolean> = ref<boolean>(true);

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
 * 公共store
*/
const publicStore: any = publicIndexStore();

/**
 * 请求报警规则列表数据接口
*/
const requestAlarmRuleListData = (params: any): void => {
    alarmRuleListAPI({
        ...params
    }).then(({data} : {data: any}) => {
        if (data?.result?.rules?.length) {
            reactiveArr.tableData = data.result.rules.map((item:any, index:number) => {
                let allOperations : any = [
                    {
                        operateName: global.t('alarm.operate.history.name'),
                        disabled: false,
                        tip: global.t('alarm.operate.history.name'),
                        value: 'history'
                    },{
                        operateName: global.t('alarm.operate.delete.name'),
                        disabled: false,
                        tip: global.t('alarm.operate.delete.name'),
                        value: 'ruleDelete'
                    }
                ]
                return {
                    ...item,
                    operations: allOperations,
                    tooltipStatus: {
                        ruleName: {
                            showTooltip: false
                        },
                        ruleId: {
                            showTooltip: false
                        },
                        resource: {
                            showTooltip: false
                        },
                        instanceCount: {
                            showTooltip: false
                        },
                        triggeringCondition: {
                            showTooltip: false
                        },
                        status: {
                            showTooltip: false
                        },
                        operations: {
                            showTooltip: false
                        },
                    }
                    
                }
            });
            totalCount.value = data.result.totalCount;
            return;
        }
        reactiveArr.tableData = [];
        totalCount.value = 0;      
    }).catch(({message} : {message: string;}) => {
        if (message.indexOf('undefined') > -1) return;
    }).finally(() => {
        isAlarmRuleLoading.value = false;
        searchOperate.hasClear.value = false;  
        (instanceMitt as any).proxy.$Bus.emit('project-name', route.query.projectName);     
    });
};

/**
 * 状态tag变化
 * @param value 
 */
 const stateChange = (value: number,) => {
    switch (value) {
        case 1:
            return 'success';
        case 2:
            return 'info';
        case 3:
            return 'danger';   
        default:
            return 'intermediate';
    }
}

/**
 * store库存储的状态类
*/
const storeUser = projectListStore();

/**
 * 处理参数
*/
const processingParameter = () => {
    const params = {
        pageSize: pageSize.value,
        pageNumber: pageNumber.value,
        userId: store.userForm.userId,
        userName: store.userForm.userName,
        projectId: route.query.projectId,
        ...reactiveArr.filterParams,
        ...searchOperate.reactiveArr.searchParams,
    };
    publicStore.deleteEmtpyData(params);
    requestAlarmRuleListData(params);
};

/**
 * 刷新接口
*/
const refreshData = (): void => {
    isAlarmRuleLoading.value = true;
    processingParameter();
    

};
refreshData()



/**
 * 每页展示条数切换
 * @param {number} size 每页展示条数
*/
const pageSizesChange = (size: number) => {
    pageNumber.value = 1;
    pageSize.value = size;
    isAlarmRuleLoading.value = true;
    processingParameter();
};

/**
 * 页数切换
 * @param {number} num 页数切换
*/
const changePage = (num: number): void => {
    pageNumber.value = num;
    isAlarmRuleLoading.value = true;
    processingParameter();
};

/**
 * 当前行值
 */
 const instanceInfo: Ref<any> = ref();


/**
 * 禁用操作打开标志
 */
const disable: Ref<boolean> = ref<boolean>(false)


/**
 * 禁用操作
 * @param value 
 */
 const clickOperateDisable: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    disable.value = !disable.value;
}

/**
 * 禁用弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
const disableCancel = (type: boolean): boolean => {
    return disable.value = type;
};

/**
 * 启用操作打开标志
 */
 const enable: Ref<boolean> = ref<boolean>(false)


/**
 * 启用操作
 * @param value 
 */
 const clickOperateEnable: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    enable.value = !enable.value;
}

/**
 * 启用弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
const enableCancel = (type: boolean): boolean => {
    return enable.value = type;
};

/**
 * 启用操作
 * @param value 
 */
 const clickOperateEdit: (value: any) => void = (value: any) => {
    router.push({
        path: `/instance/rules/create`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName,
            edit: value.ruleId
        },
    });
}

/**
 * 报警历史操作打开标志
 */
 const history: Ref<boolean> = ref<boolean>(false)


/**
 * 报警历史操作
 * @param value 
 */
 const clickOperateHistory: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    history.value = !history.value;
}

/**
 * 报警历史弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
const historyCancel = (type: boolean): boolean => {
    return history.value = type;
};

/**
 * 删除操作打开标志
 */
 const deleteRule: Ref<boolean> = ref<boolean>(false)


/**
 * 删除操作
 * @param value 
 */
 const clickOperateDelete: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    deleteRule.value = !deleteRule.value;
}

/**
 * 删除弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
const deleteCancel = (type: boolean): boolean => {
    return deleteRule.value = type;
};

/**
 * 操作下拉列表操作
 * @param value 
 * @param row 
 */
 const clickMoreOperate=(value: any, row: any) =>{
    switch (value) {
        case 'history':
            clickOperateHistory(row)
            break;
        case 'ruleDelete':
            clickOperateDelete(row)
            break;
        default:
            break;
    }
};


/**
 * 自定义列表接口
*/
publicStore.customList('monitorRulesList', reactiveArr);


</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import 'assets/css/addButton';
@import 'assets/css/icon';
@import 'assets/css/list';
@import './alarmRuleList.scss';

</style>