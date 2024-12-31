<template>
    <div class="operate-management-detail message-detail">
        <!-- 头部标题内容 -->
        <detail-header
            :name="(messageDetailOpt.setName() as unknown as string)"
            :title="'messageDetail'"
            :opt-flag="false"
            @header="messageDetailOpt.jumpRouter($defInfo.routerPath('myMessage') as unknown as string)"
        />
        <main class="operate-management-detail-info operate-management-count">
            <!-- 失败提示 -->
            <div class="delete-error">
                <p class="title">
                    {{messageDetailOpt.setTitle()}}
                </p>
                <div class="content">
                    <p class="count" v-html="messageDetailOpt.setCount()"></p>
                    <span
                        class="renew-immediately"
                        v-if="Object.is(messageDetailOpt.activeType.value, 'unexpiredExpired')"
                        @click="messageDetailOpt.jumpDeviceDetail($defInfo.routerPath('license') as unknown as string)"
                    >
                        {{$t('messageDetail.header.renewImmediately')}}
                    </span>
                </div>
                <p
                    class="desc"
                    v-if="Object.is(messageDetailOpt.activeType.value, 'inbondAlert')"
                >
                    {{$filter.withClonZh(messageDetailOpt.setDesc())}}
                    <span>{{messageDetailOpt.setIdcInfo()}}
                        <span
                            :class="setTextClass(true)"
                            @click="messageDetailOpt.jumpDeviceDetail($defInfo.routerPath('deviceDetail') as unknown as string)"
                        >
                            {{messageDetailOpt.inMonitor()}}
                        </span>
                        {{messageDetailOpt.monitorCount()}}
                    </span>
                </p>
                <p class="desc" v-else>
                    {{messageDetailOpt.setDesc()}}
                </p>
                <ui-config-provider
                    v-if="!Object.is(messageDetailOpt.activeType.value, 'inbondAlert')"
                    :locale="uiLocale.locale"
                >
                    <ui-table
                        style="width: 100%;"
                        border
                        height="100px"
                        v-loading="messageDetailOpt.isLoading.value"
                        :class="tableClass(messageDetailOpt.reactiveArr.tableData, messageDetailOpt.isLoading.value)"
                        :data="messageDetailOpt.reactiveArr.tableData"
                    >
                        <template v-for="(item, ind) in messageDetailOpt.templateData">
                            <ui-table-column
                                align="center"
                                v-if="Object.is(messageDetailOpt.activeType.value, item.hasShow)"
                                :key="ind"
                                :label="item.label"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <div>
                                        <span v-if="item.hasOther">
                                            {{$filter.emptyFilter(messageDetailOpt.setStatus(scope.row[item.prop]))}}
                                        </span>
                                        <span v-else-if="!item.hasTime && !item.hasOther">
                                            {{$filter.emptyFilter(scope.row[item.prop])}}
                                        </span>
                                        <span v-else>
                                            {{$filter.emptyFilter(getDateMinutes(scope.row[item.prop]))}}
                                        </span>
                                    </div>
                                </template>
                            </ui-table-column>
                        </template>
                    </ui-table>
                </ui-config-provider>
                <div v-if="Object.is(messageDetailOpt.activeType.value, 'optMessage')">
                    <p class="opt-count-detail">{{$t('messageDetail.optTips.title')}}</p>
                    <ui-config-provider :locale="uiLocale.locale">
                        <ui-table
                            style="width: 100%;"
                            border
                            height="100px"
                            :class="tableClass(messageDetailOpt.reactiveArr.tableData, messageDetailOpt.isLoading.value)"
                            :data="messageDetailOpt.reactiveArr.tableData"
                        >
                            <ui-table-column
                                align="center"
                                v-for="(t, i) in messageDetailOpt.instanceTemplateData"
                                :key="i"
                                :label="t.label"
                                :has-user-template="true"
                            >
                                <template #default="{scope}">
                                    <span>
                                        {{$filter.emptyFilter(scope.row[t.prop])}}
                                    </span>
                                </template>
                            </ui-table-column>
                        </ui-table>
                    </ui-config-provider>
                </div>
            </div>
        </main>
        <div class="footer-opt">
            <ui-button
                type="primary"
                class="btn"
                v-for="(item, index) in messageDetailOpt.btnData"
                :key="index"
                :disabled="messageDetailOpt.disabledSet(item.id)"
                @click="messageDetailOpt.preNextClick(item.id)"
            >
                {{item.text}}
            </ui-button>
        </div>
    </div>
</template>
<script lang="ts" setup>
import {tableClass, getDateMinutes, setTextClass} from 'utils/index.ts';
import {uiLocale} from 'utils/publicClass.ts';
import MessageDetailOpt from './msgDetail';

const messageDetailOpt = new MessageDetailOpt();

</script>
<style lang="scss">
@import 'assets/css/detail.scss';
@import './index.scss';

</style>