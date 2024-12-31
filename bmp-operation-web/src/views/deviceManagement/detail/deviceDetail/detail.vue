<template>
    <!-- <div v-if="deviceDetail.isLoading.value" style="position:absolute; width: 100%;height: 100%;" v-loading="deviceDetail.isLoading.value"></div> -->
    <main
        class="device-detail-content"
        v-loading="deviceDetail.isLoading.value"
    >
        <!-- 设备信息 -->
        <div class="operate-management-count">
            <!-- header信息 设备信息-->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 设备信息 -->
                <span class="info-title">
                    {{$t('deviceDetail.deviceInfo.deviceInfo')}}
                </span>
            </div>
            <!-- 阻止浏览器默认事件、冒泡事件 -->
            <ui-form @submit.stop.prevent>
                <!-- Layout布局 -->
                <ui-row
                    class="info-content"
                    :gutter="20"
                >
                    <!-- 布局等分 -->
                    <ui-col :span="8">
                        <!-- sn -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.sn'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.sn)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 机型名称 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.modelName'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.deviceTypeName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 管理状态 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.managementStatus'))">
                            <div class="set-wrap info-name">
                                <span style="marginRight: 10px">
                                    {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.manageStatusName)}}
                                </span>
                                <ui-tooltip
                                    placement="right"
                                    v-if="deviceDetail.reactiveArr.detail?.reason && ['putawayfail'].includes(deviceDetail.reactiveArr.detail.manageStatus)"
                                >
                                    <template #content>
                                        <div class="status-error-tip">
                                            {{deviceDetail.reactiveArr.detail?.reason}}
                                        </div>
                                    </template>
                                    <img
                                        class="warning-tips"
                                        alt=""
                                        :src="($defInfo.imagePath('uiTooltip') as unknown as string)"
                                    />
                                </ui-tooltip>
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 采集状态 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.collectionStatus'))">
                            <div class="set-wrap info-name">
                                <!-- deviceDetail.reactiveArr.detail?.collectStatus -->
                                <div>
                                    <span>
                                        {{$filter.emptyFilter(deviceDetail.setCollected(deviceDetail.reactiveArr.detail?.collectStatus))}}
                                        <ui-tooltip
                                            placement="bottom"
                                            v-if="deviceDetail.reactiveArr.detail?.collectFailReason"
                                        >
                                            <template #content>
                                                <div v-if="deviceDetail.reactiveArr.detail?.collectFailReason">
                                                    {{deviceDetail.reactiveArr.detail?.collectFailReason}}
                                                </div>
                                            </template>
                                            <img
                                                class="warning-tips"
                                                alt=""
                                                v-if="deviceDetail.reactiveArr.detail?.collectFailReason"
                                                :src="($defInfo.imagePath('uiTooltip') as unknown as string)"
                                            />
                                        </ui-tooltip>
                                    </span>
                                </div>
                                <!-- 仅在设备管理状态为【已入库】状态下执行 -->
                                <p
                                    class="collect-count"
                                    v-if="deviceDetail.reactiveArr.detail.manageStatus === 'in'"
                                    @click="collect.collectClick"
                                >
                                    <img
                                        alt=""
                                        class="collect-img"
                                        :src="($defInfo.imagePath('collect') as unknown as string)"
                                    />
                                    <span>{{$t('deviceDetail.deviceInfo.recapture')}}</span>
                                </p>
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 机房名称 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.computerRoomName'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.idcName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 机柜编码 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.cabinetCode'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.cabinet)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 所在U位 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.uBit'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.uPosition)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 备注 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.deviceInfo.remark'))">
                            <div class="info-name set-text-wrap">
                                <span>
                                    {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.description)}}
                                    <!-- <ui-tooltip
                                        placement="bottom"
                                        v-if="deviceDetail.reactiveArr?.deviceDetail.reactiveArr.detail?.description?.length > 22"
                                    >
                                        <template #content>
                                            <div class="set-tooltip-width">
                                                <div class="set-wrap info-name">
                                                    {{deviceDetail.reactiveArr.detail.description}}
                                                </div>
                                            </div>
                                        </template>
                                        <div class="single-line-omission">
                                            {{$filter.emptyFilter(deviceDetail.reactiveArr.detail.description)}}
                                        </div>
                                    </ui-tooltip>
                                    <span v-else>
                                        {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.description)}}
                                    </span> -->
                                    <img
                                        alt=""
                                        class="edit-icon"
                                        :src="($defInfo.imagePath('descEdit') as unknown as string)"
                                        @click="notesEdit.editNotesClick"
                                    />
                                    <!-- @click="editNotes.editNotesClick" -->

                                </span>
                            </div>
                        </ui-form-item>
                    </ui-col>
                </ui-row>
            </ui-form>
        </div>
        <!-- 实例信息 -->
        <div class="operate-management-count">
            <!-- header信息 -->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 实例信息 -->
                <span class="info-title">
                    {{$t('deviceDetail.instanceInfo.instanceInfo')}}
                </span>
            </div>
            <!-- 阻止浏览器默认事件、冒泡事件 -->
            <ui-form @submit.stop.prevent>
                <!-- Layout布局 -->
                <ui-row
                    class="info-content"
                    :gutter="20"
                >
                    <!-- 布局等分 -->
                    <ui-col :span="8">
                        <!-- 实例名称 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.instanceName'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.instanceName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 实例ID -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.instanceID'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.instanceId)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 实例镜像 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.instanceImage'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.imageName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 实例状态 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.instanceStatus'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.instanceStatusName)}}
                                <ui-tooltip
                                    placement="right"
                                    v-if="[
                                        'Creation failed',
                                        'Startup failed',
                                        'Shutdown failed',
                                        'Reboot failed',
                                        'Delete failed',
                                        'Reinstall failed',
                                        'Resetpasswd failed'
                                    ].includes(deviceDetail.reactiveArr.detail?.instanceStatus)"
                                >
                                    <template #content>
                                        <div class="status-error-tip">
                                            <span v-if="['Creation failed'].includes(deviceDetail.reactiveArr.detail?.instanceStatus)">
                                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.instanceReason)}}
                                            </span>
                                            <span v-else>
                                                {{$t('deviceDetail.operate.error.tip')}}
                                            </span>
                                        </div>
                                    </template>
                                    <img
                                        class="warning-tips"
                                        alt=""
                                        :src="($defInfo.imagePath('uiTooltip') as unknown as string)"
                                    />
                                </ui-tooltip>
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 监控Agent状态 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.monitorAgentStatus'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail?.statusName?.value)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 锁定状态 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.lockStatus'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.lockedname)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 实例属主 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.instanceOwner'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.userName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 创建时间 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.createTime'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.instanceCreatedTime)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 描述 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.instanceInfo.desc'))">
                            <div class="set-wrap info-name set-text-wrap">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.instanceDescription)}}
                                <!-- <ui-tooltip
                                    placement="bottom"
                                    v-if="deviceDetail.reactiveArr?.deviceDetail.reactiveArr.detail?.instanceDescription?.length > 28"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            {{deviceDetail.reactiveArr.detail.instanceDescription}}
                                        </div>
                                    </template>
                                    <div ref="desc" class="single-line-omission">
                                        {{deviceDetail.reactiveArr.detail?.instanceDescription}}
                                    </div>
                                </ui-tooltip>
                                <span v-else ref="desc">
                                    {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.instanceDescription)}}
                                </span> -->
                            </div>
                        </ui-form-item>
                    </ui-col>
                </ui-row>
            </ui-form>
        </div>
        <!-- 硬件配置 - 机型信息-->
        <div class="operate-management-count">
            <!-- header信息 -->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 硬件配置 - 机型信息 -->
                <span class="info-title">
                    {{$t('deviceDetail.hardwareConfiguration.modelInfo')}}
                </span>
            </div>
            <!-- 阻止浏览器默认事件、冒泡事件 -->
            <ui-form @submit.stop.prevent>
                <!-- Layout布局 -->
                <ui-row
                    class="info-content"
                    :gutter="20"
                >
                    <!-- 布局等分 -->
                    <ui-col :span="8">
                        <!-- 品牌 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.brand'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.brand)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 型号 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.model'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.model)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 体系架构 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.architecture'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.architecture)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- CPU -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.CPU'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.cpuInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 内存 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.storage'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.memInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 系统盘 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.sysDisc'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.svInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <ui-col :span="8">
                        <!-- adapeter_id（RAID卡）-->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.adapeterId'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(String(deviceDetail.reactiveArr.detail?.adapterId))}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- enclosure（背板号1）-->
                    <!-- <ui-col :span="8">
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.enclosure1'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.enclosure1)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- slot（槽位1）-->
                    <!-- <ui-col :span="8">
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.slot1'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(String(deviceDetail.reactiveArr.detail?.slot1))}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- enclosure（背板号2）-->
                    <!-- <ui-col :span="8">
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.enclosure2'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.enclosure2)}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- slot（槽位2）-->
                    <!-- <ui-col :span="8">
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.slot2'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(String(deviceDetail.reactiveArr.detail?.slot2))}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <!-- 数据盘 -->
                    <!-- <ui-col :span="8">
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.dataDisc'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(String(deviceDetail.reactiveArr.detail?.dvInfo))}}
                            </div>
                        </ui-form-item>
                    </ui-col> -->
                    <ui-col :span="8">
                        <!-- GPU -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.GPU'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.gpuInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 网卡 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.networkCard'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.nicInfo)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <!-- 网络设置 -->
                    <ui-col :span="8">
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.hardwareConfiguration.networkSettings'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.interfaceModeName)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                </ui-row>
            </ui-form>
        </div>
        <!-- 网络信息 -->
        <div class="operate-management-count">
            <!-- header信息 -->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 网络信息 -->
                <span class="info-title">
                    {{$t('deviceDetail.networkInfo.networkInfo')}}
                </span>
            </div>
            <!-- 阻止浏览器默认事件、冒泡事件 -->
            <ui-form @submit.stop.prevent>
                <!-- Layout布局 -->
                <ui-row
                    class="info-content"
                    :gutter="20"
                >
                    <!-- 布局等分 -->
                    <ui-col :span="8">
                        <!-- 内网IPv4(eth0) -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.intranceIPv4'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.privateIpv4)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 内网IPv4(eth1) -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.intranceIPv4First'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.privateEth1Ipv4)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 内网IPv6(eth0) -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.intranetIPv6'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.privateIpv6)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 内网IPv6(eth1) -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.intranetIPv6First'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.privateEth1Ipv6)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- MAC1（eth0）-->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.mac1'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.mac1)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- MAC2（eth1）-->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.mac2'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.mac2)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 交换机IP1（eth0）-->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.switchIPOne'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.switchIp1)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 交换机IP2（eth1）-->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.switchIPTwo'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.switchIp2)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 交换机上联端口1（eth0）-->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.switchUplinkPortOne'))">
                            <ui-tooltip
                                placement="bottom"
                                v-model="deviceDetail.reactiveArr.tooltipStatus['switchPort1'].showTooltip"
                                :disabled="!deviceDetail.reactiveArr.tooltipStatus['switchPort1'].showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        {{deviceDetail.reactiveArr.detail?.switchPort1}}
                                    </div>
                                </template>
                                <div
                                    class="info-name device-detail-tip-show"
                                    @mouseenter="hasShowTooltip($event, deviceDetail.reactiveArr.tooltipStatus['switchPort1'], deviceDetail.reactiveArr.detail?.switchPort1, 1.18)"
                                >
                                    {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.switchPort1)}}
                                </div>
                            </ui-tooltip>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 交换机上联端口2（eth1）-->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.switchUplinkPortTwo'))">
                            <ui-tooltip
                                placement="bottom"
                                v-model="deviceDetail.reactiveArr.tooltipStatus['switchPort2'].showTooltip"
                                :disabled="!deviceDetail.reactiveArr.tooltipStatus['switchPort2'].showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        {{deviceDetail.reactiveArr.detail?.switchPort2}}
                                    </div>
                                </template>
                                <div
                                    class="info-name device-detail-tip-show"
                                    @mouseenter="hasShowTooltip($event, deviceDetail.reactiveArr.tooltipStatus['switchPort2'], deviceDetail.reactiveArr.detail?.switchPort2, 1.18)"
                                >
                                    {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.switchPort2)}}
                                </div>
                            </ui-tooltip>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 子网掩码(eth0) -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.subnetMask'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.mask)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 子网掩码(eth1) -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.subnetMaskFirst'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.eth1Mask)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 网关 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.networkInfo.networkGateway'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.gateway)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                </ui-row>
            </ui-form>
        </div>
        <!-- 交换机&带外信息 -->
        <div :class="['operate-management-count', locationItem.getLocationItem === 'en_US' ? 'switch-info' : '']">
            <!-- header信息 -->
            <div class="info-header">
                <span class="info-line"></span>
                <!-- 交换机&带外信息 -->
                <span class="info-title">
                    {{$t('deviceDetail.switchOutsideInfo.switchOutsideInfo')}}
                </span>
            </div>
            <!-- 阻止浏览器默认事件、冒泡事件 -->
            <ui-form @submit.stop.prevent>
                <!-- Layout布局 -->
                <ui-row
                    class="info-content"
                    :gutter="20"
                >
                    <!-- 布局等分 -->
                    <ui-col :span="8">
                        <!-- 网口1交换机登录用户名 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.switchOutsideInfo.userName1'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.switchUser1)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 网口1交换机登录密码 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.switchOutsideInfo.password1'))">
                            <div class="set-wrap info-name">
                                <!-- 密码-明文-加密-eyeicon -->
                                <p class="login-password">
                                    <!-- 通过传入的状态，判断密码显示状态 -->
                                    <span
                                        :style="['--'].includes(deviceDetail.reactiveArr.detail.switchPassword1) ? {color: '#333', cursor: 'text'} : ''"
                                        :class="[deviceDetail.reactiveArr.internetPort1 ? 'set-size-small' : 'set-size']"
                                    >
                                        {{$filter.defaultPassword(deviceDetail.reactiveArr.detail.switchPassword1)}}
                                    </span>
                                    <!-- open-eye -->
                                    <img
                                        class="login-password-img"
                                        alt=""
                                        v-if="!deviceDetail.reactiveArr.internetPort1"
                                        :src="($defInfo.imagePath('closeEye') as unknown as string)"
                                        @click="deviceDetail.hasPasswordClick(true, 'open', 'internetPort1', 'switchPassword1')"
                                    />
                                    <!-- close-eye -->
                                    <img
                                        class="login-password-img"
                                        alt=""
                                        v-else
                                        :src="($defInfo.imagePath('openEye') as unknown as string)"
                                        @click="deviceDetail.hasPasswordClick(false, 'close', 'internetPort1', 'switchPassword1')"
                                    />
                                </p>
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 网口2交换机登录用户名 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.switchOutsideInfo.userName2'))">
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.switchUser2)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 网口2交换机登录密码 -->
                        <ui-form-item :label="$filter.withClon($t('deviceDetail.switchOutsideInfo.password2'))">
                            <div class="set-wrap info-name">
                                <!-- 密码-明文-加密-eyeicon -->
                                <p class="login-password">
                                    <!-- 通过传入的状态，判断密码显示状态 -->
                                    <span
                                        :style="['--'].includes(deviceDetail.reactiveArr.detail.switchPassword2) ? {color: '#333', cursor: 'text'} : ''"
                                        :class="[deviceDetail.reactiveArr.internetPort2 ? 'set-size-small' : 'set-size']"
                                    >
                                        {{$filter.defaultPassword(deviceDetail.reactiveArr.detail?.switchPassword2)}}
                                    </span>
                                        <!-- open-eye -->
                                    <img
                                        class="login-password-img"
                                        alt=""
                                        v-if="!deviceDetail.reactiveArr.internetPort2"
                                        :src="($defInfo.imagePath('closeEye') as unknown as string)"
                                        @click="deviceDetail.hasPasswordClick(true, 'open', 'internetPort2', 'switchPassword2')"
                                    />
                                    <!-- close-eye -->
                                    <img
                                        class="login-password-img"
                                        alt=""
                                        v-else
                                        :src="($defInfo.imagePath('openEye') as unknown as string)"
                                        @click="deviceDetail.hasPasswordClick(false, 'close', 'internetPort2', 'switchPassword2')"
                                    />
                                </p>
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 带外IP -->
                        <ui-form-item
                            class='ilo-count'
                            :label="$filter.withClon($t('deviceDetail.switchOutsideInfo.outOfBandIP'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.iloIp)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 带外登录用户名 -->
                        <ui-form-item
                            class='ilo-count'
                            :label="$filter.withClon($t('deviceDetail.switchOutsideInfo.outOfBandLoginUserName'))"
                        >
                            <div class="set-wrap info-name">
                                {{$filter.emptyFilter(deviceDetail.reactiveArr.detail?.iloUser)}}
                            </div>
                        </ui-form-item>
                    </ui-col>
                    <ui-col :span="8">
                        <!-- 带外登录密码 -->
                        <ui-form-item
                            class='ilo-count'
                            :label="$filter.withClon($t('deviceDetail.switchOutsideInfo.outOfBandLoginPassWord'))"
                        >
                            <div class="set-wrap info-name">
                                <!-- 密码-明文-加密-eyeicon -->
                                <p class="login-password">
                                    <!-- 通过传入的状态，判断密码显示状态 -->
                                    <span
                                        :style="['--'].includes(deviceDetail.reactiveArr.detail.iloPassword) ? {color: '#333', cursor: 'text'} : ''"
                                        :class="[deviceDetail.reactiveArr.outLoginPassword ? 'set-size-small' : 'set-size']"
                                    >
                                        {{$filter.defaultPassword(deviceDetail.reactiveArr.detail.iloPassword)}}
                                    </span>
                                    <!-- open-eye -->
                                    <img
                                        class="login-password-img"
                                        alt=""
                                        v-if="!deviceDetail.reactiveArr.outLoginPassword"
                                        :src="($defInfo.imagePath('closeEye') as unknown as string)"
                                        @click="deviceDetail.hasPasswordClick(true, 'open', 'outLoginPassword', 'iloPassword')"
                                    />
                                    <!-- close-eye -->
                                    <img
                                        class="login-password-img"
                                        alt=""
                                        v-else
                                        :src="($defInfo.imagePath('openEye') as unknown as string)"
                                        @click="deviceDetail.hasPasswordClick(false, 'close', 'outLoginPassword', 'iloPassword')"
                                    />
                                </p>
                            </div>
                        </ui-form-item>
                    </ui-col>
                </ui-row>
            </ui-form>
        </div>
    </main>
</template>

<script lang="ts" setup>
import {locationItem} from 'utils/publicClass.ts';
import {hasShowTooltip} from 'utils/index.ts';

// defineEmits API 来替代 emits
interface PropsType {
    deviceDetail: any;
    notesEdit: any;
    collect: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>
