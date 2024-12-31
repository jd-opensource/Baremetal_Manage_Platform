<template>
    <div class="page-position" :class="tabName==='HardwareMonitoring' || tabName==='operationLog'? 'pr7' : 'can-scroll'">
        <div class="detail-header">
            <el-button
                class="back-button"
                type= "text"
                :icon ="ArrowLeft" 
                @click="goBack"
            >
                {{$t('list.back')}}
            </el-button>
            <span class="ml32">{{ $t('instance.detail.instanceDetail') }}</span>
            <span class="ml15">/</span>
            <div style="display: inline-flex;">
                <el-tooltip  
                    :disabled="!tooltipStatus['instanceName'].showTooltip"  
                    :content= instanceInfo.detail.instanceName
                    placement="top-start"
                >
                    <span 
                        class="long-name-title ml15" 
                        @mouseenter="hasShowTooltip($event, tooltipStatus['instanceName'], instanceInfo.detail.instanceName, 179, 'detail')"
                    >
                        {{ instanceInfo.detail.instanceName || $filter.emptyFilter() }}
                    </span>
                </el-tooltip>
            </div>
            <span 
                class="ml24" 
                :class="stateChange(instanceInfo?.detail?.status)"
            >
                {{instanceInfo.detail.statusName}}
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
                            <!-- 开机 -->
                            <el-dropdown-item
                                v-if="instanceInfo?.detail?.status === 'stopped' "
                                :disabled="instanceInfo?.detail?.status !== 'stopped'"
                            >
                                <el-tooltip
                                    placement="right"
                                    :disabled="instanceInfo?.detail?.status === 'stopped'"
                                    :content="$t('instance.tip.open')"
                                >
                                    <label>
                                        <el-button
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateOpen()"
                                            :disabled="instanceInfo?.detail?.status !== 'stopped'"
                                        >
                                            <span>{{$t('list.operate.open')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 关机 -->
                            <el-dropdown-item
                                v-else
                                :disabled="instanceInfo?.detail?.status !== 'running'  ">
                                    <el-tooltip
                                        placement="right"
                                        :disabled="instanceInfo?.detail?.status === 'running' "
                                        :content="$t('instance.tip.close')"
                                    >
                                        <label>
                                            <el-button  
                                                type="text"  
                                                class="operate-btn"
                                                :disabled="instanceInfo?.detail?.status !== 'running' "
                                                @click="clickOperateClose()"
                                            >
                                                <span>{{$t('list.operate.close')}}</span>
                                            </el-button>
                                        </label>
                                    </el-tooltip>
                            </el-dropdown-item>
                            <!-- 重启 -->
                            <el-dropdown-item
                                :disabled="instanceInfo?.detail?.status !== 'running'">
                                    <el-tooltip
                                        placement="right"
                                        :disabled="instanceInfo?.detail?.status === 'running'"
                                        :content="$t('instance.tip.restart')"
                                    >
                                        <label>
                                            <el-button  
                                                type="text" 
                                                class="operate-btn"
                                                :disabled="instanceInfo?.detail?.status !== 'running'"
                                                @click="clickOperateRestart()"
                                            >
                                                <span>{{$t('list.operate.restart')}}</span>
                                            </el-button>
                                        </label>
                                    </el-tooltip>
                            </el-dropdown-item>
                            <!-- 锁定 -->
                            <el-dropdown-item
                                v-if="instanceInfo?.detail?.locked === 'unlocked' "
                                :disabled="(instanceInfo?.detail?.status === 'Creation failed')
                                ||(instanceInfo?.detail?.status === 'creating')
                                ||(instanceInfo?.detail?.status === 'Creating')
                                ||(instanceInfo?.detail?.status === 'Delete failed')
                                ||(instanceInfo?.detail?.status === 'destroying')
                                ||(instanceInfo?.detail?.status === 'Destroying')"
                            >
                                <el-tooltip
                                    placement="right"
                                    :disabled="(instanceInfo?.detail?.status !== 'Creation failed')
                                    &&(instanceInfo?.detail?.status !== 'creating')
                                    &&(instanceInfo?.detail?.status !== 'Creating')
                                    &&(instanceInfo?.detail?.status !== 'Delete failed')
                                    &&(instanceInfo?.detail?.status !== 'destroying')
                                    &&(instanceInfo?.detail?.status !== 'Destroying')"
                                    :content="$t('instance.tip.locking')"
                                >
                                    <label>
                                        <el-button  
                                            type="text" 
                                            class="operate-btn"
                                            :disabled="(instanceInfo?.detail?.status === 'Creation failed')
                                            ||(instanceInfo?.detail?.status === 'creating')
                                            ||(instanceInfo?.detail?.status === 'Creating')
                                            ||(instanceInfo?.detail?.status === 'Delete failed')
                                            ||(instanceInfo?.detail?.status === 'destroying')
                                            ||(instanceInfo?.detail?.status === 'Destroying')"
                                            @click="clickOperateLock()"
                                        >
                                            <span>{{$t('list.operate.lock')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 解锁 -->
                            <el-dropdown-item
                                v-if="instanceInfo?.detail?.locked === 'locked' "
                                :disabled="(instanceInfo?.detail?.status === 'Creation failed')
                                ||(instanceInfo?.detail?.status === 'creating')
                                ||(instanceInfo?.detail?.status === 'Creating')
                                ||(instanceInfo?.detail?.status === 'Delete failed')
                                ||(instanceInfo?.detail?.status === 'destroying')
                                ||(instanceInfo?.detail?.status === 'Destroying')"
                            >
                                <el-tooltip
                                    placement="right"
                                    :disabled="(instanceInfo?.detail?.status !== 'Creation failed')
                                    &&(instanceInfo?.detail?.status !== 'creating')
                                    &&(instanceInfo?.detail?.status !== 'Creating')
                                    &&(instanceInfo?.detail?.status !== 'Delete failed')
                                    &&(instanceInfo?.detail?.status !== 'destroying')
                                    &&(instanceInfo?.detail?.status !== 'Destroying')"
                                    :content="$t('instance.tip.unlocking')"
                                >
                                    <label>
                                        <el-button  
                                            type="text" 
                                            class="operate-btn"
                                            :disabled="(instanceInfo?.detail?.status === 'Creation failed')
                                            ||(instanceInfo?.detail?.status === 'creating')
                                            ||(instanceInfo?.detail?.status === 'Creating')
                                            ||(instanceInfo?.detail?.status === 'Delete failed')
                                            ||(instanceInfo?.detail?.status === 'destroying')
                                            ||(instanceInfo?.detail?.status === 'Destroying')"
                                            @click="clickOperateUnlock()"
                                        >
                                            <span>{{$t('list.operate.unlock')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 重置密码  -->
                            <el-dropdown-item
                                :disabled="instanceInfo?.detail?.status !== 'stopped'">
                                <el-tooltip
                                    placement="right"
                                    :disabled="instanceInfo?.detail?.status === 'stopped'"
                                    :content="$t('instance.tip.open')"
                                >
                                    <label>
                                        <el-button
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateResetPassword()"
                                            :disabled="instanceInfo?.detail?.status !== 'stopped'"
                                        >
                                            <span>{{$t('list.operate.resetPassword')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                             <!-- 重装系统  -->
                             <el-dropdown-item
                                :disabled="instanceInfo?.detail?.status !== 'stopped'"
                            >
                                <el-tooltip
                                    placement="right"
                                    :disabled="instanceInfo?.detail?.status === 'stopped'"
                                    :content="$t('instance.tip.open')"
                                >
                                    <label>
                                        <el-button
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateResystemConfirm()"
                                            :disabled="instanceInfo?.detail?.status !== 'stopped'"
                                        >
                                            <span>{{$t('list.operate.resystem')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 删除 -->
                            <el-dropdown-item
                                :disabled="
                                !(instanceInfo?.detail?.status === 'stopped'
                                ||instanceInfo?.detail?.status === 'Creation failed') 
                                || instanceInfo?.detail?.locked === 'locked'"
                            >
                                <el-tooltip
                                    placement="right"
                                    :disabled="
                                    !(!(instanceInfo?.detail?.status === 'stopped'
                                    ||instanceInfo?.detail?.status === 'Creation failed') 
                                    || instanceInfo?.detail?.locked === 'locked')"
                                    :content="instanceInfo?.detail?.locked==='unlocked' 
                                    ? $t('instance.tip.delete') : $t('instance.tip.locked')"
                                >
                                    <label>
                                        <el-button  
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateDelete()"
                                            :disabled="
                                            !(instanceInfo?.detail?.status === 'stopped'
                                            ||instanceInfo?.detail?.status === 'Creation failed')
                                            || instanceInfo?.detail?.locked === 'locked'"
                                        >
                                            <span>{{$t('list.operate.delete')}}</span>
                                        </el-button >
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div>            
        </div>
        <div class="page-content instance-detail-content t110">
            <el-tabs
				ref="pageTabs"
                class="instance-tab-info"
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
                    <!-- 实例信息 -->
                    <h3 class="section-title">{{ $t('instance.detail.instanceMessage')}}</h3>
                    <div class="ml22">
                        <el-form 
                            v-loading="isLoading"
                            @submit.stop.prevent
                        >                
                            <el-row :gutter="20">
                                <!--实例名称-->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.instanceName') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['instanceName'].showTooltip"  
                                            :content= instanceInfo.detail.instanceName
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['instanceName'], instanceInfo.detail.instanceName, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.instanceName || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                        <img
                                            class="desc-type"
                                            src="@/assets/img/describe.png"
                                            @click="clickEditInstanceName()"
                                        />
                                    </el-form-item>
                                </el-col>
                                <!-- 实例ID -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.instanceId') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['instanceId'].showTooltip"   
                                            :content= instanceInfo.detail.instanceId
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['instanceId'], instanceInfo.detail.instanceId, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.instanceId || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                        <img
                                            class="desc-type"
                                            src="@/assets/img/copy.png"
                                            @click="clickOperateCopy(instanceInfo.detail.instanceId)"
                                        />
                                    </el-form-item>   
                                </el-col>
                                <!-- 机房名称 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.computerName') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.idcName || $filter.emptyFilter()  }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 机型名称 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.deviceName') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['deviceTypeName'].showTooltip"   
                                            :content= instanceInfo.detail.deviceTypeName
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['deviceTypeName'], instanceInfo.detail.deviceTypeName, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.deviceTypeName || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- SN -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.sn') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.sn || $filter.emptyFilter()  }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 状态 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.list.state') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{instanceInfo.detail.statusName || $filter.emptyFilter()}}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 监控Agent状态 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('alarm.monitor.montiorState') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{monitorAgent || $filter.emptyFilter()}}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 锁定状态 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.lockStatus') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{lockedStatus[instanceInfo.detail.locked] || $filter.emptyFilter()}}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 镜像 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.mirror') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['imageName'].showTooltip"   
                                            :content= instanceInfo.detail.imageName
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['imageName'], instanceInfo.detail.imageName, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.imageName || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 创建时间 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.createdTime') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.createdTime || $filter.emptyFilter() }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 描述 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.description') + $filter.withClon(' ')">
                                        <el-tooltip  
                                            popper-class="desc-tooltip-popper"
                                            :disabled="!tooltipStatus['description'].showTooltip || instanceInfo.detail.description === undefined"  
                                            :content= instanceInfo.detail.description
                                            placement="bottom"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['description'], instanceInfo.detail.description, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.description}}
                                            </span>
                                        </el-tooltip>
                                        <img
                                            class="desc-type"
                                            src="@/assets/img/describe.png"
                                            @click="clickEditDescription()"
                                        />
                                    </el-form-item>
                                </el-col>
                            </el-row>
                        </el-form>
                    </div>
                    <el-divider />
                    <!-- 硬件信息 -->
                    <h3 class="section-title">{{ $t('instance.detail.hardwareInformation') }}</h3>
                    <div class="ml22">
                        <el-form 
                            v-loading="isLoading"
                            @submit.stop.prevent
                        >                
                            <el-row :gutter="20">
                                <!-- CPU -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.cpu') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['cpuInfo'].showTooltip"  
                                            :content= instanceInfo.detail.cpuInfo
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['cpuInfo'], instanceInfo.detail.cpuInfo, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.cpuInfo || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 内存 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.memory') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['memInfo'].showTooltip"  
                                            :content= instanceInfo.detail.memInfo
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['memInfo'], instanceInfo.detail.memInfo, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.memInfo || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 系统盘RAID -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.create.systemRollRaid') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.systemVolumeRaidName || $filter.emptyFilter() }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 系统盘 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.create.systemRoll') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['systemInfo'].showTooltip"  
                                            :content= instanceInfo.detail.systemInfo
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['systemInfo'], instanceInfo.detail.systemInfo, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.systemInfo || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 网卡 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.networkCard') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['nicInfo'].showTooltip"  
                                            :content= instanceInfo.detail.nicInfo
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['nicInfo'], instanceInfo.detail.nicInfo, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.nicInfo || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- GPU -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.gpu') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['gpuInfo'].showTooltip"  
                                            :content= instanceInfo.detail.gpuInfo
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['gpuInfo'], instanceInfo.detail.gpuInfo, 259, 'detail')"
                                            >
                                                {{ instanceInfo.detail.gpuInfo || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                        </el-form>
                    </div>
                    <el-divider />
                    <!-- 数据卷信息 -->
                    <h3 class="section-title">{{ $t('instance.detail.datasInformation') }}</h3>
                    <div class="table-content m0 h-max">
                        <el-table 
                            border
                            ref="datasTableRef"
                            :data="instanceInfo.datasDetail" 
                        >   
                            <!-- 卷名称 -->
                            <el-table-column 
                                prop="volume.volumeName" 
                                :label="$t('instance.create.rollName')" 
                                width="120"
                            />
                            <!-- RAID模式 -->
                            <el-table-column 
                                v-if="instanceInfo.detail.isNeedRaid === 'need_raid'"
                                prop="raid.name" 
                                :label="$t('instance.create.raidModel')" 
                                min-width="200"
                            >
                                <template v-slot="scope">
                                    <p>{{scope.row.raid?.name || $filter.emptyFilter()}}</p>
                                </template>
                            </el-table-column>
                            <!-- 硬盘数量 -->
                            <el-table-column 
                                prop="volume.volumeAmount" 
                                :label="$t('instance.create.harddiskNumber')" 
                                min-width="100"
                            />
                            <!-- 单个硬盘容量 -->
                            <el-table-column 
                                prop="volume.createdTime" 
                                :label="$t('instance.create.harddiskVolume')" 
                                min-width="200"
                            >
                                <template v-slot="scope">
                                    <p>{{scope.row.volume.volumeSize + scope.row.volume.volumeUnit}}</p>
                                </template>
                            </el-table-column>
                            <!-- 可用硬盘容量 -->
                            <el-table-column 
                                prop="volume.volumeSize" 
                                :label="$t('instance.create.useVolume')" 
                                min-width="200"
                            >
                                <template v-slot="scope">
                                    <p>{{scope.row.volume.volumeSize + scope.row.volume.volumeUnit}}</p>
                                </template>
                            </el-table-column>
                            <!-- 硬盘类型 -->
                            <el-table-column 
                                prop="volume.diskType" 
                                :label="$t('instance.create.harddiskType')" 
                                min-width="100"
                            />
                            <template #empty>
                                <div>{{$t('instance.create.noData')}}</div>
                            </template>
                        </el-table>
                    </div>
                    <el-divider />
                    <!-- 网络信息 -->
                    <h3 class="section-title">{{ $t('instance.detail.netInformation') }}</h3>
                    <div class="ml22">
                        <el-form 
                            v-loading="isLoading"
                            @submit.stop.prevent>                
                            <el-row :gutter="20">
                                <!-- 内网IPv4(eth0) -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.intranetIpv4') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.privateIpv4 || $filter.emptyFilter() }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 内网IPv4(eth1) -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.intranetIpv4eth1') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.privateEth1Ipv4 || $filter.emptyFilter() }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 内网IPv6(eth0) -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.intranetIpv6') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.privateIpv6 || $filter.emptyFilter() }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 内网IPv6(eth1) -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.intranetIpv6eth1') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.privateEth1Ipv6 || $filter.emptyFilter() }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 带外IP -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.outBandIp') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ instanceInfo.detail.iloIp || $filter.emptyFilter() }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                            </el-row>
                        </el-form>
                    </div> 
                </el-tab-pane>
                <!-- tab:硬件监控 -->
				<el-tab-pane 
					:label="$t('instance.detail.hardwareMonitoring')" 
					name="HardwareMonitoring"
					:disabled="tabName==='HardwareMonitoring'"
                >
                    <!-- 硬件设备状态 -->
                    <h3 class="section-title">{{ $t('instance.detail.hardwareState')}}</h3>
                    <hardware-status :hardWareData="hardWareData.tableData"></hardware-status>
                    <!-- 报警日志列表 -->
                    <h3 class="section-title">{{ $t('instance.detail.alarmLogList')}}</h3>
                    <alarm-log-list
                        ref="alarmLogListRef"
                        :instanceInfo="instanceInfo.detail"
                        @go-back="goBack"
                    >
                    </alarm-log-list>
                </el-tab-pane>
                <!-- tab:性能监控 -->
                <el-tab-pane
                    :label="$t('instance.detail.performanceMonitoring')" 
					name="PerformanceMonitoring"
					:disabled="tabName==='PerformanceMonitoring'">
                    <div v-if="tabName==='PerformanceMonitoring'">
                        <performance-monitoring />
                    </div>
                 
                </el-tab-pane>
                <!-- tab:操作日志 -->
                <el-tab-pane
                    :label="$t('instance.detail.operationLog')" 
					name="operationLog"
					:disabled="tabName==='operationLog'">
                    <!-- 实例信息 -->
                    <h3 class="section-title mb28">{{ $t('instance.detail.operationLog')}}
                        <span class="f14">{{$t('instance.detail.logTip')}}</span>
                    </h3>
                    <div class="operation-header">
                        <el-config-provider :locale="locale">
                            <div class="operation-search-content">
                                <div class="operation-search-items">
                                    <!-- 操作人搜索 -->
                                    <div class="operation-search-item">
                                        <label class="operation-name">{{$t('instance.detail.operator') + $filter.withClon(' ')}}</label>
                                        <el-input 
                                            v-model="operationLogReactiveArr.searchParams.userName" 
                                            maxlength="64"
                                            :placeholder="$t('instance.placeholder.operator')" 
                                        />
                                    </div>
                                    <!-- 操作时间搜索 -->
                                    <div class="operation-search-item">
                                        <label class="operation-name-date">{{$t('instance.detail.operationTime')+ $filter.withClon(' ')}}</label>
                                        <el-date-picker
                                            v-model="operationLogReactiveArr.operationTime"
                                            type="daterange"
                                            unlink-panels
                                            range-separator="-"
                                            :disabled-date="disabledDate"
                                            :start-placeholder="$t('instance.placeholder.startTime')"
                                            :end-placeholder="$t('instance.placeholder.endTime')"
                                        />
                                    </div>

                                    <div class="operation-search-button">
                                        <el-button 
                                            type="primary" 
                                            :loading="isOperationlogLoading" 
                                            @click="searchOperationLogData"
                                        >
                                            {{$t('list.search')}}
                                        </el-button>
                                        <el-button 
                                            type="primary" 
                                            :loading="isOperationlogLoading" 
                                            @click="clearOperationLogData"
                                        >
                                            {{$t('list.clear')}}
                                        </el-button>
                                    </div>
                                </div>   
                            </div>
                        </el-config-provider>
                        <div class="setting-content">
                            <div class="setting-position mr0">
                                <!-- 刷新 -->
                                <p
                                    class="operate-refresh"
                                    @click="refreshOperationLogData"
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
                    </div>  
                    <div 
                        class="table-content mlr0 mt16"
                        :class="[
                            operationLogReactiveArr.tableData.length ? '' : 'no-table-scroll',
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
                                :data="operationLogReactiveArr.tableData" 
                                v-loading="isOperationlogLoading"
                                :max-height="operationLogReactiveArr.tableData.length ? tableMaxHeight : 380" 
                                style="width: 100%"
                                @filter-change="filterChange"
                            >
                                <!-- LogID -->
                                <el-table-column     
                                    prop="logid" 
                                    v-if="operationLogReactiveArr?.hasCustomInfo['logid']?.selected"
                                    :label="$t('instance.detail.logId')"
                                    min-width=250
                                >
                                    <template v-slot="scope">
                                        {{scope.row.logid || $filter.emptyFilter()}}
                                    </template>
                                </el-table-column>
                                <!-- 操作名称 -->
                                <el-table-column     
                                    prop="operation" 
                                    key="operation"
                                    column-key="operation"
                                    v-if="operationLogReactiveArr?.hasCustomInfo['operationName']?.selected"
                                    :label="$t('instance.detail.operationName')"
                                    filter-placement="bottom-end"
                                    :filters="operationLogReactiveArr?.operationNameList"
                                    :filter-method="operationNameFilter"
                                    min-width=140
                                >
                                    <template v-slot="scope">
                                        <span>{{scope.row.operationName || $filter.emptyFilter()}}</span>
                                    </template>
                                </el-table-column>
                                <!-- 操作人 -->
                                <el-table-column     
                                    prop="userName" 
                                    v-if="operationLogReactiveArr?.hasCustomInfo['userName']?.selected"
                                    :label="$t('instance.detail.operator')" 
                                    min-width=120
                                >
                                    <template v-slot="scope">
                                        {{scope.row.userName || $filter.emptyFilter()}}
                                    </template>
                                </el-table-column>
                                <!-- 操作时间 -->
                                <el-table-column     
                                    prop="operateTime" 
                                    v-if="operationLogReactiveArr?.hasCustomInfo['operateTime']?.selected"
                                    :label="$t('instance.detail.operationTime')" 
                                    min-width=150
                                >
                                    <template v-slot="scope">
                                        {{getDateMinutes(scope.row.operateTime) || $filter.emptyFilter()}}
                                    </template>
                                </el-table-column>
                                <!-- 操作反馈 -->
                                <el-table-column     
                                    prop="result" 
                                    key="result"
                                    column-key="result"
                                    v-if="operationLogReactiveArr?.hasCustomInfo['result']?.selected"
                                    :label="$t('instance.detail.operationFeedback')" 
                                    filter-placement="bottom-end"
                                    :filters="operationLogStatus"
                                    :filter-method="operationFeedbackFilter"
                                    min-width=180
                                >
                                    <template v-slot="scope">
                                        <div :class="scope.row.result === 'doing' ? 'mr10' : ''">
                                            <div class="list-icon"  :class="scope.row.result === 'success' ? 'success-icon': scope.row.result === 'fail' ?'fail-icon':'doing'"></div>
                                            <span>{{operationLogStatusConstant[scope.row.result] || $filter.emptyFilter()}}</span>
                                        </div>
                                    </template>
                                </el-table-column>
                                <!-- 错误码 -->
                                <el-table-column     
                                    prop="failReason" 
                                    v-if="operationLogReactiveArr?.hasCustomInfo['failReason']?.selected"
                                    :label="$t('instance.detail.faultWord')" 
                                    min-width=180
                                >
                                    <template v-slot="scope">
                                        {{scope.row.failReason || $filter.emptyFilter()}}
                                    </template>
                                </el-table-column>
                                <template #empty>
                                    <div class="noData">                        
                                    </div>
                                </template>
                            </el-table>
                        </el-config-provider>
                        <div v-if="openCustom">
                            <custom
                                :page-key="'auditLogsList'"
                                :open-visible="openCustom"
                                :check-list-arr="operationLogReactiveArr.checkListArr"
                                :has-custom-info="operationLogReactiveArr.hasCustomInfo"
                                @close="openCustomCancel"
                                @determine="publicStore.customList('auditLogsList', operationLogReactiveArr)"
                            >
                            </custom>
                        </div>
                    </div>
                    <div class="footer-count">
                        <my-pagination
                            v-if="operationLogReactiveArr.tableData.length > 0 && !isOperationlogLoading"
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
                </el-tab-pane>
            </el-tabs>
            <div v-if="open">
                <instance-open
                    :openVisible="open"
                    :instanceInfo="instanceInfo.detail"
                    @close="openCancel"
                    @refresh="refreshData"
                >
                </instance-open>
            </div>
            <div v-if="close">
                <instance-close
                    :openVisible="close"
                    :instanceInfo="instanceInfo.detail"
                    @close="closeCancel"
                    @refresh="refreshData"
                >
                </instance-close>
            </div>
            <div v-if="restart">
                <instance-restart
                    :openVisible="restart"
                    :instanceInfo="instanceInfo.detail"
                    @close="restartCancel"
                    @refresh="refreshData"
                >
                </instance-restart>
            </div>
            <div v-if="resetPassword">
                <instance-reset-password
                    :openVisible="resetPassword"
                    :instanceInfo="[instanceInfo.detail]"
                    @close="resetPasswordCancel"
                    @refresh="refreshData"
                >
                </instance-reset-password>
            </div>
            <div v-if="resystemInstanceConfirm">
                <instance-resystem-confirm
                    :openVisible="resystemInstanceConfirm"
                    :instanceInfo="instanceInfo.detail"
                    @close="resystemInstanceConfirmCancel"
                    @refresh="refreshData"
                    @confirm="clickOperateResystem"
                >
                </instance-resystem-confirm>
            </div>
            <div v-if="resystemInstance">
                <instance-resystem
                    :openVisible="resystemInstance"
                    :instanceInfo="instanceInfo.detail"
                    @close="resystemInstanceCancel"
                    @refresh="refreshData"
                >
                </instance-resystem>
            </div>
            <div v-if="deleteInstance">
                <instance-delete
                    :openVisible="deleteInstance"
                    :instanceInfo="instanceInfo.detail"
                    @close="deleteInstanceCancel"
                    @refresh="goBack"
                >
                </instance-delete>
            </div>
            <div v-if="editInstanceName">
                <instance-name-edit
                    :openVisible="editInstanceName"
                    :instanceInfo="instanceInfo.detail"
                    @close="editInstanceNameCancel"
                    @refresh="refreshData"
                >
                </instance-name-edit>
            </div>
            <div v-if="editDescription">
                <description-edit
                    :openVisible="editDescription"
                    :instanceInfo="instanceInfo.detail"
                    @close="editDescriptionCancel"
                    @refresh="refreshData"
                >
                </description-edit>
            </div>
            <div v-if="lock">
                <instance-lock
                    :openVisible="lock"
                    :instanceInfo="instanceInfo.detail"
                    @close="lockCancel"
                    @refresh="refreshData"
                >
                </instance-lock> 
            </div>
            <div v-if="unlock">
                <instance-unlock
                    :openVisible="unlock"
                    :instanceInfo="instanceInfo.detail"
                    @close="unlockCancel"
                    @refresh="refreshData"
                >
                </instance-unlock> 
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
import instanceOpen from 'components/InstanceOperate/instanceOpen.vue';
import instanceClose from 'components/InstanceOperate/instanceClose.vue';
import instanceRestart from 'components/InstanceOperate/instanceRestart.vue';
import instanceResystemConfirm from 'components/InstanceOperate/instanceResystemConfirm.vue';
import instanceResystem from 'components/InstanceOperate/instanceResystem/instanceResystem.vue';
import instanceDelete from 'components/InstanceOperate/instanceDelete.vue';
import instanceResetPassword from 'components/InstanceOperate/instanceResetPassword.vue';
import descriptionEdit from 'components/InstanceOperate/editDescription.vue';
import instanceNameEdit from 'components/InstanceOperate/editName.vue';
import instanceLock from 'components/InstanceOperate/instanceLock.vue';
import instanceUnlock from 'components/InstanceOperate/instanceUnlock.vue';
import performanceMonitoring from 'components/common/performanceMonitoring/performanceMonitoring.vue';
import custom from 'components/custom/custom.vue';
import myPagination from 'components/Pagination/MyPagination.vue';
import {
    instanceDetailAPI, 
    hardWareStatusAPI, 
    alarmLogListAPI,
    operationLogListAPI,
    operationLogListExportAPI,
    operationLogTypeListAPI,
    getAgentStateAPI
} from 'api/request.ts';
import hardwareStatus from 'components/common/hardwareStatus/hardwareStatus.vue';
import alarmLogList from 'components/common/alarmLogList/alarmLogList.vue';
import allProjectStore from '@/store/modules/headerDetail.ts';
import { ElMessage, ElTable } from 'element-plus';
import {
    hasShowTooltip, // 是否显示提示
    languageSwitch, // 语言切换
    getDate, // 获取日期
    generateRandomNum, // 生成随机数
    mergeObjectsWithNonEmptyOverride, //对象合并
    convertStringsToTimestampObject, //转换时间为时间戳
    getDateMinutes
} from 'utils/index.ts';
import useClipboard from 'vue-clipboard3';
import {
    intermediate, // 中间态类型
    lockedStatus, // 锁定状态
    operationCustomInfo, // 操作列表自定义初始数据
    operationCustomIntro, // 操作列表自定义信息
    operationLogStatus, // 操作反馈状态
    operationLogStatusConstant
} from 'utils/constants.ts'; 
import i18n from 'lib/i18n/index.ts'; // 国际化
import publicIndexStore from 'store/index.ts'; // 公共store
import useProjectStore from '@/store/modules/projectId.ts';

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
 * 返回列表页
 */
const router = useRouter();
const projectStore = useProjectStore();
const goBack = () => {
    projectStore.setProject(route.query.projectId, route.query.projectName);
    router.push({
		path: `/instance/list`,
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
    datasDetail: any;
}

/**
 * 详情页信息
 */
const instanceInfo: ReactiveType = reactive<ReactiveType>({
    detail: {},
    datasDetail:[],
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
    'instanceNameInfo': {
        showTooltip: false
    },
    'instanceName': {
        showTooltip: false
    },
    'instanceId': {
        showTooltip: false
    },
    'deviceTypeName': {
        showTooltip: false
    },
    'imageName': {
        showTooltip: false
    },
    'description': {
        showTooltip: false
    },
    'cpuInfo': {
        showTooltip: false
    },
    'memInfo': {
        showTooltip: false
    },
    'systemInfo': {
        showTooltip: false
    },
    'dataInfo': {
        showTooltip: false
    },
    'nicInfo': {
        showTooltip: false
    },
    'gpuInfo': {
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
 * 开机操作打开标志
 */
const open: Ref<boolean> = ref<boolean>(false)

/**
 * 关机操作打开标志
 */
const close: Ref<boolean> = ref<boolean>(false)

/**
 * 重启操作打开标志
 */
const restart: Ref<boolean> = ref<boolean>(false)

/**
 * 重置密码操作打开标志
 */
const resetPassword: Ref<boolean> = ref<boolean>(false)

/**
 * 重装系统确认操作打开标志
 */
const resystemInstanceConfirm: Ref<boolean> = ref<boolean>(false)

/**
 * 重装系统操作打开标志
 */
const resystemInstance: Ref<boolean> = ref<boolean>(false)

/**
 * 删除操作打开标志
 */
const deleteInstance: Ref<boolean> = ref<boolean>(false)

/**
 * 实例名称编辑操作打开标志
 */
const editInstanceName: Ref<boolean> = ref<boolean>(false)  

/**
 * 描述编辑操作打开标志
 */
const editDescription: Ref<boolean> = ref<boolean>(false)  

/**
 * 锁定操作打开标志
 */
const lock: Ref<boolean> = ref<boolean>(false)

/**
 * 解锁操作打开标志
 */
const unlock: Ref<boolean> = ref<boolean>(false)

/**
 * 开机操作
 * @param value 
 */
const clickOperateOpen: () => void = () => {
    open.value = !open.value;
}

/**
 * 开机实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
const openCancel = (type: boolean): boolean => {
    return open.value = type;
};

/**
 * 关机操作
 * @param value 
 */
const clickOperateClose: () => void = () => {
    close.value = !close.value;
}

/**
 * 关机实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} close.value 弹窗关闭
 */
const closeCancel = (type: boolean): boolean => {
    return close.value = type;
};

/**
 *  重启操作
 * @param value 
 */
const clickOperateRestart: () => void = () => {
    restart.value = !restart.value;
}

/**
 * 重启实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} restart.value 弹窗关闭
 */
const restartCancel = (type: boolean): boolean => {
    return restart.value = type;
};

/**
 * 重置密码操作
 * @param value 
 */
const clickOperateResetPassword: () => void = () => {
    resetPassword.value = !resetPassword.value;
}

/**
 * 重置密码实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} resetPassword.value 弹窗关闭
 */
const resetPasswordCancel = (type: boolean): boolean => {
    return resetPassword.value = type;
};

/**
 * 重装系统确认操作
 * @param value 
 */
const clickOperateResystemConfirm: () => void = () => {
    resystemInstanceConfirm.value = !resystemInstanceConfirm.value;
}
/**
 * 重装系统确认实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} resystemInstanceConfirm.value 弹窗关闭
 */
const resystemInstanceConfirmCancel = (type: boolean): boolean => {
    return resystemInstanceConfirm.value = type;
};

/**
 * 重置系统操作
 * @param value 
 */
const clickOperateResystem: (value: object) => void = (value: object) => {
    resystemInstance.value = !resystemInstance.value;
}

/**
 * 重置系统弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} resystemInstance.value 弹窗关闭
 */
const resystemInstanceCancel = (type: boolean): boolean => {
    return resystemInstance.value = type;
};

/**
 * 删除操作
 * @param value 
 */
const clickOperateDelete: () => void = () => {
    deleteInstance.value = !deleteInstance.value;
}

/**
 * 删除实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} deleteInstance.value 弹窗关闭
 */
const deleteInstanceCancel = (type: boolean): boolean => {
    return deleteInstance.value = type;
};

/**
 * 实例名称编辑操作
 * @param value 
 */
const clickEditInstanceName: () => void = () => {
    editInstanceName.value = !editInstanceName.value;
}
/**
 * 实例名称编辑弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} editInstanceName.value 弹窗关闭
 */
const editInstanceNameCancel = (type: boolean): boolean => {
    return editInstanceName.value = type;
};

/**
 * 描述编辑操作
 * @param value 
 */
const clickEditDescription: () => void = () => {
    editDescription.value = !editDescription.value;
}
/**
 * 描述编辑弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} editDescription.value 弹窗关闭
 */
const editDescriptionCancel = (type: boolean): boolean => {
    return editDescription.value = type;
};


/**
 * 锁定操作
 * @param value 
 */
const clickOperateLock: () => void = () => {
        lock.value = true;
        unlock.value = false;       
}

/**
 * 解锁操作
 * @param value 
 */
const clickOperateUnlock: () => void = () => {
        lock.value = false;
        unlock.value = true;       
}
/**
 * 重启实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} lock.value 弹窗关闭
 */
const lockCancel = (type: boolean): boolean => {
    return lock.value = type;
};
/**
 * 重启实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} unlock.value 弹窗关闭
 */
const unlockCancel = (type: boolean): boolean => {
    return unlock.value = type;
};

/**
 * 状态tag变化
 * @param value 
 */
const stateChange = (value: string,) => {
    switch (value) {
        case 'running':
        case 'Running':
            return 'success';
        case 'stopped':
        case 'Stopped':
            return 'info';
        case 'Creation failed':
        case 'Startup failed':
        case 'Shutdown failed':
        case 'Reboot failed':
        case 'Delete failed':
        case 'Reinstall failed':
        case 'Resetpasswd failed':  
        case 'destroyed':
        case 'Destroyed':
            return 'danger';   
        default:
            return 'intermediate';
    }
}

let timer: null | number = null;
let isIntervalRequest: boolean = false;
onMounted(() => clearTimeout((timer as number)));
/**
 * 页面销毁时-触发，清空延时器
*/
onUnmounted(() => clearTimeout((timer as number)));
onBeforeUnmount(() => clearTimeout((timer as number)));

/**
 * 判断硬件监控授权
 */
const hasHardwareMonitoring: Ref<boolean> = ref<boolean>(false);
/**
 * 判断性能监控授权
 */
const hasPerformanceMonitoring: Ref<boolean> = ref<boolean>(false);   
/**
 * 请求实例详情数据接口
*/
const requestInstanceDetailData = (): void => {
    instanceDetailAPI({
        instanceId: route.query.instanceId
    }).then(({data} : {data: any}) => {
        if (data?.result?.instance) {
            instanceInfo.detail = data.result.instance;
            instanceInfo.datasDetail = data.result.instance.volumeRaid.filter((item: any) => item.volume.volumeType === 'data');
            if ([...intermediate].some((item: string) => item === instanceInfo.detail.status)) {
                timer = setTimeout(() => {
                    requestInstanceDetailData();
                    isIntervalRequest = true;
                }, 5000)
                return;
            }
            isIntervalRequest = false;
            return;
        }
        return Promise.reject();
    }).catch(({message} : {message: string;}) => {      
        if (message === '找不到对象' || message === 'Not found') { 
            goBack();
            return;
        }
        if (message.indexOf('undefined') > -1) return;
        if (!isIntervalRequest) {
            typeof message === 'string' && ElMessage.error(message);
            instanceInfo.detail = {};
            return;
        }
        timer = setTimeout(() => {
            requestInstanceDetailData();
        }, 5000)
    }).finally(() => {
        isLoading.value = false;
        if(!isIntervalRequest) {
            store.requestObject();
            store.requestUser();
            //sessionStorage.setItem('userName', store.userForm.userName);
       
            hasHardwareMonitoring.value = store.monitor; 
            hasPerformanceMonitoring.value = store.inMonitor; 
            //userName.value = sessionStorage.getItem('userName'); 
            (instanceMitt as any).proxy.$Bus.emit('project-name', route.query.projectName);
        }     
    });
};

/**
 * 监控agent状态
 */
const monitorAgent: Ref<string> = ref<string>('')
const getAgentState = (): void => {
    getAgentStateAPI({
        instanceId: route.query.instanceId
    }).then(({data} : {data: any}) => {
        if(data?.result?.agentStatus) {    
            monitorAgent.value = data?.result?.agentStatus[0].statusName
        }
    }).catch(({message} : {message: string;}) => {      

    }).finally(() => {
   
    });
}


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
 * 报警日志列表ref
 */
const alarmLogListRef = ref<any>();

const monitorChartRef = ref<any>();

/**
 * 改变标签页动作
 */
const changeTab = (): void => {
    if(tabName.value === 'HardwareMonitoring') {
        requestHardWareStatusData(); //获取硬件设备状态
        alarmLogListRef.value.refreshAlarmLoglistData() // 获取报警日志列表
    }

    else if(tabName.value === 'operationLog') {
        publicStore.customList('auditLogsList', operationLogReactiveArr);
        getOperationLogListData();
        refreshOperationLogData();
    }
}

/**
 * 硬件设备状态item类型
 */
type HardwareType = {
    name: string;
    status: number;
};

/**
 * 硬件设备状态类型
 */
interface hardware {
    tableData: HardwareType[]
}

/**
 * 硬件设备状态列表数据
 */
const hardWareData: hardware = reactive<hardware>({
    tableData: []
})

/**
 * 请求硬件设备状态接口
 */
const requestHardWareStatusData = (): void => {
    hardWareData.tableData = [];
    hardWareStatusAPI({
        sn: instanceInfo.detail.sn,
        idcId: instanceInfo.detail.idcId
    }).then(({data} : {data: any}) => {
        let status = data.result.detail[0]
        hardWareData.tableData.push(
            {
                name: 'CPU',
                status: status.cpu_status
            },
            {
                name: global.t('instance.detail.memory'),
                status: status.mem_status
            },
            {
                name: global.t('instance.detail.hardDisk'),
                status: status.disk_status
            },
            {
                name: global.t('instance.detail.networkCard'),
                status: status.nic_status
            },
            {
                name: global.t('instance.detail.power'),
                status: status.power_status
            },
            {
                name: global.t('instance.detail.other'),
                status: status.other_status
            },
        )
    }).catch(({message} : {message: string;}) => {      
        if (message === '找不到对象' || message === 'Not found') { 
            goBack();
            return;
        }
        if (message.indexOf('undefined') > -1) return;
    }).finally(() => {

    });
};

onMounted(() => {
    tabName.value = 'basicInfo';
    requestInstanceDetailData();
    getAgentState();
    cookie?.set && cookie.set('X-Jdcloud-Username', store.userForm.userName);
    if(route.query.monitor === 'PerformanceMonitoring') {
        tabName.value = 'PerformanceMonitoring'
    }
});

/**
 * 刷新接口
*/
const refreshData = (): void => {
    isLoading.value = true;
    clearTimeout(timer as number)
    requestInstanceDetailData();
};

/**
 * 语言
*/
const locale: any = languageSwitch();

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();

/**
 * 操作item类型
 */
 type OperationsType = {
    logid: string,
    operation: string,
    userName: string,
    operationTime: string,
    result: string,
    failReason: string,
}

/**
 * 项目类型
 */
 interface operation {
    tableData: OperationsType[],
    operationNameList: any,
    filterParams: any,
    searchParams: any,
    operationTime: any,
    hasCustomInfo: {},
    checkListArr: {},
}

/**
 * 项目列表数据
 */
 const operationLogReactiveArr: operation = reactive<operation>({
    tableData: [],
    operationNameList: [],
    filterParams: {
        operation: '',
        result: ''
    },
    searchParams: {
        userName: '',
        
    },
    operationTime: [],
    hasCustomInfo: operationCustomIntro,
    checkListArr: operationCustomInfo,
})

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
 * 操作日志loadig态
*/
const isOperationlogLoading: Ref<boolean> = ref<boolean>(true);

/**
 * 禁用的日期
 */
const disabledDate = (time: Date) => {
    return time.getTime() > Date.now() || time.getTime() < Date.now() - 3600 * 1000 * 24 * 90;
}

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
 * 刷新接口
*/
const refreshOperationLogData = (): void => {
    isOperationlogLoading.value = true;
    pageNumber.value = 1;
    pageSize.value = 20;
    processingParameter();
};

/**
 * 搜索结果
*/
const searchOperationLogData = (): void => {
    isOperationlogLoading.value = true;
    pageNumber.value = 1;
    pageSize.value = 20;
    processingParameter();
    if(operationLogReactiveArr.searchParams['userName'] || operationLogReactiveArr.operationTime.length) {
        searchResultMark.value = true;
    }
    
};

/**
 * 搜索清空
*/
const clearOperationLogData = (): void => {
    isOperationlogLoading.value = true;
    pageNumber.value = 1;
    pageSize.value = 20;
    operationLogReactiveArr.searchParams ={
        userName: '',
    };
    operationLogReactiveArr.operationTime = [];
    processingParameter();
    if(!operationLogReactiveArr.filterParams['operation'] && !operationLogReactiveArr.filterParams['result'] && !operationLogReactiveArr.searchParams['userName'] && !operationLogReactiveArr.operationTime.length) {
        searchResultMark.value = false
    }
};

/**
 * 表格ref
*/
const tableRef = ref<InstanceType<typeof ElTable>>()

/**
 * 返回列表
 */
 const backToList = () =>{
    
    operationLogReactiveArr.filterParams['operation'] = '';
    operationLogReactiveArr.filterParams['result'] = '';
    tableRef.value!.clearFilter();
    clearOperationLogData();
    searchResultMark.value = false
}

watch(() =>  operationLogReactiveArr.filterParams, (): void => {
    if(!operationLogReactiveArr.filterParams['operation'] && !operationLogReactiveArr.filterParams['result'] && !operationLogReactiveArr.searchParams['userName'] && !operationLogReactiveArr.operationTime.length) {
        searchResultMark.value = false
    }
},{deep:true});
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
        operationLogListExportAPI(
            {
                exportType: '1',
                instanceId: instanceInfo.detail.instanceId,
                isAll: '1'
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
            link.download = `project_list_${getDate()}_${generateRandomNum(6)}`;
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

// 计算表格的最大高度
const tableMaxHeight: Ref<number> = ref<number>(window.innerHeight - 480)
// 监听窗口大小变化，以便动态调整表格高度
const updateSize = () => {
  // 触发响应式更新
  tableMaxHeight.value = window.innerHeight - 480
};

onMounted(() => {
  window.addEventListener('resize', updateSize);
});

onUnmounted(() => {
  window.removeEventListener('resize', updateSize);
});

/**
 * 每页展示条数切换
 * @param {number} size 每页展示条数
*/
const pageSizesChange = (size: number) => {
    pageNumber.value = 1;
    pageSize.value = size;
    isOperationlogLoading.value = true;
    processingParameter();
};

/**
 * 页数切换
 * @param {number} num 页数切换
*/
const changePage = (num: number): void => {
    pageNumber.value = num;
    isOperationlogLoading.value = true;
    processingParameter();
};



/**
 * 处理参数
*/
const processingParameter = () => {
    let timeObject
    if(operationLogReactiveArr.operationTime === null) {
        operationLogReactiveArr.operationTime =[]
    }
    if(operationLogReactiveArr.operationTime.length) {
        timeObject = convertStringsToTimestampObject(operationLogReactiveArr.operationTime, 'startTime', 'endTime'); 
        timeObject['endTime'] = timeObject['endTime'] + (24 * 60 * 60 - 1);
    }
    
    const params = {
        pageSize: pageSize.value,
        pageNumber: pageNumber.value,
        instanceId: instanceInfo.detail.instanceId,
        ...operationLogReactiveArr.filterParams,
        ...timeObject,
        ...operationLogReactiveArr.searchParams
    };
    publicStore.deleteEmtpyData(params);
    requestOperationLogListData(params);
};

/**
 * 搜索结果提示
 */
 const searchResultMark: Ref<boolean> = ref<boolean>(false)

/**
 * 请求操作日志列表数据接口
*/
const requestOperationLogListData = (params: any): void => {
    operationLogListAPI({
        ...params
    }).then(({data} : {data: any}) => {
        if(data?.result?.messages?.length) {
            operationLogReactiveArr.tableData = data.result.messages.map((item: any) => {
                return {
                    ...item,
                    showTooltip: false
                }
            });
            totalCount.value = data.result.totalCount;
            return;
    }
    operationLogReactiveArr.tableData  = [];
    totalCount.value = 0;
    }).catch (({message} : {message: string;})=>{
        if (message.indexOf('undefined') > -1) return;
        operationLogReactiveArr.tableData  = [];
        totalCount.value = 0;
    })
    .finally(() => {
        isOperationlogLoading.value = false;
    });
};

/**
 * filter参数类
*/
interface FilterParamsType {
    operation: string,
    result: string;
};

const filterChange = (filter: {[x: string]: any;}) => {
    const filterParams: FilterParamsType = {
        operation: 'operation',
        result: 'result',
    };
    const tag: string[] = [];
    for (const key in filter) {
        if (key === filterParams[key]) {
            for (const item of filter[key]) {
                tag.push(item);    
            }
            operationLogReactiveArr.filterParams[filterParams[key]] = tag.join(',');
        }
    }
    refreshOperationLogData();
    searchResultMark.value = true 
};

/**
 * 处理筛选数据
 */
 const processingData = (data: any) => {
    let arrData: any = []
    data.map((item: any) => {
        arrData.push({
            filterParams: item.operation,
            text: item.name,
            value: item.operation
        })
        
    });
    return arrData
}
/**
 * 获取操作日志筛选下拉框列表
 */
 const getOperationLogListData = (): void => {
    operationLogTypeListAPI({
    }).then(({data} : {data: any}) => {
        if (data?.result) {
            operationLogReactiveArr.operationNameList = processingData(data.result);
            return;
        }
        operationLogReactiveArr.operationNameList = [];
 
    }).catch (()=>{
        operationLogReactiveArr.operationNameList = [];
    })
    .finally(() => {

    });
};

const operationNameFilter = (value: string, row: any): boolean => {
    return row.operation === value;
};

const operationFeedbackFilter = (value: string, row: any): boolean => {  
    return row.result === value;
};

</script>

<style lang="scss">
@import 'assets/css/page';
@import './instanceDetail.scss';
@import 'assets/css/icon';
</style>