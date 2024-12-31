<template>
    <div class="page-position" :class="tabName==='relateUser' ? 'pr7':'can-scroll'">
        <div class="detail-header pb5">
            <el-button
                class="back-button"
                type= "text"
                :icon ="ArrowLeft" 
                @click="goBack"
            >
                {{$t('list.back')}}
            </el-button>
            <span class="ml32">{{ $t('personCentre.detail.projectDetail') }}</span>
            <span class="ml15">/</span>
            <div style="display: inline-flex;">
                <el-tooltip  
                    :disabled="!tooltipStatus['projectName'].showTooltip"  
                    :content= projectInfo.detail.projectName
                    placement="top-start"
                >
                    <span 
                        class="long-name-title ml15" 
                        @mouseenter="hasShowTooltip($event, tooltipStatus['projectName'], projectInfo.detail.projectName, 179, 'detail')"
                    >
                        {{ projectInfo.detail.projectName || $filter.emptyFilter() }}
                    </span>
                </el-tooltip>
            </div>
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
                            <!-- 编辑 -->
                            <el-dropdown-item>
                                <el-tooltip
                                    placement="right"
                                    :disabled="projectInfo.detail.owned === 1"
                                    :content="$t('list.tip.ownedOnly')"
                                >
                                    <label>
                                        <el-button
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateEdit()"
                                            :disabled="projectInfo.detail.owned !== 1"
                                        >
                                            <span>{{$t('personCentre.operate.edit')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 转移 -->
                            <el-dropdown-item>
                                <el-tooltip
                                    placement="right"
                                    :disabled="projectInfo.detail.owned === 1"
                                    :content="$t('list.tip.ownedOnly')"
                                >
                                    <label>
                                        <el-button
                                            type="text" 
                                            class="operate-btn"
                                            @click="clickOperateTransfer()"
                                            :disabled="projectInfo.detail.owned !== 1"
                                        >
                                            <span>{{$t('personCentre.operate.transfer')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                            <!-- 删除 -->
                            <el-dropdown-item>
                                <el-tooltip
                                    placement="right"
                                    :disabled="projectInfo.detail.owned === 1 && !projectInfo.detail.instanceCount"
                                    :content="deleteChangeTip()"
                                >
                                    <label>
                                        <el-button
                                            type="text"
                                            class="operate-btn"
                                            @click="clickOperateDelete()"
                                            :disabled="(projectInfo.detail.instanceCount ? true : false) || projectInfo.detail.owned !== 1"
                                        >
                                            <span>{{$t('personCentre.operate.delete')}}</span>
                                        </el-button>
                                    </label>
                                </el-tooltip>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </div> 
        </div>
        <div class="project-detail-content">
            <el-tabs
				ref="pageTabs"
                class="tab-info"
				type="border-card" 
				v-model="tabName"
            >
				<!-- tab:资源信息 -->
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
                                <!-- 项目名称 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('personCentre.detail.projectName') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['projectName'].showTooltip"   
                                            :content= projectInfo.detail.projectName
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['projectName'], projectInfo.detail.projectName, 259, 'detail')"
                                            >
                                                {{ projectInfo.detail.projectName || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                        <img
                                            v-if="projectInfo.detail.owned === 1"
                                            class="desc-type"
                                            src="@/assets/img/describe.png"
                                            @click="clickEditProjectName()"
                                        />
                                    </el-form-item>
                                </el-col>
                                <!-- 项目ID -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('personCentre.detail.projectId') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['projectId'].showTooltip"   
                                            :content= projectInfo.detail.projectId
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['projectId'], projectInfo.detail.projectId, 259, 'detail')"
                                            >
                                                {{ projectInfo.detail.projectId || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 项目协作状态 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('personCentre.detail.projectStatus') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ ownedChange(projectInfo.detail.owned) || $filter.emptyFilter()  }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 创建者 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('personCentre.detail.creater') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['creater'].showTooltip"   
                                            :content= projectInfo.detail.createdBy
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['creater'], projectInfo.detail.creater, 259, 'detail')"
                                            >
                                                {{ projectInfo.detail.createdBy || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 拥有者 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('personCentre.detail.owner') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            :disabled="!tooltipStatus['owner'].showTooltip"   
                                            :content= projectInfo.detail.ownedBy
                                            placement="top-start"
                                        >
                                            <span 
                                                class="long-name"
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['owner'], projectInfo.detail.owner, 259, 'detail')"
                                            >
                                                {{ projectInfo.detail.ownedBy || $filter.emptyFilter() }}
                                            </span>
                                        </el-tooltip>
                                    </el-form-item>
                                </el-col>
                                <!-- 实例数 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('list.project.instanceNumber') + $filter.withClon(' ')"
                                    >
                                        <span
                                            class="f12 default-color mouse-point"
                                            @click="jumpToInstance(projectInfo.detail)"
                                        >
                                            {{ projectInfo.detail.instanceCount || $filter.numberEmptyFilter()}}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 创建时间 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('personCentre.detail.createdTime') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ projectInfo.detail.createdTime || $filter.emptyFilter()  }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 更新时间 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('personCentre.detail.updatedTime') + $filter.withClon(' ')"
                                    >
                                        <span class="f12">
                                            {{ projectInfo.detail.updatedTime || $filter.emptyFilter()  }}
                                        </span>
                                    </el-form-item>
                                </el-col>
                                <!-- 描述 -->
                                <el-col :span="8">
                                    <el-form-item 
                                        :label="$t('instance.detail.description') + $filter.withClon(' ')"
                                    >
                                        <el-tooltip  
                                            popper-class="desc-tooltip-popper"
                                            :disabled="!tooltipStatus['description'].showTooltip || projectInfo.detail.description === undefined"  
                                            :content= projectInfo.detail.description
                                            placement="bottom"
                                        >
                                            <span 
                                                class="long-name" 
                                                @mouseenter="hasShowTooltip($event, tooltipStatus['description'], projectInfo.detail.description, 259, 'detail')">
                                                {{ projectInfo.detail.description}}
                                            </span>
                                        </el-tooltip>
                                        <img
                                            v-if="projectInfo.detail.owned === 1"
                                            class="desc-type"
                                            src="@/assets/img/describe.png"
                                            @click="clickEditDescription()"
                                        />
                                    </el-form-item>
                                </el-col>
                            </el-row>  
                        </el-form>    
                    </div>
                </el-tab-pane>
                <!-- tab: 关联共享用户 -->
                <el-tab-pane 
                    v-if="projectInfo.detail.owned === 1"
					:label="$t('personCentre.detail.relatedUser')" 
					name="relateUser"
					:disabled="tabName==='relateUser'"
                >
                    <div class="pane-content bl-none"> 
                        <el-tooltip
                            placement="bottom"
                            :disabled="projectInfo.detail.owned === 1"
                            :content="$t('list.tip.ownedOnly')"
                        >
                            <label>   
                                <el-button 
                                    type="primary" 
                                    class="add-button-style ml0" 
                                    :disabled="projectInfo.detail.owned !== 1"
                                    @click="userOperateAdd"
                                >
                                    {{$t('personCentre.detail.addUpdatedUser') }}
                                </el-button>
                            </label>
                        </el-tooltip>
                        <div class="table-content ml0 mr0">
                            <el-table
                                border 
                                :data="relateUserData.tableData" 
                                :max-height="relateUserData.tableData.length ? tableMaxHeight : 390" 
                                style="width: auto"
                            >
                                <!-- 用户名 -->
                                <el-table-column 
                                    prop="sharedUserName" 
                                    :label="$t('personCentre.detail.userName')" 
                                    min-width="100"
                                >
                                    <template v-slot="scope">
                                        <el-tooltip
                                            v-model="scope.row.showTooltip"
                                            :disabled="!scope.row.showTooltip"
                                            :content= scope.row.sharedUserName
                                        >
                                            <p 
                                                class="long-row"
                                                @mouseenter="hasShowTooltip($event, scope.row, scope.row.name, 1.17, 'list')"
                                            >
                                                {{scope.row.sharedUserName || $filter.emptyFilter()}}
                                            </p>
                                        </el-tooltip>
                                    </template>
                                </el-table-column>
                                <!-- 用户ID -->
                                <el-table-column 
                                    prop="sharedUserId" 
                                    :label="$t('personCentre.detail.userId')" 
                                    min-width="100"
                                />
                                <!-- 权限 -->
                                <el-table-column 
                                    prop="power" 
                                    :label="$t('personCentre.detail.power')" 
                                    min-width="100"
                                >
                                    {{$t('personCentre.apiPower.writeAndRead') }}
                                </el-table-column>
                                <!-- 操作 -->
                                <el-table-column :label="$t('personCentre.form.operation')">
                                    <template v-slot="scope">
                                        <operate-unit
                                            :operateName="$t('personCentre.detail.delete')"
                                            :disabled="projectInfo.detail.owned !== 1"
                                            @handleClick="clickRelatedUserDelete(scope.row)"
                                            :instanceInfo="scope.row"
                                            :operateTip="$t('list.tip.ownedOnly')"
                                        >
                                        </operate-unit>
                                    </template>
                                </el-table-column>
                                <template #empty>
                                    <div class="noData">                        
                                    </div>
                                    <p>
                                        <span class="c333">{{$t('personCentre.detail.noData')}}</span>
                                        <el-button 
                                            type="text" 
                                            class="f12 pt7" 
                                            :disabled="projectInfo.detail.owned !== 1"
                                            @click="userOperateAdd"
                                        >
                                            {{$t('personCentre.detail.addUpdatedUser') }}
                                        </el-button>
                                    </p>
                                </template>
                            </el-table>
                        </div>
                        <div v-if="deleteRelatedUser">
                            <!-- 移除 -->
                            <related-user-delete
                                :openVisible="deleteRelatedUser"
                                :projectInfo="projectRowInfo"
                                :userName="userName"
                                @close="deleteRelatedUserCancel"
                                @refresh="refreshData"
                            >
                            </related-user-delete>
                        </div>
                    </div>  
                </el-tab-pane>
            </el-tabs>
            <!-- 项目删除 -->
            <div v-if="deleteProject">
                <project-delete
                    :openVisible="deleteProject"
                    :instanceInfo="projectInfo.detail"
                    @close="deleteCancel"
                    @refresh="refreshData"
                >
                </project-delete>
            </div>  
            <!-- 项目编辑 -->
            <div v-if="editProject">
                <project-create
                    :openVisible="editProject"
                    :instanceInfo="projectInfo.detail"
                    :operateState="'edit'"
                    @close="editCancel"
                    @refresh="refreshData"
                >
                </project-create>
            </div> 
            <!-- 项目转移 -->
            <div v-if="transferProject">
                <project-transfer
                    :openVisible="transferProject"
                    :projectInfo="projectInfo.detail"
                    :userName="userName"
                    @close="transferCancel"
                    @refresh="goBack"
                >
                </project-transfer>
            </div> 
            <div v-if="editProjectName">
                <project-name-edit
                    :openVisible="editProjectName"
                    :instanceInfo="projectInfo.detail"
                    :operateState="'edit'"
                    @close="editProjectNameCancel"
                    @refresh="refreshData"
                >
                </project-name-edit>
            </div>
            <!-- 项目描述修改 -->
            <div v-if="editDescription">
                <description-edit
                    :openVisible="editDescription"
                    :projectInfo="projectInfo.detail"
                    @close="editDescriptionCancel"
                    @refresh="refreshData"
                >
                </description-edit>
            </div>
            <!-- 添加共享用户 -->
            <div v-if="addRelateUser">
                <relate-user-add
                    :openVisible="addRelateUser"
                    :projectInfo="projectInfo.detail"
                    :userName="userName"
                    @close="addRelateUserCancel"
                    @refresh="refreshData"
                >
                </relate-user-add>
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
} from 'vue';
import { useRouter, useRoute } from 'vue-router';
import {
    ArrowLeft,
} from '@element-plus/icons-vue';
import {
    hasShowTooltip // 是否显示提示
} from 'utils/index.ts';
import {projectDetailAPI, projectListAPI} from 'api/request.ts';
import operateUnit from 'components/OperateUnit/operateUnit.vue';
import allProjectStore from '@/store/modules/headerDetail.ts';
import projectCreate from 'components/ProjectOperate/projectCreate.vue';
import projectTransfer from 'components/ProjectOperate/ProjectTransfer/projectTransfer.vue';
import projectDelete from 'components/ProjectOperate/projectDelete.vue';
import projectNameEdit from 'components/ProjectOperate/editName.vue';
import descriptionEdit from 'components/ProjectOperate/editDescription.vue';
import relateUserAdd from 'components/ProjectOperate/RelateUserAdd/relateUserAdd.vue';
import relatedUserDelete from 'components/ProjectOperate/relatedUserDelete.vue';
import i18n from 'lib/i18n/index.ts'; // 国际化
import useProjectStore from '@/store/modules/projectId.ts';

// 计算表格的最大高度
const tableMaxHeight: Ref<number> = ref<number>(window.innerHeight - 310)
// 监听窗口大小变化，以便动态调整表格高度
const updateSize = () => {
  // 触发响应式更新
  tableMaxHeight.value = window.innerHeight - 310
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
 * 路由带值
 */
 const route = useRoute();
 /**
 * store库存储的项目列表
*/
const store = allProjectStore();
/**
 * 项目管理tab跳转标志
 */
let tabName: Ref<any> = ref('');
/**
 * 项目管理tab点击跳转到指定页面
 */
 onMounted(() => {
    tabName.value = 'basicInfo';
    requestProjectDetailData();
});
/**
 * 操作下拉状态
 */
const selectStatus = ref(false);
const selectHover = (val: boolean) => {
    selectStatus.value = val;
};

interface ReactiveType {
    detail: any;
} 

/**
 * 详情页信息
 */
 const projectInfo: ReactiveType = reactive<ReactiveType>({
    detail: {},
})

/**
 * 个人关联共享用户item类型
 */
 type UserType = {
    userName: string,
    userId: string,
    power: string,
    operation: string,
}
/**
 * 关联共享用户类型
 */
 interface user {
    tableData: UserType[],
}

/**
 * 关联共享用户数据
 */
 const relateUserData: user = reactive<user>({    
    tableData: [],
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
    'projectNameInfo': {
        showTooltip: false
    },
    'projectName': {
        showTooltip: false
    },
    'projectId': {
        showTooltip: false
    },
    'creater': {
        showTooltip: false
    },
    'owner': {
        showTooltip: false
    },
    'description': {
        showTooltip: false
    }
})

/**
 * 返回列表页
 */
const router = useRouter();
const goBack = () => {
    router.push({
		path: `/management`,
	});
}

const projectStore = useProjectStore();
/**
 * 跳入实例列表页
 */
 const jumpToInstance = (project: any) => {  
    projectStore.setProject(project.projectId, project.projectName);
    router.push({
        path: `/instance/list`,
        query: {
            projectId: project.projectId,
            projectName: project.projectName
        }
    });
    
}
/**
 * 项目名称编辑操作打开标志
 */
 const editProjectName: Ref<boolean> = ref<boolean>(false)  
/**
 * 项目名称编辑操作
 * @param value 
 */
 const clickEditProjectName: () => void = () => {
    editProjectName.value = !editProjectName.value;
}
/**
 * 项目名称编辑弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} editProjectName.value 弹窗关闭
 */
 const editProjectNameCancel = (type: boolean): boolean => {
    return editProjectName.value = type;
};
/**
 * 删除操作置灰提示
 * @return {string} 置灰提示文案
 */
const deleteChangeTip = (): any => {
    if(projectInfo.detail.owned !== 1) {
        return global.t('list.tip.ownedOnly')
    }   
    if(projectInfo.detail.instanceCount) {
        return global.t('list.tip.delete')
    }   
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
 * loadig态
*/
const isLoading: Ref<boolean> = ref<boolean>(true);
/**
 * 用户名
 */
const userName: Ref<any> = ref<any>('')
/**
 * 请求实例详情数据接口
*/
const requestProjectDetailData = (): void => {
    projectDetailAPI({
        projectId: route.query.projectId
    }).then(({data} : {data: any}) => {
        if (data?.result) {
            projectInfo.detail = data.result;
            relateUserData.tableData = data.result.shareProjects
            return;
        }
        return Promise.reject();
    }).catch(({message} : {message: string;}) => {       
        if (message === '找不到对象' || message === 'Not found') { 
            goBack();
            return;
        }
        if (message.indexOf('undefined') > -1) return;
    }).finally(() => {
        isLoading.value = false;
        store.requestObject();
        //projectInfo.repeatData = store.projectList.map((item: any) => item.projectName);
        //store.requestUser();
        if(store.userForm.userName) {
            sessionStorage.setItem('userName', store.userForm.userName);
        }
        userName.value = sessionStorage.getItem('userName')
    });
};

/**
 * 总条数
*/
const totalCount: Ref<number> = ref<number>(0);


/**
 * 编辑操作打开标志
 */
 const editProject: Ref<boolean> = ref<boolean>(false)

/**
 * 编辑操作
 * @param value 
 */
 const clickOperateEdit: () => void = () => {
    editProject.value = !editProject.value;
}

/**
 * 编辑项目弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
 const editCancel = (type: boolean): boolean => {
    return editProject.value = type;
};

/**
 * 转移操作打开标志
 */
 const transferProject: Ref<boolean> = ref<boolean>(false)
/**
 * 转移操作
 * @param value 
 */
 const clickOperateTransfer: () => void = () => {
    transferProject.value = !transferProject.value;
}
/**
 * 转移项目弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
 const transferCancel = (type: boolean): boolean => {
    return transferProject.value = type;
};

/**
 * 删除操作打开标志
 */
 const deleteProject: Ref<boolean> = ref<boolean>(false)

/**
 * 删除操作
 * @param value 
 */
 const clickOperateDelete: () => void = () => {
    deleteProject.value = !deleteProject.value;
}

/**
 * 删除项目弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} open.value 弹窗关闭
 */
 const deleteCancel = (type: boolean): boolean => {
    return deleteProject.value = type;
};

/**
 * 描述编辑操作打开标志
 */
 const editDescription: Ref<boolean> = ref<boolean>(false)   
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
 * 共享用户添加操作打开标志
 */
 const addRelateUser: Ref<boolean> = ref<boolean>(false);
/**
 * 添加共享用户操作
 * @param value 
 */
 const userOperateAdd: () => void = () => {
    addRelateUser.value = !addRelateUser.value;
}
/**
 * 添加共享用户弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} addRelateUser.value 弹窗关闭
 */
 const addRelateUserCancel = (type: boolean): boolean => {
    return addRelateUser.value = type;
}; 

/**
 * 移除共享用户打开标志
 */
 const deleteRelatedUser: Ref<boolean> = ref<boolean>(false)
/**
 * 当前行值
 */
 const projectRowInfo: Ref<any> = ref();
/**
 * 移除共享用户操作
 * @param value 
 */
 const clickRelatedUserDelete: (value: object) => void = (value: object) => {
    projectRowInfo.value = value 
    deleteRelatedUser.value = !deleteRelatedUser.value;
}
/**
 * 移除共享用户弹窗取消
 * @param type false 弹窗关闭
 * @return {boolean} deleteRelatedUser.value 弹窗关闭
 */
 const deleteRelatedUserCancel = (type: boolean): boolean => {
    return deleteRelatedUser.value = type;
};

/**
 * 刷新接口
*/
const refreshData = (): void => {
    isLoading.value = true;
    requestProjectDetailData();
};
</script>

<style lang="scss">
@import 'assets/css/page';
@import './projectManagementDetail.scss';
@import 'assets/css/table';
@import 'assets/css/page';
@import 'assets/css/addButton';
</style>