<template>
    <ui-form-item
        prop="noConfigurationData"
        id="volumeManager"
        :label="$t('modelForm.label.volumeManage.title')"
        :show-message="false"
    
    >
        <ui-config-provider :locale="uiLocale.locale">
            <ui-table
                class="volume-manager-table"
                style="width: 100%"
                :data="formRulesOperate.ruleForm.noConfigurationData"
                border
                :class="tableClass(formRulesOperate.ruleForm.noConfigurationData, false)"
                v-loading="false"
            >
                <ui-table-column
                    align="center"
                    fixed
                    :label="$t('modelForm.label.volumeManage.label.name')"
                    :min-width="setDiffClass('260', '170')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div>
                            <span v-if="scope.row.editFlag">
                                {{$filter.emptyFilter(scope.row.volumeName)}}
                            </span>
                            <ui-form v-else>
                                <ui-form-item
                                    prop=""
                                    class="ipt-height"
                                    :label="''"
                                >
                                <ui-input
                                    v-model="scope.row.volumeName"
                                    :placeholder="$t('modelForm.label.volumeManage.placeholder.tip1')"
                                    :class="[
                                        ['success', void 0].includes(scope.row[`volumeNameFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                    ]"
                                    @blur="formRulesOperate.volumeNameBlur(scope.row, scope.row.id)"
                                />
                                <p
                                    class="error-tip"
                                    v-if="['error', ''].includes(scope.row[`volumeNameFlag${scope.row.id}`])"
                                >
                                    {{formRulesOperate.volumeNameTip.value}}
                                </p>
                                </ui-form-item>
                            </ui-form>
                        </div>
                    </template>
                </ui-table-column>
                <ui-table-column
                    align="center"
                    :label="$t('modelForm.label.volumeManage.label.type')"
                    :min-width="setDiffClass('180', '140')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div>
                            <span v-if="scope.row.editFlag">
                                {{$filter.emptyFilter(scope.row.volumeType)}}
                            </span>
                            <ui-form v-else>

                                <ui-form-item
                                    prop=""
                                    class="ipt-height"
                                    :label="''"
                                >
                                    <ui-select
                                        v-model="scope.row.volumeType"
                                        :class="[
                                            ['success', void 0].includes(scope.row[`volumeTypeFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                        ]"
                                        @change="formRulesOperate.volumeTypeChange(scope.row, scope.row.id)"
                                    >
                                        <ui-option
                                            style="fontSize: 12px"
                                            v-for="(item, index) in dataOperate.reactiveData.volumeTypeData"
                                            :key="index"
                                            :label="item"
                                            :value="item"
                                            
                                        />
                                    </ui-select>
                                    <p
                                        class="error-tip"
                                        v-if="['error', ''].includes(scope.row[`volumeTypeFlag${scope.row.id}`])"
                                    >
                                        {{formRulesOperate.volumeTypeTip.value}}
                                    </p>
                                </ui-form-item>
                            </ui-form>
                        </div>
                    </template>
                </ui-table-column>
                <ui-table-column
                    align="center"
                    :label="$t('modelForm.label.volumeManage.label.disk')"
                    :min-width="setDiffClass('180', '140')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div>
                            <span v-if="scope.row.editFlag">
                                {{$filter.emptyFilter(formRulesOperate.setDiskInterfaceType(scope.row.diskType))}}
                            </span>
                            <ui-form v-else>
                                <ui-form-item
                                    prop=""
                                    class="ipt-height"
                                    :label="''"
                                >
                                    <ui-select
                                        v-model="scope.row.diskType"
                                        :class="[
                                            ['success', void 0].includes(scope.row[`diskTypeFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                        ]"
                                        @change="formRulesOperate.diskTypeChange(scope.row, scope.row.id)"
                                    >
                                        <ui-option
                                            style="fontSize: 12px"
                                            v-for="(item, index) in dataOperate.reactiveData.diskTypeData"
                                            :key="index"
                                            :label="item.label"
                                            :value="item.value"
                                        />
                                    </ui-select>
                                    <p
                                        class="error-tip"
                                        v-if="['error', ''].includes(scope.row[`diskTypeFlag${scope.row.id}`])"
                                    >
                                        {{formRulesOperate.diskTypeTip.value}}
                                    </p>
                                </ui-form-item>
                            </ui-form>
                        </div>
                    </template>
                </ui-table-column>
                <ui-table-column
                    align="center"
                    :label="$t('modelForm.label.volumeManage.label.interfaceType')"
                    :min-width="setDiffClass('180', '140')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div>
                            <span v-if="scope.row.editFlag">
                                {{$filter.emptyFilter(formRulesOperate.setDiskInterfaceType(scope.row.interfaceType))}}
                            </span>
                            <ui-form v-else>
                                <ui-form-item
                                    prop=""
                                    class="ipt-height"
                                    :label="''"
                                >
                                    <ui-select
                                        v-model="scope.row.interfaceType"
                                        :class="[
                                            ['success', void 0].includes(scope.row[`interfaceTypeFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                        ]"
                                        @change="formRulesOperate.interfaceTypeChange(scope.row, scope.row.id)"
                                    >
                                        <ui-option
                                            style="fontSize: 12px"
                                            v-for="(item, index) in dataOperate.reactiveData.interfaceData"
                                            :key="index"
                                            :label="item.label"
                                            :value="item.value"
                                        />
                                    </ui-select>
                                    <p
                                        class="error-tip"
                                        v-if="['error', ''].includes(scope.row[`interfaceTypeFlag${scope.row.id}`])"
                                    >
                                        {{formRulesOperate.interfaceTypeTip.value}}
                                    </p>
                                </ui-form-item>
                            </ui-form>
                        </div>
                    </template>
                </ui-table-column>
                <ui-table-column
                    align="center"
                    :label="$t('modelForm.label.volumeManage.label.minNum')"
                    :min-width="setDiffClass('245', '230')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div>
                            <span v-if="scope.row.editFlag">
                                {{$filter.emptyFilter(scope.row.minSizeNum)}}{{ scope.row.minSize }}
                            </span>
                            <ui-form
                                class="min-size-count"
                                v-else
                            >
                                <ui-form-item
                                    prop=""
                                    :label="''"
                                    :class="[
                                        'ipt-height',
                                        setDiffClass('set-data-sys-style', 'sys-single-capacity input-width')
                                    ]"
                                >
                                    <ui-input
                                        :placeholder="$t('modelForm.label.volumeManage.placeholder.tip1')"
                                        :class="[
                                            ['success', void 0].includes(scope.row[`minSizeNumFlag${scope.row.id}`]) ? 'default-ipt' : 'active-ipt'
                                        ]"
                                        v-model="scope.row.minSizeNum"
                                        @blur="formRulesOperate.minSizeNumBlur(scope.row, scope.row.id)"
                                    >
                                    </ui-input>
                                    <ui-select
                                        v-model="scope.row.minSize"
                                        @change="formRulesOperate.minSizeChange(scope.row)"
                                    >
                                        <ui-option
                                            style="fontSize: 12px"
                                            v-for="(item, index) in ['GB', 'TB']"
                                            :key="index"
                                            :label="item"
                                            :value="item"
                                        />
                                    </ui-select>
                                    <p
                                        class="error-tip"
                                        v-if="['error', ''].includes(scope.row[`minSizeNumFlag${scope.row.id}`])"
                                    >
                                        {{formRulesOperate.minSizeNumTip.value}}
                                    </p>
                                </ui-form-item>
                            </ui-form>
                        </div>
                    </template>
                </ui-table-column>
                <ui-table-column
                    align="center"
                    :min-width="setDiffClass('170', '150')"
                    :label="$t('modelForm.label.volumeManage.label.amount')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <div>
                            <span v-if="scope.row.editFlag">
                                {{$filter.emptyFilter(scope.row.amount)}}
                            </span>
                            <ui-form v-else>
                                <ui-form-item
                                    prop=""
                                    :label="''"
                                    :class="[
                                        'ipt-height',
                                    ]"
                                >
                                    <ui-input-number
                                        v-model="scope.row.amount"
                                        :min="1"
                                        :max="1"
                                        :precision="0"
                                    />
                                    <p
                                        class="error-tip"
                                        v-if="['error', ''].includes(scope.row[`amountFlag${scope.row.id}`])"
                                    >
                                        {{formRulesOperate.amountTip.value}}
                                    </p>
                                </ui-form-item>
                            </ui-form>
                        </div>
                    </template>
                </ui-table-column>
                <ui-table-column
                    align="center"
                    min-width="100"
                    fixed="right"
                    :label="$t('modelForm.label.volumeManage.label.operate')"
                    :has-user-template="true"
                >
                    <template #default="{scope}">
                        <p :class="['operate-btn', scope.row.editFlag ? 'edit-btn' : 'save-btn']">
                            <span v-if="!scope.row.editFlag" @click="formRulesOperate.noConfigurationSaveClick(scope.row, scope.row.id)">
                                {{$t('modelForm.label.volumeManage.btn.save')}}
                            </span>
                            <span  v-else @click="formRulesOperate.noConfigurationEditClick(scope.row)">
                                {{$t('modelForm.label.volumeManage.btn.edit')}}
                            </span>
                            <span @click="formRulesOperate.noConfigurationDeleteClick(scope.$index)">
                                {{$t('modelForm.label.volumeManage.btn.delete')}}
                            </span>
                        </p>
                    </template>
                </ui-table-column>
            </ui-table>
        </ui-config-provider>
    </ui-form-item>
</template>

<script setup lang="ts">
// setDiffClass, hasShowTooltip, 
import {tableClass, setDiffClass} from 'utils/index.ts';
// import {RuleFormRefType} from '@utils/publicType';
import {uiLocale} from 'utils/publicClass.ts';

interface PropsType {
    formRulesOperate: any
    dataOperate: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
// import {RuleFormType, RulesType} from './typeManagement'
// import reactiveData from './utils/reactiveData';
// import FormRulesEvent from './utils/formRulesEvent';
// import formRules from './utils/newFormRules';
// import {regExpCheck} from './utils/regExpCheck';
// // import store from 'store/index.ts';
// import DataOperate from './utils/dataOperate';
// import ModelStaticData from 'staticData/model/index.ts';

// const formRulesOperate = formRules();

// const reactiveDataOperate = reactiveData();

// const formRulesEvent = new FormRulesEvent(formRulesOperate as any);
// const dataOperate = new DataOperate(formRulesOperate.ruleForm);


</script>
<style lang="scss">
@import 'assets/css/detail.scss';
@import './index';
</style>
