<template>
    <!-- 安全验证操作 -->
    <div
        class="security-verification"
        v-if="diaLog"
    >
        <!-- 安全验证弹窗 -->
        <custom-dia-log
            :dia-log="diaLog"
            :header-title="$t('securityVerification.header.verification')"
            :is-loading="verificationOperate.isLoading"
            :src="($defInfo.imagePath('password') as unknown as string)"
            @dia-log-close="verificationOperate.cancelClick"
            @determine-click="verificationOperate.determineClick"
        >
            <!-- 主体内容-表单操作 -->
            <ui-form
                label-width="80px"
                label-position="left"
                class="security-verification-ruleForm"
                :model="formRulesOperate.ruleForm"
                :rules="formRulesOperate.rules"
                @submit.native.prevent
                @rule-form-ref="formRulesOperate.getFormRef"
            >
                <!-- 用户名 -->
                <ui-form-item
                    prop="username"
                    :label="$t('securityVerification.label.userName')"
                    :show-message="false"
                >
                    <ui-input
                        type="text"
                        maxlength="64"
                        v-if="!formRulesOperate.ruleForm.username"
                        v-model.trim="formRulesOperate.ruleForm.username"
                        :placeholder="$t('securityVerification.placeholder.userName')"
                    >
                    </ui-input>
                    <span v-else>
                        {{formRulesOperate.ruleForm.username}}
                    </span>
                </ui-form-item>
                <!-- 密码 -->
                <ui-form-item
                    prop="password"
                    :label="$t('securityVerification.label.passWord')"
                >
                    <ui-input
                        v-model.trim="formRulesOperate.ruleForm.password"
                        type="password"
                        maxlength="64"
                        :placeholder="$t('securityVerification.placeholder.passWord')"
                    >
                    </ui-input>
                </ui-form-item>
            </ui-form>
        </custom-dia-log>
    </div>
</template>
  
<script lang="ts" setup>
import {encrypt, customTip} from 'utils/index.ts';
import {securityVerificationAPI} from 'api/public/request.ts'; // 安全验证接口
import {language} from 'utils/publicClass.ts';
import {RulesType, RuleFormType, RuleFormRefType2} from './typeManagement';

/**
 * props类
*/
interface PropsType {
    diaLog: boolean;
    hasEditOperate?: boolean;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
 * @param {boolean} diaLog 弹窗状态
 * @param {boolean} hasEditOperate 是否进行编辑操作 默认true
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
    hasEditOperate: true
});

// defineEmits API 来替代emit，进行事件回传
const emitValue = defineEmits(['dia-log-close', 'determine-click', 'has-edit-operate']);

class FormRulesOperate {
    // ref 操作
    ruleFormRef: Ref<RuleFormRefType2 | undefined> = ref<RuleFormRefType2>();
    // 表单输入项
    ruleForm: RuleFormType = reactive<RuleFormType>({
        username: '', // 用户名
        password: '' // 密码
    });
    rules: RulesType = reactive<RulesType>({
        username: [ // 用户名
            {
                required: true,
                trigger: 'click'
            }
        ],
        password: [ // 密码
            {
                required: true,
                trigger: 'click',
                message: language.t('securityVerification.empty')
            }
        ]
    });

    /**
     * 获取表单ref
     * @param {Object} formEl 表单的信息，可以进行清空、验证是否输入正确的操作
    */
    getFormRef = (formEl: {value: RuleFormRefType2}): void => {
        this.ruleFormRef.value = formEl.value;
    };
};
const formRulesOperate: FormRulesOperate = new FormRulesOperate();

/**
 * 验证操作
*/
class VerificationOperate {
    // loading加载态，默认false，用于发送请求之前的loading
    isLoading: Ref<boolean> = ref<boolean>(false);
    // 构造器
    constructor () {
        // 登录名称
        let loginUserName: string = localStorage.getItem('userName') || '';
        loginUserName = loginUserName ? window.atob(loginUserName) : '';
        formRulesOperate.ruleForm.username = loginUserName;
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.determineClick();
        };
    };

    /**
     * 确定按钮弹窗
    */
    determineClick = async (): Promise<void> => {
        // 判断输入项是否符合条件
        await nextTick(() => {
            formRulesOperate.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.isLoading.value = true;
                    this.requestVerificationInterface();
                }
            });
        });
    };

    /**
     * 请求安全验证接口，成功后把事件回传，关闭弹窗
    */
    requestVerificationInterface = (): void => {
        securityVerificationAPI(
            {
                username: formRulesOperate.ruleForm.username,
                password:  encrypt(formRulesOperate.ruleForm.password)
            }
        ).then(({data} : {data: {result: {success: boolean;}}}) => {
            if (data?.result?.success) {
                this.cancelClick();
                this.isLoading.value = false;
                if (!props.hasEditOperate) {
                    emitValue('has-edit-operate', props.hasEditOperate);
                }
                else {
                    emitValue('determine-click');
                }
            }
        })
        .catch(({message}: {message: string}) => {
            customTip('error', message, 800, () => this.isLoading.value = false);
        });
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 清空输入内容、异常提示状态
        formRulesOperate.ruleFormRef.value!.resetFields();
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};
// 实例化
const verificationOperate: VerificationOperate = new VerificationOperate();

</script>
<style lang="scss">
@import './securityVerification.scss';
</style>
