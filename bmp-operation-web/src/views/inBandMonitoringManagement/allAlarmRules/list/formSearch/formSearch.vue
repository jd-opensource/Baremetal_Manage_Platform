<template>
    <div ref="searchRef">
        <ui-config-provider :locale="uiLocale.locale">
            <!-- 主体内容-表单操作 -->
            <ui-form
                class="public-form set-ipt-error"
                label-position="left"
                :model="ruleFormOperate.ruleForm"
                :rules="rulesOpt.rules"
                @submit.native.prevent
                @rule-form-ref="ruleFormOperate.getFormRef"
            >
                <!-- 用户名 -->
                <ui-form-item
                    v-for="(item, index) in Object.freeze(ruleFormOperate.formData)"
                    :prop="item.prop"
                    :key="index"
                    :label="$filter.withClon(item.label)"
                >
                    <ui-input
                        class="custom-width"
                        type="text"
                        clearable
                        v-model.trim="ruleFormOperate.ruleForm[item.model]"
                        :placeholder="item.placeholder"
                        @change="ruleFormOperate.iptChange(item.change)"
                        @clear="ruleFormOperate.clearIpt(item.clear)"
                    >
                    </ui-input>
                </ui-form-item>
                <ui-form-item>
                    <ui-button
                        type="primary"
                        v-for="(item, index) in ruleFormOperate.btnData"
                        :key="index"
                        @click="ruleFormOperate.searchClearClick(item.type)"
                    >
                        {{item.text}}
                    </ui-button>
                </ui-form-item>
            </ui-form>
        </ui-config-provider>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale} from 'utils/publicClass.ts';
import RulesOpt from './rules';

// props类型
type PropsType = {
    ruleFormOperate: any;
};

/**
 * props参数
 * @param {boolean} diaLog 弹窗状态
*/
withDefaults(defineProps<PropsType>(), {});

// emit 事件
const emitValue = defineEmits(['search-ref']);

const rulesOpt = new RulesOpt();

const searchRef = ref(0);

onMounted(() => {
    emitValue('search-ref', searchRef)
})

</script>
