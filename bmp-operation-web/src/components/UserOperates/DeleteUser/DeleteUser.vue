<template>
    <!-- 删除用户操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'currency delete-user'"
        :is-loading="deleteUserOperate.isLoading"
        :header-title="$t('deleteUser.header.deleteUser')"
        :src="($defInfo.imagePath('user') as unknown as string)"
        @dia-log-close="deleteUserOperate.cancelClick"
        @determine-click="deleteUserOperate.determineClick"
    >
        <!-- 主体内容 -->
        <div class="currency-count">
            <!-- 删除提示 -->
            <p class="currency-count-title">
                {{$t('deleteUser.deleteTip')}}
            </p>
            <!-- 表格数据 -->
            <div class="currency-count-table">
                <ui-table
                    style="width: 100%"
                    :class="tableClass(deleteUserOperate.tableData, false)"
                    :data="deleteUserOperate.tableData"
                >
                    <!-- 用户名 -->
                    <ui-table-column
                        prop="userName"
                        width="170"
                        align="center"
                        :label="$t('deleteUser.label.userName')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <ui-tooltip
                                placement="bottom"
                                v-model="scope.row.showTooltip"
                                :disabled="!scope.row.showTooltip"
                            >
                                <template #content>
                                    <div class="set-tooltip-width">
                                        <span>
                                            {{scope.row.userName}}
                                        </span>
                                    </div>
                                </template>
                                <div
                                    class="more-text-ellipsis"
                                    @mouseenter="hasShowTooltip($event, scope.row, scope.row.userName)"
                                >
                                    <span>
                                        {{$filter.emptyFilter(scope.row.userName)}}
                                    </span>    
                                </div>
                            </ui-tooltip>
                        </template>
                    </ui-table-column>
                    <!-- 角色 -->
                    <ui-table-column
                        prop="roleName"
                        min-width="100"
                        align="center"
                        :label="$t('deleteUser.label.role')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <span>
                                {{$filter.emptyFilter(scope.row.roleName)}}
                            </span>
                        </template>
                    </ui-table-column>
                    <!-- 手机 -->
                    <ui-table-column
                        prop="phoneNumber"
                        width="120"
                        align="center"
                        :label="$t('deleteUser.label.cellPhone')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <span>
                                {{$filter.emptyFilter(scope.row.phoneNumber)}}
                            </span>
                        </template>
                    </ui-table-column>
                    <!-- 邮箱 -->
                    <ui-table-column
                        prop="email"
                        min-width="130"
                        align="center"
                        :label="$t('deleteUser.label.email')"
                        :has-user-template="true"
                    >
                        <template #default="{scope}">
                            <span>
                                {{$filter.emptyFilter(scope.row.email)}}
                            </span>
                        </template>
                    </ui-table-column>
                </ui-table>
            </div>
        </div>
    </custom-dia-log>
</template>
  
<script setup lang="ts">
import {ListDataType} from './type';
import {language, CurrentInstance} from 'utils/publicClass.ts'; // 国际化
import {msgTip, hasShowTooltip, tableClass, methodsTotal} from 'utils/index.ts'; // message提示

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    listData: ListDataType;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false
});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

class DeleteUserOperate {
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(false);
    tableData: ListDataType[] = [props.listData];
    proxy = new CurrentInstance().proxy;

    /**
     * 确定按钮弹窗
    */
    determineClick = async () => {
        this.isLoading.value = true;
        try {
            const deleteApi = await this.proxy.$userApi.deleteUserAPI({userId: props.listData.userId});
            const status: boolean = (deleteApi?.data?.result?.success)?? false;
            if (status) {
                this.#dealWithResponse();
                methodsTotal.sendMsg('delete-user-info', props.listData.userId);
            }
        }
        finally {
            this.#dealWithFinally();
        }
    };

    #dealWithResponse = () => {
        msgTip.success(language.t('operate.success'));
    };

    #dealWithFinally = () => {
        this.isLoading.value = false;
        this.cancelClick();
        emitValue('determine-click');
    };

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};
const deleteUserOperate: DeleteUserOperate = new DeleteUserOperate();

</script>
<style lang="scss">
@import './deleteUser.scss';
</style>
