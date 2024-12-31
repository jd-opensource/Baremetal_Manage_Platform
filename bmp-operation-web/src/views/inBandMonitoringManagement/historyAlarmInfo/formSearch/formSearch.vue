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
                    prop="userName"
                    :label="$filter.withClon($t('historyAlarmInfo.search.label.userName'))"
                >
                    <ui-input
                        class="custom-width"
                        type="text"
                        clearable
                        v-model.trim="ruleFormOperate.ruleForm.userName"
                        :placeholder="$t('historyAlarmInfo.search.placeholder.userName')"
                        @change="ruleFormOperate.userNameChange"
                        @clear="ruleFormOperate.clearUserName"
                    >
                    </ui-input>
                </ui-form-item>
                <!-- 报警时间 -->
                <ui-form-item
                    prop="alarmTime"
                    :label="$filter.withClon($t('historyAlarmInfo.search.label.alarmTime'))"
                >
                    <!-- 请选择报警时间 -->
                    <ui-date-picker
                        class="custom-width custom-width4"
                        type="daterange"
                        v-model="ruleFormOperate.ruleForm.alarmTime"
                        :disabled-date="ruleFormOperate.disabledDate"
                        :default-time="[new Date(0, 0, 0, 0, 0, 0), new Date(0, 0, 0, 23, 59, 59)]"
                        :start-placeholder="$t('historyAlarmInfo.search.placeholder.startDate')"
                        :end-placeholder="$t('historyAlarmInfo.search.placeholder.endDate')"
                        @change="ruleFormOperate.dateChange"
                        @blur="ruleFormOperate.handleBlur"
                    />
                </ui-form-item>
                <ui-form-item>
                    <ui-button
                        type="primary"
                        v-for="(item, index) in ruleFormOperate.btnData"
                        :key="index"
                        @click="ruleFormOperate.searchClear(item.type)"
                    >
                        {{item.text}}
                    </ui-button>
                </ui-form-item>
            </ui-form>
        </ui-config-provider>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale, language} from 'utils/publicClass.ts';

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

const searchRef = ref(0);

class RulesOpt {
    rules = reactive({
        userName: [
            {
                trigger: 'blur',
                required: true,
                message: language.t('historyAlarmInfo.search.errorTip.userName')
            }
        ],
        alarmTime: [
            {
                trigger: 'blur',
                required: true,
                message: language.t('historyAlarmInfo.search.errorTip.alarmTime')
            }
        ]
    })
}

const rulesOpt = new RulesOpt();
onMounted(() => {
    emitValue('search-ref', searchRef)
})
</script>
