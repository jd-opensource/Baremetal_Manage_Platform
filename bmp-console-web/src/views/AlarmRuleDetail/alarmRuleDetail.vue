<template>
    <div class="page-position can-scroll">
        <div class="detail-header">
            <el-button
                class="back-button"
                type= "text"
                :icon ="ArrowLeft" 
                @click="goBack"
            >
                {{$t('list.back')}}
            </el-button>
            <span class="ml32">{{ $t('alarm.alarmRule.alarmDetail') }}</span>
            <span class="ml15">/</span>
            <div style="display: inline-flex;">
                <el-tooltip  
                    :disabled="!tooltipStatus['ruleNameInfo'].showTooltip"  
                    :content= alarmRuleInfo.detail.ruleName
                    placement="top-start"
                >
                    <span 
                        class="long-name-title ml15" 
                        @mouseenter="hasShowTooltip($event, tooltipStatus['ruleNameInfo'], alarmRuleInfo.detail.ruleName, 179, 'detail')"
                    >
                        {{ alarmRuleInfo.detail.ruleName || $filter.emptyFilter() }}
                    </span>
                </el-tooltip>
            </div>
            <span 
                class="ml24" 
                :class="stateChange(alarmRuleInfo?.detail?.status)"
            >
                {{alarmRuleInfo.detail.statusName}}
            </span>
            <div class="operation-position">
                <el-dropdown 
                    :class="[selectStatus ? 'active-operate' : 'inactive-operate']"
                    size="small"
                    @visible-change="selectHover"
                >
                    <el-button type="primary">
                        {{$t('personCentre.form.operation')}}<el-icon class="el-icon--right"><arrow-down /></el-icon>
                    </el-button>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <!-- 禁用 -->
                            <el-dropdown-item
                                v-if="alarmRuleInfo?.detail?.status === 1 || alarmRuleInfo?.detail?.status === 3 "
                            >
                                <el-tooltip
                                    placement="right"
                                    :disabled="true"
                                    :content="$t('alarm.operate.disable.tip')"
                                >
                                    <label>
                                        <el-button
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateDisable()"
                                            :disabled="false"
                                        >
                                            <span>{{$t('alarm.operate.disable.name')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 启用 -->
                            <el-dropdown-item
                                v-else-if="alarmRuleInfo?.detail?.status === 2"
                            >
                                <el-tooltip
                                    placement="right"
                                    :disabled="true"
                                    :content="$t('alarm.operate.enable.tip')"
                                >
                                    <label>
                                        <el-button  
                                            type="text"  
                                            class="operate-btn"
                                            :disabled="false"
                                            @click="clickOperateEnable()"
                                        >
                                            <span>{{$t('alarm.operate.enable.name')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 编辑 -->
                            <el-dropdown-item>
                                <el-tooltip
                                    placement="right"
                                    :disabled="true"
                                    :content="$t('alarm.operate.edit.tip')"
                                >
                                    <label>
                                        <el-button  
                                            type="text" 
                                            class="operate-btn"
                                            :disabled="false"
                                            @click="clickOperateREdit()"
                                        >
                                            <span>{{$t('alarm.operate.edit.name')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                             <!-- 报警历史  -->
                             <el-dropdown-item>
                                <el-tooltip
                                    placement="right"
                                    :disabled="true"
                                    :content="$t('alarm.operate.history.tip')"
                                >
                                    <label>
                                        <el-button
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateHistory()"
                                            :disabled="false"
                                        >
                                            <span>{{$t('alarm.operate.history.name')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 删除 -->
                            <el-dropdown-item>
                                <el-tooltip
                                    placement="right"
                                    :disabled="true"
                                    :content="$t('alarm.operate.delete.tip')"
                                >
                                    <label>
                                        <el-button  
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateDelete()"
                                            :disabled="false"
                                        >
                                            <span>{{$t('alarm.operate.delete.name')}}</span>
                                        </el-button >
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div>            
        </div>
        <div class="page-content alarm-detail-content t110">
            <el-tabs
				ref="pageTabs"
                class="tab-info"
				type="border-card" 
				v-model="tabName"
                @tab-click="changeTab"
            >
                <!-- tab:基本信息 -->
				<el-tab-pane 
					:label="$t('personCentre.form.basicInformation')" 
					name="basicInfo"
					:disabled="tabName==='basicInfo'"
                >
                    <!-- 基本信息 -->
                    <h3 class="section-title">{{ $t('personCentre.form.basicInformation')}}</h3>
                    <div class="ml22">
                        <el-form 
                            v-loading="isLoading"
                            @submit.stop.prevent
                        >                
                            <el-row :gutter="20">
                                <!--规则ID-->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.ruleId') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['ruleId'].showTooltip"  
                                            :content= alarmRuleInfo.detail.ruleId
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['ruleId'], alarmRuleInfo.detail.ruleId, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.ruleId || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                        <img
                                            class="desc-type"
                                            src="@/assets/img/copy.png"
                                            @click="clickOperateCopy(alarmRuleInfo.detail.ruleId)"
                                        />
                                    </el-form-item>
                                </el-col>
                                <!-- 规则名称 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.ruleName') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['ruleName'].showTooltip"   
                                            :content= alarmRuleInfo.detail.ruleName
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['ruleName'], alarmRuleInfo.detail.ruleName, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.ruleName || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                        <img
                                            class="desc-type"
                                            src="@/assets/img/copy.png"
                                            @click="clickOperateCopy(alarmRuleInfo.detail.ruleName)"
                                        />
                                    </el-form-item>   
                                </el-col>
                                <!-- 资源类型 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.resourceType') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['resourceType'].showTooltip"   
                                            :content= alarmRuleInfo.detail.resourceName
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['resourceType'], alarmRuleInfo.detail.resourceName, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.resourceName || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>   
                                </el-col>
                                <!-- 维度 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.dimension') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['dimension'].showTooltip"   
                                            :content= alarmRuleInfo.detail.dimension
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['dimension'], alarmRuleInfo.detail.dimension, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.dimension || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                        </el-form>
                    </div>
                    <el-divider />
                    <!-- 报警资源 -->
                    <h3 class="section-title">{{ $t('alarm.alarmRule.alarmResource') }}</h3>
                    <div class="table-content m0 h-max">
                        <el-table 
                            border
                            ref="datasTableRef"
                            :data="alarmRuleInfo.alarmResourceList" 
                        >   
                            <!-- 实例名称 -->
                            <el-table-column 
                                prop="instanceName" 
                                :label="$t('instance.detail.instanceName')" 
                                min-width="120"
                            >
                                <template v-slot="scope">
                                    <el-tooltip
                                        v-model="scope.row.tooltipStatus['instanceName'].showTooltip"
                                        :disabled="!scope.row.tooltipStatus['instanceName'].showTooltip"
                                        :content= scope.row.instanceName
                                    >
                                        <div
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['instanceName'], scope.row.instanceName, 1.17, 'list')"
                                        >
                                            <span>{{scope.row.instanceName || $filter.emptyFilter()}}</span>
                                        </div>
                                    </el-tooltip>
                                </template>
                            </el-table-column>
                            <!-- 实例ID -->
                            <el-table-column 
                                prop="instanceId" 
                                :label="$t('instance.detail.instanceId')" 
                                min-width="200"
                            >
                                <template v-slot="scope">
                                    <el-tooltip
                                        v-model="scope.row.tooltipStatus['instanceId'].showTooltip"
                                        :disabled="!scope.row.tooltipStatus['instanceId'].showTooltip"
                                        :content= scope.row.instanceId
                                    >
                                        <div
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['instanceId'], scope.row.instanceId, 1.17, 'list')"
                                        >
                                            <span>{{scope.row.instanceId || $filter.emptyFilter()}}</span>
                                        </div>
                                    </el-tooltip>
                                </template>
                            </el-table-column>
                            <!-- 内网IPv4 -->
                            <el-table-column 
                                prop="privateIpv4" 
                                :label="$t('instance.list.intranetIpv4')" 
                                min-width="100"
                            >
                                <template v-slot="scope">
                                    <el-tooltip
                                        v-model="scope.row.tooltipStatus['privateIpv4'].showTooltip"
                                        :disabled="!scope.row.tooltipStatus['privateIpv4'].showTooltip"
                                        :content= scope.row.privateIpv4
                                    >
                                        <div
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['privateIpv4'], scope.row.privateIpv4, 1.17, 'list')"
                                        >
                                            <span>{{scope.row.privateIpv4 || $filter.emptyFilter()}}</span>
                                        </div>
                                    </el-tooltip>
                                </template>
                            </el-table-column>

                            <!-- 内网IPv6 -->
                            <el-table-column 
                                prop="privateIpv6" 
                                :label="$t('instance.list.intranetIpv6')" 
                                min-width="100"
                            >
                                <template v-slot="scope">
                                    <el-tooltip
                                        v-model="scope.row.tooltipStatus['privateIpv6'].showTooltip"
                                        :disabled="!scope.row.tooltipStatus['privateIpv6'].showTooltip"
                                        :content= scope.row.privateIpv6
                                    >
                                        <div
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['privateIpv6'], scope.row.privateIpv6, 1.17, 'list')"
                                        >
                                            <span>{{scope.row.privateIpv6 || $filter.emptyFilter()}}</span>
                                        </div>
                                    </el-tooltip>
                                </template>
                            </el-table-column>
                            <template #empty>
                                <div>{{$t('instance.create.noData')}}</div>
                            </template>
                        </el-table>
                    </div>
                    <el-divider />
                    <!-- 触发条件 -->
                    <h3 class="section-title">{{ $t('alarm.alarmRule.triggeringCondition') }}</h3>
                    <div class="triggering-condition-detail">
                        <span v-for="(item, index) in alarmRuleInfo.detail.triggerDescription" :key="index">
                            <el-tag><template v-if="index !== 0">{{$t('alarm.alarmRule.or')}}</template>
                            {{ item.item || $filter.emptyFilter() }}</el-tag><br>
                        </span>
                    </div>
                    <el-divider />
                    <!-- 通知策略 -->
                    <h3 class="section-title">{{ $t('alarm.alarmRule.notificationPolicy') }}</h3>
                    <div class="ml22">
                        <el-form 
                            v-loading="isLoading"
                            @submit.stop.prevent>                
                            <el-row :gutter="20">
                                <!-- 通知周期 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.noticePeriod') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['noticePeriod'].showTooltip"   
                                            :content= alarmRuleInfo.detail?.noticePeriod
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['noticePeriod'], alarmRuleInfo.detail?.noticePeriod, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail?.noticePeriod || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 有效时段 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.validPeriod') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['validPeriod'].showTooltip"   
                                            :content= alarmRuleInfo.detail.validPeriod
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['validPeriod'], alarmRuleInfo.detail.validPeriod, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.validPeriod || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 通知条件 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.noticeCondition') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['noticeCondition'].showTooltip"   
                                            :content= alarmRuleInfo.detail.noticeCondition
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['noticeCondition'], alarmRuleInfo.detail.noticeCondition, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.noticeCondition || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 接受渠道 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.acceptChannel') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['acceptChannel'].showTooltip"   
                                            :content= alarmRuleInfo.detail.acceptChannel
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['acceptChannel'], alarmRuleInfo.detail.acceptChannel, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.acceptChannel || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 通知对象 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.alarmRule.noticeObject') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['noticeObject'].showTooltip"   
                                            :content= alarmRuleInfo.detail.noticeOption?.userName
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['noticeObject'], alarmRuleInfo.detail.noticeOption?.userName, 259, 'detail')"
                                            >
                                                {{ alarmRuleInfo.detail.noticeOption?.userName || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                        </el-form>
                    </div> 
                </el-tab-pane>
            </el-tabs>
            <div v-if="disable">
                <alarm-rule-disable
                    :openVisible="disable"
                    :instanceInfo="alarmRuleInfo.detail"
                    :project-id="route.query.projectId"
                    @close="disableCancel"
                    @refresh="refreshData"
                >
                </alarm-rule-disable> 
            </div>
            <div v-if="enable">
                <alarm-rule-enable
                    :openVisible="enable"
                    :instanceInfo="alarmRuleInfo.detail"
                    :project-id="route.query.projectId"
                    @close="enableCancel"
                    @refresh="refreshData"
                >
                </alarm-rule-enable> 
            </div>
            <div v-if="history">
                <alarm-rule-history
                    :openVisible="history"
                    :instanceInfo="alarmRuleInfo.detail"
                    :project-id="route.query.projectId"
                    @close="historyCancel"
                    @refresh="refreshData"
                >
                </alarm-rule-history>
            </div>
            <div v-if="deleteRule">
                <alarm-rule-delete
                    :openVisible="deleteRule"
                    :instanceInfo="alarmRuleInfo.detail"
                    :project-id="route.query.projectId"
                    @close="deleteCancel"
                    @refresh="refreshData"
                >
                </alarm-rule-delete>
            </div>                       
        </div>       
    </div>
</template>

<script setup lang="ts">
import { 
    ref,
    Ref,
    reactive,
    onMounted, 
    onUnmounted,
    onBeforeUnmount, 
    getCurrentInstance,
    nextTick,
    watch
} from 'vue';
import { useRouter, useRoute } from 'vue-router';
import {
    ArrowLeft,
} from '@element-plus/icons-vue';
import VueCookies from 'vue-cookies';
import {
    alarmRuleDetailAPI, 
} from 'api/request.ts';
import allProjectStore from '@/store/modules/headerDetail.ts';
import { ElMessage, ElTable } from 'element-plus';
import {
    hasShowTooltip, // 是否显示提示
    findValueByKey,
    findValuesByKey,
    renderItemsWithBreaks
} from 'utils/index.ts';
import {
    dimentionState,
    resourceData, // 报警资源类型
    metricInstanceData,
    metricDiskData,
    metricMountpointData,
    metricNetworkData,
    periodData,
    calculationData,
    operationData, 
    calculationUnitData,  
    timesData,
    noticeLevelData,
    noticePeriodData,
    noticeConditionData,
    noticeWayData,
    dimensionData
} from 'utils/constants.ts'; 
import useClipboard from 'vue-clipboard3';
import {

} from 'utils/constants.ts'; 
import i18n from 'lib/i18n/index.ts'; // 国际化
import publicIndexStore from 'store/index.ts'; // 公共store
import useProjectStore from '@/store/modules/projectId.ts';
import dropDownOperation from 'components/DropdownOperation/DropdownOperation.vue';
import alarmRuleDisable from 'components/AlarmRuleOperate/alarmRuleDisable.vue';
import alarmRuleEnable from 'components/AlarmRuleOperate/alarmRuleEnable.vue';
import alarmRuleHistory from 'components/AlarmRuleOperate/alarmRuleHistory.vue';
import alarmRuleDelete from 'components/AlarmRuleOperate/alarmRuleDelete.vue';

/**
 * 国际化
*/
const {global} = i18n;

/**
 * 路由带值
 */
const route = useRoute();

/**
 * cookie ts规范校验
*/
const cookie: {
    [x: string]: unknown;
    get?: Function;
    set?: Function;
} = VueCookies;

/**
 * store库存储的项目列表
*/
const store = allProjectStore();

/**
 * 设置类
*/
type SetType<T> = T extends {} ? any : T;

/**
 * 使用mitt传值
*/
const instanceMitt: Exclude<Required<SetType<{}>> | null, never> = getCurrentInstance();

/**
 * 实例tab跳转标志
 */
 let tabName: Ref<any> = ref('');

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
 * 返回列表页
 */
const router = useRouter();
const projectStore = useProjectStore();
const goBack = () => {
    projectStore.setProject(route.query.projectId, route.query.projectName);
    router.push({
		path: `/instance/rules/list`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName
        }
	});
}

/**
 * loadig态
*/
const isLoading: Ref<boolean> = ref<boolean>(true); 

/**
 * 操作下拉状态
 */
const selectStatus = ref(false);
const selectHover = (val: boolean) => {
    selectStatus.value = val;
};

interface ReactiveType {
    detail: any;
    alarmResourceList: any;
}

/**
 * 详情页信息
 */
const alarmRuleInfo: ReactiveType = reactive<ReactiveType>({
    detail: {},
    alarmResourceList:[],
})

/**
 * tip提示信息
 */
interface TipStatusType {
        [x: string]: {
            showTooltip: boolean;
        }
};

/**
 * tip提示信息显示状态
 */
const tooltipStatus: TipStatusType = reactive<TipStatusType>({
    'ruleNameInfo': {
        showTooltip: false
    },
    'ruleName': {
        showTooltip: false
    },
    'ruleId': {
        showTooltip: false
    },
    'resourceType': {
        showTooltip: false
    },
    'dimension': {
        showTooltip: false
    },
    'noticePeriod': {
        showTooltip: false
    },
    'validPeriod': {
        showTooltip: false
    },
    'noticeCondition': {
        showTooltip: false
    },
    'acceptChannel': {
        showTooltip: false
    },
    'noticeObject': {
        showTooltip: false
    },
})

/**
 *  实例ID复制操作
 * @param value 
 */
const {toClipboard} = useClipboard();
const clickOperateCopy: (value: any) => void = (value: any) => {
    toClipboard(value);  
    ElMessage({
        message: global.t('personCentre.content.copySuccess'),
        type: 'success'
    })
}
    
/**
 * 请求实例详情数据接口
*/
const requestAlarmRuleDetailData = (): void => {
    alarmRuleDetailAPI({
        ruleId: route.query.ruleId,
        projectId: route.query.projectId
    }).then(({data} : {data: any}) => {
        if (data?.result) {
            alarmRuleInfo.detail = data.result;
            alarmRuleInfo.alarmResourceList = data.result.instances.map((item:any) => {
                return {
                    ...item,
                    tooltipStatus: {
                        'instanceName': {
                            showTooltip: false
                        },
                        'instanceId': {
                            showTooltip: false
                        },
                        'privateIpv4': {
                            showTooltip: false
                        },
                        'privateIpv6': {
                            showTooltip: false
                        },

                    }
                }
            });
            alarmRuleInfo.detail.dimension = findValueByKey( dimensionData, 'value', alarmRuleInfo.detail?.dimension, 'text');
            alarmRuleInfo.detail.noticePeriod = findValueByKey( noticePeriodData, 'value', alarmRuleInfo.detail?.noticeOption?.noticePeriod, 'text');
            alarmRuleInfo.detail.validPeriod = alarmRuleInfo.detail?.noticeOption?.effectiveIntervalStart + ' - ' + alarmRuleInfo.detail?.noticeOption?.effectiveIntervalEnd;
            alarmRuleInfo.detail.noticeCondition = findValuesByKey( noticeConditionData, 'value', alarmRuleInfo.detail?.noticeOption?.noticeCondition, 'text').join(', ');
            alarmRuleInfo.detail.acceptChannel = findValuesByKey( noticeWayData, 'value', alarmRuleInfo.detail?.noticeOption?.noticeWay, 'text').join(', ');
            alarmRuleInfo.detail.triggerDescription = renderItemsWithBreaks(alarmRuleInfo.detail.triggerDescription)
            return;
        }
        return Promise.reject();
    }).catch(({message} : {message: string;}) => {      
        if (message === '找不到对象' || message === 'record not found') { 
            goBack();
            return;
        }
        if (message.indexOf('undefined') > -1) return;
    }).finally(() => {
        store.requestObject();
        isLoading.value = false; 
        (instanceMitt as any).proxy.$Bus.emit('project-name', route.query.projectName);    
    });
};



/**
 * 数据卷item类型
 */
 type DataType =  {
    volumeName: string,
    volumeAmount: 0,
    deviceTypeId: string,
    diskType: string,
    raidModelChoose: {}
}

/**
 * 数据卷类型
 */
interface data {
    tableData: DataType[]
}

/**
 * 系统盘数据
 */
const datas: data = reactive<data>({
    tableData: []
})

/**
 * 改变标签页动作
 */
const changeTab = (): void => {
    
}


onMounted(() => {
    tabName.value = 'basicInfo';
    requestAlarmRuleDetailData();
});


/**
 * 禁用操作打开标志
 */
 const disable: Ref<boolean> = ref<boolean>(false)


/**
 * 禁用操作
 * @param value 
 */
 const clickOperateDisable: () => void = () => {
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
 const clickOperateEnable: () => void = () => {
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

const clickOperateREdit: () => void = () => {
    router.push({
        path: `/instance/rules/create`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName,
            edit: route.query.ruleId
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
 const clickOperateHistory: () => void = () => {
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
 const clickOperateDelete: () => void = () => {
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
 * 刷新接口
*/
const refreshData = (): void => {
    isLoading.value = true;
    requestAlarmRuleDetailData();
};



</script>

<style lang="scss">
@import 'assets/css/page';
@import './alarmRuleDetail.scss';
@import 'assets/css/icon';
@import 'assets/css/table';
</style>