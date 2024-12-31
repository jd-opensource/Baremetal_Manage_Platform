<template>
    <div class="middle-dialog dialog-pading-change" v-if="openVisible">
        <div class="dialog-width952">
            <el-dialog
                v-model="openVisible"
                center
                :close-on-click-modal="false"
                :show-close="false"
                destroy-on-close
            >
                <template #title>
                    <div>
                        <div class="folder dialog-icon"></div>
                        <span class="my-title">{{$t('personCentre.detail.addUpdatedUser')}}</span>
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
                    <!-- 共享项目注意事项 -->
                    <div class="error-tip-content share-tip-position">
                        <div class="mt4">
                            <div class="error error-icon"></div>
                            <span class="warning-tips">{{$t('personCentre.toolTip.relateUserTip0')}}</span>
                        </div>
                        <div>
                            <p class="ml20 lh24">{{$t('personCentre.toolTip.relateUserTip1')}}</p>
                            <p class="ml20 lh24">{{$t('personCentre.toolTip.relateUserTip2')}}</p>
                            <p class="ml20 lh24">{{$t('personCentre.toolTip.relateUserTip3')}}</p>  
                        </div>             
                    </div>
                    <div class="dialog-body-content">
                        <el-form
                            label-position="left"
                            :model="relateUserForm.relateUserData"
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
                            <!-- 选择共享用户 -->
                            <el-form-item 
                                prop="user"
                                :label="$t('personCentre.detail.chooseRelateUser')"
                            >
                                <el-input 
                                    v-model="relateUserForm.relateUserData.user" 
                                    type="text"
                                    maxlength="64"
                                    @change="requestCheck"
                                    :placeholder="$t('list.placeholder.chooseRelateUser')"
                                />
                            </el-form-item>
                            <!-- 权限 -->
                            <el-form-item
                                prop="power"
                                :label="$t('personCentre.form.power') "
                            >
                                <el-select 
                                    v-model="relateUserForm.relateUserData.power" 
                                    :placeholder="$t('personCentre.placeholder.select')"
                                >
                                    <el-option 
                                        v-for="item in relateUserPowerData"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value"
                                    />
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
                            :allChooseTip="false"                          
                            :projectId=props.projectInfo.projectId
                            :nextStepTip="nextStepTip"
                            :instanceData=" reactiveArr.tableData"
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
import {relateUserPowerType} from 'utils/constants.ts'; // 共享用户权限类型数据
import {
    relatedUserAddAPI, 
    relatedUserCheckAPI, 
    shareInstanceListAPI
} from 'api/request.ts'; // 接口
import chooseInstance from 'components/common/chooseInstance/chooseInstance.vue'; //实例选择表格

/**
 * relateUser item类型
 */
 type RelateUserType ={
    projectName: string, // 项目名称
    user: string, // 选择共享用户
    power: string, // 权限
};

/**
 * relateUser类型
 */
interface relateUser {
    relateUserData: RelateUserType
}

/**
 * relateUser表单
 */
const relateUserForm: relateUser = reactive<relateUser>({
    relateUserData: {
        projectName: '', // 项目名称
        user: '', // 选择共享用户
        power: '', // 权限
    }
})

/**
 * relateUser数据类型类
*/
interface DataType {
    value: string;
    label: string;
};

/**
 * relateUser类型数据
*/
const relateUserPowerData: DataType[] = reactive<DataType[]>(relateUserPowerType);

/**
 * loading 加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 国际化
*/
const {t} = useI18n();

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
    relateUser: UserInfoType[];
    power: UserInfoType[];
};

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 判断共享用户是否不能选择
*/
const isNotRelatedUser: Ref<boolean> = ref<boolean>(false);

/**
 * 不可选择的共享用户错误信息
*/
const relatedUserErrorMessage: Ref<string> = ref<string>('');  

/**
 * 选择共享用户正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const relateUserCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        callback(new Error(t('personCentre.toolTip.emptyRelateUser')));
    }
    else if (isNotRelatedUser.value) {
        callback(new Error(relatedUserErrorMessage.value));
    }
    else {
        callback();
    }
};

/**
 * 规则
*/
const rules: RulesType = reactive<FormRules>({
    user: [ // 共享用户
        {
            required: true,
            trigger: 'blur',
            validator: relateUserCheck
        }
    ],
    power: [ // 权限
        {
            required: true,
            trigger: 'change',
            message: t('personCentre.toolTip.emptyPower')
        }
    ],
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
const closeDialog = () => {  
    stepKey.value = 0;  
    relateUserForm.relateUserData.user = '';
    relateUserForm.relateUserData.power = '';
    ruleFormRef.value.resetFields();  
    emitClose('close', false)
}

/**
 * 请求用户验证接口
*/
const requestCheck = (): void => {
    relatedUserCheckAPI({
        checkedUserName: relateUserForm.relateUserData.user,
        projectId: props.projectInfo.projectId,
        operation: 'share'
    }).then((data: any) => {
        if(data?.data?.result?.success) {
            isNotRelatedUser.value = false;
            ruleFormRef.value.clearValidate(['user'])
        } else {
            isNotRelatedUser.value = true
        }
    }).catch(({message} : {message: string;})=>{
        if(message) {
            isNotRelatedUser.value = true
            relatedUserErrorMessage.value = message
            ruleFormRef.value.validateField(['user']) //即时显示验证信息
            
        }
    }).finally(() => {

    })
}

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
 * 实例资源选择ref
 */
 const chooseInstanceRef = ref<any>();

/**
 * 获取全部实例列表
 */
 const getAllInstanceData = (): void => {
    shareInstanceListAPI({
        projectId: props.projectInfo.projectId,
        ownerName: props.userName,
        sharerName: relateUserForm.relateUserData.user
    }).then(({data} : {data: any}) => {
        if (data?.result?.length) {
            reactiveArr.tableData = data.result.map((item: any) => {
                return {
                    'instanceName': item.instanceName,
                    "instanceId": item.instanceId,
                    tooltipStatus: {
                        'hasShared': item.hasShared 
                    }
                    
                }
            });
        }
    }).catch (()=>{
        reactiveArr.tableData = [];
    })
    .finally(() => {       
        // 已共享的资源设为选中状态
        if(reactiveArr.tableData.length) {
            reactiveArr.tableData.map((item: any) => {
                if (item.tooltipStatus.hasShared) {
                    chooseInstanceRef.value.selectChange(item)
                }    
            }) 
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
 * 已点击下一步提示
 */
 const nextStepTip: Ref<boolean> = ref<boolean>(false);

/**
 * 点击上一步
 */
const previousStep=()=> {  
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
 * 请求添加用户接口
*/
const requestAdd = (): void => {
    let instanceIds = reactiveArr.selectedTableData.map((item: any) => item.instanceId).join(',')
    let params = {
        projectId: props.projectInfo.projectId,
        ownerName: props.userName,
        sharerName: relateUserForm.relateUserData.user,
    }
    if(instanceIds) {
        params['instanceIDs'] = instanceIds
    }
    relatedUserAddAPI({
        ...params
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            closeDialog();
            emitClose('refresh');

        }
    }).catch(() => {
        closeDialog();
        emitClose('refresh')
    }).finally(() => {
        isLoading.value = false;
    });
};


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
    if (!reactiveArr.tableData.length) {
        return false
    } else {
        if (reactiveArr.selectedTableData.length) {
            return false
        }
        return true
    }   
})

/**
 * 点击确认键
 */
 const confirmDialog = () :void => {
    isLoading.value = true;
    requestAdd();
};
</script>
<style lang="scss" scope>
@import 'assets/css/middleDialog';
@import 'assets/css/icon';
@import 'assets/css/bmp-step';
@import './relateUserAdd.scss';

</style>