<template>
    <!-- 操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'secondary-confirmation'"
        :header-title="$t('resetSystemConfirm.header.title')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="emitValue('dia-log-close', false)"
        @determine-click="sureClick"
    >

        <div class="secondary-confirmation-content">
            <div class="warnning-tip">
                <!-- 提示icon -->
                <div style="display: flex; align-items: center; margin-top: 5px;">
                    <img
                        alt=""
                        :src="($defInfo.imagePath('uiTooltip') as unknown as string)"
                    />
                    <p style="width: 40px;font-size:12px;color: #ff4D4f; text-align: left;">{{$t('resetSystemConfirm.header.tip0')}}</p>
                </div>
                <!-- 提示 -->
                <div class="tip-count">
                    <p class="reset-system-tip" v-html="$t('resetSystemConfirm.header.tip')"></p>
                </div>
            </div>
            <p class="tip-title set-word">
                {{$t('resetSystemConfirm.header.tip2', [name])}}
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
// props 类
interface PropsType {
    name: string;
    diaLog: boolean;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    name: '', // 标题
    diaLog: false, // 蒙层
});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// 确定按钮弹窗
const sureClick = () => {
    // 回传父组件当前页码，可以处理相关操作
    emitValue('dia-log-close', false);
    emitValue('determine-click');
};

</script>
<style lang="scss">
@import './secondaryConfirmation.scss';
</style>
