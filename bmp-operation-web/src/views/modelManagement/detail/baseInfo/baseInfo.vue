<template>
    <!-- 阻止浏览器默认事件以及阻止冒泡事件 -->
    <div
        class="model-detail-basic-content"
        v-loading="loading"
    >
        <div class="operate-management-count">
            <!-- header信息 规格信息-->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 规格信息 -->
                <span class="info-title">
                    {{$t('modelDetail.header.specifications')}}
                </span>
            </div>
            <ui-form @submit.stop.prevent>
                <!-- Layout布局 -->
                <ui-row
                    class="info-content"
                    :gutter="20"
                >
                    <!-- 机型名称 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.modelName'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.name)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 机房名称 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.computerRoomName'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.idcName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 机型类型 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.modelType'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.deviceSeriesName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 体系架构 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.architecture'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.architecture)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 引导模式 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.bootMode'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.bootMode)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 机型规格 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.modelSpecifications'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.deviceType)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- CPU -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.CPU'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.cpuInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 内存 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.storage'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.memInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 系统盘 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.sysDisc'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.svInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- RAID配置 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.raidConfig'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.raidCan)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 系统盘RAID -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.sysRAID'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.raid)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 数据盘 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.dataDisc'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.dvInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 网卡 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.networkCard'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.nicInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 网络设置 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.networkSettings'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.interfaceModeName)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- GPU -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.GPU'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.gpuInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 高度U -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.height'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.height)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 创建时间 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.createTime'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.createdTime)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 描述 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.desc'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.description)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                </ui-row>
            </ui-form>
        </div>
        <div class="operate-management-count">
            <!-- header信息-->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 配置信息 -->
                <span class="info-title">
                    {{$t('modelDetail.header.configParams')}}
                </span>
            </div>
            <ui-form @submit.stop.prevent>
                <!-- Layout布局 -->
                <ui-row
                    class="info-content"
                    :gutter="20"
                >
                    <!-- CPU -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.CPU'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.cpuInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 内存 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.storage'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.memInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 系统盘 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.sysDisc'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.svInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- RAID配置 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.raidConfig'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.raidCan)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 系统盘RAID -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.sysRAID'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.raid)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 数据盘 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.dataDisc'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.dvInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 网卡 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.networkCard'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.nicInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 网络设置 -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.networkSettings'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.interfaceModeName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- GPU -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.GPU'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.gpuInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 高度U -->
                    <ui-col :span="8">
                        <ui-form-item
                            :label="$filter.withClon($t('modelDetail.tabs.basicInfo.label.height'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(reactiveArr.detail.height)}}
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
                    {{$t('modelDetail.header.volumeManagement.title')}}
                </span>
            </div>
            <ui-config-provider :locale="locale">
                <ui-table
                    border
                    style="width: 100%"
                    :class="[tableClass(reactiveArr.detail.volumes, false)]"
                    :data="reactiveArr.detail.volumes"
                >
                    <!-- 卷名称 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        fixed
                        :label="$t('modelDetail.header.volumeManagement.label.name')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeName)}}
                        </template>
                    </ui-table-column>
                    <!-- 卷类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('modelDetail.header.volumeManagement.label.type')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeType)}}
                        </template>
                    </ui-table-column>
                    <!-- RAID配置 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        v-if="reactiveArr.detail.isNeedRaid !== 'no_need_raid'"
                        :label="$t('modelDetail.header.volumeManagement.label.raidCan')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(baseInfo.setRaidCan(scope.row.raidCan))}}
                        </template>
                    </ui-table-column>
                    <!-- RAID模式 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        v-if="reactiveArr.detail.isNeedRaid !== 'no_need_raid'"
                        :label="$t('modelDetail.header.volumeManagement.label.raid')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <p>
                                {{$filter.emptyFilter(scope.row.raid)}}
                            </p>
                        </template>
                    </ui-table-column>
                    <!-- 磁盘类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('modelDetail.header.volumeManagement.label.disk')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(baseInfo.setDiskInterfaceType(scope.row.diskType))}}
                        </template>
                    </ui-table-column>
                    <!-- 接口类型 -->
                    <ui-table-column
                        align="center"
                        min-width="100"
                        :label="$t('modelDetail.header.volumeManagement.label.interfaceType')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(baseInfo.setDiskInterfaceType(scope.row.interfaceType))}}
                        </template>
                    </ui-table-column>
                    <!-- 单盘最小容量 -->
                    <ui-table-column
                        align="center"
                        :min-width="setDiffClass('170', '100')"
                        :label="$t('modelDetail.header.volumeManagement.label.minNum')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeSize)}}{{scope.row.volumeUnit}}
                        </template>
                    </ui-table-column>
                    <!-- 最低数量（块） -->
                    <ui-table-column
                        align="center"
                        fixed="right"
                        :min-width="setDiffClass('170', '100')"
                        :label="$t('modelDetail.header.volumeManagement.label.amount')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            {{$filter.emptyFilter(scope.row.volumeAmount)}}
                        </template>
                    </ui-table-column>
                </ui-table>
            </ui-config-provider>
        </div>
    </div>
</template>
<script lang="ts" setup>
import {tableClass, languageSwitch, setDiffClass} from 'utils/index.ts';
import { CurrencyType } from '@utils/publicType';
import BasicInfo from './base';

interface PropsType {
    reactiveArr: {
        detail: CurrencyType;
    };
    loading: boolean;
    baseInfo: BasicInfo;
    // imgSrc?: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

const locale = languageSwitch();
</script>