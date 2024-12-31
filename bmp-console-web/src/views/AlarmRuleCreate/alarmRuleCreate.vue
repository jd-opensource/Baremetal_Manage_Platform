<template>
    <div class="page-position can-scroll">
        <div class="detail-header">
            <el-button
                class="back-btn"
                type= "text"
                :icon ="ArrowLeft"
                @click="jumpToList"
            >
            </el-button>
            <span class="ml10">{{isEdit ? $t('alarm.alarmRule.editRule') : $t('alarm.alarmRule.addRule') }}</span>
        </div>
        <div class="page-content t120 alarm-rule-create">
            <el-config-provider :locale="locale">
                <div class="alarm-form-content">
                    <el-form
                        label-position="left"
                        :model="alarmRuleForm"
                        scroll-to-error
                        ref="ruleFormRef"
                        :rules="rules"
                        label-width="150px"
                        @submit.native.prevent
                    >
                        <!-- 基本信息 -->
                        <h3 class="section-title">{{ $t('personCentre.form.basicInformation')}}</h3>
                        <div class="ml20">
                            <!-- 规则名称 -->
                            <el-form-item 
                                prop="ruleName"
                                :label="$t('alarm.alarmRule.ruleName') "
                            >
                                <el-input 
                                    v-model="alarmRuleForm.ruleName" 
                                    maxlength="128"
                                    :placeholder="$t('alarm.placeHolder.ruleName')"
                                />                           
                            </el-form-item>  
                            <!-- 资源类型 -->
                            <el-form-item 
                                prop="resource"
                                required
                                :label="$t('alarm.alarmRule.resourceType') "
                            >
                                <el-select
                                    v-model="alarmRuleForm.resource" 
                                    :disabled= 'isEdit'
                                >
                                    <el-option 
                                        v-for="item in resourceData" 
                                        :key="item.value"
                                        :label="item.text" 
                                        :value="item.value"
                                    /> 
                                </el-select>                        
                            </el-form-item> 
                            <!-- 维度 -->
                            <el-form-item 
                                prop="dimension"
                                required
                                :label="$t('alarm.alarmRule.dimension') "
                            >
                                <el-radio-group v-model="alarmRuleForm.dimension" @change="dimensionChange" :disabled='isEdit'>
                                    <el-radio value="instance" label="instance">{{$t('alarm.alarmRule.instance')}} </el-radio>
                                    <el-radio value="disk" label="disk">{{$t('alarm.alarmRule.disk')}}</el-radio>
                                    <el-radio value="mountpoint" label="mountpoint">{{$t('alarm.alarmRule.mountpoint')}}</el-radio>
                                    <el-radio value="nic" label="nic">{{$t('alarm.alarmRule.nic')}}</el-radio>
                                </el-radio-group>
                            </el-form-item> 
                            <!-- 报警资源 -->
                            <el-form-item 
                                prop="instanceNames"
                                :label="$t('alarm.alarmRule.alarmResource') "
                            >
                                <el-select
                                    multiple
                                    collapse-tags
                                    :max-collapse-tags="3"
                                    v-model="alarmRuleForm.instanceNames" 
                                    class="choose-instance-name"
                                    :placeholder="$t('alarm.placeHolder.chooseInstance')"
                                >
                                    <template #empty>
                                        <div  class="instance-content" >
                                            <div class="instance-message-list-input">
                                                <el-input
                                                    v-model="instanceNameSearch"
                                                    :placeholder="$t('list.placeholder.instanceName')"
                                                />
                                            </div>
                                            <div class="table-content m0">
                                                <el-table 
                                                    ref="instanceTableRef"
                                                    :data="tableData" 
                                                    max-height="300"
                                                    @selection-change="handleSelectionChange"
                                                    >
                                                    <!-- 选择多行 -->
                                                    <el-table-column type="selection" width="55" />
                                                    <el-table-column prop="instanceName" :label="$t('instance.detail.instanceName')" >
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
                                                    <el-table-column prop="instanceId" :label="$t('instance.detail.instanceId')" >
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
                                                    <el-table-column prop="privateIpv4" :label="$t('instance.list.intranetIpv4')" >
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
                                                </el-table>
                                            </div>
                                        </div>
                                    </template>     
                                </el-select>              
                            </el-form-item>
                            <!-- 维度下拉 deviceTag -->
                            <el-form-item 
                                v-if="alarmRuleForm.dimension != 'instance'"
                                prop="deviceTag"
                                :label="dimentionState[alarmRuleForm.dimension]"
                            >
                                <el-select
                                    v-if="alarmRuleForm.instanceIds.length === 1"
                                    v-model="alarmRuleForm.deviceTag" 
                                >
                                    <el-option 
                                        v-for="item in deviceTags" 
                                        :key="item.value"
                                        :label="item.text" 
                                        :value="item.value"
                                    /> 
                                </el-select>  
                                <p v-else class="f12">{{$t('alarm.alarmRule.allDevice')}}</p>                         
                            </el-form-item> 
                        </div>
                        <!-- 触发条件 -->
                        <h3 class="section-title">{{ $t('alarm.alarmRule.triggeringCondition')}}</h3> 
                        <div class="ml20">
                            <div v-for="(card, index) in alarmRuleForm.triggerOption" :key="card.id">
                                <el-card class="trigger-card" shadow="never" >
                                    <div v-if="alarmRuleForm.triggerOption.length > 1">
                                        <el-button     
                                            circle 
                                            class="trigger-close-button" 
                                            @click="removeCondition(index)" 
                                        >
                                        <el-icon size="16"><CircleCloseFilled /></el-icon>
                                        </el-button>
                                    </div>
                                    <!-- 触发条件 -->
                                    <el-form-item 
                                        :prop="'triggerOption.' + index"
                                        required
                                        :label="$t('alarm.alarmRule.triggeringCondition')"
                                        :rules="[{ validator: triggerOptionCheck(index), trigger: 'blur' }]"
                                    >   
                                        <div class="trigger-condition-content">
                                            <el-select v-model="card.metric" class="long150" @change="triggerOptionItemChange(index)">
                                                <el-option 
                                                    v-for="item in metricData" 
                                                    :key="item.value"
                                                    :label="item.text" 
                                                    :value="item.value"
                                                /> 
                                            </el-select>
                                            <el-select v-model="card.period" class="long100" @change="triggerOptionItemChange(index)">
                                                <el-option 
                                                    v-for="item in periodData" 
                                                    :key="item.value"
                                                    :label="item.text" 
                                                    :value="item.value"
                                                /> 
                                            </el-select>
                                            <el-select v-model="card.calculation" class="long100" @change="triggerOptionItemChange(index)">
                                                <el-option 
                                                    v-for="item in calculationData" 
                                                    :key="item.value"
                                                    :label="item.text" 
                                                    :value="item.value"
                                                /> 
                                            </el-select>
                                            <el-select v-model="card.operation" class="long90" @change="triggerOptionItemChange(index)">
                                                <el-option 
                                                    v-for="item in operationData" 
                                                    :key="item.value"
                                                    :label="item.text" 
                                                    :value="item.value"
                                                />
                                            </el-select>
                                            <div :class="getLocationItem === 'zh_CN' ? 'threshold-content' : 'threshold-en-content'">
                                                <el-form-item 
                                                    :prop="`triggerOption.${index}.threshold`"
                                                    :rules="getUnitsById(card.metric).text === '%' ? triggerOptionItemPercentRules  : triggerOptionItemRules"
                                                >        
                                                    <el-input v-model="card.threshold" :placeholder="$t('alarm.placeHolder.threshold')" @change="triggerOptionItemChange(index)"></el-input>     
                                                </el-form-item>
                                            </div>
                                            <span class="calculation-unit">{{getUnitsById(card.metric).text}}</span>
                                            <el-select v-model="card.times" class="long150" @change="triggerOptionItemChange(index)">
                                                <el-option 
                                                    v-for="item in timesData" 
                                                    :key="item.value"
                                                    :label="item.text" 
                                                    :value="item.value"
                                                /> 
                                            </el-select>
                                        </div>
                                    </el-form-item>
                                    <!-- 报警级别 -->
                                    <el-form-item     
                                        required
                                        :label="$t('alarm.alarmRule.alarmLevel') "
                                    >
                                        <div class="trigger-condition-content">
                                            <el-select v-model="card.noticeLevel" :placeholder="$t('personCentre.placeholder.select')" class="long150">
                                                <el-option 
                                                    v-for="item in noticeLevelData" 
                                                    :key="item.value"
                                                    :label="item.text" 
                                                    :value="item.value"
                                                /> 
                                            </el-select>
                                        </div>
                                    </el-form-item>
                                    
                                </el-card>
                            </div>
                            <div class="trigger-button">
                                <el-button 
                                    type="text" 
                                    @click="addCondition" 
                                    class="trigger-button-content">
                                    + &nbsp;&nbsp;{{$t('alarm.alarmRule.addTriggeringCondition')}}
                                </el-button>
                            </div>
                        </div>
                        <!-- 通知策略 -->
                        <h3 class="section-title">{{ $t('alarm.alarmRule.notificationPolicy')}}</h3> 
                        <div class="ml20">
                            <!-- 通知周期 -->
                            <el-form-item 
                                prop="noticeOption.noticePeriod"
                                class="mb42"
                                required
                                :label="$t('alarm.alarmRule.noticePeriod')"
                            >
                                <el-select
                                    v-model="alarmRuleForm.noticeOption.noticePeriod" 
                                >
                                    <el-option 
                                        v-for="item in noticePeriodData" 
                                        :key="item.value"
                                        :label="item.text" 
                                        :value="item.value"
                                    /> 
                                </el-select>  
                                <div 
                                    class="notice-period-tip el-form-item__error">
                                    {{$t('alarm.tip.noticePeriodTip')}}
                                </div>                      
                            </el-form-item> 
                            <!-- 有效时段 -->
                            <el-form-item 
                                prop="validPeriod"
                                class="no-requrie-label mb42"
                                :label="$t('alarm.alarmRule.validPeriod') "
                            >
                                <el-time-picker
                                    v-model="validTime"
                                    is-range
                                    range-separator="-"
                                    start-placeholder="Start time"
                                    end-placeholder="End time"
                                    /> 
                                    <div 
                                        class="notice-period-tip el-form-item__error">
                                        {{$t('alarm.tip.validPeriodTip')}}
                                    </div>                     
                            </el-form-item> 
                            <!-- 通知条件 -->
                            <el-form-item 
                                prop="noticeOption.noticeCondition"
                                :label="$t('alarm.alarmRule.noticeCondition') "
                            >   
                                <el-checkbox-group v-model="alarmRuleForm.noticeOption.noticeCondition">
                                    <el-checkbox :label= 1 :value= 1 >{{$t('alarm.alarmRule.alarm')}}</el-checkbox>
                                    <el-checkbox :label= 2 :value= 2  >{{$t('alarm.alarmRule.beNormal')}}</el-checkbox>
                                </el-checkbox-group>                        
                            </el-form-item> 
                            <!-- 接受渠道 -->
                            <el-form-item 
                                prop="noticeOption.noticeWay"
                                :label="$t('alarm.alarmRule.acceptChannel') "
                            >   
                                <el-checkbox-group v-model="alarmRuleForm.noticeOption.noticeWay">
                                    <el-checkbox :label= 1 :value= 1 >{{$t('alarm.alarmRule.inEmail')}}</el-checkbox>
                                    <el-checkbox :label= 2 :value= 2  >{{$t('alarm.alarmRule.email')}}</el-checkbox>
                                </el-checkbox-group>                        
                            </el-form-item> 
                            <!-- 通知对象-->
                            <el-form-item 
                                prop="noticeOption.userId"
                                required
                                :label="$t('alarm.alarmRule.noticeObject') "
                            >   
                                <el-radio-group v-model="contactValue">
                                    <el-radio :value="1" :label="1">{{$t('alarm.alarmRule.currentContacts')}}</el-radio>
                                </el-radio-group>                          
                            </el-form-item> 
                            <div class="table-content contact-content">
                                <el-table 
                                    ref="contactTableRef"
                                    :data="reactiveUserArr.tableData"
                                    width="775" 
                                    max-height="400"
                                    >
                                    <el-table-column prop="userName" :label="$t('personCentre.form.user')" />
                                    <el-table-column prop="phoneNumber" :label="$t('personCentre.form.phoneNumber')" />
                                    <el-table-column prop="email" :label="$t('personCentre.form.email')" />
                                </el-table> 
                                <div class="create-key-tip">
                                    <span> {{$t('alarm.tip.contact1')}}</span>
                                    <a 
                                        class="mouse-point"
                                        @click="jumpToUser"
                                        target="_blank">
                                        {{$t('alarm.tip.myUser')}}
                                    </a> 
                                    <span> {{$t('alarm.tip.contact2')}}</span>
                                </div> 
                            </div>
                        </div>
                        <div class="create-confirm mt50">
                            <el-button 
                                type="primary" 
                                :loading="isLoading"
                                class="create-confirm-button" 
                                @click="addRule"
                            >
                                {{$t('list.confirm')}}
                            </el-button>
                            <el-button 
                                class="create-confirm-button cancel-button" 
                                @click="jumpToList"
                            >
                                {{$t('instance.create.cancel')}}
                            </el-button>
                        </div>
                    </el-form>
                </div>
            </el-config-provider>
        </div>
    </div>
</template>

<script setup lang="ts">
import { useRouter, useRoute} from 'vue-router';
import {
  ArrowLeft,
  ArrowDown,
    ArrowUp,
    Search
} from '@element-plus/icons-vue';
import { 
    reactive,
    ref,
    Ref,
    toRefs,
    onUnmounted,
    onMounted,
    computed,
    nextTick,
    watch,
    getCurrentInstance
} from 'vue';
import type {
    FormInstance, // element-plus ui类
    FormRules, // element-plus ui类
    ElTable
} from 'element-plus';
import {ElMessage} from 'element-plus';
import {
    CircleCloseFilled,
} from '@element-plus/icons-vue';
import { useI18n } from 'vue-i18n'; // 国际化`
import VueCookies from 'vue-cookies'; // cookie
import useProjectStore from '@/store/modules/projectId.ts';
import allProjectStore from '@/store/modules/headerDetail.ts';
import {
    projectNameReg, // 规则名称正则
    ruleThresholdReg,
    ruleThresholdPrecentReg
} from 'utils/regular.ts';
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
    noticePeriodData
} from 'utils/constants.ts'; 
import {
    languageSwitch, // 语言切换
    hasShowTooltip, // 是否显示提示
    createTimeObjectFromDateStringArray,
    addUniqueId,
    removeIdFromItems,
    convertBracketedTimeStringsToDateArray,
    convertArrayToObjectArray,
    removeObjectKey,
} from 'utils/index.ts';
import publicIndexStore from 'store/index.ts'; // 公共store
import { instanceListAPI, userAPI, addRuleAPI, editRuleAPI, deviceTagAPI, alarmRuleDetailAPI } from 'api/request.ts'; // 实例列表数据接口
import maxCollapseTagsSelect from 'components/common/maxCollapseTagsSelect/maxCollapseTagsSelect.vue'


/**
 * 路由
 */
 const route = useRoute();
 const router = useRouter();

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: any = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * 语言
*/
const locale: any = languageSwitch();

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 是否为编辑
*/
const isEdit: Ref<boolean> = ref<boolean>(false);

/**
 * 实例ref
*/
const instanceTableRef: Ref<any> = ref<FormInstance>();

onMounted(() => {

})

/**
 * 请求实例详情数据接口
*/
const requestAlarmRuleDetailData = (): void => {
    isEdit.value = true;
    alarmRuleDetailAPI({
    ruleId: route.query.edit,
    projectId: route.query.projectId
    }).then(({data} : {data: any}) => {
        if (data?.result) { 
            alarmRuleForm.ruleName = data?.result.ruleName;
            alarmRuleForm.dimension = data?.result.dimension;
            alarmRuleForm.resource = data?.result.resource;
            alarmRuleForm.deviceTag = data?.result.deviceTag;
            alarmRuleForm.noticeOption.noticePeriod = data?.result.noticeOption.noticePeriod;
            alarmRuleForm.noticeOption.noticeCondition = data?.result.noticeOption.noticeCondition;
            alarmRuleForm.noticeOption.noticeWay = data?.result.noticeOption.noticeWay;
            reactiveInstanceArr.tableData.forEach(obj => {
                if (data.result.instanceIds.includes(obj.instanceId)) {
                    instanceTableRef.value!.toggleRowSelection(obj, true)
                }
             });
             alarmRuleForm.triggerOption = addUniqueId(data?.result?.triggerOption);
             validTime.value = convertBracketedTimeStringsToDateArray(data?.result?.noticeOption?.effectiveIntervalStart, data?.result?.noticeOption?.effectiveIntervalEnd)
             
            return;
        }
        return Promise.reject();
    }).catch(({message} : {message: string;}) => {      
        if (message === '找不到对象' || message === 'Not found') { 
            return;
        }
        if (message.indexOf('undefined') > -1) return;
    }).finally(() => {
        store.requestObject(); 
        (instanceMitt as any).proxy.$Bus.emit('project-name', route.query.projectName);    
    });
};

/**
 * 跳转到列表页
 */
const projectStore = useProjectStore();
const jumpToList = () => {
    ruleFormRef.value.resetFields(); 
    projectStore.setProject(route.query.projectId, route.query.projectName);
    router.push({
        path: `/instance/rules/list`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName
        }
    });
};

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
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 监控指标
 */
const metricData: Ref<any> = ref<any>(metricInstanceData)

/**
 * 根据指标获取单位
 * @param {string} id 监控指标
 */
const getUnitsById = (id: any | null) => {
    const item = calculationUnitData.find((d:any)=> d.id === id);
    return item ? { value: item.value, text: item.text } : null;
}

/**
 * 通知策略
 */
interface NoticeOption {
    noticePeriod: number; // 通知周期(分钟) [5 10 15 30 60 180 360 720 1440]
    effectiveIntervalEnd: string; // 有效时段end
    effectiveIntervalStart: string; //又是时段start
    noticeCondition: number[]; // 通知条件,可多选 [1表示报警， 2表示恢复正常]  
    noticeWay: number[]; // 接收渠道，可多选 [1表示站内信， 2表示邮件]
    userId: string; // 通知对象
}

/**
 * 触发条件
 */
interface TriggerOption {
    id: number;
    metric: string; // 监控指标 cps.cpu.util
    calculation: string; // 计算方式 [min max avg sum]
    calculationUnit: string; // 计算结果单位 [对于使用量，有Bytes,KB,MB,GB,TB，对于使用率，是%，对于连接个数，是count，对于网络包量，是pps,Kpps,Mpps,Gpps,Tpps，对于网络速率，是bps,Kbps,Mbps,Gbps,Tbps 对于负载，没有单位]  
    noticeLevel: number; // 告警级别 [1表示一般，2表示严重，3表示紧急]
    operation: string; // 比较方式 [> >= < <= == !=]或者[gt gte lt lte eq neq]
    period: number; // 周期(分钟) [1, 2, 5, 15, 30, 60]
    threshold: string; // 阈值
    times: number; // 持续周期数 [1, 2, 3, 4, 5, 10, 15, 30, 60]
}

/**
 * 报警规则type
 */
interface Alarm {
    ruleName: string; // 规则名称
    resource: string; // 资源类型
    dimension: string; // 维度 [instance、disk、mountpoint、nic]
    deviceTag: string; // 盘符、挂载点、网口列表
    instanceIds: any; //报警资源
    instanceNames: any //报警资源名称
    ruleId: any;
    triggerOption: TriggerOption[]; // 触发条件
    noticeOption: NoticeOption; // 通知策略
    
}

/**
 * 报警规则表单
 */
const alarmRuleForm: Alarm = reactive({
    deviceTag: "",
    dimension: "instance",
    instanceIds: [],
    instanceNames: [],
    ruleId: '',
    noticeOption: {
        effectiveIntervalEnd: "",
        effectiveIntervalStart: "",
        noticeCondition: [1],
        noticePeriod: 180,
        noticeWay: [1, 2],
        userId: ""
    },
    resource: "instance",
    ruleName: "",
    triggerOption: [
        {   
            id: Date.now(),
            calculation: "avg",
            calculationUnit: '',
            metric: metricData.value[0].value,
            noticeLevel: 1,
            operation: "gt",
            period: 5,
            threshold: '',
            times: 1
        }
    ]
})

/**
 * 规则名称正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const ruleNameCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('alarm.tip.emptyRuleName')));
    }
    else if(!projectNameReg.test(value)) {
        callback(new Error(t('personCentre.content.stringChineseVerification')));
    }else {
        callback();
    }
};

/**
 * 触发条件正则校验
 * @param {Number} currentIndex 触发条件index
*/
const triggerOptionCheck = (currentIndex: number) => {
    return (_: any, value: string, callback: (arg0?: Error | string) => void, source: any,) => {
        const duplicate = alarmRuleForm.triggerOption.some((item, index) => {
        return index !== currentIndex && !!item.threshold &&
        item.metric === alarmRuleForm.triggerOption[currentIndex].metric &&
        item.period === alarmRuleForm.triggerOption[currentIndex].period &&
        item.calculation === alarmRuleForm.triggerOption[currentIndex].calculation &&
        item.operation === alarmRuleForm.triggerOption[currentIndex].operation &&
        item.threshold === alarmRuleForm.triggerOption[currentIndex].threshold &&
        item.times === alarmRuleForm.triggerOption[currentIndex].times;
    });

    if (duplicate) {
        callback(new Error(t('alarm.tip.triggerOptionRepeat')));
    } else {
        callback();
    }
    };
};

/**
 * 阈值正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const ruleThresholdCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('alarm.placeHolder.threshold')));
    }
    else if(!ruleThresholdReg.test(value)) {
        callback(new Error(t('alarm.tip.ruleThresholdVerification')));
    }
    else {
        callback();
    }
};

/**
 * 阈值%正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const ruleThresholdPercentCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('alarm.placeHolder.threshold')));
    }
    else if(!ruleThresholdReg.test(value)) {
        callback(new Error(t('alarm.tip.ruleThresholdVerification')));
    }
    else if(!ruleThresholdPrecentReg.test(value)) {
        callback(new Error(t('alarm.tip.ruleThresholdPercentVerification')));
    }
    else {
        callback();
    }
};

/**
 * 用户信息类
*/
type UserInfoType = {
    required: boolean;
    message: string;
    trigger: string;
    validator?: unknown;
};

/**
 * 规则类
*/
interface RulesType {
    ruleName: UserInfoType[], // 规则名称
    instanceNames: UserInfoType[], // 报警资源
    triggerOption: UserInfoType[], // 触发条件
    'noticeOption.noticeCondition': UserInfoType[], // 通知条件
    'noticeOption.noticeWay': UserInfoType[], // 接受渠道
};

/**
 * 规则
*/
const rules: RulesType = reactive<FormRules>({
    ruleName: [ // 规则名称
        {
            required: true,
            trigger: 'blur',
            validator: ruleNameCheck
        }
    ],
    instanceNames: [ // 报警资源
        {
            required: true,
            trigger: 'change',
            message: t('instance.tip.chooseInstance')
        }
    ],
    triggerOption: [ // 触发条件
        {   
            
            threshold: [ // 触发条件阈值
                { 
                    required: true, 
                    message: 'Please input content', 
                    trigger: 'blur',
                }
            ]
        }
    ],
    'noticeOption.noticeCondition': [ // 通知条件
        {
            required: true,
            trigger: 'change',
            message: t('alarm.tip.chooseNoticeCondition')
        }
    ],
    'noticeOption.noticeWay': [ // 接受渠道
        {
            required: true,
            trigger: 'change',
            message: t('alarm.tip.chooseAcceptChannel')
        }
    ],

});

/**
 * 触发条件阈值规则
*/
const triggerOptionItemRules = [
    { 
        required: true, 
        trigger: 'blur' ,
        validator: ruleThresholdCheck
    }
];

/**
 * 触发条件阈值为率的规则
*/
const triggerOptionItemPercentRules = [
    { 
        required: true, 
        trigger: 'blur' ,
        validator: ruleThresholdPercentCheck
    }
];


/**
 * 实例item类型
 */
 type InstancesType = {
    instanceName: string,
    instanceId: string,
    privateIpv4: string,
}

/**
 * 实例类型
 */
 interface instance {
    tableData: InstancesType[],
}

/**
 * 实例列表数据
 */
 const reactiveInstanceArr: instance = reactive<instance>({
    tableData: [],
})

/**
 * 实例查询key
 */
const instanceNameSearch: Ref<string> = ref<string>(''); 

/**
 * 获取实例列表
 */
 const getinstanceListData = (): void => {
    instanceListAPI({
        isAll:'1',
        isInstallAgent: '1',
        projectId: route.query.projectId ? route.query.projectId : '',
    }).then(({data} : {data: any}) => {
        if (data?.result?.instances?.length) {
            reactiveInstanceArr.tableData = data.result.instances.map((item:any) => {
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

                    }
                }
            });
            return;
        }
        reactiveInstanceArr.tableData = [];
 
    }).catch (()=>{
        reactiveInstanceArr.tableData = [];
    })
    .finally(() => {
        if(route.query.edit) {   
            requestAlarmRuleDetailData()
        }
    });
};
getinstanceListData()


/**
 * 取得tag数据
 * @params -参数
 * @dimensionChange -判断是否为维度改变的判断值
 */
const deviceTags: Ref<any> = ref<any>([]); 
const getDeviceTagData = (params: any, dimensionChange: boolean): void => {
    deviceTagAPI({
        ...params
    }).then(({data} : {data: any}) => {
        if (data?.result?.length) {
            deviceTags.value = convertArrayToObjectArray(data.result, ['text', 'value']);  
            if(!alarmRuleForm.deviceTag || dimensionChange) {
                alarmRuleForm.deviceTag = deviceTags.value[0].value
            }             
            return;
        }
        deviceTags.value = [];
 
    }).catch (()=>{
        deviceTags.value = [];
    })
    .finally(() => {

    });
}

/**
 * 筛选后实例数据
 */
 const tableData = computed(() => {  
    if(instanceNameSearch.value){
        return reactiveInstanceArr.tableData.filter(function(dataNews:any){
            return Object.keys(dataNews).some(function(key){
                if(key === 'instanceName') {
                    return String(dataNews[key]).toLowerCase().indexOf(instanceNameSearch.value) > -1
                }   
            })
        })
    }
    return reactiveInstanceArr.tableData
});

/**
 * 实例选择
 */
const handleSelectionChange = (val: InstancesType[]) => {
    alarmRuleForm.instanceIds = extractInstanceIds(val);
    alarmRuleForm.instanceNames = extractInstanceNames(val)
    if(val.length === 1 && alarmRuleForm.dimension !== 'instance') {
        getDeviceTagData({tagName: alarmRuleForm.dimension, instanceId: val[0].instanceId}, false)
    } else {
        alarmRuleForm.deviceTag = ''
    }

}

/**
 * 根据给定的实例名称数组和实例对象数组，
 * 找到名称匹配的实例对象，并提取它们的 instanceId 组成新的数组。
 * 
 * @param instanceNames - 需要查找的实例名称数组
 * @param instances - 包含实例对象的数组
 * @returns 匹配的实例对象的 instanceId 数组
 */
 const getInstanceIds = (instanceNames: string[], instances: InstancesType[]): string[] => {
    return instances
        // 过滤出名称在 instanceNames 中的实例对象
        .filter(instance => instanceNames.includes(instance.instanceName))
        // 从过滤后的实例对象中提取 instanceId
        .map(instance => instance.instanceId);
};

/**
 * 监听报警资源实例，根据选项来选择实例
 */
watch(() => alarmRuleForm.instanceNames, (): void => {
    alarmRuleForm.instanceIds = getInstanceIds(alarmRuleForm.instanceNames, reactiveInstanceArr.tableData);
    reactiveInstanceArr.tableData.forEach(obj => {
        if (alarmRuleForm.instanceIds.includes(obj.instanceId)) {
            instanceTableRef.value!.toggleRowSelection(obj, true)
        } else {
            instanceTableRef.value!.toggleRowSelection(obj, false)
        }
    });
});

const extractInstanceIds = (instances: InstancesType[]): string[] => {
    return instances.map(instance => instance.instanceId);
};

const extractInstanceNames = (instances: InstancesType[]): string[] => {
    return instances.map(instance => instance.instanceName);
};

/**
 * 触发条件item改变，阈值验证触发，触发条件验证触发
 * @param index 条件标识
 */
const triggerOptionItemChange = async (index: number): Promise<void> => {
    await nextTick();
    ruleFormRef.value.validateField(`triggerOption.${index}.threshold`);
    ruleFormRef.value.validateField(`triggerOption.${index}`);
};

/**
 * 添加触发条件
 */
const addCondition = () => {
    alarmRuleForm.triggerOption.push({
        id: Date.now(),
        calculation: "avg",
        calculationUnit: '',
        metric: metricData.value[0].value,
        noticeLevel: 1,
        operation: "gt",
        period: 5,
        threshold: '',
        times: 1
    });
};

/**
 * 删减触发条件
 */
const removeCondition = (index: number) => {
    alarmRuleForm.triggerOption.splice(index, 1);
};

/**
 * 改变metric触发条件变化
 */
const changeTriggerOption = (metricValue: string) => ({
    id: Date.now(),
    calculation: "avg",
    calculationUnit: '',
    metric: metricValue,
    noticeLevel: 1,
    operation: "gt",
    period: 5,
    threshold: '',
    times: 1
});

/**
 * 维度变化
 * @param val 维度值
 */
 const dimensionChange = (val: string) => {
    if(alarmRuleForm.instanceIds?.length === 1 && val != 'instance') {
        getDeviceTagData({tagName: val, instanceId: alarmRuleForm.instanceIds[0]}, true)
    } else {
        alarmRuleForm.deviceTag = "";
    }
    switch (val) {
        case 'instance':
            metricData.value = metricInstanceData;
            break;
        case 'disk':
            metricData.value = metricDiskData;
            break;
        case 'mountpoint':
            metricData.value = metricMountpointData;
            break;
        case 'nic':
            metricData.value = metricNetworkData;
            break;
        default:
            metricData.value = [];
            alarmRuleForm.triggerOption = []; // 如果需要在 default 情况下清空 triggerOption
            return;
    }

    // 对于非 default 情况，设置 triggerOption
    if (metricData.value.length > 0) {
        alarmRuleForm.triggerOption = [changeTriggerOption(metricData.value[0].value)];
    }

}

/**
 * 有效时段
 */
const validTime = ref<[Date, Date]>([
  new Date(0, 0, 0, 0, 0, 0),
  new Date(0, 0, 0, 23, 59, 59),
])

/**
 * 联系人选项
 */
const contactValue = ref(1) 

/**
 * 联系人item类型
 */
 type UserType = {
    userName: string,
    phoneNumber: string,
    email: string,
}

/**
 * 联系人类型
 */
 interface user {
    tableData: UserType[],
}

/**
 * 联系人列表数据
 */
 const reactiveUserArr: user = reactive<user>({
    tableData: [],
})

/**
 * 请求用户接口,获取projectId
*/
const requestUser = (): void => {
    userAPI({
    }).then(({data} : {data: any}) => {
        if(data?.result){
            alarmRuleForm.noticeOption.userId = data?.result.userId;
            reactiveUserArr.tableData = [data?.result]
        }                  
    }).finally(() => { 
       
    });
};
requestUser()

/**
 * 跳转到我的账户
 */
 const jumpToUser = () => {
    window.open('/user?type=account', '_blank')
};

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();

/**
 * 处理参数
*/
const processingParameter = () => {
    // 根据metric获取单位(calculationUnit)
    alarmRuleForm.triggerOption .map((item: any) => {
        item.calculationUnit = getUnitsById(item.metric)?.value
    })
    // 传递有效时段开始和结束的特定格式
    alarmRuleForm.noticeOption.effectiveIntervalStart = createTimeObjectFromDateStringArray(validTime.value).start;
    alarmRuleForm.noticeOption.effectiveIntervalEnd = createTimeObjectFromDateStringArray(validTime.value).end;
    // 删除触发条件自加的id（id用来区别每个触发条件）
    alarmRuleForm.triggerOption = removeIdFromItems(alarmRuleForm.triggerOption);
    // 将触发条件中的值变为阈值
    alarmRuleForm.triggerOption.map((item: any) => {
        if(item.threshold) {
            item.threshold = parseFloat(item.threshold)
        }  
    })
    if( isEdit.value ) {
        alarmRuleForm.ruleId = route.query.edit;
    } else {
        alarmRuleForm.ruleId = '';
    }
    let params = {
        ...alarmRuleForm
    }
    params  = removeObjectKey(alarmRuleForm, 'instanceNames')
    publicStore.deleteEmtpyData(params);
    if( isEdit.value ) {
        requestEditRule(params)
    } else {
        requestAddRule(params)
    }
    

};
/**
 * 添加报警规则
 */
const requestAddRule = (params: any): void => {
    addRuleAPI({
        ...params,
        projectId: route.query.projectId
    }).then(({data} : {data: any}) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.createSuccess'),
                type:'success'
            });
            jumpToList()
        }
    }).finally(() => {
        isLoading.value = false;
    });
}

/**
 * 编辑报警规则
 */
const requestEditRule = (params: any): void => {
    editRuleAPI({
        ...params,
        projectId: route.query.projectId
    }).then(({data} : {data: any}) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.createSuccess'),
                type:'success'
            });
            jumpToList()
        }
    }).finally(() => {
        isLoading.value = false;
    });
}

/**
 * 点击立即创建
 */
 const addRule = async (): Promise<void> => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        ruleFormRef.value.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                processingParameter()
                isLoading.value = true;
            }
        });
    });
};

</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import 'assets/css/icon';
@import './alarmRuleCreate.scss';
</style>