<template>
    <!-- 镜像描述操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'image-desc'"
        :is-loading="editDesc.isLoading"
        :header-title="$t('editDesc.header.title')"
        :src="($defInfo.imagePath('imageManage') as unknown as string)"
        @dia-log-close="editDesc.cancelClick"
        @determine-click="editDesc.determineClick"
    >
        <div class="image-desc-count">
            <p>{{$t('deviceDetail.editDesc.count.title')}}</p>
            <ui-input
                class="image-desc-textarea"
                type="textarea"
                maxlength="256"
                show-word-limit
                v-model.trim="editDesc.description.value"
                :placeholder="$t('deviceDetail.editDesc.placeholder')"
            />
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip} from 'utils/index.ts'
import {language} from 'utils/publicClass.ts';
import {SuccessType} from '@utils/publicType';
import {imageEditAPI} from 'api/image/request.ts';

interface PropsType {
    diaLog: boolean;
    imageId: string;
    description: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

onUnmounted (() => document.onkeyup = () => {return;});

/**
 * 编辑描述
*/
class EditDesc {
    // 描述内容
    description: Ref<string> = ref<string>('');
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    // 构造器
    constructor () {
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.determineClick();
        };
        this.description.value = props.description;
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = () => {
        // 关闭蒙层
        this.diaLogClose();
    };

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    diaLogClose = () => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestEditDesc();
    };

    /**
     * 请求编辑描述接口
    */
    requestEditDesc = () => {
        imageEditAPI(
            {
                imageId: props.imageId,
                description: this.description.value
            }
        )
        .then(({data} : {data: SuccessType}) => {
            if (data?.result?.success) {
                msgTip.success(language.t('operate.success'));
            }
        })
        .finally(() => {
            this.isLoading.value = false;
            emitValue('determine-click');
            this.cancelClick();
        })
    };
};
// 实例化-编辑描述
const editDesc: EditDesc = new EditDesc();

</script>
<style lang="scss">
@import './imageDesc.scss';
</style>
