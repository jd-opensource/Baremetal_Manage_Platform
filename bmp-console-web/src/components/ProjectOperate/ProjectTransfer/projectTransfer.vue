<template>
    <div class="middle-dialog dialog-trans" v-if="openVisible">
        <div class="dialog-width952">
            <el-dialog
                v-model="openVisible"
                center
                :close-on-click-modal="false"
                :show-close="false"
                destroy-on-close
            >
                <template #title>
                    <div >
                        <div class="folder dialog-icon"></div>
                        <span class="my-title">{{$t('list.operate.transferProject')}}</span>
                        <el-button  
                            circle 
                            class="close-button" 
                            :icon="Close"
                            @click="closeDialog()"
                        >
                        </el-button>
                    </div>
                </template>
                <!-- 步骤展示 -->
                <el-steps 
                    class="bmp-step"
                    :active="stepKey" 
                    align-center 
                    finish-status="success"
                >
                    <!-- 选择用户 -->
                    <el-step :title="$t('list.chooseUser')"/>
                    <!-- 选择资源 -->
                    <el-step :title="$t('list.chooseResource')"/>
                </el-steps>
                <!-- 第一步 选择用户 -->
                <div v-show="stepKey === 0">
                    <!-- 转移项目注意事项 -->
                    <div class="error-tip-content share-tip-position">
                        <div class="mt4">
                            <span class="warning-tips">
                                <span class="error error-icon mr0"></span>
                                {{$t('list.tip.transferUserTip0')}}
                            </span>
                        </div>
                        <div>
                            <p class="ml20 lh24">{{$t('list.tip.transferUserTip1')}}</p>
                            <p class="ml20 lh24">{{$t('list.tip.transferUserTip2')}}</p>
                            <p class="ml20 lh24">{{$t('list.tip.transferUserTip3')}}</p> 
                        </div>
                                      
                    </div>
                    <div class="dialog-body-content">
                        <el-form
                            label-position="left"
                            :model="transferUserForm.transferUserData"
                            ref="ruleFormRef"
                            :rules="rules"
                            label-width="auto"
                            :hide-required-asterisk="true"
                            @submit.native.prevent
                        >
                            <!-- 项目名称 -->
                            <el-form-item 
                                prop="projectName"
                                :label="$t('personCentre.detail.projectName')"
                            >
                                <span class="f12">{{projectInfo.projectName}}</span>  
                            </el-form-item>
                            <!-- 资源选择 -->
                            <el-form-item 
                                prop="resourceChoose"
                                class="mb42"
                                :label="$t('list.project.resourceChoose')"
                            >
                                <el-radio-group 
                                    v-model="transferUserForm.transferUserData.resourceChoose" 
                                    class="password-size" @change="setResourceChoose()"
                                >
                                    <el-radio label="1">{{$t('list.project.allTransfer')}}</el-radio>
                                    <el-radio label="2">{{$t('list.project.partTransfer')}}</el-radio>
                                </el-radio-group>
                                <div 
                                class="el-form-item__error">
                                {{transferUserForm.transferUserData.resourceChoose === '1' ? $t('list.content.transformAllTip') : $t('list.content.transformPartTip') }}
                            </div>
                            </el-form-item>
                             
                            <!-- 选择转移用户 -->
                            <el-form-item 
                                prop="user"
                                :label="$t('list.project.choosetransferUser')"
                            >
                                <el-input 
                                    v-model="transferUserForm.transferUserData.user" 
                                    type="text"
                                    maxlength="64"
                                    @change="setResourceChoose()"
                                
                                    :placeholder="$t('list.placeholder.transferUser')"
                                />
                            </el-form-item>
                            <!-- 选择项目 -->
                            <el-form-item 
                                v-if="transferUserForm.transferUserData.resourceChoose === '2'"
                                prop="chooseResource"
                                :label="$t('list.chooseProject')"
                                
                            >
                                <el-select 
                                    v-model="transferUserForm.transferUserData.chooseResource" 
                                    :placeholder="$t('personCentre.form.selectProject') "
                                >
                                    <el-option 
                                        v-for="item in projectList.tableData" 
                                        :key="item.value"
                                        :label="item.label" 
                                        :value="item.value"
                                    />
                                    <template #empty>
                                        <p class="el-select-dropdown__empty">{{$t('instance.create.noData')}}</p>
                                    </template>
                                </el-select>
                            </el-form-item>
                        </el-form>
                    </div>
                </div>
                <!-- 第二步 选择资源 -->
                <div v-show="stepKey === 1">
                    <div class="dialog-body-content">
                        <choose-instance
                            ref="chooseInstanceRef"
                            :allChooseTip="transferUserForm.transferUserData.resourceChoose === '1'? true : false"
                            :nextStepTip="nextStepTip"
                            :instanceData="reactiveArr.tableData"
                            @getData="getAllInstanceData()"
                            @selectedData="getSelectedData"
                        ></choose-instance>
                    </div>
                </div>
                <template #footer>
                    <span v-if="stepKey === 0 ">
                        <!-- 取消按钮 -->
                        <el-button 
                            type="primary" 
                            class="cancel-button"
                            @click="closeDialog()" 
                        >
                            {{$t('list.cancel')}}
                        </el-button>
                        <!-- 下一步 -->
                        <el-button 
                            type="primary" 
                            @mouseup="nextStep()"
                        >
                            {{$t('list.nextStep')}}
                        </el-button>
                    </span>
                    <span v-else-if="stepKey === 1 ">
                        <!-- 上一步 -->
                        <el-button 
                            type="primary" 
                            @click="previousStep()" 
                        >
                            {{$t('list.previousStep')}}
                        </el-button>
                        <!-- 确认按钮 -->
                        <el-button 
                            type="primary" 
                            :loading="isLoading" 
                            :disabled="confirmDisable"
                            @mouseup="confirmDialog()"
                        >
                            {{$t('list.confirm')}}
                        </el-button>
                    </span>
                </template>
            </el-dialog>
        </div>
    </div>
</template>
<script setup lang="ts">
import {
    Ref, 
    ref, 
    reactive, 
    watch, 
    nextTick, 
    onMounted,
    computed
} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import type {
    FormInstance, // element-plus ui类
    FormRules // element-plus ui类
} from 'element-plus';
import {
    transferUserAPI,
    transferPartUserAPI, 
    transferUserCheckAPI, 
    projectListAPI,
    instanceListAPI
} from 'api/request.ts'; // 接口
import chooseInstance from 'components/common/chooseInstance/chooseInstance.vue'; //实例选择表格

/**
 * transferUser item类型
 */
 type transferUserType ={
    projectName: string, // 项目名称
    resourceChoose: string, // 资源选择
    user: string, // 选择转移用户
    chooseResource: string, //选择资源
};

/**
 * transferUser类型
 */
interface transferUser {
    transferUserData: transferUserType
}

/**
 * transferUser表单
 */
const transferUserForm: transferUser = reactive<transferUser>({
    transferUserData: {
        projectName: '', // 项目名称
        resourceChoose: '1', // 资源选择
        user: '', // 选择转移用户
        chooseResource: '' // 选择资源
    }
})

/**
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * 资源选择切换
 */
 const setResourceChoose= () => {  
    transferUserForm.transferUserData.chooseResource = ''; // 选择资源重置

}

/**
 * 项目菜单类型类
*/
interface ProjectType {
    value:string;
    label:string;
};

/**
 * 项目类型
 */
 interface project {
    tableData: ProjectType[],
}

/**
 * 项目列表数据
 */
 const projectList: project = reactive<project>({
    tableData: [],
})

/**
 * 请求项目列表数据接口
*/
const requestProjectListData = (): void => {
    projectList.tableData  = [];
    projectListAPI({
        isAll:'1',
        sharerName: transferUserForm.transferUserData.user
    }).then(({data} : {data: any}) => {
        if (data?.result?.projects?.length) {
            projectList.tableData = data.result.projects.map((item: any) => {
                return {
                    value: item.projectId,
                    label: item.projectName,    
                }
            });
    }
    }).catch(({message} : {message: string;})=>{
        if (message.indexOf('undefined') > -1) return;
        projectList.tableData  = [];
    })
    .finally(() => {
    });
};


/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    projectInfo?: any;
    userName?: string;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    projectInfo: {},
    userName:''
});

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
    transferUser: UserInfoType[];
};

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 不可选择的转移用户错误信息
*/
const transferdUserErrorMessage: Ref<string> = ref<string>(''); 

/**
 * 选择转移用户正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const requestCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('list.tip.emptyTransferUser')));       
    } else {
        transferUserCheckAPI({
            checkedUserName: transferUserForm.transferUserData.user,
            projectId: props.projectInfo.projectId,
            operation: 'move'
        }).then((data: any) => {
            if(data?.data?.result?.success) {
                requestProjectListData(); // 获取资源项目
                callback();
            }
        }).catch(({message} : {message: string;})=>{
            if(message) {
                transferdUserErrorMessage.value = message
                callback(new Error(transferdUserErrorMessage.value));               
            }
        }).finally(() => {

        })
    }
    
}

/**
 * 选择资源校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const resourceCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('list.tip.emptyResoure')));
    }
    else {
        callback();
    }
};

// const transferUserCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
//     if (value === '' || undefined) {
//         callback(new Error(t('list.tip.emptyTransferUser')));
//     }
//     else if (isNotTransferdUser.value) {
//         callback(new Error(transferdUserErrorMessage.value));
//     }
//     else {
//         callback();
//     }
// };

/**
 * 规则
*/
const rules: RulesType = reactive<FormRules>({
    user: [ // 转移用户
        {
            required: true,
            trigger: 'blur',
            validator: requestCheck
        }
    ],
    chooseResource: [ // 选择资源
        {
            required: true,
            trigger: 'change',
            validator: resourceCheck
        }
    ]
});

/**
 * 回传对话框关闭
 */
const emitClose = defineEmits(['close', 'refresh']);

/**
 * 监听蒙层开关
 * @param {boolean} newValue 蒙层显示隐藏 true false
*/
watch(() => props.openVisible, (newValue: Required<boolean>): Readonly<void> => { 
    emitClose('close', newValue);
});

/**
 * 点击取消
 */
const closeDialog=()=> {  
    transferUserForm.transferUserData.user = '';
    ruleFormRef.value.resetFields();  
    stepKey.value = 0;
    emitClose('close', false)
}

/**
 * 请求用户验证接口
*/
// const requestCheck = (): void => {
//     isNotTransferdUser.value = true
//     transferUserCheckAPI({
//         checkedUserName: transferUserForm.transferUserData.user,
//         projectId: props.projectInfo.projectId,
//     }).then((data: any) => {
//         if(data?.data?.result?.success) {
//             isNotTransferdUser.value = false;
//             ruleFormRef.value.clearValidate(['user']) //即时显示验证信息
//         } else {
//             isNotTransferdUser.value = true
//         }
//     }).catch(({message} : {message: string;})=>{
//         if(message) {
//             isNotTransferdUser.value = true
//             transferdUserErrorMessage.value = message
//             ruleFormRef.value.validateField(['user']) //即时显示验证信息
            
//         }
//     }).finally(() => {

//     })
// }

/**
 * 请求转移全部实例接口
*/
const requestAllTransfer = (): void => {
    transferUserAPI({
        projectId: props.projectInfo.projectId,
        ownerName: props.userName,
        moverName: transferUserForm.transferUserData.user,
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            emitClose('refresh');
            closeDialog();
        }
    }).catch(() => {
        closeDialog();
        emitClose('refresh');
    }).finally(() => {
        isLoading.value = false;
    });
};

/**
 * 请求转移部分实例接口
*/
const requestPartTransfer = (): void => {
    let instanceIds = reactiveArr.selectedTableData.map((item: any) => item.instanceId).join(',')
    let params = {
        ownerProjectID: props.projectInfo.projectId,
        ownerName: props.userName,
        moverName: transferUserForm.transferUserData.user,
        moverProjectID: transferUserForm.transferUserData.chooseResource,
    }
    if(instanceIds) {
        params['instanceIDs'] = instanceIds
    }
    transferPartUserAPI({
        ...params
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            emitClose('refresh');
            closeDialog();
        }
    }).catch(() => {
        closeDialog();
        emitClose('refresh');
    }).finally(() => {
        isLoading.value = false;
    });
};
/**
 * 实例item类型
 */
 type InstancesType = {
    instanceName: string,
    instanceId: string,
}

/**
 * 实例类型
 */
 interface instance {
    tableData: InstancesType[],
    selectedTableData: InstancesType[],

}

/**
 * 实例列表数据
 */
 const reactiveArr: instance = reactive<instance>({
    tableData: [],
    selectedTableData: []
})

/**
 * 获取全部实例列表
 */
 const getAllInstanceData = (): void => {
    instanceListAPI({
        isAll:'1',
        projectId: props.projectInfo.projectId,
    }).then(({data} : {data: any}) => {
        if (data?.result?.instances?.length) {
            reactiveArr.tableData = data.result.instances.map((item: any) => {
                return {
                    'instanceName': item.instanceName,
                    'instanceId': item.instanceId,
                }
            });
        }
    }).catch (()=>{
        reactiveArr.tableData = [];
    })
    .finally(() => {
        if(transferUserForm.transferUserData.resourceChoose === '1') {
            chooseInstanceRef.value.handleCheckAllChange()
        }
    });
};

/**
 * 得到已选实例数据
 * @param {Array} data 已选实例数据
 */
const getSelectedData = (data: any): any=> {
    reactiveArr.selectedTableData = data
};

/**
 * 步骤key
 */
 const stepKey: Ref<number> = ref<number>(0);

/**
 * 实例资源选择ref
 */
const chooseInstanceRef = ref<any>();

/**
 * 点击上一步
 */
const previousStep = () => {  
    if ( stepKey.value > 0 ) {
        stepKey.value -- ;
        nextStepTip.value = false;
        chooseInstanceRef.value.handlePreStep() // 实例选择清空
    }
}

/**
 * 回车下一步
 */
document.onkeydown = (event: {keyCode: number;}) => {
    if (stepKey.value === 0) {
        event.keyCode === 13 && nextStep();
    } else if (stepKey.value === 1 && !confirmDisable.value) {
        event.keyCode === 13 && confirmDialog();
    }
    
};

/**
 * 已点击下一步提示
 */
const nextStepTip: Ref<boolean> = ref<boolean>(false);

/**
 * 点击下一步
 */
 const nextStep = async (): Promise<void> => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        ruleFormRef.value.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                if (stepKey.value < 1) {
                    stepKey.value ++ ;
                    nextStepTip.value = true

                }
            }
        });        
    });
};

/**
 * 确认键置灰逻辑
 */
const confirmDisable = computed(() => { 
    if (reactiveArr.selectedTableData.length || transferUserForm.transferUserData.resourceChoose === '1') {
        return false
    } else {
        return true
    }   
})

/**
 * 点击确认
 */
 const confirmDialog = () :void => {
    if(transferUserForm.transferUserData.resourceChoose === '1') {
        isLoading.value = true;
        requestAllTransfer();
    } else if (transferUserForm.transferUserData.resourceChoose === '2') {
        isLoading.value = true;
        requestPartTransfer();
    }
    
};
</script>

<style lang="scss" scope>
@import 'assets/css/middleDialog';
@import 'assets/css/icon';
@import 'assets/css/bmp-step';
@import './projectTransfer.scss';
</style>