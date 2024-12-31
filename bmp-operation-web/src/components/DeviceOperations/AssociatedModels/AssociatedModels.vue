<template>
    <!-- 关联机型操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'associated-models'"
        :is-loading="associatedModelsOpt.isLoading"
        :header-title="$t('associatedModel.header.title')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="associatedModelsOpt.cancelClick"
        @determine-click="associatedModelsOpt.determineClick"
    >
        <div class="currency-count">
            <div class="associated-models-count  currency-count-table">
                <form-rules
                    :upFormOpt="upFormOpt"
                    :associatedModelsStatus="associatedModelsOpt.errorFlag.value"
                />
                <ui-config-provider :locale="uiLocale.locale">
                    <ui-table
                        border
                        v-loading="upFormOpt.isTableLoading.value"
                        max-height="300"
                        :data="upFormOpt.reactiveArr.tableData"
                        :class="tableClass(upFormOpt.reactiveArr.tableData, upFormOpt.isTableLoading.value)"
                    >
                        <ui-table-column
                            align="center"
                            prop="volumeName"
                            min-width="130"
                            fixed
                            :label="$t('associatedModel.content.table.label.volumeName')"
                        >
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="volumeType"
                            min-width="100"
                            :label="$t('associatedModel.content.table.label.volumeType')"
                        >
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            v-if="upFormOpt.arrayCard.value !== 'no_need_raid'"
                            :min-width="setDiffClass('160', '120')"
                            :label="$t('associatedModel.content.table.label.raidCan')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span>
                                    {{$filter.emptyFilter(upFormOpt.setRaidCan(scope.row.raidCan))}}
                                </span>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            v-if="upFormOpt.arrayCard.value !== 'no_need_raid'"
                            :min-width="upFormOpt.reactiveArr.tableData.length ? '180' : '110'"
                            :label="$t('associatedModel.content.table.label.raid')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <p>
                                    {{$filter.emptyFilter(upFormOpt.setRaid(scope.row.raids))}}
                                </p>
                                <!-- raids -->
                                <!-- <span>{{$filter.emptyFilter(scope.row.raidCan)}}</span> -->
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="diskType"
                            min-width="100"
                            :label="$t('associatedModel.content.table.label.diskType')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <p>
                                    {{$filter.emptyFilter(upFormOpt.setDiskInterfaceType(scope.row.diskType))}}
                                </p>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="interfaceType"
                            min-width="100"
                            :label="$t('associatedModel.content.table.label.interfaceType')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <p>
                                    {{$filter.emptyFilter(upFormOpt.setDiskInterfaceType(scope.row.interfaceType))}}
                                </p>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            :min-width="setDiffClass('210', '120')"
                            :label="$t('associatedModel.content.table.label.minNum')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span>{{$filter.emptyFilter(scope.row.volumeSize)}}{{scope.row.volumeUnit}}</span>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            prop="amount"
                            :min-width="setDiffClass('180', '120')"
                            :label="$t('associatedModel.content.table.label.minimumQuantity')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span>{{$filter.emptyFilter(scope.row.volumeAmount)}}</span>
                            </template>
                        </ui-table-column>
                        <ui-table-column
                            align="center"
                            fixed="right"
                            :min-width="upFormOpt.reactiveArr.tableData.length ? '380' : '140'"
                            :label="$t('associatedModel.content.table.label.associatedDisk')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-form v-if="['RAID0-stripping', 'NO RAID'].indexOf(scope.row.raidCan) == -1">
                                    <ui-form-item :class="['multiple-choice', scope.row.disksDataFlag ? 'error-tip' : 'def-tip']">
                                        <ui-select
                                            v-model="upFormOpt.reactiveArr.selectObj[scope.row.volumeId]"
                                            collapse-tags
                                            multiple
                                            popper-class="select-seleted-icon"
                                            :placeholder="scope.row.disksDataFlag ? $t('upDownFrame.placeholder.tip2') : $t('associatedModel.content.table.diskPlaceholder')"
                                            @change="upFormOpt.disksDataChange(scope.row)"
                                        >
                                            <ui-option
                                                style="fontSize: 12px"
                                                v-for="(item, index) in upFormOpt.reactiveArr?.disksData[scope.$index]"
                                                :key="index"
                                                :label="upFormOpt.setDiskLabel(item)"
                                                :value="item.diskId"
                                                :disabled="upFormOpt?.getAllSelectDisks(scope.row.volumeId).includes(item.diskId)"
                                            />
                                        </ui-select>
                                    </ui-form-item>
                                </ui-form>
                                <ui-form v-else>
                                    <ui-form-item :class="['def-select', scope.row.disksDataFlag ? 'error-tip' : 'def-tip']">
                                        <ui-select
                                            v-model="upFormOpt.reactiveArr.selectObj[scope.row.volumeId]"
                                            clearable
                                            :placeholder="scope.row.disksDataFlag ? $t('upDownFrame.placeholder.tip2') : $t('associatedModel.content.table.diskPlaceholder')"
                                            @change="upFormOpt.disksDataChange(scope.row)"
                                        >
                                            <ui-option
                                                style="fontSize: 12px"
                                                v-for="(item, index) in upFormOpt.reactiveArr.disksData[scope.$index]"
                                                :key="index"
                                                :label="upFormOpt.setDiskLabel(item)"
                                                :value="item.diskId"
                                                :disabled="upFormOpt.getAllSelectDisks(scope.row.volumeId).includes(item.diskId)"
                                            />
                                        </ui-select>
                                    </ui-form-item>
                                </ui-form>
                            </template>
                        </ui-table-column>
                    </ui-table>
                </ui-config-provider>
            </div>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {methodsTotal, tableClass, setDiffClass} from 'utils/index.ts';
import {uiLocale} from 'utils/publicClass.ts';
import store from 'store/index.ts';
import FormRules from './formRules.vue';
import UpFormOpt from './utils/upForm';
import AssociatedModelsOpt from './utils/associatedModelsOpt';

// props类
interface PropsType {
    diaLog: boolean;
    deviceId: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

const upFormOpt = new UpFormOpt(props);

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

const associatedModelsOpt = new AssociatedModelsOpt(upFormOpt, emitValue, props.deviceId);


</script>
<style lang="scss">
@import './index.scss';
</style>
