<template>
    <div class="operate-management-detail all-alarm-rules-detail">
        <!-- 头部信息 -->
        <detail-header
            :name="allAlarmRulesDetailOpt.setName()"
            :title="'allAlarmRulesDetail'"
            @header="router.push($defInfo.routerPath('alarmRulesList'))"
        >
            <ui-dropdown-item>
                <span
                    class="drop-down-operate-content"
                    @click="alaramHistoryOpt.alarmHistoryClick(ruleId)"
                >
                    {{$t('allAlarmRulesDetail.header.operate.alarmHistory')}}
                </span>
            </ui-dropdown-item>
        </detail-header>
        <div
            class="operate-management-detail-info"
            v-loading="allAlarmRulesDetailOpt.isLoading.value"
        >
            <div class="operate-management-count">
                <!-- header信息 基本信息-->
                <div class="info-header">
                    <span class="info-line"></span>
                    <!-- 基本信息 -->
                    <span class="info-title">
                        {{$t('allAlarmRulesDetail.baseInfo.title')}}
                    </span>
                </div>
                <ui-form @submit.stop.prevent>
                    <!-- Layout布局 -->
                    <ui-row
                        class="info-content"
                        :gutter="20"
                    >
                        <!-- 规则ID -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.baseInfo.label.ruleId'))"
                            >
                                <div class="set-wrap info-name">
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt.reactiveArr.detail.ruleId)}}
                                    <!-- <img
                                        alt=""
                                        class="common-img copy-img"
                                        src="@/assets/img/uiImg/copy.png"
                                        @click="methodsTotal.copyInfo(allAlarmRulesDetailOpt.reactiveArr.detail.ruleId)"
                                    /> -->
                                </div>
                            </ui-form-item>
                        </ui-col>
                        <!-- 规则名称 -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.baseInfo.label.ruleName'))"
                            >
                                <div class="set-wrap info-name">
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt.reactiveArr.detail.ruleName)}}
                                    <!-- <img
                                        alt=""
                                        class="common-img copy-img"
                                        src="@/assets/img/uiImg/copy.png"
                                        @click="methodsTotal.copyInfo(allAlarmRulesDetailOpt.reactiveArr.detail.ruleName)"
                                    /> -->
                                </div>
                            </ui-form-item>
                        </ui-col>
                        <!-- 资源类型 -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.baseInfo.label.resourceType'))"
                            >
                                <div class="set-wrap info-name">
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt.reactiveArr.detail.resourceName)}}
                                </div>
                            </ui-form-item>
                        </ui-col>
                    </ui-row>
                </ui-form>
            </div>
            <!-- 不需要表格图表空样式，去掉 operate-management-count 样式 -->
            <div class="no-table-img-url">
                <!-- header信息-->
                <div class="info-header">
                    <span class="info-line"></span>
                    <!-- 配置信息 -->
                    <span class="info-title">
                        {{$t('allAlarmRulesDetail.alarmResources.title')}}
                    </span>
                </div>
                <ui-config-provider :locale="locale">
                    <ui-table
                        border
                        style="width: 100%"
                        :class="[tableClass(allAlarmRulesDetailOpt.reactiveArr.detail.instances, false)]"
                        :data="allAlarmRulesDetailOpt.reactiveArr.detail.instances"
                    >
                        <!-- 实例名称 -->
                        <ui-table-column
                            align="center"
                            min-width="110"
                            fixed
                            :label="$t('allAlarmRulesDetail.alarmResources.table.label.instanceName')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.instanceName)}}
                            </template>
                        </ui-table-column>
                        <!-- 实例ID -->
                        <ui-table-column
                            align="center"
                            min-width="150"
                            :label="$t('allAlarmRulesDetail.alarmResources.table.label.instanceId')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.instanceId)}}
                            </template>
                        </ui-table-column>
                        <!-- 内网IPV4(eth0) -->
                        <ui-table-column
                            align="center"
                            min-width="110"
                            :label="$t('allAlarmRulesDetail.alarmResources.table.label.ipv4')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.privateIpv4)}}
                            </template>
                        </ui-table-column>
                        <!-- 内网IPV6(eth0) -->
                        <ui-table-column
                            align="center"
                            min-width="150"
                            :label="$t('allAlarmRulesDetail.alarmResources.table.label.ipv6')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                {{$filter.emptyFilter(scope.row.privateIpv6)}}
                            </template>
                        </ui-table-column>
                    </ui-table>
                </ui-config-provider>
            </div>
            <div class="no-table-img-url" style="padding-bottom: 5px;">
                <!-- header信息-->
                <div class="info-header">
                    <span class="info-line"></span>
                    <!-- 配置信息 -->
                    <span class="info-title">
                        {{$t('allAlarmRulesDetail.triggeringConditions.title')}}
                    </span>
                </div>
                <div>
                    <div
                        v-for="(item, index) in allAlarmRulesDetailOpt.reactiveArr.detail.triggerDescription"
                        :key="index"
                        :style="{
                            marginBottom: allAlarmRulesDetailOpt.reactiveArr.detail.triggerDescription.length - 1 === index ? '0px' : '10px'
                        }"
                    >
                        <ui-tag effect="light">
                            {{item}}
                        </ui-tag>
                    </div>
                </div>
            </div>
            <div class="operate-management-count">
                <!-- header信息 通知策略-->
                <div class="info-header">
                    <span class="info-line"></span>
                    <!-- 通知策略 -->
                    <span class="info-title">
                        {{$t('allAlarmRulesDetail.notificationStrategy.title')}}
                    </span>
                </div>
                <ui-form @submit.stop.prevent>
                    <!-- Layout布局 -->
                    <ui-row
                        class="info-content"
                        :gutter="20"
                    >
                        <!-- 通知周期 -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.notificationStrategy.label.notificationCycle'))"
                            >
                                <div class="set-wrap info-name">
                                    {{allAlarmRulesDetailOpt.setNoticePeriod()}}
                                </div>
                            </ui-form-item>
                        </ui-col>
                        <!-- 有效时段 -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.notificationStrategy.label.effectivePeriod'))"
                            >
                                <div class="set-wrap info-name">
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt.reactiveArr.detail?.noticeOption?.effectiveIntervalStart)}}-
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt.reactiveArr.detail?.noticeOption?.effectiveIntervalEnd)}}
                                </div>
                            </ui-form-item>
                        </ui-col>
                        <!-- 通知条件 -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.notificationStrategy.label.notificationConditions'))"
                            >
                                <div class="set-wrap info-name">
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt.reactiveArr.detail?.statusName)}}
                                </div>
                            </ui-form-item>
                        </ui-col>
                        <!-- 接收渠道 -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.notificationStrategy.label.receivingChannel'))"
                            >
                                <div class="set-wrap info-name">
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt?.setNoticeWay())}}
                                </div>
                            </ui-form-item>
                        </ui-col>
                        <!-- 通知对象 -->
                        <ui-col :span="8">
                            <ui-form-item
                                :label="$filter.withClon($t('allAlarmRulesDetail.notificationStrategy.label.notificationObject'))"
                            >
                                <div class="set-wrap info-name">
                                    {{$filter.emptyFilter(allAlarmRulesDetailOpt.reactiveArr.detail?.noticeOption?.userName)}}
                                </div>
                            </ui-form-item>
                        </ui-col>
                    </ui-row>
                </ui-form>
            </div>
        </div>
        <history-alarm :alaram-history-opt="alaramHistoryOpt"/>
    </div>
</template>
<script lang="ts" setup>
import {RouterOperate} from 'utils/publicClass.ts';
import {languageSwitch, tableClass, methodsTotal} from 'utils/index.ts';
import AllAlarmRulesDetailOpt from './detail';
import AlarmHistoryOpt from './historyAlarm/historyAlarm';
import HistoryAlarm from './historyAlarm/historyAlarm.vue';

const locale = languageSwitch();

// 路由
const router = new RouterOperate().router;

const alaramHistoryOpt = new AlarmHistoryOpt();

const ruleId = router.currentRoute.value.query.ruleId;

const allAlarmRulesDetailOpt = new AllAlarmRulesDetailOpt(ruleId, (path: string) => router.push(path));
</script>
<style lang="scss">
@import 'assets/css/detail.scss';

</style>