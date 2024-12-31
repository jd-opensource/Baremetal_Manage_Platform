<template>
    <div class="page-position instance-list">
        <div class="page-content pos-fixed">
            <div class="operation-header">
                <el-button 
                    type="primary" 
                    class="add-button-style" 
                    :class="getLocationItem === 'zh_CN' ? '' : 'w140'" 
                    @click="jumpToCreate"
                >
                     + &nbsp;&nbsp;{{$t('instance.list.create') }}
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
                                $t('instance.placeholder.instanceName'),
                                $t('instance.placeholder.instanceId'),
                                $t('instance.placeholder.outBandIp'),
                                $t('instance.placeholder.intranetIpv4'),
                                $t('instance.placeholder.intranetIpv6')
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
                    reactiveArr.tableData.length ? '' : 'no-table-scroll',
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
                        v-loading="isLoading"
                        style="width: 100%"
                        :max-height="reactiveArr.tableData.length ? tableMaxHeight : 380" 
                        @row-click="rowClick"
                        @filter-change="filterChange"
                        @selection-change="handleSelectionChange"
                    >
                        <!-- 选择多行 -->
                        <el-table-column type="selection" width="55" />
                        <!-- 实例名称/ID -->
                        <el-table-column 
                            fixed
                            prop="instanceName" 
                            v-if="reactiveArr?.hasCustomInfo['instanceName']?.selected"
                            :label="$t('instance.list.instanceName')" 
                            width = 170
                        >
                            <template v-slot="scope">
                                <div class="instance-name-height">
                                    <el-tooltip
                                        v-model="scope.row.tooltipStatus['instanceName'].showTooltip"
                                        :disabled="!scope.row.tooltipStatus['instanceName'].showTooltip"
                                        :content= scope.row.instanceName
                                        placement="right"
                                        :show-after="500"
                                    >
                                        <div
                                            class="long-instance-name"
                                            @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['instanceName'], scope.row.instanceName, 1.17, 'list')"
                                        >
                                            <span @click.stop="jumpToDetail(scope.row.instanceId)">
                                                {{scope.row.instanceName || $filter.emptyFilter()}}
                                            </span>   
                                        </div>                                   
                                    </el-tooltip>
                                    <el-tooltip
                                        v-model="scope.row.tooltipStatus['instanceId'].showTooltip"
                                        :disabled="!scope.row.tooltipStatus['instanceId'].showTooltip"
                                        :content= scope.row.instanceId
                                        placement="right"
                                        :show-after="500"
                                    >
                                        <div
                                            class="long-row"
                                            @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['instanceId'], scope.row.instanceId, 1.17, 'list')"
                                        >
                                            <span>{{scope.row.instanceId}}</span>
                                        </div>
                                    </el-tooltip>
                                </div>    
                            </template>
                        </el-table-column>
                        <!-- 监控 -->
                        <el-table-column 
                            prop="monitor" 
                            :label="$t('alarm.monitor.monitor')" 
                            min-width=80
                        >
                            <template v-slot="scope">
                                <el-tooltip :content= "$t('alarm.monitor.viewMonitor')" placement="right">
                                    <div class="monitor instance-list-icon" @click.stop="jumpToDetailMonitor(scope.row.instanceId)"></div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 状态 -->
                        <el-table-column 
                            prop="status" 
                            key="status"
                            column-key="status"
                            v-if="reactiveArr?.hasCustomInfo['status']?.selected"
                            :label="$t('instance.list.state')" 
                            filter-placement="bottom-end"
                            :filters="ossStore.status"
                            :filter-method="statusFilter"
                            :min-width="getLocationItem === 'zh_CN' ? 160 : 190"
                        >
                            <template v-slot="scope">
                                <el-tag
                                    effect="plain"
                                    :class="getLocationItem === 'zh_CN' ? 'mw78' : 'mw114'"
                                    class="tag-content"
                                    :type="stateChange(scope.row.status)"
                                >
                                    {{scope.row.statusName || $filter.emptyFilter()}} 
                                    <el-icon 
                                        v-if="[...intermediate].some((item: string) => item === scope.row.status)" 
                                        class="is-loading"
                                    >
                                        <Loading />
                                    </el-icon>
                                </el-tag>
                                <el-tooltip
                                    placement="right"
                                    v-if="scope?.row?.status === 'Shutdown failed' 
                                    || scope?.row?.status === 'Startup failed' 
                                    || scope?.row?.status === 'Reboot failed' 
                                    || scope?.row?.status === 'Reinstall failed'
                                    || scope?.row?.status === 'Resetpasswd failed'
                                    || scope?.row?.status === 'Delete failed'"
                                >
                                    <template #content>
                                        <div>
                                            {{$t('list.tip.statusError')}}
                                        </div>
                                    </template>
                                    <img
                                        style="marginLeft: 10px; marginTop: 6px;"
                                        class="warning-tip"
                                        src="@/assets/img/failedTip.png"
                                        alt=""
                                    />
                                </el-tooltip>
                                <el-tooltip
                                    placement="right"
                                    v-if="scope?.row?.reason && scope?.row?.status === 'Creation failed'"
                                >
                                    <template #content>
                                        <div>
                                            {{scope?.row?.reason}}
                                        </div>
                                    </template>
                                    <img
                                        style="marginLeft: 10px; marginTop: 6px;"
                                        class="warning-tip"
                                        src="@/assets/img/failedTip.png"
                                        alt=""
                                    />
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 机型 -->
                        <el-table-column     
                            prop="deviceTypeName" 
                            key="deviceTypeName"
                            column-key="deviceTypeId"
                            v-if="reactiveArr?.hasCustomInfo['deviceTypeName']?.selected"
                            :label="$t('instance.list.model')"
                            filter-placement="bottom-end"
                            :filters="reactiveArr.deviceSeriesList"
                            :filter-method="deviceSeriesFilter"
                            min-width=140
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['deviceTypeName'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['deviceTypeName'].showTooltip"
                                    :content= scope.row.deviceTypeName
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['deviceTypeName'], scope.row.deviceTypeName, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.deviceTypeName || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 镜像 -->
                        <el-table-column 
                            prop="imageName" 
                            v-if="reactiveArr?.hasCustomInfo['imageName']?.selected"
                            :label="$t('instance.list.mirror')" 
                            min-width=150
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['imageName'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['imageName'].showTooltip"
                                    :content= scope.row.imageName
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['imageName'], scope.row.imageName, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.imageName || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 带外IP -->
                        <el-table-column     
                            prop="iloIp" 
                            v-if="reactiveArr?.hasCustomInfo['iloIp']?.selected"
                            :label="$t('instance.list.outBandIp')"
                            min-width=130
                        >
                            <template v-slot="scope">
                                {{scope.row.iloIp || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 内网IPv4(eth0) -->
                        <el-table-column     
                            prop="privateIpv4" 
                            v-if="reactiveArr?.hasCustomInfo['privateIpv4']?.selected"
                            :label="$t('instance.list.intranetIpv4')" 
                            min-width=120
                        >
                            <template v-slot="scope">
                                {{scope.row.privateIpv4 || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 内网IPv4(eth1) -->
                        <el-table-column     
                            prop="privateEth1Ipv4" 
                            v-if="reactiveArr?.hasCustomInfo['privateEth1Ipv4']?.selected"
                            :label="$t('instance.list.intranetIpv4eth1')" 
                            min-width=120
                        >
                            <template v-slot="scope">
                                {{scope.row.privateEth1Ipv4 || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 内网IPv6(eth0) -->
                        <el-table-column     
                            prop="privateIpv6" 
                            v-if="reactiveArr?.hasCustomInfo['privateIpv6']?.selected"
                            :label="$t('instance.list.intranetIpv6')" 
                            min-width=180
                        >
                            <template v-slot="scope">
                                {{scope.row.privateIpv6 || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 内网IPv6(eth1) -->
                        <el-table-column     
                            prop="privateEth1Ipv6" 
                            v-if="reactiveArr?.hasCustomInfo['privateEth1Ipv6']?.selected"
                            :label="$t('instance.list.intranetIpv6eth1')" 
                            min-width=180
                        >
                            <template v-slot="scope">
                                {{scope.row.privateEth1Ipv6 || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 配置 -->
                        <el-table-column     
                            prop="config" 
                            v-if="reactiveArr?.hasCustomInfo['config']?.selected"
                            :label="$t('instance.list.config')" 
                            min-width=275
                        >
                            <template v-slot="scope">
                                <p>{{$t('instance.detail.cpu') + $filter.withClon('')  + scope.row.cpuInfo || $filter.emptyFilter()}}</p>
                                <p>{{$t('instance.detail.memory') + $filter.withClon('') + scope.row.memInfo || $filter.emptyFilter()}}</p>
                            </template>
                        </el-table-column>
                        <!-- 创建时间 -->
                        <el-table-column     
                            prop="createdTime" 
                            v-if="reactiveArr?.hasCustomInfo['createdTime']?.selected"
                            :label="$t('instance.list.createdTime')"
                            min-width=160
                        >
                            <template v-slot="scope">
                                {{scope.row.createdTime || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 操作 -->
                        <el-table-column 
                            fixed='right'
                            prop="operation" 
                            :label="$t('personCentre.form.operation')" width="215">
                            <template v-slot="scope"
                        >
                                <operate-unit
                                    v-if="['stopped'].includes(scope.row.status)"
                                    :operateName="$t('list.operate.open')"
                                    :operateTip="$t('instance.tip.open')"
                                    :disabled="checkOperationDisable(scope.row.status,'stopped')"
                                    :instanceInfo="scope.row"
                                    @handleClick="clickOperateOpen(scope.row)"
                                >
                                </operate-unit>
                                <operate-unit
                                    v-else
                                    :operateName="$t('list.operate.close')"
                                    :disabled="checkOperationDisable(scope.row.status,'running')"
                                    @handleClick="clickOperateClose(scope.row)"
                                    :instanceInfo="scope.row"
                                    :operateTip="$t('instance.tip.close')"
                                >
                                </operate-unit> 
                                <operate-unit
                                    :operateName="$t('list.operate.restart')"
                                    :disabled="checkOperationDisable(scope.row.status,'running')"
                                    @handleClick="clickOperateRestart(scope.row)"
                                    :instanceInfo="scope.row"
                                    :operateTip="$t('instance.tip.restart')"
                                >
                                </operate-unit> 
                                <!-- 更多下拉菜单 -->
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
                            <div class="no-data-list">                        
                            </div>
                        </template>
                    </el-table>
                    <p 
                        v-if="!reactiveArr.tableData.length && !isLoading"
                        class="no-data-jump"
                    >
                        <span class="c333 f12">{{$t('instance.list.noData')}}</span>
                        <el-button type="text" class="f12 pt7" @click="jumpToCreate">{{$t('instance.list.create') }}</el-button>
                    </p>
                </el-config-provider>
                <div v-if="openCustom">
                    <custom
                        :page-key="'instanceList'"
                        :open-visible="openCustom"
                        :check-list-arr="reactiveArr.checkListArr"
                        :has-custom-info="reactiveArr.hasCustomInfo"
                        @close="openCustomCancel"
                        @determine="publicStore.customList('instanceList', reactiveArr)"
                    >
                    </custom>
                </div>
                <div v-if="open">
                    <instance-open
                        :openVisible="open"
                        :instanceInfo="instanceInfo"
                        @close="openCancel"
                        @refresh="refreshData"
                    >
                    </instance-open>
                </div>
                <div v-if="close">
                    <instance-close
                        :openVisible="close"
                        :instanceInfo="instanceInfo"
                        @close="closeCancel"
                        @refresh="refreshData"
                    >
                    </instance-close>
                </div>
                <div v-if="restart">
                    <instance-restart
                        :openVisible="restart"
                        :instanceInfo="instanceInfo"
                        @close="restartCancel"
                        @refresh="refreshData"
                    >
                    </instance-restart>
                </div>
                <div v-if="lock">
                    <instance-lock
                        :openVisible="lock"
                        :instanceInfo="instanceInfo"
                        @close="lockCancel"
                        @refresh="refreshData"
                    >
                    </instance-lock> 
                </div>
                <div v-if="unlock">
                    <instance-unlock
                        :openVisible="unlock"
                        :instanceInfo="instanceInfo"
                        @close="unlockCancel"
                        @refresh="refreshData"
                    >
                    </instance-unlock> 
                </div> 
                <div v-if="resetPasswordInstance">
                    <instance-reset-password
                        :openVisible="resetPasswordInstance"
                        :instanceInfo="instanceInfo"
                        @close="resetPasswordInstanceCancel"
                        @refresh="refreshData"
                    >
                    </instance-reset-password>
                </div> 
                <div v-if="resystemInstanceConfirm">
                    <instance-resystem-confirm
                        :openVisible="resystemInstanceConfirm"
                        :instanceInfo="instanceInfo"
                        @close="resystemInstanceConfirmCancel"
                        @refresh="refreshData"
                        @confirm="clickOperateResystem"
                    >
                    </instance-resystem-confirm>
                </div>
                <div v-if="resystemInstance">
                    <instance-resystem
                        :openVisible="resystemInstance"
                        :instanceInfo="instanceInfo"
                        @close="resystemInstanceCancel"
                        @refresh="refreshData"
                    >
                    </instance-resystem>
                </div>
                <div v-if="deleteInstance">
                    <instance-delete
                        :openVisible="deleteInstance"
                        :instanceInfo="instanceInfo"
                        @close="deleteInstanceCancel"
                        @refresh="refreshData"
                    >
                    </instance-delete>
                </div>
            </div>
            <div class="footer-count">
                <div class="batch-operate">
                    <el-checkbox @change="handleCheckAllChange" v-model="hasCheckAll" class="mr20"></el-checkbox>
                    <!-- 批量关机 -->
                    <el-tooltip 
                        placement="bottom" 
                        :content= "multipleSelection.length ? $t('instance.tip.close') : $t('instance.tip.chooseInstance')"
                        :disabled="isAllowBatchClose"
                    >
                        <span>
                            <el-button 
                                type="primary" 
                                :disabled="!isAllowBatchClose"    
                                @click="batchCloseDialog"
                            >
                                {{$t('list.operate.close')}}
                            </el-button>
                        </span>
                    </el-tooltip>
                    <!-- 批量开机 -->
                    <el-tooltip 
                        placement="bottom"
                        :content="multipleSelection.length ? $t('instance.tip.open') : $t('instance.tip.chooseInstance')"
                        :disabled="isAllowBatchOpen">
                        <span>
                            <el-button 
                                type="primary" 
                                :disabled="!isAllowBatchOpen"
                                @click="batchOpenDialog"
                            >
                                {{$t('list.operate.open')}}
                            </el-button>
                         </span>
                    </el-tooltip>
                    <!-- 更多下拉菜单 -->
                    <div class="batch-drop-down">
                        <batch-dropdown
                            :operations="batchOperations"
                            :instanceInfo="multipleSelection"
                            @clickSelectItem="clickBatchOperate($event)"
                        >
                        </batch-dropdown>    
                    </div>
                </div>
                <div class="page-change">
                    <my-pagination
                        v-if="reactiveArr.tableData.length > 0"
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
                <!--批量关机-->
                <batch-close 
                    :openVisible="multipleClose"
                    :instanceInfo="multipleSelection"
                    @close="batchCloseCancel"
                    @refresh="refreshData"
                >		
                </batch-close>
                <!--批量开机-->
                <batch-open
                    :openVisible="multipleOpen"
                    :instanceInfo="multipleSelection"
                    @close="batchOpenCancel"
                    @refresh="refreshData"
                >		
                </batch-open>
                <!--批量重启-->
                <batch-restart
                    :openVisible="multipleRestart"
                    :instanceInfo="multipleSelection"
                    @close="batchRestartCancel"
                    @refresh="refreshData"
                >		
                </batch-restart>
                <!--批量编辑实例名称-->
                <batch-rename
                    :openVisible="multipleRename"
                    :instanceInfo="multipleSelection"
                    :allInstance="reactiveArr.tableData"
                    @close="batchRenameCancel"
                    @refresh="refreshData"
                >		
                </batch-rename>
                <!--批量重置密码名称-->
                <batch-repassword
                    :openVisible="multipleResetPassword"
                    :instanceInfo="multipleSelection"
                    :allInstance="reactiveArr.tableData"
                    @close="batchResetPasswordCancel"
                    @refresh="refreshData"
                >		
                </batch-repassword>
                <!--批量删除-->
                <batch-delete
                    :openVisible="multipleDelete"
                    :instanceInfo="multipleSelection"
                    @close="batchDeleteCancel"
                    @refresh="refreshData"
                >		
                </batch-delete>
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
import {
    getDate, // 获取日期
    generateRandomNum, // 生成随机数
    filterData, // 过滤重复数据
    languageSwitch, // 语言切换
    deepCopy, //深拷贝
    hasShowTooltip // 是否显示提示
} from 'utils/index.ts';
//import useCurrentInstance from "hooks/useCurrentInstance.ts";
import operateUnit from 'components/OperateUnit/operateUnit.vue';
import dropDownOperation from 'components/DropdownOperation/DropdownOperation.vue';
import instanceOpen from 'components/InstanceOperate/instanceOpen.vue';
import instanceClose from 'components/InstanceOperate/instanceClose.vue';
import instanceRestart from 'components/InstanceOperate/instanceRestart.vue';
import instanceLock from 'components/InstanceOperate/instanceLock.vue';
import instanceUnlock from 'components/InstanceOperate/instanceUnlock.vue';
import instanceResetPassword from 'components/InstanceOperate/instanceResetPassword.vue';
import instanceResystemConfirm from 'components/InstanceOperate/instanceResystemConfirm.vue';
import instanceResystem from 'components/InstanceOperate/instanceResystem/instanceResystem.vue';
import instanceDelete from 'components/InstanceOperate/instanceDelete.vue';
import batchDropdown from 'components/multipleOperate/batchDropdown.vue';
import batchClose from 'components/multipleOperate/batchClose.vue';
import batchOpen from 'components/multipleOperate/batchOpen.vue';
import batchRestart from 'components/multipleOperate/batchRestart.vue';
import batchRename from 'components/multipleOperate/batchRename.vue';
import batchRepassword from 'components/multipleOperate/batchRepassword.vue';
import batchDelete from 'components/multipleOperate/batchDelete.vue';
import searchInput from 'components/SearchInput/searchInput.vue';
import custom from 'components/custom/custom.vue';
import myPagination from 'components/Pagination/MyPagination.vue';
import allProjectStore from '@/store/modules/headerDetail.ts';
import modelDetailStore from 'store/modules/modelDetail.ts'; // store库里的oss数据
import {intermediate} from 'utils/constants.ts'; // 中间态类型
import {instanceListAPI, instanceListExportAPI} from 'api/request.ts'; // 实例列表数据接口
import { ElMessage } from 'element-plus';
import {
    instanceCustomInfo, // 实例列表自定义初始数据
    instanceCustomIntro, // 项目列表自定义信息
} from 'utils/constants.ts';
import publicIndexStore from 'store/index.ts'; // 公共store
import i18n from 'lib/i18n/index.ts'; // 国际化
import { ElTable } from 'element-plus'
/**
 * 当前this
 */
 //const { globalProperties } = useCurrentInstance();

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
 * 国际化
*/
const {global} = i18n;

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();

/**
 * store库存储的oss数据类
*/
const ossStore: any = modelDetailStore();

/**
 * 语言
*/
const locale: any = languageSwitch();

/**
 * 路由带值
 */
const route = useRoute();

/**
 * 选择不同项目获取数据
 */
watch(() => route.query.projectId, (): Readonly<void> => { 
    if(route.query.projectId) {
        refreshData() 
    }
   
});

/**
 * store库存储的项目列表
*/
const store = allProjectStore();

/**
 * 实例item类型
 */
type InstancesType = {
    name: string,
    instanceId: string,
    status: string,
    model: string,
    osType: string,
    outBandIp: string,
    intranetIpv4: string,
    locked: string,
    deviceTypeName: string,
    deviceTypeId: string,
    rowChange: boolean,
    operations: any
}

/**
 * 项目类型
 */
interface instance {
    tableData: InstancesType[],
    multipleSelectionInstanceInfo: {},
    deviceSeriesList: any,
    filterParams: any,
    hasCustomInfo: {},
    checkListArr: {},
}

/**
 * 项目列表数据
 */
const reactiveArr: instance = reactive<instance>({
    tableData: [],
    multipleSelectionInstanceInfo: {},
    deviceSeriesList: [],
    filterParams: {
        deviceTypeId: '',
        status: ''
    },
    hasCustomInfo: instanceCustomIntro,
    checkListArr: instanceCustomInfo,
})

/**
 * 设置类
*/
type SetType<T> = T extends {} ? any : T;

/**
 * 使用mitt传值
*/
const instanceMitt: Exclude<Required<SetType<{}>> | null, never> = getCurrentInstance();

/**
 * 当前行值
 */
const instanceInfo: Ref<any> = ref();
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
 * loadig态
*/
const isLoading: Ref<boolean> = ref<boolean>(true);

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: any = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * 自定义列表操作打开标志
 */
const openCustom: Ref<boolean> = ref<boolean>(false)

/**
 * 搜索结果提示
 */
const searchResultMark: Ref<boolean> = ref<boolean>(false)

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
 * 锁定操作打开标志
 */
const lock: Ref<boolean> = ref<boolean>(false)

/**
 * 解锁操作打开标志
 */
const unlock: Ref<boolean> = ref<boolean>(false)

/**
 * 重置密码操作打开标志
 */
const resetPasswordInstance: Ref<boolean> = ref<boolean>(false)

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
 * 自定义列表接口
*/
publicStore.customList('instanceList', reactiveArr);

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
 * 开机操作
 * @param value 
 */
const clickOperateOpen: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
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
const clickOperateClose: (value: object) => void = (value: object) => {
    instanceInfo.value = value;
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
const clickOperateRestart: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
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
 *  锁定操作
 * @param value 
 */
const clickOperateLock: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    lock.value = !lock.value;
}

/**
 * 锁定实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} lock.value 弹窗关闭
 */
const lockCancel = (type: boolean): boolean => {
    return lock.value = type;
};

/**
 *  解锁操作
 * @param value 
 */
const clickOperateUnlock: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    unlock.value = !unlock.value;
}

/**
 * 解锁实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} unlock.value 弹窗关闭
 */
const unlockCancel = (type: boolean): boolean => {
    return unlock.value = type;
};

/**
 * 重置密码操作
 * @param value 
 */
const clickOperateResetPassword: (value: object) => void = (value: object) => {
    instanceInfo.value = [value] 
    resetPasswordInstance.value = !resetPasswordInstance.value;
}

/**
 * 重置密码实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} resetPasswordInstance.value 弹窗关闭
 */
const resetPasswordInstanceCancel = (type: boolean): boolean => {
    return resetPasswordInstance.value = type;
};

/**
 * 重置系统确认操作
 * @param value 
 */
const clickOperateResystemConfirm: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    resystemInstanceConfirm.value = !resystemInstanceConfirm.value;
}

/**
 * 重置系统确认弹窗取消
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
    instanceInfo.value = value 
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
const clickOperateDelete: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
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
 * 操作根据状态置灰
 * @param status string 当前状态
 * @param target string 对比状态
 * @return {boolean} 操作是否置灰
 */
const checkOperationDisable = (status: string, target: string) => {
    if(!(status === target)) {
        return true;
    }
    return false;  
}

/**
 * 操作下拉列表操作
 * @param value 
 * @param row 
 */
const clickMoreOperate=(value: any, row: any) =>{
    switch (value) {
        case 'lock':
            clickOperateLock(row)
            break;
        case 'unlock':
            clickOperateUnlock(row)
            break;
        case 'resetPassword':
            clickOperateResetPassword(row)
            break;
        case 'resystem':
            clickOperateResystemConfirm(row)
            break;
        case 'delete':
            clickOperateDelete(row)
            break;
        default:
            break;
    }
};

/**
 * 表格ref
*/
const tableRef = ref<InstanceType<typeof ElTable>>()

/**
* 列表选择实例
*/
const multipleSelection = ref<InstancesType[]>([])

/**
* 底部全选标志
*/
const hasCheckAll: Ref<boolean> = ref<boolean>(false)

/**
* 批量关机置灰状态
*/
const isAllowBatchClose: Ref<boolean> = ref<boolean>(false)

/**
* 批量开机置灰状态
*/
const isAllowBatchOpen: Ref<boolean> = ref<boolean>(false)

/**
* 批量重启置灰状态
*/
const isAllowBatchRestart: Ref<boolean> = ref<boolean>(false)

/**
* 批量重置密码置灰状态
*/
const isAllowBatchResetPassword: Ref<boolean> = ref<boolean>(false)

/**
* 批量编辑实例名称置灰状态
*/
const isAllowBatchRename: Ref<boolean> = ref<boolean>(false)

/**
* 批量删除置灰状态
*/
const isAllowBatchDelete: Ref<boolean> = ref<boolean>(false)

/**
* 列表全选 
*/
const handleCheckAllChange = () => {
    tableRef.value!.toggleAllSelection()
}

/**
* 列表选中项 
*/
const handleSelectionChange = (val: InstancesType[]) => {
    // 选中项
    multipleSelection.value = val
    // 全选标识
    hasCheckAll.value = val.length === reactiveArr.tableData.length ? true : false;
    // 刷新存储实例选择项
    let multipleSelectionInstanceInfo = {};
    multipleSelection.value.map((item: any) => {
        multipleSelectionInstanceInfo[item.instanceId] = item;
    })
    reactiveArr.multipleSelectionInstanceInfo = multipleSelectionInstanceInfo;
    let isBatchClose = multipleSelection.value.length ? multipleSelection.value.every((item: any) => item.status === 'running') : false;
	isAllowBatchClose.value = isBatchClose;
    let isBatchOpen = multipleSelection.value.length ? multipleSelection.value.every((item: any) => item.status === 'stopped') : false;
	isAllowBatchOpen.value = isBatchOpen;
    let isBatchRestart = multipleSelection.value.length ? multipleSelection.value.every((item: any) => item.status === 'running') : false;
	isAllowBatchRestart.value = isBatchRestart;
    let isBatchResetPassword = multipleSelection.value.length ? multipleSelection.value.every((item: any) => item.status === 'stopped') : false;
	isAllowBatchResetPassword.value = isBatchResetPassword;
    let isBatchRename = multipleSelection.value.length ? true : false;
	isAllowBatchRename.value = isBatchRename;
    let isBatchDelete = multipleSelection.value.length ? multipleSelection.value.every((item: any) => (item.status === 'stopped' || item.status === 'Creation failed') && item.locked === 'unlocked') : false;
	isAllowBatchDelete.value = isBatchDelete;
}

/**
 * filter参数类
*/
interface FilterParamsType {
    statusName: string;
};

/**
 * 批量关机操作打开标志
 */
const multipleClose: Ref<boolean> = ref<boolean>(false)

/**
 * 批量关机操作
 * @param value 
 */
const batchCloseDialog: () => void = () => {
    multipleClose.value = !multipleClose.value;
}

/**
 * 批量关机实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} multipleClose.value 弹窗关闭
 */
const batchCloseCancel = (type: boolean): boolean => {
    return multipleClose.value = type;
};

/**
 * 批量开机操作打开标志
 */
const multipleOpen: Ref<boolean> = ref<boolean>(false)

/**
 * 批量开机操作
 * @param value 
 */
const batchOpenDialog: () => void = () => {
    multipleOpen.value = !multipleOpen.value;
}

/**
 * 批量开机实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} multipleOpen.value 弹窗关闭
 */
const batchOpenCancel = (type: boolean): boolean => {
    return multipleOpen.value = type;
};

/**
 * 批量重启操作打开标志
 */
const multipleRestart: Ref<boolean> = ref<boolean>(false)

/**
 * 批量重启操作
 * @param value 
 */
const batchRestartDialog: () => void = () => {
    multipleRestart.value = !multipleRestart.value;
}

/**
 * 批量重启实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} multipleRestart.value 弹窗关闭
 */
const batchRestartCancel = (type: boolean): boolean => {
    return multipleRestart.value = type;
};

/**
 * 批量编辑实例名称操作打开标志
 */
const multipleRename: Ref<boolean> = ref<boolean>(false)

/**
 * 批量编辑实例名称操作
 * @param value 
 */
const batchRenameDialog: () => void = () => {
    multipleRename.value = !multipleRename.value;
}

/**
 * 批量编辑实例名称实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} multipleRename.value 弹窗关闭
 */
const batchRenameCancel = (type: boolean): boolean => {
    return multipleRename.value = type;
};

/**
 * 批量重置密码操作打开标志
 */
const multipleResetPassword: Ref<boolean> = ref<boolean>(false)

/**
 * 批量编辑实例名称操作
 * @param value 
 */
const batchResetPasswordDialog: () => void = () => {
    multipleResetPassword.value = !multipleResetPassword.value;
}

/**
 * 批量编辑实例名称实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} multipleResetPassword.value 弹窗关闭
 */
const batchResetPasswordCancel = (type: boolean): boolean => {
    return multipleResetPassword.value = type;
};

/**
 * 批量删除操作打开标志
 */
const multipleDelete: Ref<boolean> = ref<boolean>(false)

/**
 * 批量删除操作
 * @param value 
 */
const batchDeleteDialog: () => void = () => {
    multipleDelete.value = !multipleDelete.value;
}

/**
 * 批量删除实例弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} multipleDelete.value 弹窗关闭
 */
const batchDeleteCancel = (type: boolean): boolean => {
    return multipleDelete.value = type;
};

/**
 * 批量开更多操作
 */
const batchOperations = computed(() => {
    const operations = [
        {
            operateName: global.t('list.operate.restart'),
            disabled: !isAllowBatchRestart.value,
            tip: multipleSelection.value.length ? global.t('instance.tip.restart') : global.t('instance.tip.chooseInstance'),
            value: 'restart'
	    },
        {
            operateName: global.t('list.operate.resetPassword'),
            disabled: !isAllowBatchResetPassword.value,
            tip: multipleSelection.value.length ? global.t('instance.tip.open') : global.t('instance.tip.chooseInstance'),
            value: 'resetPassword'
	    },
        {
            operateName: global.t('list.operate.rename'),
            disabled: !isAllowBatchRename.value,
            tip: multipleSelection.value.length ? '' : global.t('instance.tip.chooseInstance'),
            value: 'rename'
	    },     
        {
            operateName: global.t('list.operate.delete'),
            disabled: !isAllowBatchDelete.value,
            tip: !multipleSelection.value.length ? global.t('instance.tip.chooseInstance') : ((multipleSelection.value.every((item: any) => ( item.locked === 'locked')) ? global.t('instance.tip.locked') : global.t('instance.tip.delete'))),
            value: 'delete'
	    }]
	        return operations
});

const batchDeleteTipsChange = () => {
    if(!multipleSelection.value.length) {
        return global.t('instance.tip.chooseInstance')
    } else {
        if(multipleSelection.value.every((item: any) => ( item.locked === 'locked'))) {
            return global.t('instance.tip.locked')
        } else {
            return global.t('instance.tip.delete')
        }
    }
}

/**
 * 操作下拉列表操作
 * @param value 
 * @param row 
 */
const clickBatchOperate = (value: any) =>{
    switch (value) {
        case 'restart':
            batchRestartDialog()
            break;
        case 'rename':
            batchRenameDialog()
            break;
        case 'resetPassword':
            batchResetPasswordDialog()
            break;
        case 'delete':
            batchDeleteDialog()
            break;
        default:
            break;
    }
};

const filterChange = (filter: {[x: string]: any;}) => {
    const filterParams: FilterParamsType | any = {
        status: 'status',
        deviceTypeId: 'deviceTypeId'
    };
    publicStore.filterParams(filter, ossStore, reactiveArr, filterParams, reactiveArr.deviceSeriesList).then(() => {
        pageNumber.value = 1;
        refreshData();
    });
};

/**
 * 表格点击某一行触发选中事件
 * @param {Object} row 当前选中的这一项
 */
const rowClick = (row: any) => {
    row.rowChange = !row.rowChange;
    tableRef.value!.toggleRowSelection(row, row.rowChange);
};

/**
 * 机型filter
 * @param {number} value 当前选中的value值
 * @param {Object} row 当前选中的这一项
 * @return {boolean} xxx 返回对应信息
*/
const deviceSeriesFilter = (value: number, row: any): boolean => {
    return row.deviceTypeId === reactiveArr.deviceSeriesList[value - 1]?.filterParams;
};

/**
 * 管理状态filter
 * @param {number} value 当前选中的value值
 * @param {Object} row 当前选中的这一项
 * @return {boolean} xxx 返回对应信息
*/
const statusFilter = (value: number, row: any): boolean => {
    return row.status === ossStore?.status[value - 1]?.filterParams;
};

interface SearchArrType {
    searchParams: {
        name?: string;
        instanceId?: string;
        iloIp?: string;
        privateIp?: string;
        privateIpv6?: string;
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
                label: global.t('instance.detail.instanceName')
            },
            {
                value: 1,
                label: global.t('instance.detail.instanceId')
            },
            {
                value: 2,
                label: global.t('instance.detail.outBandIp')
            },
            {
                value: 3,
                label: global.t('instance.detail.intranetIpv4')
            },
            {
                value: 4,
                label: global.t('instance.detail.intranetIpv6')
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
                this.reactiveArr.searchParams = {name: iptValue}
                break; 
            case 1 :
                this.reactiveArr.searchParams = {instanceId: iptValue}
                break; 
            case 2 :
                this.reactiveArr.searchParams = {iloIp: iptValue}
                break;  
            case 3 :
                this.reactiveArr.searchParams = {privateIp: iptValue}
                break; 
            case 4 :
                this.reactiveArr.searchParams = {privateIpv6: iptValue}
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
            privateIpv6,
            privateIp,
            iloIp,
            instanceId, 
            name
        } : {
            privateIpv6?: string; 
            privateIp?: string; 
            iloIp?: string; 
            instanceId?: string; 
            name?: string;
        } = this.reactiveArr.searchParams;
        if (privateIpv6 || privateIp || iloIp || instanceId || name) {
            this.reactiveArr.searchParams = {};
            this.request();
        }
    };

    request = () => {
        if (pageNumber.value > 1) {
            pageNumber.value = 1;
        }
        refreshData();
    };

};
const searchOperate: SearchOperate = new SearchOperate();

/**
 * 返回列表
 */
const backToList = () =>{
    searchOperate.hasClear.value = true;
    searchOperate.reactiveArr.searchParams = {};
    searchOperate.request();
    searchResultMark.value = false
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
        instanceListExportAPI(
            {
                exportType: '1',
                projectId: route.query.projectId,
                isAll: '1',
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
            link.download = `instance_list_${getDate()}_${generateRandomNum(6)}`;
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
            return '';
    }
}

const processingData = (data: any, type: string, id: string) => {
    return filterData(data, type).map((item: any, index: number) => {
        return {
            ...item,
            text: item[type],
            value: index + 1,
            filterParams: item[id]
        };
    });
}

/**
 * 跳入详情页
 */
const router = useRouter();
const jumpToDetail = (id: string) => {
    router.push({
        path: `/instance/detail`,
        query: {
            instanceId: id,
            projectId: route.query.projectId,
            projectName: route.query.projectName
        }
    });
}

const jumpToDetailMonitor = (id: string,) => {
    router.push({
        path: `/instance/detail`,
        query: {
            instanceId: id,
            projectId: route.query.projectId,
            projectName: route.query.projectName,
            monitor: 'PerformanceMonitoring'
        }
    });
}

/**
 * 跳入创建页
 */
const jumpToCreate = () => {
    router.push({
        path: `/instance/create`,
        query: {
            projectId: route.query.projectId,
            projectName: route.query.projectName
        }
    });
}

/**
 * 获取机型筛选列表
 */
const getDeviceSeriesListData = (): void => {
    instanceListAPI({
        isAll:'1',
        projectId: route.query.projectId ? route.query.projectId : '',
        //status: selectStatus
    }).then(({data} : {data: any}) => {
        if (data?.result?.instances?.length) {
            reactiveArr.deviceSeriesList = processingData(data.result.instances, 'deviceTypeName', 'deviceTypeId');
            return;
        }
        reactiveArr.deviceSeriesList = [];
 
    }).catch (()=>{
        reactiveArr.deviceSeriesList = [];
    })
    .finally(() => {

    });
};

/**
 * 页面生成，销毁时-触发，清空延时器
*/
let timer: null | number = null;
const isIntervalRequest: Ref<boolean> = ref<boolean>(false);
onMounted(() => clearTimeout((timer as number)));
onUnmounted(() => clearTimeout((timer as number)));
onBeforeUnmount(() => clearTimeout((timer as number)));
/**
 * 判断性能监控授权
 */
const hasInMonitor: Ref<boolean> = ref<boolean>(false);  

/**
 * 请求实例列表数据接口
*/
const requestInstanceListData = (params: any): void => {
    clearTimeout((timer as number))
    instanceListAPI({
        ...params
        //status: selectStatus
    }).then(({data} : {data: any}) => {
        if (data?.result?.instances?.length) {
            reactiveArr.tableData = data.result.instances.map((item:any) => {
                let lockOperation : any = {}
                if(item.locked === 'unlocked') {
                    lockOperation.value = {
                        operateName: global.t('list.operate.lock'),
                        disabled: (item.status === 'Creation failed')
                                    ||(item.status === 'creating')
                                    ||(item.status === 'Creating')
                                    ||(item.status === 'Delete failed')
                                    ||(item.status === 'destroying')
                                    ||(item.status === 'Destroying'),
                        tip: global.t('instance.tip.locking'),
                        value: 'lock'
                    }
                } else {
                    lockOperation.value = {
                        operateName: global.t('list.operate.unlock'),
                        disabled: (item.status === 'Creation failed')
                                    ||(item.status === 'creating')
                                    ||(item.status === 'Creating')
                                    ||(item.status === 'Delete failed')
                                    ||(item.status === 'destroying')
                                    ||(item.status === 'Destroying'),
                        tip: global.t('instance.tip.unlocking'),
                        value: 'unlock'
                    }
                };
                let allOperations : any = [
                    {
                        operateName: global.t('list.operate.resetPassword'),
                        disabled: !(item.status === 'stopped'),
                        tip: global.t('instance.tip.open'),
                        value: 'resetPassword'
                    },{
                        operateName: global.t('list.operate.resystem'),
                        disabled: !(item.status === 'stopped'),
                        tip: global.t('instance.tip.open'),
                        value: 'resystem'
                    },{
                        operateName: global.t('list.operate.delete'),
                        disabled: !(item.status === 'stopped' || item.status === 'Creation failed') || item.locked === 'locked',
                        tip: item.locked === 'locked' ? global.t('instance.tip.locked') : global.t('instance.tip.delete'),
                        value: 'delete'
                    }
                ]
                allOperations.splice(0, 0, lockOperation.value)
                return {
                    ...item,
                    rowChange: false,
                    operations: allOperations,
                    tooltipStatus: {
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
                    }
                }
            });
            // 防止列表实例选择在中间态刷新中清除
            let multipleSelectionInstanceInfo:any = deepCopy(reactiveArr.multipleSelectionInstanceInfo)
            nextTick(() => {
                reactiveArr.tableData.map(item => {
                    let instanceId = item.instanceId;
                    if (multipleSelectionInstanceInfo[instanceId]) {
                        item.rowChange = true;
                        tableRef.value!.toggleRowSelection(item,true)
                    }
                })
            })
            totalCount.value = data.result.totalCount;
            for (const key of reactiveArr.tableData) {
                if ([...intermediate].some((item: string) => item === key.status)) {
                    timer = setTimeout(() => {
                        isIntervalRequest.value = true;
                        requestInstanceListData(params);    
                    }, 5000)
                    return;
                }               
            }
            isIntervalRequest.value = false;
            return;
        }
        reactiveArr.tableData = [];
        totalCount.value = 0;      
    }).catch(({message} : {message: string;}) => {
        if (message.indexOf('undefined') > -1) return;
        if (!isIntervalRequest.value) {
            typeof message === 'string' && ElMessage.error(message);
            reactiveArr.tableData = [];
            totalCount.value = 0;
            return;
        }
        timer = setTimeout(() => {
            requestInstanceListData(params);
        }, 5000)
    }).finally(() => {
        isLoading.value = false;
        searchOperate.hasClear.value = false;
        if(!isIntervalRequest.value) {
            store.requestObject();
            store.requestUser();
            hasInMonitor.value = store.inMonitor; 
            getDeviceSeriesListData();
        }     
        (instanceMitt as any).proxy.$Bus.emit('project-name', route.query.projectName);
        
    });
};

/**
 * 处理参数
*/
const processingParameter = () => {
    const params = {
        pageSize: pageSize.value,
        pageNumber: pageNumber.value,
        projectId: route.query.projectId ? route.query.projectId : '',
        ...reactiveArr.filterParams,
        ...searchOperate.reactiveArr.searchParams,
    };
    publicStore.deleteEmtpyData(params);
    requestInstanceListData(params);
};

/**
 * 刷新接口
*/
const refreshData = (): void => {
    isLoading.value = true;
    reactiveArr.multipleSelectionInstanceInfo = {};
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
    refreshData();
};

/**
 * 页数切换
 * @param {number} num 页数切换
*/
const changePage = (num: number): void => {
    pageNumber.value = num;
    refreshData();
};
</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import 'assets/css/addButton';
@import 'assets/css/icon';
@import 'assets/css/tagColor';
@import 'assets/css/list';
@import './instanceList.scss';

</style>