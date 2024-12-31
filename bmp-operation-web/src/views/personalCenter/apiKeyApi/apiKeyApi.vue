<template>
    <div class="user-center-content">
        <header-info
            :type="'apiKey'"
            @header-ref="tableS.getHeaderRef"
        />
        <main class="operate-management-content">
            <div class="api-key">
                <ui-button
                    class="create-key"
                    type="primary"
                    :disabled="apiKey.isBtnDisabled.value"
                    @click="createKey1.createKeyClick"
                >
                    <ui-tooltip
                        placement="bottom"
                        v-if="apiKey.isBtnDisabled.value"
                        :content="$t('userCenter.tooltip.tip')"
                    >
                        +  {{$t('userCenter.btn.createKey')}}
                    </ui-tooltip>
                    <span v-else> +  {{$t('userCenter.btn.createKey')}}</span>
                </ui-button>
                <!-- 表格数据-数据列表 -->
                <article class="operate-management-count user-centet-table">
                    <ui-table
                        ref="tableRef"
                        border
                        style="width: 100%"
                        v-loading="apiKey.isLoading.value"
                        :max-height="tableS.tableMaxHeight.value || 'auto'"
                        :class="tableClass(apiKey.reactiveArr.tableData, apiKey.isLoading.value)"
                        :data="apiKey.reactiveArr.tableData"
                    >
                        <!-- 密钥名称 -->
                        <ui-table-column
                            prop="name"
                            min-width="130"
                            align="center"
                            :border="false"
                            :label="$t('userCenter.table.label.keyName')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <ui-tooltip
                                    placement="bottom"
                                    v-model="scope.row[`nameTip${scope.$index}`].showTooltip"
                                    :disabled="!scope.row[`nameTip${scope.$index}`].showTooltip"
                                >
                                    <template #content>
                                        <div class="set-tooltip-width">
                                            <span>
                                                {{scope.row.name}}
                                            </span>
                                        </div>
                                    </template>
                                    <div
                                        class="more-text-ellipsis"
                                        @mouseenter="hasShowTooltip($event, scope.row[`nameTip${scope.$index}`], scope.row.name)"
                                    >
                                        <span :class="setTextClass(false)">
                                            {{$filter.emptyFilter(scope.row.name)}}
                                        </span>
                                    </div>
                                </ui-tooltip>
                            </template>
                        </ui-table-column>
                        <!-- token -->
                        <ui-table-column
                            prop="token"
                            min-width="130"
                            align="center"
                            :label="$t('userCenter.table.label.token')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span :class="setTextClass(false)">
                                    {{$filter.emptyFilter(scope.row.token)}}
                                </span>
                            </template>
                        </ui-table-column>
                        <!-- 权限 -->
                        <ui-table-column
                            prop="readOnly"
                            align="center"
                            :label="$t('userCenter.table.label.permissions')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span :class="setTextClass(false)">
                                    {{scope.row.readOnly ? $t('userCenter.table.select.read') : $t('userCenter.table.select.title')}}
                                </span>
                            </template>
                        </ui-table-column>
                        <!-- 创建时间 -->
                        <ui-table-column
                            prop="createdTime"
                            align="center"
                            :label="$t('userCenter.table.label.createTime')"
                            :has-user-template="true"
                        >
                            <template #default="{scope}">
                                <span :class="setTextClass(false)">
                                    {{$filter.emptyFilter(scope.row.createdTime)}}
                                </span>
                            </template>
                        </ui-table-column>
                        <!-- 操作 -->
                        <ui-table-column
                            prop="operate"
                            align="center"
                            :has-user-template="true"
                            :label="$t('userCenter.table.label.operate.name')"
                        >
                            <template #default="{scope}">
                                <!-- 复制 -->
                                <span
                                    :class="setTextClass(true, 'currency-style-right')"
                                    @click="methodsTotal.copyInfo(scope.row.token)"
                                >
                                    {{$t('userCenter.table.label.operate.copy')}}
                                </span>
                                <!-- 删除 -->
                                <span
                                    :class="setTextClass(true)"
                                    @click="deleteKey1.deleteKeyClick(scope.row)"
                                >
                                    {{$t('userCenter.table.label.operate.delete')}}
                                </span>
                            </template>
                        </ui-table-column>
                    </ui-table>
                </article>
                <!-- 分页组件 -->
                <api-key-list-pagination :pagination-opt="apiKey"/>
            </div>
        </main>
        <api-key-create
            :api-key="apiKey"
            :create-key="createKey1"
        >
        </api-key-create>
        <api-key-delete
            :api-key="apiKey"
            :delete-key="deleteKey1"
        >
        </api-key-delete>
    </div>
</template>
<script lang="ts" setup>
// import {paginationOperate} from 'utils/publicClass.ts';
import {tableClass, hasShowTooltip, setTextClass, methodsTotal} from 'utils/index.ts';
import ApiKey from './apikey';
import SetEmpty from './setEmpty';
// import PaginationWatch from './pagination';
import createKey from './apiKeyCreate/operate'
import deleteKey from './apiKeyDelete/operate'
import tableScrollOperate from './tableScroll';
import ApiKeyCreate from './apiKeyCreate/apiKeyCreate.vue';
import ApiKeyDelete from './apiKeyDelete/apiKeyDelete.vue';
import usePagination from 'hooks/pagination/usePagination.ts';
import ApiKeyListPagination from 'hooks/pagination/paginationSelect.vue';

const apiKey = new ApiKey();

usePagination(apiKey.getApiKey);

const createKey1 = createKey(apiKey);

const deleteKey1 = deleteKey(apiKey);

// const emitValue = defineEmits(['empty-click']);

new SetEmpty(createKey1)


const tableS = tableScrollOperate();
</script>
<style scoped lang="scss">
@import './apiKey';
</style>