<template>
    <div ref="searchRef">
        <ui-config-provider :locale="uiLocale.locale">
            <!-- 主体内容-表单操作 -->
            <ui-form
                class="operate-log-form"
                label-position="left"
                :model="ruleFormOperate.ruleForm"
                @submit.native.prevent
                @rule-form-ref="ruleFormOperate.getFormRef"
            >
                <!-- 操作人 -->
                <ui-form-item
                    prop="userName"
                    :label="$t('deviceDetail.operatLog.search.label.operatePeople')"
                >
                    <ui-input
                        class="custom-width custom-width1"
                        type="text"
                        clearable
                        v-model.trim="ruleFormOperate.ruleForm.userName"
                        :placeholder="$t('deviceDetail.operatLog.search.placeholder.operatePeople')"
                        @change="ruleFormOperate.userNameChange"
                        @clear="ruleFormOperate.clearUserName"
                    />
                </ui-form-item>
                <ui-form-item
                    prop="operateTime"
                    :label="$t('deviceDetail.operatLog.search.label.operateTime')"
                >
                    <ui-date-picker
                        class="custom-width custom-width4"
                        type="daterange"
                        v-model="ruleFormOperate.ruleForm.operateTime"
                        :disabled-date="ruleFormOperate.disabledDate"
                        :default-time="[new Date(0, 0, 0, 0, 0, 0), new Date(0, 0, 0, 23, 59, 59)]"
                        :start-placeholder="$t('deviceDetail.operatLog.search.placeholder.startDate')"
                        :end-placeholder="$t('deviceDetail.operatLog.search.placeholder.endDate')"
                        @change="ruleFormOperate.dateChange"
                    />
                </ui-form-item>
                <ui-form-item>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.searchClick"
                    >
                        {{$t('deviceDetail.operatLog.search.btn.search')}}
                    </ui-button>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.clearClick"
                    >
                        {{$t('deviceDetail.operatLog.search.btn.clear')}}
                    </ui-button>
                </ui-form-item>
        
            </ui-form>
        </ui-config-provider>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale} from 'utils/publicClass.ts';

interface PropsType {
    ruleFormOperate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

const searchRef = ref();
// defineEmits API 用来代替emit
const emitValue = defineEmits(['search-ref']);

onMounted(() => {
    emitValue('search-ref', searchRef)
})
</script>