<template>
    <!-- 删除apikey操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'delete-api-key'"
        :is-loading="deleteApiKey.isLoading"
        :header-title="$t('deleteKey.header.title')"
        :src="($defInfo.imagePath('apiKey') as unknown as string)"
        @dia-log-close="deleteApiKey.cancelClick"
        @determine-click="deleteApiKey.determineClick"
    >
        <!-- 主体内容 -->
        <div class="delete-keycontent">
            <span class="delete-keyscope">
                {{$t('deleteKey.content.title', [name])}}
            </span>
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {deleteApiKeyAPI} from 'api/userCenter/request.ts';

/**
 * props类
*/
interface PropsType {
    diaLog: boolean;
    name: string;
    apikeyId: string;
};

/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
 * @param {boolean} diaLog 弹窗状态
 * @param {boolean} hasEditOperate 是否进行编辑操作 默认true
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    diaLog: false,
});

// defineEmits API 来替代emit，进行事件回传
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// 删除apikey
class DeleteApiKey {
    isLoading: Ref<boolean> = ref<boolean>(false);

    determineClick = () => {
        this.isLoading.value = true;
        this.deleteApiKey();
    };

    deleteApiKey = () => {
        deleteApiKeyAPI(
            {
                apiKey: props.apikeyId
            }
        )
        .then(({data} : {data: {result: {success: boolean}}}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
                this.cancelClick();
                emitValue('determine-click');
            }
        })
        .finally(() => {
            this.isLoading.value = false;
        })
        .catch(() => {
            this.cancelClick();
            emitValue('determine-click');
        })
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };
};
// 实例化
const deleteApiKey: DeleteApiKey = new DeleteApiKey();

</script>
<style lang="scss">
@import './deleteApiKey.scss';
</style>
