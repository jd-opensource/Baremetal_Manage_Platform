<template>
    <div class="page-position project-list">
        <div class="page-content pos-fixed">
            <div class="operation-header">
                <el-button 
                    type="primary" 
                    class="add-button-style" 
                    @click="operateCreate"
                >
                    + &nbsp;&nbsp;{{$t('list.createdProject') }}
                </el-button>
                <div class="setting-position">
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
            <div class="table-content">
                <el-config-provider :locale="locale">
                    <el-table 
                        border
                        :data="reactiveArr.tableData" 
                        v-loading="isLoading"
                        style="width: 100%"
                        :max-height="reactiveArr.tableData.length ? tableMaxHeight : 380" 
                        @filter-change="filterChange"
                    >
                        <!-- 项目名称 -->
                        <el-table-column 
                            prop="projectName" 
                            v-if="reactiveArr?.hasCustomInfo['projectName']?.selected"
                            :label="$t('list.project.projectName')" 
                        >
                            <template v-slot="scope">  
                                <el-tooltip
                                    v-model="scope.row.showTooltip"
                                    :disabled="!scope.row.showTooltip"
                                    :content= scope.row.projectName
                                >
                                    <div 
                                        class="long-row" 
                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.projectName, 1.17, 'list')" 
                                    >
                                        <span
                                            class="project-name-content"
                                            @click="jumpToDetail(scope.row)"
                                        >
                                            {{scope.row.projectName || $filter.emptyFilter()}}
                                        </span>
                                    </div>
                                </el-tooltip>
                            </template>                        
                        </el-table-column>
                        <!-- 项目ID -->
                        <el-table-column 
                            prop="projectId" 
                            v-if="reactiveArr?.hasCustomInfo['projectId']?.selected"
                            :label="$t('list.project.projectId')" 
                        >
                            <template v-slot="scope">
                                <el-tooltip    
                                    v-model="scope.row.showTooltip"
                                    :disabled="!scope.row.showTooltip"                           
                                    :content= scope.row.projectId
                                    placement="top-start"
                                >
                                    <p 
                                        class="long-row"
                                        @mouseenter="hasShowTooltip($event, scope.row, scope.row.projectId, 1.17, 'list')"
                                    >
                                    {{scope.row.projectId || $filter.emptyFilter()}}</p>
                                </el-tooltip>
                            </template>
                        </el-table-column>
                        <!-- 项目协作状态状态 -->
                        <el-table-column 
                            v-if="reactiveArr?.hasCustomInfo['owned']?.selected"
                            prop="owned" 
                            key="owned"
                            column-key="owned"
                            :label="$t('list.project.relatedState')" 
                            filter-placement="bottom-end"
                            :filters="statusStore.status"
                            :filter-method="statusFilter"
                        >
                            <template v-slot="scope">
                                {{ ownedChange(scope.row.owned) || $filter.emptyFilter()}} 
                            </template>
                        </el-table-column>
                        <!-- 实例数 -->
                        <el-table-column 
                            prop="instanceCount" 
                            v-if="reactiveArr?.hasCustomInfo['instanceCount']?.selected"
                            :label="$t('list.project.instanceNumber')" 
                            width="90"
                        >
                            <template v-slot="scope">
                                <el-button 
                                    type="text"
                                    class="f12"
                                    @click="jumpToInstance(scope.row)"
                                >
                                    {{ scope.row.instanceCount || $filter.numberEmptyFilter()}}
                                </el-button>
                            </template>
                        </el-table-column>
                        <!-- 创建时间 -->
                        <el-table-column 
                            prop="createdTime" 
                            v-if="reactiveArr?.hasCustomInfo['createdTime']?.selected"
                            :label="$t('personCentre.form.createdTime')" 
                        >
                            <template v-slot="scope">
                                {{scope.row.createdTime || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 更新时间 -->
                        <el-table-column 
                            prop="updatedTime" 
                            v-if="reactiveArr?.hasCustomInfo['updatedTime']?.selected"
                            :label="$t('personCentre.form.updatedTime')" 
                        >
                            <template v-slot="scope">
                                {{scope.row.updatedTime || $filter.emptyFilter()}}
                            </template>
                        </el-table-column>
                        <!-- 操作 -->
                        <el-table-column 
                            prop="operation" 
                            :label="$t('personCentre.form.operation')" 
                        >
                            <template v-slot="scope">
                                <operate-unit
                                    :operateName="$t('personCentre.operate.edit')"
                                    :operateTip="$t('list.tip.ownedOnly')"
                                    :disabled="scope.row.owned !== 1"
                                    :instanceInfo="scope.row"
                                    @handleClick="operateEdit(scope.row)"
                                >
                                </operate-unit>
                                <operate-unit
                                    :operateName="$t('personCentre.operate.transfer')"
                                    :operateTip="$t('list.tip.ownedOnly')"
                                    :disabled="scope.row.owned !== 1"
                                    :instanceInfo="scope.row"
                                    @handleClick="operateTransfer(scope.row)"
                                >
                                </operate-unit>
                                <operate-unit
                                    :operateName="$t('personCentre.operate.delete')"
                                    :operateTip="deleteChangeTip(scope.row)"
                                    :disabled="(scope.row.instanceCount || totalCount <= 1 ? true : false) || scope.row.owned !== 1"
                                    :instanceInfo="scope.row"
                                    @handleClick="operateDelete(scope.row)"
                                >
                                </operate-unit>
                            </template>
                        </el-table-column>
                        <template #empty>
                            <div class="noData">                        
                            </div>
                            <p>
                                <span class="c333">{{$t('list.project.noData')}}</span>
                                <el-button
                                    type="text"
                                    class="f12 pt7" 
                                    @click="operateCreate"
                                >
                                    {{$t('list.operate.createProject') }}
                                </el-button>
                            </p>
                        </template>
                    </el-table>
                </el-config-provider>   
            </div>
            <div class="footer-count">
                <my-pagination
                    v-if="reactiveArr.tableData.length > 0 && !isLoading"
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
            <div v-if="openCustom">
                <custom
                    :page-key="'projectList'"
                    :open-visible="openCustom"
                    :check-list-arr="reactiveArr.checkListArr"
                    :has-custom-info="reactiveArr.hasCustomInfo"
                    @close="openCustomCancel"
                    @determine="publicStore.customList('projectList', reactiveArr)"
                >
                </custom>
            </div> 
            <div v-if="deleteProject">
                <project-delete
                    :openVisible="deleteProject"
                    :instanceInfo="instanceInfo"
                    @close="deleteProjectCancel"
                    @refresh="refreshData"
                >
                </project-delete>
            </div>                      
            <div v-if="createProject">
                <project-create
                    :openVisible="createProject"
                    :instanceInfo="instanceInfo"
                    :operateState="operateState"
                    @close="createProjectCancel"
                    @refresh="refreshData"
                >
                </project-create>
            </div>  
            <div v-if="transferProject">
                <project-transfer
                    :openVisible="transferProject"
                    :projectInfo="instanceInfo"
                    :userName="userName"
                    @close="transferProjectCancel"
                    @refresh="refreshData"
                >
                </project-transfer>
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
    nextTick
} from 'vue';
import { useRouter } from 'vue-router';
import {
    getDate, // 获取日期
    hasShowTooltip, // 是否显示提示
    languageSwitch, // 语言切换
    generateRandomNum // 生成随机数
} from 'utils/index.ts';
import projectCreate from 'components/ProjectOperate/projectCreate.vue';
import projectDelete from 'components/ProjectOperate/projectDelete.vue';
import projectTransfer from 'components/ProjectOperate/ProjectTransfer/projectTransfer.vue';
import operateUnit from 'components/OperateUnit/operateUnit.vue';
import custom from 'components/custom/custom.vue';
import myPagination from 'components/Pagination/MyPagination.vue';
import {projectListStore} from 'store/loginInfo.ts'; // store库存储的用户名
import VueCookies from 'vue-cookies'; // cookie
import {projectListAPI, projectListExportAPI} from 'api/request.ts'; // 项目列表数据接口
import allProjectStore from '@/store/modules/headerDetail.ts';
import publicIndexStore from 'store/index.ts'; // 公共store
import projectDetailStore from 'store/modules/projectDetail.ts'; // store库里的oss数据
import useProjectStore from '@/store/modules/projectId.ts';
import {
    porjectCustomInfo, // 项目列表自定义初始数据
    porjectCustomIntro
} from 'utils/constants.ts';
import i18n from 'lib/i18n/index.ts'; // 国际化

// 计算表格的最大高度
const tableMaxHeight: Ref<number> = ref<number>(window.innerHeight - 290)
// 监听窗口大小变化，以便动态调整表格高度
const updateSize = () => {
  // 触发响应式更新
  tableMaxHeight.value = window.innerHeight - 290
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
 * 语言
*/
const locale: any = languageSwitch();

/**
 * 公共store
*/
const publicStore: any = publicIndexStore();

/**
 * store库存储的状态类
*/
const storeUser = projectListStore();

/**
 * store库存储的oss数据类
*/
const statusStore: any = projectDetailStore();
/**
 * store库存储的项目列表
*/
const store = allProjectStore();
/**
 * 用户ID
*/
let userId: string = (VueCookies as any).get('user_id');

/**
 * 用户ID解码
 */
userId = window.atob(userId);

/**
 * 项目item类型
 */
 type ProjectType = {
    createdTime: string;
    projectName: string;
    instanceCount: number;
    projectId: string;
    updatedTime: string;
};

/**
 * 项目类型
 */
 interface project {
    tableData: ProjectType[],
    filterParams: any,
    hasCustomInfo: {},
    checkListArr: {}
}

/**
 * 项目列表数据
 */
const reactiveArr: project = reactive<project>({
    tableData: [],
    filterParams: {
        owned: 0
    },
    hasCustomInfo: porjectCustomIntro,
    checkListArr: porjectCustomInfo
})
/**
 * 自定义列表接口
*/
publicStore.customList('projectList', reactiveArr);
/**
 * 当前行值
 */
 const instanceInfo: Ref<any> = ref();
/**
 * 自定义列表操作打开标志
 */
 const openCustom: Ref<boolean> = ref<boolean>(false)
/**
 * 项目创建操作打开标志
 */
 const createProject: Ref<boolean> = ref<boolean>(false)
/**
 * 项目转移操作打开标志
 */
 const transferProject: Ref<boolean> = ref<boolean>(false)
/**
 * 项目删除操作打开标志
 */
 const deleteProject: Ref<boolean> = ref<boolean>(false)
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
 * 判断是操作编辑还是创建
 */
const operateState: Ref<string> = ref<string>('')
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
        projectListExportAPI(
            {
                exportType: '1',
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

/**
 * 管理状态filter
 * @param {number} value 当前选中的value值
 * @param {Object} row 当前选中的这一项
 * @return {boolean} xxx 返回对应信息
*/
const statusFilter = (value: number, row: any): boolean => {
    return row.owned === statusStore?.status[value - 1]?.value;
};

/**
 * filter参数类
*/
interface FilterParamsType {
    owned: string;
};

const filterChange = (filter: {[x: string]: any;}) => {
    const filterParams: FilterParamsType | any = {
        owned: 'owned',
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
};

/**
 * 项目创建操作
 * @param value 
 */
 const operateCreate: () => void = () => {
    instanceInfo.value = {};
    operateState.value = 'create'
    createProject.value = !createProject.value;
} 
/**
 * 项目创建弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} createProject.value 弹窗关闭
 */
const createProjectCancel = (type: boolean): boolean => {
    return createProject.value = type;
};
/**
 * 项目编辑操作
 * @param value 
 */
 const operateEdit: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    operateState.value = 'edit'
    createProject.value = !createProject.value;
} 
/**
 * 项目转移操作
 * @param value 
 */
 const operateTransfer: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    transferProject.value = !transferProject.value;
} 
/**
 * 项目转移弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} transferProject.value 弹窗关闭
 */
 const transferProjectCancel = (type: boolean): boolean => {
    return transferProject.value = type;
};  
/**
 * 项目删除操作
 * @param value 
 */
 const operateDelete: (value: object) => void = (value: object) => {
    instanceInfo.value = value 
    deleteProject.value = !deleteProject.value;
}
/**
 * 项目删除弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} deleteProject.value 弹窗关闭
 */
 const deleteProjectCancel = (type: boolean): boolean => {
    return deleteProject.value = type;
};
/**
 * 跳入实例列表页
 */
 const router = useRouter();
 const projectStore = useProjectStore();
 const jumpToInstance = (row: any) => {  
    projectStore.setProject(row.projectId, row.projectName);
    router.push({
        path: `/instance/list`,
        query: {
            projectId: row.projectId,
            projectName: row.projectName
        }
    });
    
}

/**
 * 跳入项目详情页
 */
const jumpToDetail = (row: any) => {
    router.push({
        path: `/management/detail`,
        query: {
            projectId: row.projectId,
            type: 'management'
        }
    });
}
/**
 * 项目协作状态修改
 * @param  row  状态key
 * @return {string} 状态value
 */
 const ownedChange = (row: number): string => {
    switch (row) {
        case 1:
        return global.t('list.project.owner');
        case 2:
        return global.t('list.project.cooperate');
        default:
        return '';
    }
};
/**
 * 删除操作置灰提示
 * @param  row  置灰条件
 * @return {string} 置灰提示文案
 */
const deleteChangeTip = (row: any): any => {
    if(row.owned !== 1) {
        return global.t('list.tip.ownedOnly')
    } 
    if(totalCount.value <= 1) {
        return global.t('list.tip.deleteOnly')
    }  
    if(row.instanceCount) {
        return global.t('list.tip.delete')
    }   
}
/**
 * 用户名
 */
 const userName: Ref<any> = ref<any>('')

/**
 * 请求项目列表数据接口
*/
const requestProjectListData = (params: any): void => {
    store.requestObject();
    projectListAPI({
        ...params
    }).then(({data, userId} : {data: any, userId: any}) => {
        
        if(data?.result?.projects?.length) {
            storeUser.userId = window.btoa(userId);
            (VueCookies as any).set('user_id', storeUser.userId);
            reactiveArr.tableData = data.result.projects.map((item: any) => {
                return {
                    ...item,
                    showTooltip: false
                }
            });
            totalCount.value = data.result.totalCount;
            return;
    }
    reactiveArr.tableData  = [];
    totalCount.value = 0;
    
    }).catch (({message} : {message: string;})=>{
        if (message.indexOf('undefined') > -1) return;
        reactiveArr.tableData  = [];
        totalCount.value = 0;
    })
    .finally(() => {
        isLoading.value = false;
        
        //store.requestUser();
        if(store.userForm.userName) {
            sessionStorage.setItem('userName', store.userForm.userName);
        }
        userName.value = sessionStorage.getItem('userName')       
    });
};

/**
 * 获取全部项目数据
*/
// const getAllProjectList = (): void => {
//     projectListAPI({
//         isAll:1,
//     }).then(({data} : {data: any}) => {
//         if(data?.result?.projects?.length) {
//             reactiveArr.repeatData = data?.result?.projects.map((item: any) => item.projectName);
//         } else {
//             reactiveArr.repeatData = [];
//         }
//     }).catch(() => {
//         reactiveArr.repeatData = [];
//     })
    
// };
onMounted(() => {
    processingParameter();
});

/**
 * 刷新接口
*/
const refreshData = (): void => {
    isLoading.value = true;
    pageNumber.value = 1;
    pageSize.value = 20;
    processingParameter();
};

/**
 * 每页展示条数切换
 * @param {number} size 每页展示条数
*/
const pageSizesChange = (size: number) => {
    pageNumber.value = 1;
    pageSize.value = size;
    isLoading.value = true;
    processingParameter();
};

/**
 * 页数切换
 * @param {number} num 页数切换
*/
const changePage = (num: number): void => {
    pageNumber.value = num;
    isLoading.value = true;
    processingParameter();
};

/**
 * 处理参数
*/
const processingParameter = () => {
    const params = {
        pageSize: pageSize.value,
        pageNumber: pageNumber.value,
        ...reactiveArr.filterParams,
    };
    publicStore.deleteEmtpyData(params);
    requestProjectListData(params);
};

</script>

<style lang="scss">
@import 'assets/css/page';
@import 'assets/css/table';
@import 'assets/css/addButton';
@import 'assets/css/icon';
@import 'assets/css/list';
@import './projectManagement.scss';
</style>