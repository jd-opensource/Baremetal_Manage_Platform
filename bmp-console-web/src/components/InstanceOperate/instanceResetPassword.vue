<template>
    <div class="middle-dialog" v-if="openVisible">
        <el-dialog
            v-model="openVisible"
            center
            :close-on-click-modal="false"
            :show-close="false"
        >
            <template #title>
                <div>
                    <div class="list dialog-icon"></div>
                    <span class="my-title">{{$t('list.operate.resetPassword')}}</span>
                    <el-button     
                        circle 
                        class="close-button" 
                        :icon="Close"
                        @click="closeDialog" 
                    >
                    </el-button>
                </div>
            </template>
            <div class="batch-table-content m0 reset-password">
                <div class="error-tip-content batch-tip-error">
                    <div class="error error-icon"></div>
                    {{$t('list.content.repasswordTip')}}
                </div>
                <batch-table
                    v-bind="$attrs" 
                    batchTip="resetPassword"
                    :multipleSelection="instanceInfo"
                >
                </batch-table>
                <el-form
                    label-position="left"
                    :model="passwordForm.passwordData"
                    ref="ruleFormRef"
                    :rules="rules"
                    :hide-required-asterisk="false"
                    label-width="auto"
                    style="margin-top: 20px;"
                    @submit.native.prevent
                >
                    <!-- 新密码 -->
                    <div class="mb18">
                        <el-form-item  
                            prop="password"
                            ref="password"
                            :label="$t('personCentre.form.newPassword') " 
                        >
                            <el-input 
                                v-model="passwordForm.passwordData.password" 
                                maxlength="30"
                                minlength="8"
                                type="password"
                            />
                            <div 
                                class="f12 c999 el-form-item__error" 
                                v-if="!newPasswordState"
                            >
                                {{$t('personCentre.content.passwordVerification')}}
                            </div>
                        </el-form-item>
                    </div>
                   
                    <!-- 确认密码 -->
                    <el-form-item  
                        prop="confirmPassword"
                        ref="confirmPassword"
                        :label="$t('instance.create.confirmPassword')"
                    >
                        <el-input 
                            v-model="passwordForm.passwordData.confirmPassword" 
                            type="password"
                        />
                    </el-form-item>
                </el-form>
            </div>
            <template #footer>
                <span>
                    <el-button                          
                        class="cancel-button"
                        @click="closeDialog"
                    >
                        {{$t('list.cancel')}}
                    </el-button>
                    <el-button 
                        type="primary" 
                        :loading="isLoading"
                        @click="openDialog"
                    >
                        {{$t('list.confirm')}}
                    </el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>
<script setup lang="ts">
import {
    Ref, 
    ref, 
    watch, 
    reactive, 
    nextTick
} from 'vue';
import {
  Close,
} from '@element-plus/icons-vue';
import {ElMessage} from 'element-plus';
import {useI18n} from 'vue-i18n'; // 国际化
import {
    passwordReg // 密码正则
} from 'utils/regular.ts';
import type {
    FormInstance, // element-plus ui类
    FormRules // element-plus ui类
} from 'element-plus';
import {instanceRepasswordAPI} from 'api/request.ts'; // 重置密码接口
import batchTable from 'components/multipleOperate/batchTable.vue'; //实例选择表格
/**
 * 密码状态
 */
 const newPasswordState: Ref<boolean> = ref<boolean>(false)

/**
 * password item类型
 */
 type PasswordType ={
    password: string, //密码
    confirmPassword: string, //确认密码
};

/**
 * password类型
 */
interface password {
    passwordData:PasswordType
}

/**
 * 密码表单
 */
const passwordForm: password = reactive<password>({
    passwordData: {
        password: '',
        confirmPassword: '',
    }
})

/**
 * 国际化
*/
const {t} = useI18n();

/**
 * loading加载态
*/
const isLoading: Ref<boolean> = ref<boolean>(false);

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    openVisible: boolean;
    instanceInfo: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    openVisible: false,
    instanceInfo: []
});

/**
 * 用户信息类
*/
type UserInfoType = {
    required: boolean;
    message?: string;
    trigger: string;
    validator?: unknown;
};

/**
 * 规则类
*/
interface RulesType {
    password: UserInfoType[];
    confirmPassword:UserInfoType[]
};

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 密码正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const passwordCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    newPasswordState.value = true
    if (value === ''||undefined) {
        callback(new Error(t('personCentre.toolTip.emptyCurrentPassword')));
    }
    else if (!passwordReg.test(value)) {
        callback(new Error(t('personCentre.content.passwordVerification')));
    }
    else {
        callback();
    }
};

/**
 * 确认密码正则校验
 * @param {Object} _ 占位符
 * @param {string} value 输入项
 * @param {Function} callback 回调函数，返回对应状态
*/
const confirmPasswordCheck: unknown = (_: any, value: string, callback: (arg0?: Error | string) => void) => {
    if (value === ''||undefined) {
        callback(new Error(t('personCentre.toolTip.emptyconfirmPassword')));
    }
    else if (value !== passwordForm.passwordData.password) {
        callback(new Error(t('personCentre.toolTip.passNotSame')));
    }
    else {
        callback();
    }
};

/**
 * 规则
*/
const rules: RulesType = reactive<FormRules>({
    password: [ // 密码
        {
            required: false,
            trigger: 'blur',
            validator: passwordCheck
        }
    ],
    confirmPassword: [ // 确认密码
        {
            required: false,
            trigger: 'blur',
            validator: confirmPasswordCheck
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
const closeDialog:  () => void = () => {
    newPasswordState.value = false
    passwordForm.passwordData.password = '',
    passwordForm.passwordData.confirmPassword = '',
    emitClose('close', false)
}

/**
 * 请求重置密码接口
*/
const requestResetPassword = (): void => {
    instanceRepasswordAPI({
        instanceId: props.instanceInfo[0].instanceId,
        idcId: props.instanceInfo[0].idcId,
        password: passwordForm.passwordData.confirmPassword
    }).then(({data}: any) => {
        if(data?.result?.success) {
            ElMessage({
                message: t('personCentre.content.operateSuccess'),
                type: 'success'
            })
            closeDialog()
            emitClose('refresh');    
        }
    }).catch(() => {
        closeDialog()
        emitClose('refresh'); 
    }).finally(() => {
        isLoading.value = false;
    });
};

/**
 * 点击确定
 */
const openDialog: () => Promise<void> = async () => {
    // 判断输入项是否符合条件
    await nextTick(() => {
        ruleFormRef.value.validate((valid: boolean) => {
            // 输入符合条件后才可以继续后续操作
            if (valid) {
                isLoading.value = true;
                requestResetPassword();
            }
        });
    });
};
</script>
<style lang="scss">
@import 'assets/css/middleDialog';
@import 'assets/css/batch-table';
@import 'assets/css/icon';
.batch-tip-error {
    font-size: 12px;
    margin-bottom: 20px;
    text-align: center;
}

.reset-password {
    .el-input {
        width: 440px !important;
    }
    .el-form-item.is-error {
        margin-bottom: 0 !important;
    }
    .el-form-item__error {
        position: static;
        display: flex;
        align-items: center;
        word-break: keep-all;
        word-wrap: break-word;
    }

}
</style>