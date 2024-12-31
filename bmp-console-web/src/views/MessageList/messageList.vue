<template>
    <div class="page-position message-list">
        <div class="page-content pos-fixed">
            <div class="message-list-title">
                <!-- 消息列表 -->
                <h3 class="section-title ">{{ $t('messageCentre.messageList') }}</h3>
                <span class="unread-tip">{{$t('messageCentre.list.toolTip.unreadTip1', [totalMessageCount ? totalMessageCount : 0])}}</span>
                <span class="default-color mouse-point" @click="changeUnreadGroup">{{totalUnreadCount ? totalUnreadCount : 0}}</span>
                <span class="unread-tip">{{$t('messageCentre.list.toolTip.unreadTip2')}}</span>
            </div>
            
            <div class="operation-header">  
                <div class="group-list-choose">
                    <el-radio-group v-model="groupingMessageList" @change="readChange">
                        <el-radio-button label = '2'   
                        >
                            {{$t('messageCentre.list.all')}}
                        </el-radio-button>
                        <el-radio-button label = '0'   
                        >
                            {{$t('messageCentre.list.unread')}}
                        </el-radio-button>
                        <el-radio-button label = '1'   
                        >
                            {{$t('messageCentre.list.read')}}
                        </el-radio-button>
                    </el-radio-group> 
                </div>           
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
                                $t('messageCentre.list.toolTip.messageContent'),
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
                    <!-- 消息子类型筛选弹框 -->
                    <div 
                        v-if="messageSubTypeShow"
                        class="filter-content"
                        :class="[
                            getLocationItem === 'zh_CN' ? 'right-zh' : 'right-en'
                        ]"
                    >
                        <div class="el-popper is-pure is-light  filter-content-fixed">
                            <div 
                                class="filter-message-list-input"
                                :class="[
                                    getLocationItem === 'zh_CN' ? 'input-zh' : 'input-en'
                                ]"
                            >
                                <el-input
                                    v-model="messageSubTypeSearch"
                                    :placeholder="$t('messageCentre.list.toolTip.inputSubtype')"
                                    :suffix-icon="Search"
                                    @click.stop
                                    />
                            </div>
                            <div class="filter-message-list-checkbox">
                                <el-checkbox-group
                                    v-model="messageSubTypeChoose"
                                    >
                                    <el-checkbox
                                        v-for="(item, index) in messageSubTypeList"
                                        :key="item.text"
                                        :label="item.text"
                                        :value="item.text"
                                        :title="item.text"
                                        @click.stop
                                    />
                                </el-checkbox-group>
                            </div>
                            <div class="filter-message-list-button">
                                <el-button
                                    type="text"
                                    :disabled="messageSubTypeChoose.value"
                                    @click="messageSubTypeConfirm"
                                >
                                {{$t('messageCentre.list.confirm')}}</el-button>
                                <el-button
                                    type="text"
                                    @click="messageSubTypeReset"
                                >
                                {{$t('messageCentre.list.reset')}}</el-button>
                            </div>
                        </div>
                    </div>
                    <el-table 
                        border
                        ref="tableRef"
                        :data="reactiveArr.tableData" 
                        :header-cell-class-name="cellHeaderClassName"
                        :cell-class-name="cellClassName"
                        style="width: 100%"
                        v-loading="isLoading"
                        @row-click="rowClick"
                        :max-height="reactiveArr.tableData.length ? tableMaxHeight : 380"
                        @filter-change="filterChange"
                        @selection-change="handleSelectionChange"
                    >
                        <!-- 选择多行 -->
                        <el-table-column type="selection" width="55" />
                        <el-table-column width="36">  
                            <template v-slot="scope">
                                <div class="dot" v-if="!scope.row.has_read"></div>
                            </template>  
                        </el-table-column>
                        <!-- 消息内容 -->
                        <el-table-column 
                            prop="detail" 
                            v-if="reactiveArr?.hasCustomInfo['detail']?.selected"
                            :label="$t('messageCentre.list.messageContent')" 
                            min-width = 475
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-if="scope.row.message_type === $t('messageCentre.detail.operateMessage')"
                                    v-model="scope.row.tooltipStatus['messageContent'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['messageContent'].showTooltip"
                                    :content= "scope.row.message_type + ' - 【' +  scope.row.message_sub_type + '】' + messageContentConstant[scope.row.result] + ' (' + scope.row.detail + ')'"
                                    placement="right"
                                >
                                    <div
                                        :class="scope.row.has_read === 1 ? 'read-message-content' : 'unread-message-content'"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['messageContent'], scope.row.message_type + ' - 【' +  scope.row.message_sub_type + '】' + messageContentConstant[scope.row.result] + ' (' + scope.row.detail + ')', 1.17, 'list')"
                                    >   
                                        <span 
                                        class="text-underline"
                                        @click.stop="jumpToDetail(scope.row.message_id)">
                                            {{scope.row.message_type + ' - 【' +  scope.row.message_sub_type + '】' + messageContentConstant[scope.row.result] + ' (' + scope.row.detail + ')' || $filter.emptyFilter()}}
                                        </span>   
                                    </div>                                   
                                </el-tooltip> 
                                <el-tooltip
                                    v-else-if="scope.row.message_type === $t('messageCentre.detail.faultMessage')"
                                    v-model="scope.row.tooltipStatus['messageContent'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['messageContent'].showTooltip"
                                    :content= "scope.row.message_type + ' - 【' +  scope.row.fault_type + '】' + $t('messageCentre.detail.remindContent') + ' (' + scope.row.instance_name +'/' + scope.row.sn + ')'"
                                    placement="right"
                                >
                                    <div
                                        :class="scope.row.has_read === 1 ? 'read-message-content' : 'unread-message-content'"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['messageContent'], scope.row.message_type + ' - 【' +  scope.row.message_sub_type + '】' + messageContentConstant[scope.row.result] + ' (' + scope.row.instance_name +'/' + scope.row.sn + ')', 1.17, 'list')"
                                    >   
                                        <span 
                                        class="text-underline"
                                        @click.stop="jumpToDetail(scope.row.message_id)">
                                            {{scope.row.message_type + ' - 【' +  scope.row.fault_type + '】' + $t('messageCentre.detail.remindContent') + ' (' + scope.row.instance_name +'/' + scope.row.sn + ')' || $filter.emptyFilter()}}
                                        </span>   
                                    </div>                                   
                                </el-tooltip>  
                                <el-tooltip
                                    v-else-if="scope.row.message_type === $t('messageCentre.detail.alarmMessage')"
                                    v-model="scope.row.tooltipStatus['messageContent'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['messageContent'].showTooltip"
                                    :content= "scope.row.message_type + ' - 【' +  scope.row.message_sub_type + '】' +  ' (' + scope.row.instance_name +'/'+ scope.row.instance_id + ')'"
                                    placement="right"
                                >
                                    <div
                                        :class="scope.row.has_read === 1 ? 'read-message-content' : 'unread-message-content'"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['messageContent'], scope.row.message_type + ' - 【' +  scope.row.message_sub_type + '】' + ' (' + scope.row.instance_name +'/'+ scope.row.instance_id + ')', 1.11, 'list')"
                                    >   
                                        <span 
                                        class="text-underline"
                                        @click.stop="jumpToDetail(scope.row.message_id)">
                                            {{scope.row.message_type + ' - 【' +  scope.row.message_sub_type + '】' + ' (' + scope.row.instance_name +'/'+ scope.row.instance_id + ')' || $filter.emptyFilter()}}
                                        </span>   
                                    </div>                                   
                                </el-tooltip>   
                            </template>
                        </el-table-column>
                        <!-- 接收时间 -->
                        <el-table-column     
                            prop="finish_time" 
                            v-if="reactiveArr?.hasCustomInfo['finish_time']?.selected"
                            :label="$t('messageCentre.list.receivingTime')"
                            min-width=140
                        >
                            <template v-slot="scope">
                                {{scope.row.finish_time || $filter.emptyFilter()}}
                            </template>
                        </el-table-column> 
                        <!-- 消息类型 -->
                        <el-table-column     
                            prop="message_type" 
                            key="messageType"
                            column-key="messageType"
                            v-if="reactiveArr?.hasCustomInfo['message_type']?.selected"
                            :label="$t('messageCentre.list.messageType')"
                            filter-placement="bottom-end"
                            :filter-multiple="false"
                            :filters="reactiveArr.messageTypeList"
                            :filter-method="messageTypeFilter"
                            min-width=120
                        >
                            <template v-slot="scope">
                                <el-tooltip
                                    v-model="scope.row.tooltipStatus['message_type'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['message_type'].showTooltip"
                                    :content= scope.row.message_type
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['message_type'], scope.row.message_type, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.message_type || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                         <!-- 消息子类型 -->
                        <el-table-column     
                            prop="message_sub_type" 
                            key="messageSubType"
                            column-key="messageSubType"
                            v-if="reactiveArr?.hasCustomInfo['message_sub_type']?.selected"
                            :label="$t('messageCentre.list.messageSubtype')"
                            width=200
                        >   
                            <template #header>
                                <span :class="changeBlue()">{{$t('messageCentre.list.messageSubtype')}}</span>
                                <span 
                                    class="el-table__column-filter-trigger el-none-outline el-tooltip__trigger el-tooltip__trigger"
                                    
                                    @click.stop="messageSubTypeDialog"
                                    ></span>
                                
                            </template>
                            <template v-slot="scope">
                                <el-tooltip
                                    v-if="scope.row.message_type === $t('messageCentre.detail.operateMessage')"
                                    v-model="scope.row.tooltipStatus['message_sub_type'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['message_sub_type'].showTooltip"
                                    :content= scope.row.message_sub_type
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['message_sub_type'], scope.row.message_sub_type, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.message_sub_type || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                                <el-tooltip
                                    v-if="scope.row.message_type === $t('messageCentre.detail.faultMessage')"
                                    v-model="scope.row.tooltipStatus['message_sub_type'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['message_sub_type'].showTooltip"
                                    :content= scope.row.fault_type
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['message_sub_type'], scope.row.message_sub_type, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.fault_type || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                                <el-tooltip
                                    v-if="scope.row.message_type === $t('messageCentre.detail.alarmMessage')"
                                    v-model="scope.row.tooltipStatus['message_sub_type'].showTooltip"
                                    :disabled="!scope.row.tooltipStatus['message_sub_type'].showTooltip"
                                    :content= scope.row.message_sub_type
                                >
                                    <div
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row.tooltipStatus['message_sub_type'], scope.row.message_sub_type, 1.17, 'list')"
                                    >
                                        <span>{{scope.row.message_sub_type || $filter.emptyFilter()}}</span>
                                    </div>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 列表为空 -->
                        <template #empty>
                            <div class="no-data-list">                        
                            </div>
                        </template>
                    </el-table>
                    <p 
                        v-if="!reactiveArr.tableData.length && !isLoading"
                        class="no-data-jump"
                    >
                    </p>
                </el-config-provider>
                <div v-if="openCustom">
                    <custom
                        :page-key="'messageList'"
                        :open-visible="openCustom"
                        :check-list-arr="reactiveArr.checkListArr"
                        :has-custom-info="reactiveArr.hasCustomInfo"
                        @close="openCustomCancel"
                        @determine="publicStore.customList('messageList', reactiveArr)"
                    >
                    </custom>
                </div>
            </div>
            <div class="footer-count">
                <div class="batch-operate">
                    <el-checkbox @change="handleCheckAllChange" v-model="hasCheckAll" class="mr20"></el-checkbox>
                    <!-- 删除 -->
                    <el-button 
                        type="primary" 
                        :disabled="!isAllowBatchDelete"    
                        @click="batchDeleteOperate"
                    >
                        {{$t('messageCentre.operate.delete')}}
                    </el-button>   
                    <!-- 标为已读 -->
                    <el-button 
                        type="primary" 
                        :disabled="!isAllowBatchRead"
                        @click="batchReadOperate"
                    >
                        {{$t('messageCentre.operate.markedRead')}}
                    </el-button>        
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
import i18n from 'lib/i18n/index.ts'; // 国际化
import searchInput from 'components/SearchInput/searchInput.vue';
import {
    messageCustomInfo, // 消息列表自定义初始数据
    messageCustomIntro, // 消息列表自定义信息
    groupingMessageListConstant // 类型数据
} from 'utils/constants.ts';
import {
    getDate, // 获取日期
    generateRandomNum, // 生成随机数
    filterData, // 过滤重复数据
    languageSwitch, // 语言切换
    deepCopy, //深拷贝
    getDateMinutes,
    hasShowTooltip // 是否显示提示
} from 'utils/index.ts';
import { 
    ElTable,
    ElMessage
 } from 'element-plus'
import publicIndexStore from 'store/index.ts'; // 公共store
import allProjectStore from '@/store/modules/headerDetail.ts';
import { useRouter, useRoute } from 'vue-router';
import custom from 'components/custom/custom.vue';
import myPagination from 'components/Pagination/MyPagination.vue';
import { Search } from '@element-plus/icons-vue'
import {
    messageContentConstant, // 消息通知状态
} from 'utils/constants.ts';
import {
    messageDeleteAPI,
    messageTypeAPI,
    messageReadAPI,
    messageListAPI,
    getMessageStatisticAPI,
    messageListExportAPI
} from 'api/request.ts'; // 请求接口

 // 计算表格的最大高度
 const tableMaxHeight: Ref<number> = ref<number>(window.innerHeight - 380)
// 监听窗口大小变化，以便动态调整表格高度
const updateSize = () => {
  // 触发响应式更新
  tableMaxHeight.value = window.innerHeight - 380
};

onMounted(() => {
  window.addEventListener('resize', updateSize);
});

onUnmounted(() => {
  window.removeEventListener('resize', updateSize);
});

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: any = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';

/**
 * 语言
*/
const locale: any = languageSwitch();

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();

/**
 * 国际化
*/
const {global} = i18n;

/**
 * 消息item类型
 */
 type MessagesType = {
    detail: string,
    finish_time: string,
    message_type: string,
    message_sub_type: string,
    rowChange: boolean,
    message_id: string
}

/**
 * 消息类型
 */
 interface message {
    tableData: MessagesType[],
    multipleSelectionMessageInfo: {},
    messageTypeList: any,
    messageSubtypeList: any,
    filterParams: {},
    hasCustomInfo: {},
    checkListArr: {},
}

/**
 * 消息列表数据
 */
 const reactiveArr: message = reactive<message>({
    tableData: [],
    multipleSelectionMessageInfo: {},
    messageTypeList: [],
    messageSubtypeList: [],
    filterParams: {
        messageType: '',
        messagesubType: '',
    },
    hasCustomInfo: messageCustomIntro,
    checkListArr: messageCustomInfo,
})

/**
 * 消息列表分组
 */
 const groupingMessageList: Ref<string> = ref<string>(groupingMessageListConstant[0].value)

/**
 * 多选列表分组，每次改变，重新获取对应列表信息
 */
const readChange = () => {
    pageNumber.value = 1; 
    refreshData()
    if(groupingMessageList.value !== '2') {
        searchResultMark.value = true;
    } 
    if(groupingMessageList.value === '2' && !reactiveArr.filterParams['messageType'] && !reactiveArr.filterParams['messageSubType'] && !searchOperate.reactiveArr.searchParams['detail'] ) {
        searchResultMark.value = false;
    }
}

/**
 * 列表消息内容一列表头靠左排列
 */
 const cellHeaderClassName = ({row, column}: any) => {
    if(column.no === 1) {
        return 'border-none'
    } 
    if(column.no === 2) {
        return 'column-left'
    } 
}

/**
 * 列表消息内容一列表内容靠左排列
 */
 const cellClassName = ({row, column}: any) => {
    if(column.no === 2) {
        return 'column-left'
    } 
}

const changeUnreadGroup = () => {
    groupingMessageList.value = '0'
};

/**
 * 搜索type
 */
interface SearchArrType {
    searchParams: {
        detail?: string;
    };
    selectOption: {
        value: number;
        label: string;
    }[];
};

/**
 * 搜索结果提示
 */
 const searchResultMark: Ref<boolean> = ref<boolean>(false)


/**
 * 搜索框类
 */
class SearchOperate {
    hasClear: Ref<boolean> = ref<boolean>(false);
    // 搜索框值
    selectVal: Ref<number> = ref<number>(0);
    reactiveArr: SearchArrType = reactive<SearchArrType>({
        searchParams: {},
        selectOption: [
            {
                value: 0,
                label: global.t('messageCentre.list.messageContent')
            },
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
                this.reactiveArr.searchParams = {detail: iptValue}
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
            detail
        } : {
            detail?: string;
        } = this.reactiveArr.searchParams;
        if (detail) {
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
 * 自定义列表操作打开标志
 */
 const openCustom: Ref<boolean> = ref<boolean>(false)

/**
 * 自定义列表接口
*/
publicStore.customList('messageList', reactiveArr);

/**
 * loadig态
*/
const isLoading: Ref<boolean> = ref<boolean>(true);

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
        messageListExportAPI(
            {
                exportType: '1',
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
            link.download = `message_list_${getDate()}_${generateRandomNum(6)}`;
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
 * 表格ref
*/
const tableRef = ref<InstanceType<typeof ElTable>>()

/**
 * 返回列表
 */
 const backToList = () =>{
    searchOperate.hasClear.value = true;
    searchOperate.reactiveArr.searchParams = {};
    //searchOperate.request();
    messageSubTypeChoose.value = [];
    reactiveArr.filterParams['messageSubType'] = '';
    reactiveArr.filterParams['messageType'] = '';
    groupingMessageList.value = groupingMessageListConstant[0].value;
    // @ts-ignore
    tableRef.value!.clearFilter();
    refreshData()
    searchResultMark.value = false
}



/**
 * 表格点击某一行触发选中事件
 * @param {Object} row 当前选中的这一项
 */
const rowClick = (row: any) => {
    row.rowChange = !row.rowChange;
    tableRef.value!.toggleRowSelection(row, row.rowChange);
};

/**
 * 处理筛选数据
 */
const processingData = (data: any) => {
    let arrData: any = []
    data.map((item: string) => {
        arrData.push({
            filterParams: item,
            text: item,
            value: item
        })
        
    });
    return arrData
}

/**
 * 获取消息类型筛选下拉框列表
 */
 const getMessageTypeListData = (): void => {
    messageTypeAPI({
    }).then(({data} : {data: any}) => {
        if (data?.result) {
            reactiveArr.messageTypeList = processingData(Object.keys(data.result));
            //flatMessageSubTypeList = flattenAndJoinArray(Object.values(data.result));
            reactiveArr.messageSubtypeList = processingData(Object.values(data.result).flat());
            return;
        }
        reactiveArr.messageTypeList = [];
        reactiveArr.messageSubtypeList = [];
 
    }).catch (()=>{
        reactiveArr.messageTypeList = [];
        reactiveArr.messageSubtypeList = [];
    })
    .finally(() => {

    });
};

/**
 * 筛选子类型下拉框展示
 */
const messageSubTypeShow: Ref<boolean> = ref<boolean>(false);

/**
 * 筛选子类型下拉框出现
 */
const messageSubTypeDialog = () => {
    messageSubTypeShow.value = !messageSubTypeShow.value 
}

/**
 * 选子类型下拉框消失
 */
document.body.addEventListener('click', () => {
    if(messageSubTypeShow.value) {
        messageSubTypeShow.value = false;
        messageSubTypeSearch.value = ''
    }    
})

/**
 * 筛选子类型下拉框搜索关键字
 */
 const messageSubTypeSearch: Ref<string> = ref<string>('');

/**
 * 筛选后实例数据
 */
const messageSubTypeList = computed(() => {  
    if(messageSubTypeSearch.value){
        return reactiveArr.messageSubtypeList.filter(function(dataNews:any){
            return Object.keys(dataNews).some(function(key){
              return String(dataNews[key]).toLowerCase().indexOf(messageSubTypeSearch.value) > -1
            })
        })
    }
    return reactiveArr.messageSubtypeList
});

/**
 * 筛选子类型下拉框选择关键字
 */
const messageSubTypeChoose: Ref<any> = ref<any>([]);

/**
 * 筛选子类型下拉框选择重置
 */
const messageSubTypeReset = () => {
    messageSubTypeChoose.value = [];
    reactiveArr.filterParams['messageSubType'] = '';
    refreshData()
}

/**
 * 筛选子类型下拉框选择筛选
 */
const messageSubTypeConfirm = () => {
    if(messageSubTypeChoose.value) {
        reactiveArr.filterParams['messageSubType'] = Object.values(messageSubTypeChoose.value).join(',');
    }
    refreshData() 
    searchResultMark.value = true 
}

const changeBlue = () => {
    if(Object.keys(messageSubTypeChoose.value).length) {
        return 'change-blue'
    }
    return   
}



/**
 * filter参数类
*/
interface FilterParamsType {
    messageType: string;
    messageSubType: string;
};


const filterParamsArray = (
    filter: {[x: string]: number[]},
    reactiveArr: any,
    filterParams: any,
) => {
    if (!Object.keys(filter).length) {
            return Promise.reject();
        }
        else {
            for (const key in filter) {
                if (key === filterParams[key]) {  
                    // 找到匹配的对应参数，如果多选，数组转字符串，用逗号分隔
                    reactiveArr.filterParams[key] = Object.values(filter[key]).join(',');
                }
            }
            return Promise.resolve();
    }
}

const filterChange = (filter: {[x: string]: any;}) => {
    const filterParams: FilterParamsType  = {
        messageType: 'messageType',
        messageSubType: 'messageSubType'
    };
    filterParamsArray(filter, reactiveArr, filterParams).then(() => {
        pageNumber.value = 1;
        refreshData();
        searchResultMark.value = true 
    });
};



/**
 * 消息类型filter
 * @param {number} value 当前选中的value值
 * @param {Object} row 当前选中的这一项
 * @return {boolean} xxx 返回对应信息
*/
const messageTypeFilter = (value: number, row: any): boolean => {
    return row.message_type === value;
};

/**
 * 消息子类型filter
 * @param {number} value 当前选中的value值
 * @param {Object} row 当前选中的这一项
 * @return {boolean} xxx 返回对应信息
*/
const messageSubtypeFilter = (value: number, row: any): boolean => {
    return row.messageSubtypeId === reactiveArr.messageSubtypeList[value - 1]?.filterParams;
};


/**
* 列表全选 
*/
const handleCheckAllChange = () => {
    tableRef.value!.toggleAllSelection()
}

/**
* 列表选择实例
*/
const multipleSelection = ref<MessagesType[]>([])

/**
* 底部全选标志
*/
const hasCheckAll: Ref<boolean> = ref<boolean>(false)

/**
* 批量删除置灰状态
*/
const isAllowBatchDelete: Ref<boolean> = ref<boolean>(false)

/**
* 批量已读置灰状态
*/
const isAllowBatchRead: Ref<boolean> = ref<boolean>(false)

/**
* 列表选中项 
*/
const handleSelectionChange = (val: MessagesType[]) => {
    // 选中项
    multipleSelection.value = val
    // 全选标识
    hasCheckAll.value = val.length === reactiveArr.tableData.length ? true : false;
    // 刷新存储实例选择项
    let multipleSelectionMessageInfo = {};
    multipleSelection.value.map((item: any) => {
        multipleSelectionMessageInfo[item.message_id] = item;
    })
    reactiveArr.multipleSelectionMessageInfo = multipleSelectionMessageInfo;
    let isBatchDelete = multipleSelection.value.length ? true : false;
	isAllowBatchDelete.value = isBatchDelete;
    let isBatchRead = multipleSelection.value.length ? true : false;
	isAllowBatchRead.value = isBatchRead;
}

/**
 * 页面生成，销毁时-触发，清空延时器
*/
let timer: null | number = null;
const isIntervalRequest: Ref<boolean> = ref<boolean>(false);
onMounted(() => clearTimeout((timer as number)));
onUnmounted(() => clearTimeout((timer as number)));
onBeforeUnmount(() => clearTimeout((timer as number)));

/**
 * store库存储的项目列表
*/
const store = allProjectStore();

/**
 * 请求信息列表数据接口
*/
const requestMessageListData = (params: any): void => {
    clearTimeout((timer as number))
    messageListAPI({
        ...params

    }).then(({data} : {data: any}) => {
        if(data?.result?.messages?.length) {
            reactiveArr.tableData = data.result.messages.map((item:any) => {
                item.finish_time = getDateMinutes(item.finish_time)
                return {
                    ...item,
                    rowChange: false,
                    tooltipStatus: {
                        'messageContent': {
                            showTooltip: false
                        },
                        'finish_time': {
                            showTooltip: false
                        },
                        'message_type': {
                            showTooltip: false
                        },
                        'message_sub_type': {
                            showTooltip: false
                        },
                    }
                }
            });
            // 防止列表实例选择在中间态刷新中清除
            let multipleSelectionMessageInfo:any = deepCopy(reactiveArr.multipleSelectionMessageInfo)
            nextTick(() => {
                reactiveArr.tableData.map(item => {
                    let message_id = item.message_id;
                    if (multipleSelectionMessageInfo[message_id]) {
                        item.rowChange = true;
                        tableRef.value!.toggleRowSelection(item,true)
                    }
                })
            })
            totalCount.value = data.result.totalCount;
        } 
        else {
            reactiveArr.tableData = [];
            totalCount.value = 0;  
        }
        
        // 每五秒请求一次接口，轮询
        timer = setTimeout(() => {
            isIntervalRequest.value = true;
            requestMessageListData(params);  
            requestGetMessageStatistic()
        }, 5000)  
        return            
        // isIntervalRequest.value = false;
        // return;
        // reactiveArr.tableData = [];
        // totalCount.value = 0;  
        // totalUnreadCount.value = 0;  
        // totalMessageCount.value = 0  
    }).catch(({message} : {message: string;}) => {
        if (message.indexOf('undefined') > -1) return;
        if (!isIntervalRequest.value) {
            typeof message === 'string' && ElMessage.error(message);
            reactiveArr.tableData = [];
            totalCount.value = 0;
            totalUnreadCount.value = 0; 
            totalMessageCount.value = 0
            return;
        }
        timer = setTimeout(() => {
            requestMessageListData(params);
        }, 5000)
    }).finally(() => {
        isLoading.value = false;
        searchOperate.hasClear.value = false;
        if(!isIntervalRequest.value) {
            store.requestObject();
            getMessageTypeListData();
        }     
        
    });
};

/**
 * 处理参数
*/
const processingParameter = () => {    
    const params = {
        pageSize: pageSize.value,
        pageNumber: pageNumber.value,
        ...reactiveArr.filterParams,
        ...searchOperate.reactiveArr.searchParams,
    };
    if(groupingMessageList.value !== '2') {
        params['hasRead'] = groupingMessageList.value
    }
    publicStore.deleteEmtpyData(params);
    requestMessageListData(params);
};

/**
 * 刷新接口
*/
const refreshData = (): void => {
    isLoading.value = true;
    reactiveArr.multipleSelectionMessageInfo = {};
    processingParameter();
};

/**
 * 请求信息未读数接口
*/
const requestGetMessageStatistic = (): void => {
    getMessageStatisticAPI({  
    }).then(({data}: any) => {
        if(data?.result) {
            totalUnreadCount.value = data.result.unreadCount
            totalMessageCount.value = data.result.totalCount
            return
        }
        totalUnreadCount.value = 0
        totalMessageCount.value = 0
    }).catch(() => {
        totalUnreadCount.value = 0
        totalMessageCount.value = 0
    }).finally(() => {

    });
}; 

onMounted(() => {
    requestGetMessageStatistic()
    refreshData()
})


/**
 * 信息批量删除操作
 */
const batchDeleteOperate = () => {
    requestDelete(multipleSelection.value)
}

/**
 * 请求删除接口
 * @param {MessagesType[]} messages 已选信息
*/
const requestDelete = (messages: MessagesType[]): void => {
    messageDeleteAPI({
        messageIds: messages.map((item: any) => item.message_id)
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: global.t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            refreshData()
        }
    }).catch(() => {
        refreshData()
    }).finally(() => {
        isLoading.value = false;
    });
}; 

/**
 * 信息标为已读操作
 */
 const batchReadOperate = () => {
    requestRead(multipleSelection.value)
}

/**
 * 请求已读接口
 * @param {MessagesType[]} messages 已选信息
*/
const requestRead = (messages: MessagesType[]): void => {
    messageReadAPI({
        messageIds: messages.map((item: any) => item.message_id)
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: global.t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            refreshData()
        }
    }).catch(() => {
        refreshData()
    }).finally(() => {
        requestGetMessageStatistic()
        isLoading.value = false;
    });
};

/**
 * 跳入详情页
 */
 const router = useRouter();
const jumpToDetail = (id: string) => {
    router.push({
        path: `/message/detail`,
    });
    localStorage.setItem('messageInfo', window.btoa(JSON.stringify(id)));
}

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
 * 总消息数量条数
*/
const totalMessageCount: Ref<number> = ref<number>(0);  

/**
 * 未读消息数量条数
*/
const totalUnreadCount: Ref<number> = ref<number>(0);  

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
@import 'assets/css/icon';
@import 'assets/css/list';
@import './messageList.scss';

</style>