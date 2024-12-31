<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'batch-edit-instance-name'"
        :is-loading="editInstanceName.isLoading"
        :header-title="$t('batchOperate.instanceName.title')"
        :src="($defInfo.imagePath('descEdit') as unknown as string)"
        @dia-log-close="editInstanceName.cancelClick"
        @determine-click="editInstanceName.determineClick"
    >
        <div class="batch-edit-instance-name-content">
            <p class="warnning-tip">
                <!-- 提示icon -->
                <img
                    alt=""
                    :src="($defInfo.imagePath('warningTip') as unknown as string)"
                />
                <!-- 提示 -->
                <span>{{$t('batchOperate.instanceName.tip')}}</span>
            </p>
            <p class="tip-title">
                {{ $t('batchOperate.instanceName.tip1') }}
                <span>{{selectArr.length}}</span>
                {{ $t('batchOperate.instanceName.tip2') }}
            </p>
            <!-- 表格数据 -->
            <edit-instance-name-table :select-arr="selectArr"/>
            <edit-instance-name-form :form-opt="formOpt"/>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {language} from 'utils/publicClass.ts';
import {CurrencyType, SuccessType} from '@utils/publicType';
import FormRulesOpt from './formTable/ruleForm';
import {msgTip, customTip} from 'utils/index.ts';
import {BatchEditInstanceNameAPI} from 'api/device/request.ts';

/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    selectArr: CurrencyType[];
};


const formOpt = new FormRulesOpt();

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

/**
 * defineEmits API 用来代替emit
*/
const emitValue = defineEmits(['dia-log-close', 'determine-click']);


class EditInstanceName {
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    constructor() {
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.determineClick();
        };

        onUnmounted(() => document.onkeyup = (() => {return}));
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        // 判断输入项是否符合条件
        await nextTick(() => {
            formOpt.ruleFormRef.value!.validate((valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid) {
                    this.isLoading.value = true;
                    this.requestInstanceName();
                }
            });
        });
    };

    /**
     * 请求接口，成功后把事件回传，关闭弹窗
    */
    requestInstanceName = (): string | void => {
        BatchEditInstanceNameAPI(
            {
                instanceIds: props.selectArr.map((item: CurrencyType) => item.instanceId),
                instanceName: formOpt.ruleForm.instanceName
            }
        )
        .then(({data}: {data: SuccessType}) => {
            if (data?.result?.success) {
               this.#dealWithResponse();
            }
        })
        .catch(({message}: {message: string}) => {
            if (['实例名称', 'instanceName'].indexOf(message) > -1) {
                this.#setMsg(message);
                return;
            }
            this.#setMsg(message);
            emitValue('determine-click');
            this.cancelClick();
        });
    };

    #dealWithResponse = () => {
        msgTip.success(language.t('operate.success'));
        emitValue('determine-click');
        this.cancelClick();
        this.isLoading.value = false;
    };

    #setMsg = (message: string) => {
        customTip('error', message, 1000, () => this.isLoading.value = false);
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};

const editInstanceName = new EditInstanceName();
</script>
<style lang="scss">
@import './batchEditInstanceName.scss';
</style>
