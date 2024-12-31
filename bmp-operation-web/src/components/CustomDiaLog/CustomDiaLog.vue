<template>
    <!-- 自定义通用dialog操作 -->
    <div :class="['currency custom-dialog', customClass]">
        <!-- 自定义通用dialog弹窗 -->
        <ui-dialog
            v-model="diaLog"
            center
            :show-close="false"
            :close-on-click-modal="false"
            :destroy-on-close="true"
        >
            <template #title>
                <!-- 自定义弹窗标题 -->
                <div class="custom-dialog-header">
                    <!-- 标题icon 名称 -->
                    <p class="title">
                        <img
                            class="common-img"
                            alt=""
                            :src="src"
                        />
                        <span>{{headerTitle}}</span>
                    </p>
                    <!-- 标题关闭icon -->
                    <ui-button
                        class="close-button"
                        circle
                        :icon="Close"
                        @click="operate?.diaLogClose"
                    >
                    </ui-button>
                </div>
            </template>
            <slot></slot>
            <!-- 底部操作 -->
            <template #footer>
               <!-- 取消、确定操作 -->
               <div class="custom-dialog-footer">
                    <!-- 取消 -->
                    <ui-button
                        class="custom-dialog-footer-cancel-btn"
                        v-if="!hasSetUp"
                        @click="operate?.cancelClick"
                    >
                        {{$t('operate.btn.cancel')}}
                    </ui-button>
                    <ui-button
                        class="custom-dialog-footer-cancel-btn"
                        v-if="hasSetUp"
                        @click="operate?.preClick"
                    >
                        {{setTitle}}
                    </ui-button>
                    <!-- 确定 -->
                    <ui-button
                        type="primary"
                        class="custom-dialog-footer-confirm-btn"
                        :loading="nextLoading?.value || isLoading?.value"
                        v-if="hasSetUp"
                        @click="operate?.nextClick"
                    >
                        {{setUpTitle}}
                    </ui-button>
                    <ui-button
                        type="primary"
                        class="custom-dialog-footer-confirm-btn"
                        v-if="!hasSetUp"
                        :loading="isLoading?.value??false"
                        :disabled="disabledStatus"
                        @click="operate?.determineClick"
                    >
                        {{$t('operate.btn.sure')}}
                    </ui-button>
                   
                </div>
            </template>
        </ui-dialog>
    </div>
</template>
  
<script lang="ts" setup>
import {Close} from '@element-plus/icons-vue';
import {language} from 'utils/publicClass.ts';

// props类型
type PropsType = {
    diaLog: boolean;
    src: string;
    isLoading?: {
        value: boolean;
    };
    nextLoading?: {
        value: boolean;
    };
    hasSetUp?: boolean;
    setTitle?: string;
    headerTitle?: string;
    customClass?: any;
    setUpTitle?: string;
    disabledStatus?: boolean;
};

/**
 * props参数
 * @param {boolean} diaLog 弹窗状态
*/
withDefaults(defineProps<PropsType>(), {
    diaLog: false, // 蒙层
    setTitle: '',
    setUpTitle: '',
    hasSetUp: false,
    headerTitle: language.t('loginOut.title')
});

// emit 事件
const emitValue = defineEmits(['dia-log-close', 'determine-click', 'next-click', 'pre-click']);

/**
 * 退出登录
*/
class OperateEvent {

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 关闭蒙层
        this.diaLogClose();
    };

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    diaLogClose = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close');
    };

    /**
     * 确定按钮弹窗
    */
    determineClick = () => {
        emitValue('determine-click');
    };

    nextClick = () => {
        emitValue('next-click');
    }

    preClick = () => {
        emitValue('pre-click')
    }
};
// new 实例化
const operate: OperateEvent = new OperateEvent();

</script>
<style lang="scss">
@import './customDiaLog.scss';
</style>
