<template>
    <div class="small-dialog no-center" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="folder dialog-icon"></div>
                    <span class="my-title">
                        {{Object.keys(props.instanceInfo).length ? $t('list.operate.editProject') : $t('list.operate.createProject')}}
                    </span>
                    <el-button                      
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog(ruleFormRef)"
                    >
                    </el-button>
                </div>
            </template>
            <div class="dialog-body-content">
                <el-form
                    label-position="left"
                    :model="projectForm"
                    ref="ruleFormRef"
                    :rules="rules"
                    label-width="auto"
                    @submit.native.prevent
                >
                    <el-form-item 
                        prop="projectName"
                        :label="$t('list.project.projectName') "
                    >
                        <el-input 
                            v-model="projectForm.projectName"
                            maxlength="64" 
                            :placeholder="$t('personCentre.placeholder.projectName')"
                        />
                        <p class="f12 c999 w390 el-form-item__error" v-if="!projectNameState">{{$t('personCentre.content.stringChineseVerification')}}</p>
                    </el-form-item>
                </el-form>
            </div>
            <template #footer>
                <span>
                    <el-button      
                        class="cancel-button"
                        @click="closeDialog(ruleFormRef)"
                    >
                        {{$t('list.cancel')}}
                    </el-button>
                    <el-button 
                        type="primary" 
                        :loading="isLoading"
                        @click="openDialog(ruleFormRef)"
                    >
                        {{$t('list.confirm')}}
                    </el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>
<script setup lang="ts">
import {Ref, ref, reactive, watch, nextTick, onMounted} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {
    projectNameReg // 项目名称正则
} from 'utils/regular.ts';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import type {
    FormInstance, // element-plus ui类
    FormRules // element-plus ui类
} from 'element-plus';
import {deepCopy} from 'utils/index.ts';
import {projectCreateAPI, projectEditAPI} from 'api/request.ts'; // 项目创建接口
import VueCookies from 'vue-cookies'; // cookie
import allProjectStore from '@/store/modules/headerDetail.ts';

/**
 * 国际化
*/
const {t} = useI18n();
const projectForm = reactive({
    projectName: '',
})
/**
 * 用户ID
*/
let userId: string = (VueCookies as any).get('user_id');

/**
 * 用户ID解码
 */
userId = window.atob(userId);

/**
 * loading加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    operateState: string;
    instanceInfo?: any;
};
/**
 * prop组件传值
 */
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    operateState: '',
    instanceInfo: {}
});

/**
 * 项目名称状态
 */
 const projectNameState: Ref<boolean> = ref<boolean>(false)

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
    projectName: UserInfoType[];
};

/**
 * store库存储的项目列表
*/
const store = allProjectStore();
/**
 * 项目名称正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const projectNameCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === '' || undefined) {
        projectNameState.value = true
        callback(new Error(t('personCentre.toolTip.emptyProjectName')));
    }
    else if (store.projectListName.includes(value)) {
        projectNameState.value = true
        callback(new Error(t('personCentre.content.projectNameExist')));
    }
    else if (!projectNameReg.test(value)) {
        projectNameState.value = true
        callback(new Error(t('personCentre.content.stringChineseVerification')));
    }
    else {
        projectNameState.value = true
        callback();
    }
};
/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();
/**
 * 项目规则
*/
const rules: RulesType = reactive<FormRules>({
    projectName: [ // 邮箱
        {
            required: false,
            trigger: 'blur',
            validator: projectNameCheck
        }
    ],
});
// 名称重复除自己以外的数组
const repeatDataNotSelf = reactive<any>([])
/**
 * 编辑项目时，项目名称赋值
 */
onMounted(() => {
    projectForm.projectName = Object.keys(props.instanceInfo).length ? props.instanceInfo.projectName:'';
    const repeatDeepCopyData = deepCopy(store.projectListName);
    repeatDeepCopyData.map((item:any, index:number)=>{
        if(item === props.instanceInfo.projectName){
            repeatDeepCopyData.splice(index, 1)
        }
    })
    repeatDataNotSelf.value = repeatDeepCopyData;
})
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
const closeDialog = (formEl: FormInstance | undefined)=> {
    formEl.resetFields();
    projectNameState.value = false; 
    isLoading.value = false;
    emitClose('close', false)
}

/**
 * 请求创建接口
*/
const requestCreate = (): void => {
    projectCreateAPI({
        projectId: props.instanceInfo.projectId,
        projectName: projectForm.projectName,
        userId: userId
    }).then((res: any) => {
        if (res) {
            projectNameState.value = false;
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            
        }
    }).finally(() => {
        emitClose('refresh');
        emitClose('close', false)
        isLoading.value = false;
    });
};

/**
 * 请求编辑接口
*/
const requestEdit = (): void => {
    isLoading.value = true;   
    projectEditAPI({
        projectId: props.instanceInfo.projectId,
        projectName: projectForm.projectName,
    }).then(({data}: any) => {
        if (data?.result?.success) {
            projectNameState.value = false;
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            emitClose('refresh');
            emitClose('close', false)
        }
    }).catch(() => {
        emitClose('refresh');
        emitClose('close', false)
    }).finally(() => {
        isLoading.value = false;
    });
};


document.onkeydown = (event: {keyCode: number;}) => {
    event.keyCode === 13 && openDialog(ruleFormRef.value);
};

/**
 * 点击确定
 */
const openDialog = async (formEl: FormInstance | undefined): Promise<void> => {
    
    // 判断输入项是否符合条件
    await nextTick(() => {
        formEl.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                if(props.operateState === 'create') {
                    requestCreate();
                } else if(props.operateState === 'edit') {
                    requestEdit();
                }
                
            }
        });
    });
};
</script>
<style lang="scss">
@import 'assets/css/smallDialog';
@import 'assets/css/icon';
</style>