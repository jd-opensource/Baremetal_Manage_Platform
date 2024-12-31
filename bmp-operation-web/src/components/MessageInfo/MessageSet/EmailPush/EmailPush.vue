<template>
    <!-- 关联、取消关联邮箱推送 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'currency email-push-count'"
        :is-loading="emailPushOpt.isLoading"
        :header-title="emailPushOpt.tipTitle()"
        :src="($defInfo.imagePath('messageLight') as unknown as string)"
        @dia-log-close="emailPushOpt.cancelClick"
        @determine-click="emailPushOpt.determineClick"
    >
        <!-- 主体内容 -->
        <div class="currency-count">
            <!-- 关联、取消关联提示 -->
            <p class="currency-count-title">
                {{emailPushOpt.tipContent()}}
            </p>
        </div>
    </custom-dia-log>
</template>
  
<script setup lang="ts">
import {msgTip, methodsTotal} from 'utils/index.ts';
// import { SuccessType } from '@utils/publicType';
import { CurrentInstance } from 'utils/publicClass.ts';
import {language} from 'utils/publicClass.ts'; // 国际化
// import {SuccessType} from '@utils/publicType';


/**
 * props 类
*/
interface PropsType {
    diaLog: boolean;
    operate: string;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 来替代 emits
const emitValue = defineEmits(['dia-log-close', 'determine-click', 'error-click']);

class EmailPushOpt {
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(false);
    proxy = new CurrentInstance().proxy;

    /**
     * 确定按钮弹窗
    */
    determineClick = (): void => {
        this.isLoading.value = true;
        this.requestHasPush();
    };

    tipTitle = () => {
        const info = new Map([
            ['push', language.t('emailPush.start.title')]
        ]);
        return info.get(props.operate)?? language.t('emailPush.close.title')
    };

    tipContent = () => {
        const info = new Map([
            ['push', language.t('emailPush.start.content')]
        ]);
        return info.get(props.operate)?? language.t('emailPush.close.content')
    };

    /**
     * 请求是否关联邮箱推送，成功后把事件回传，关闭弹窗
    */
    requestHasPush = async () => {
        const params = methodsTotal.toLine({
            isPush: props.operate === 'push' ? '1' : '0'
        })
        try {
            const res = await this.proxy.$messageApi.savelsPushMailAPI({...params});
            if (res?.data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                emitValue('determine-click');
                return;
            }
            throw new Error('');
        }
        catch {
            emitValue('error-click');
        }
        finally {
            this.isLoading.value = false;
            this.cancelClick();
            
        }
    };

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};
const emailPushOpt = new EmailPushOpt();

</script>
<style lang="scss">
@import './index.scss';
</style>
