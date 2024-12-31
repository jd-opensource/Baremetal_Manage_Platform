<template>
    <!-- 登录页面 -->
    <div class="operate-content">
        <!-- 登录页标题 -->
        <div class="operate-content-header">
            <!-- 登录页大标题 -->
            <p class="operate-content-title">{{$t('login.header.title2')}}</p>
        </div>
        <!-- 登录页登录窗口 -->
        <div class="operate-login">
            <!-- 登录标题 -->
            <p class="operate-login-title">{{ $t('login.formSubmit.title')}}</p>
            <!-- 表单输入提交 -->
            <el-form
                class="operate-login-ruleForm"
                :model="ruleForm"
                :rules="rules"
                ref="ruleFormRef"
                @submit.native.prevent
            >
                <!-- 登录用户名 -->
                <el-form-item
                    prop="loginUserName"
                    :show-message="false"
                >
                    <div class="operate-content-ipt">
                        <!-- 用户icon -->
                        <img
                            class="operate-content-ipt-user-img"
                            src="@/assets/img/user-login.png"
                        />
                        <!-- 用户名输入框 -->
                        <el-input
                            v-model.trim="ruleForm.loginUserName"
                            :placeholder="$t('login.formSubmit.placeholder.userName')"
                        >
                        </el-input>
                        <!-- 清空icon -->
                        <img
                            class="operate-content-ipt-clear"
                            src="@/assets/img/clear.png"
                            alt=""
                            v-if="ruleForm.loginUserName.length"
                            @click="clearInput('loginUserName')"
                        />
                    </div>
                </el-form-item>
                <!-- 登录密码 -->
                <el-form-item prop="loginPassword">
                    <div class="operate-content-ipt">
                        <!-- open eye icon -->
                        <img
                            class="operate-content-ipt-open-eye"
                            src="@/assets/img/open-eye.png"
                            v-if="passwordType !== 'password'"
                            @click="passwordStatusChange('password')"
                        />
                        <!-- close eye icon -->
                        <img
                            class="operate-content-ipt-close-eye"
                            src="@/assets/img/close-eye.png"
                            v-else
                            @click="passwordStatusChange('text')"
                        />
                        <!-- 密码输入框 -->
                        <el-input
                            v-model.trim="ruleForm.loginPassword"
                            :type="passwordType"
                            :placeholder="$t('login.formSubmit.placeholder.passWord')"
                        >
                        </el-input>
                        <!-- 清空icon -->
                        <img
                            v-if="ruleForm.loginPassword.length" 
                            class="operate-content-ipt-clear"
                            src="@/assets/img/clear.png"
                            alt=""   
                            @click="clearInput('loginPassword')"
                        />
                    </div>
                </el-form-item>
                <!-- 异常提示 -->
                <div
                    class="error-tooltip"
                    v-if="errMessage"
                >
                    {{errMessage}}
                </div>
            </el-form>
            <!-- 登录按钮 -->
            <el-button
                class="operate-login-button"
                type="primary"
                :loading="isLoading"
                @click="loginOperate(ruleFormRef)"
            >
                {{$t('login.formSubmit.login')}}
            </el-button>
        </div>
        <div></div>
    </div>
</template>

<script setup lang="ts">
import {useI18n} from 'vue-i18n'; // 使用国际化
import {ref, Ref, reactive} from 'vue'; // 实现响应式数据的方法、ts类
import {loginAPI, userAPI} from 'api/request.ts'; // 登录接口
import {encrypt} from 'utils/index.ts';
import VueCookies from 'vue-cookies'; // cookie
import type {FormInstance, FormRules} from 'element-plus'; // ui ts
import { ElMessage } from 'element-plus';
import useProjectStore from '@/store/modules/projectId.ts';

/**
 * 默认ts类
*/
type DefStrType = string;
type DefBolType = boolean;
type DefVoiType = void;
type DefAnyType = any;

/**
 * 用户输入的类
*/
type UserInfoType = {
    required: DefBolType;
    message: DefStrType;
    trigger: DefStrType;
};

/**
 * 表单提交类
*/
interface RuleFormType {
    loginUserName: DefStrType;
    loginPassword: DefStrType;
};

/**
 * 表单规则类
*/
interface RulesType {
    loginUserName: UserInfoType[];
    loginPassword: UserInfoType[];
};

/**
 * 使用国际化
*/
const {t} = useI18n();

/**
 * 错误信息
 */
const errMessage: Ref<string> = ref<string>('')
    
/**
 * 表单输入项
*/
const ruleForm: RuleFormType = reactive<RuleFormType>({
    loginUserName: '', // 用户登录账号
    loginPassword: '' // 用户登录密码
});

/**
 * 表单ref
*/
const ruleFormRef: Ref<any> = ref<FormInstance>();

/**
 * 规则类型
*/
const rules: RulesType = reactive<FormRules>({
    loginUserName: [ // 用户登录账号
        {
            required: true,
            message: t('personCentre.toolTip.formSubmit'),
            trigger: 'click'
        }
    ],
    loginPassword: [ // 用户登录密码
        {
            required: true,
            message: t('personCentre.toolTip.formSubmit'),
            trigger: 'click'
        }
    ]
});

/**
 * 加载态-用于登录请求时，按钮的加载态
*/
const isLoading: Ref<DefBolType> = ref<DefBolType>(false);

/**
 * 清空输入项
 * @param {string} inputValue 输入框输入的值
 * @return {string} ruleForm[inputValue] 当前输入项内容清空
*/
const clearInput = (inputValue: DefStrType): DefStrType => {
    errMessage.value = '';
    return (ruleForm as any)[inputValue] = '';
};

/**
 * 密码type切换
*/
const passwordType: Ref<DefStrType> = ref<DefStrType>('password');

/**
 * 密码是否明文显示，状态切换
 * @param {string} type 密码明文显示的type值
 * @return {string} passwordType.value 密码明文显示的type值
*/
const passwordStatusChange = (type: DefStrType): DefStrType => {
    return passwordType.value = type;
};

const hasKeyDown = ref(false);
document.onkeydown = (event: {keyCode: number;}) => {
    if (hasKeyDown.value) return;
    event.keyCode === 13 && loginOperate(ruleFormRef.value);
};
/**
 * 登录操作，校验输入项
*/
const loginOperate = async (formEl: FormInstance | undefined): Promise<DefVoiType> =>{
    await formEl.validate((valid: DefBolType) => {
        // 输入项符合条件
        if (valid) {
            isLoading.value = true;
            requestLogin();
        }
    });
};

/**
 * 请求登录接口
*/
const requestLogin = (): DefVoiType => {
    hasKeyDown.value = true;
    errMessage.value = '';
    loginAPI({
        username: ruleForm.loginUserName,
        password: encrypt(ruleForm.loginPassword, 'ABCDEFGHIJKLMNOP', '0123456789ABCDEF')
    }).then(({data} : {data: DefAnyType}) => {
        if(data?.result?.success || data!==undefined) {
            requestUser()
                    
        }      
    }).catch(({message, status} : {message: string; status: number}) => {
        if (status === 400) {
            loadingKeyDownStatus();
            errMessage.value = message;
        }
        else {
            ElMessage({
                type: 'error',
                message,
                onClose: () => {
                    loadingKeyDownStatus();
                }
            });
        }
    }).finally(() => {
        isLoading.value = false;
    });
};

const loadingKeyDownStatus = () => {
    isLoading.value = false;
    hasKeyDown.value = false;
};

const projectStore = useProjectStore();
/**
 * 请求用户接口,获取projectId
*/
const requestUser = (): void => {
    userAPI({
    }).then(({data} : {data: any}) => {
        if(data?.result){
            (VueCookies as any).set('X-Jdcloud-Language', data.result.language); 
            projectStore.setProject(data.result.defaultProjectId, data.result.DefaultProjectName);
            window.open(`/instance/list?projectId=${data.result.defaultProjectId}&projectName=${data.result.DefaultProjectName}`, '_self');
        }                  
    }).finally(() => { 
       
    });
};

</script>
<style lang="scss">
@import './login.scss';
</style>
