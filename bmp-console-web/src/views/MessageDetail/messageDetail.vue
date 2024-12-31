<template>
    <div  class="page-position can-scroll">        
        <div class="page-content message-detail">
            <div class="header-detail fw600">
                <el-button
                    type= "text"
                    :icon ="ArrowLeft" 
                    @click="goBack"
                >
                </el-button>
                <span class="ml15">{{ $t('messageCentre.detail.messageDetail') }}</span>
                <span class="ml15">/</span>
                <span 
                    class="ml15"
                    v-if="reactiveArr.detail.message_type === $t('messageCentre.detail.operateMessage')"
                >
                    {{'【'+ reactiveArr.detail.message_sub_type + '】' + messageContentConstant[reactiveArr.detail.result] + ' (' + reactiveArr.detail.detail + ')' || $filter.emptyFilter()}}
                </span>
                <span 
                    class="ml15"
                    v-if="reactiveArr.detail.message_type === $t('messageCentre.detail.faultMessage')"
                >
                    {{'【'+ reactiveArr.detail.fault_type + '】' + $t('messageCentre.detail.remindContent') + ' (' + reactiveArr.detail.instance_name +'/' + reactiveArr.detail.sn + ')'
                    || $filter.emptyFilter()}}
                </span>
                <span 
                    class="ml15"
                    v-if="reactiveArr.detail.message_type === $t('messageCentre.detail.alarmMessage')"
                >
                    {{'【'+ reactiveArr.detail.message_sub_type + '】' + $t('messageCentre.detail.alarmRemindContent') + ' (' + reactiveArr.detail.instance_name +'/' + reactiveArr.detail.instance_id + ')'
                    || $filter.emptyFilter()}}
                </span>
            </div>
            <div class="message-content">
                <!-- 消息提示 -->
                <div class="message-tip" v-if="reactiveArr?.detail?.message_type === $t('messageCentre.detail.operateMessage')">
                    {{
                    '【'+ reactiveArr?.detail?.message_sub_type + '】' + messageContentConstant[reactiveArr.detail.result]
                    || $filter.emptyFilter()}}
                </div>
                <div class="message-tip" v-if="reactiveArr?.detail?.message_type === $t('messageCentre.detail.faultMessage')">
                    {{
                    '【'+ reactiveArr?.detail?.fault_type + '】' + $t('messageCentre.detail.remindContent') 
                    || $filter.emptyFilter()}}
                </div>
                <div class="message-tip" v-if="reactiveArr?.detail?.message_type === $t('messageCentre.detail.alarmMessage')">
                    {{
                    '【'+ reactiveArr?.detail?.message_sub_type + '】' + $t('messageCentre.detail.alarmRemindContent') 
                    || $filter.emptyFilter()}}
                </div>
                <!-- 用户问好 -->
                <div class="message-greeting">
                    {{$t('messageCentre.detail.greetingMessage',[store.userForm.userName])}}
                    <span class="ml10" v-if="reactiveArr.detail.message_type === $t('messageCentre.detail.operateMessage')">{{$t('list.tip.statusError')}}</span>
                </div>
                <!-- 操作消息 -->
                <div class="operate-content table-content" v-if="reactiveArr.detail.message_type === $t('messageCentre.detail.operateMessage')">
                    <!-- 以下为操作概览 -->
                    <p class="f12 mb10">{{$t('messageCentre.detail.operateOverview')}}</p>
                    <!-- 操作概览 -->
                    <el-table
                        :data="reactiveArr.instanceTableData"
                        v-loading="isLoading"
                    >
                        <!-- 操作内容 -->
                        <el-table-column     
                            prop="message_sub_type" 
                            :label="$t('messageCentre.detail.operateContent')"
                        >
                            <template v-slot="scope">
                                {{scope.row.message_sub_type || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 操作人 -->
                        <el-table-column     
                            prop="user_name" 
                            :label="$t('messageCentre.detail.operator')"
                        >
                            <template v-slot="scope">
                                {{scope.row.user_name || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 操作时间 -->
                        <el-table-column     
                            prop="finish_time" 
                            :label="$t('messageCentre.detail.operateTime')"
                        >
                            <template v-slot="scope">
                                {{scope.row.finish_time || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 

                    </el-table>
                    <!-- 以下为操作内容详情 -->
                    <p class="f12 mb10">{{$t('messageCentre.detail.operateDetail')}}</p>
                    <!-- 操作内容详情 -->
                    <el-table
                        :data="reactiveArr.instanceTableData"
                        v-loading="isLoading"
                    >
                        <!-- 实例名称 -->
                        <el-table-column     
                            prop="instance_name" 
                            :label="$t('messageCentre.detail.instanceName')"
                        >
                            <template v-slot="scope">
                                {{scope.row.instance_name || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 实例ID -->
                        <el-table-column     
                            prop="instance_id" 
                            :label="$t('messageCentre.detail.instanceId')"
                        >
                            <template v-slot="scope">
                                {{scope.row.instance_id || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                    </el-table>
                </div>
                <!-- 故障消息 -->
                <div class="operate-content table-content" v-if="reactiveArr.detail.message_type === $t('messageCentre.detail.faultMessage')">
                    <!-- 以下为故障详情 -->
                    <p class="f12 mb10">{{$t('messageCentre.detail.faultDetail')}}</p>
                    <!-- 故障详情 -->
                    <el-table
                        :data="reactiveArr.instanceTableData"
                        v-loading="isLoading"
                    >
                        <!-- 序号 -->
                        <el-table-column     
                            prop="logid" 
                            :label="$t('messageCentre.detail.number')"
                        >
                            <template v-slot="scope">
                                {{scope.row.logid || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 故障类型 -->
                        <el-table-column     
                            prop="fault_type" 
                            :label="$t('messageCentre.detail.faultType')"
                        >
                            <template v-slot="scope">
                                {{scope.row.fault_type || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 故障描述 -->
                        <el-table-column     
                            prop="content" 
                            :label="$t('messageCentre.detail.faultDes')"
                            min-width="180"
                        >
                            <template v-slot="scope">
                                {{scope.row.content || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 故障报警时间 -->
                        <el-table-column     
                            prop="alert_time" 
                            :label="$t('messageCentre.detail.faultTime')"
                        >
                            <template v-slot="scope">
                                {{scope.row.alert_time || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 故障次数 -->
                        <el-table-column     
                            prop="alert_count" 
                            :label="$t('messageCentre.detail.faultTimes')"
                        >
                            <template v-slot="scope">
                                {{scope.row.alert_count || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                    </el-table>
                </div>
                <!-- 报警信息 -->
                <div class="f12 mt20" v-if="reactiveArr.detail.message_type === $t('messageCentre.detail.alarmMessage')">
                    {{ '【'+ $t('alarm.monitor.alarmMonitor')  + '】：' + reactiveArr.detail.idc_name + $t('alarm.message.instance')}}
                    <span class="point-click" @click="jumpToMoitor()">{{reactiveArr.detail.instance_name +'/'+ reactiveArr.detail.instance_id}}</span>
                    <span>{{$t('alarm.message.current')+ reactiveArr.detail.message_sub_type + '-' + reactiveArr.detail.detail }}</span>
                    <span>{{$t('alarm.message.rule', [reactiveArr.detail.rule_name]) + $t('alarm.message.check')}}</span>

                </div>
            </div>
            <!-- 上一条 下一条 -->
            <div class="message-detail-button">
                <el-button 
                    @click="jumpToLast()"
                    :disabled="!reactiveArr.prevMessageId"
                >
                    {{$t('messageCentre.detail.lastStep')}}
                </el-button>
                <el-button 
                    @click="jumpToNext()"
                    :disabled="!reactiveArr.nextMessageId"
                >
                    {{$t('messageCentre.detail.nextStep')}}
                </el-button>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { 
    ref,
    Ref,
    watch,
    reactive,
    onMounted, 
    onUnmounted,
    onBeforeUnmount, 
    getCurrentInstance
} from 'vue';
import { useRouter, useRoute } from 'vue-router';
import {ElMessage} from 'element-plus';
import i18n from 'lib/i18n/index.ts'; // 国际化
import {
    ArrowLeft,
} from '@element-plus/icons-vue';
import allProjectStore from '@/store/modules/headerDetail.ts';
import useProjectStore from '@/store/modules/projectId.ts';
import {
    messageContentConstant, // 消息通知状态
} from 'utils/constants.ts';
import {
    getDateMinutes,
} from 'utils/index.ts';
import {
    messageDetailAPI, 
    messageReadAPI
} from 'api/request.ts';

/**
 * store库存储的项目列表
*/
const store = allProjectStore();

/**
 * 国际化
*/
const {global} = i18n;

/**
 * 路由带值
 */
 const route = useRoute();

/**
 * 返回列表页
 */
const router = useRouter();
const goBack = () => {
    router.push({
		path: `/message/list`,
	});
    localStorage.removeItem('messageInfo');
}

/**
 * 消息操作类型
 */
 type OperateType = {
    message_sub_type: string,
    user_name: string,
    finish_time: string,
}

/**
 * 消息实例类型
 */
 type InstanceType = {
    instance_name: string,
    instance_id: string,
}

/**
 * 消息故障类型
 */
 type FaultType = {
    logid: string,
    fault_type: string,
    content: string,
    alert_time: string,
    alert_countv: string,
}

/**
 * 消息类型
 */
 type tableType = {
    message_type: string,
    user_name: string,
    finish_time: string,
    instance_name: string,
    instance_id: string,
    logid: string,
    fault_type: string,
    content: string,
    alert_count: string,
    alert_time: string,
}

/**
 * 消息类型
 */
 interface message {
    tableData: tableType[],
    operateTableData: OperateType[],
    instanceTableData: InstanceType[],
    faultTableData: FaultType[],
    detail: any,
    nextMessageId: string,
    prevMessageId: string

}

/**
 * 消息详情数据
 */
 const reactiveArr: message = reactive<message>({
    tableData: [],
    operateTableData: [],
    instanceTableData: [],
    faultTableData: [],
    detail: {},
    nextMessageId: '',
    prevMessageId: ''


})

/**
 * 请求已读接口
*/
const requestRead = (): void => {
    messageReadAPI({
        messageIds: [msgId.value]
    }).then(({data}: any) => {

    }).catch(() => {
    }).finally(() => {
        isLoading.value = false;
    })
};
/**
 * loadig态
*/
const isLoading: Ref<boolean> = ref<boolean>(true);
/**
 * 信息ID
 */
const msgId: Ref<string> = ref<string>('');
/**
 * 请求消息详情数据接口
*/
const requestInstanceDetailData = async() => {
    const res =  await messageDetailAPI({
        messageId: msgId.value
    }).then(({data} : {data: any}) => {
        isLoading.value = true;
        return data;
    }).catch(({message} : {message: string;}) => {      
        if (message === '找不到对象' || message === 'Not found') { 
            goBack();
            return;
        }
        if (message.indexOf('undefined') > -1) return;

    }).finally(() => {
        requestRead()
    });
    if (res?.result?.message) {
            reactiveArr.detail = res.result.message;
            reactiveArr.detail.finish_time = getDateMinutes(reactiveArr.detail.finish_time)
            reactiveArr.detail.alert_time = getDateMinutes(reactiveArr.detail.alert_time)
            reactiveArr.instanceTableData[0] = res.result.message;
            reactiveArr.prevMessageId = res.result.prevMessageId;
            reactiveArr.nextMessageId = res.result.nextMessageId;
            localStorage.setItem('messageInfo', window.btoa(JSON.stringify(msgId.value)));
            return;
        }

};

onMounted(() => {
    const obj = localStorage?.getItem('messageInfo')?? '';   
    if (!obj?.length) {
        goBack()
        return
    }
    const id = JSON.parse(window.atob(obj));
    msgId.value = id;
    requestInstanceDetailData()
})

onUnmounted(() => localStorage.removeItem('messageInfo'));


/**
 * 跳入上一页
 */
const jumpToLast = () => {
    msgId.value = reactiveArr.prevMessageId;
    requestInstanceDetailData()
}

/**
 * 跳入下一页
 */
 const jumpToNext = () => {
    msgId.value = reactiveArr.nextMessageId;
    requestInstanceDetailData()
}

/**
 * 判断性能监控授权
 */
 const hasInMonitor: Ref<boolean> = ref<boolean>(false);  

/**
 * 返回实例列表页
 */
const projectStore = useProjectStore();
hasInMonitor.value = true; 
const jumpToMoitor = () => {
    if (hasInMonitor.value) {
        projectStore.setProject(reactiveArr.detail.project_id, reactiveArr.detail.project_name);
        router.push({
            path: `/instance/detail`,
            query: {
                instanceId: reactiveArr.detail.instance_id,
                projectId: reactiveArr.detail.project_id,
                projectName: reactiveArr.detail.project_name,
                monitor: 'PerformanceMonitoring'
            }
        });
    } else {
        ElMessage({
                message: global.t('personCentre.content.operateNoModel'),
                type: 'warning'
            })
    }
    
}
</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import './messageDetail.scss';

</style>