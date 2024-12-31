<template>
    <div ref="searchRef">
        <ui-config-provider :locale="uiLocale.locale">
            <ui-form
                class="fault-form"
                label-position="left"
                :model="ruleFormOperate.ruleForm"
                @submit.native.prevent
                @rule-form-ref="ruleFormOperate.getFormRef"
            >
                <!-- 机房名称 -->
                <ui-form-item
                    prop="idcName"
                    :label="$filter.withClon($t('hardwareStatusList.search.label.idcName'))"
                >
                    <!-- 请选择机房 -->
                    <ui-select
                        class="custom-width custom-width1"
                        v-model="ruleFormOperate.ruleForm.idcId"
                        :loading="store.idcInfo.loading"
                        :placeholder="$t('hardwareStatusList.search.placeholder.idcName')"
                    >
                        <ui-option
                            style="fontSize: 12px"
                            v-for="item in store.idcInfo.idcData"
                            :key="item.name"
                            :label="item.newIdcName"
                            :value="item.idcId"
                        />
                    </ui-select>
                </ui-form-item>
                <!-- sn -->
                <ui-form-item
                    prop="sn"
                    :label="$filter.withClon($t('hardwareStatusList.search.label.sn'))"
                >
                    <ui-input
                        class="custom-width custom-width2"
                        type="text"
                        v-model.trim="ruleFormOperate.ruleForm.sn"
                        :placeholder="$t('hardwareStatusList.search.placeholder.sn')"
                        :clearable="true"
                        @change="ruleFormOperate.snChange"
                        @clear="ruleFormOperate.clearSn"
                    />
                </ui-form-item>
                <ui-form-item>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.searchClick"
                    >
                        {{$t('hardwareStatusList.search.btn.search')}}
                    </ui-button>
                    <ui-button
                        type="primary"
                        @click="ruleFormOperate.clearClick"
                    >
                        {{$t('hardwareStatusList.search.btn.clear')}}
                    </ui-button>
                </ui-form-item>
            </ui-form>
        </ui-config-provider>
    </div>
</template>
<script lang="ts" setup>
import {uiLocale, CurrentInstance} from 'utils/publicClass.ts';
import store from 'store/index.ts';

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

const proxy = new CurrentInstance().proxy;

store.idcInfo.getIdc(proxy.$idcApi.idcListAPI);

const searchRef: Ref = ref(0);

onMounted(() => {
    emitValue('search-ref', searchRef)
})
</script>
