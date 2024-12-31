<template>
    <!-- 创建密钥操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'create-api-key'"
        :is-loading="createApiKeyOperate.isLoading"
        :header-title="$t('createKey.header.title')"
        :src="($defInfo.imagePath('apiKey') as unknown as string)"
        @dia-log-close="createApiKeyOperate.cancelClick"
        @determine-click="createApiKeyOperate.determineClick"
    >
        <!-- 主体内容 -->
        <div class="create-apicontent">
            <ui-form
                label-position="left"
                class="apikey-ruleForm"
                :label-width="setDiffClass('85px', '70px')"
                :model="formRulesOperate.ruleForm"
                :rules="regExpValidate.rules"
                @submit.native.prevent
                @rule-form-ref="formRulesOperate.getFormRef"
            >
                <!-- 密钥 -->
                <ui-form-item
                    prop="name"
                    class="api-key-name"
                    :label="$t('createKey.form.label.keyName')"
                    :show-message="false"
                >
                    <ui-input
                        type="text"
                        maxlength="64"
                        v-model.trim="formRulesOperate.ruleForm.name"
                        :placeholder="$t('createKey.form.placeholder.keyName')"
                    />
                    <span
                        :class="[
                            regExpValidate.hasErrorStatus.value ? 'error-tip' :'name-tip',
                            locationItem.getLocationItem === 'en_US'
                                && regExpValidate.keyNameTip.value.length > 20
                                ? 'set-error-tip'
                                : ''
                        ]"
                    >
                        {{regExpValidate.keyNameTip.value}}
                    </span>
                </ui-form-item>
                <!-- 权限 -->
                <ui-form-item
                    prop="readOnly"
                    class="read-only"
                    :label="$t('createKey.form.label.permissions')"
                >
                    <ui-select 
                        v-model="formRulesOperate.ruleForm.readOnly"
                        :placeholder="$t('createKey.form.placeholder.permissions')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in formRulesOperate.reactiveArr.readOnlyData"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value"
                        />
                    </ui-select>
                </ui-form-item>
            </ui-form>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts';
import {locationItem} from 'utils/publicClass.ts';
import regExpCheck from './utils/regExp';
import formRules from './utils/formRules';
import createApiKey from './utils/createApiKey';

/**
 * props类
*/
interface PropsType {
    diaLog: boolean;
    keyNameData: string[];
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
 * @param {boolean} diaLog 弹窗状态
 * @param {boolean} hasEditOperate 是否进行编辑操作 默认true
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false
});

// defineEmits API 来替代emit，进行事件回传
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const regExpValidate = regExpCheck(props);


const formRulesOperate = formRules();

const createApiKeyOperate = createApiKey(formRulesOperate, emitValue);

</script>
<style lang="scss">
@import './createApiKey.scss';
</style>
