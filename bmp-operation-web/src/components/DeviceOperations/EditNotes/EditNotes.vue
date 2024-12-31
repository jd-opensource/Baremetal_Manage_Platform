<template>
    <!-- 编辑备注操作 -->
    <custom-dia-log
        :dia-log="diaLog"
        :custom-class="'edit-notes'"
        :is-loading="editNotes.isLoading"
        :header-title="$t('deviceDetail.editDesc.header.title2')"
        :src="($defInfo.imagePath('device') as unknown as string)"
        @dia-log-close="editNotes.cancelClick"
        @determine-click="editNotes.determineClick"
    >
        <div class="edit-notes-count">
            <p>{{$t('deviceDetail.editDesc.count.title2')}}</p>
            <ui-input
                class="edit-notes-textarea"
                type="textarea"
                maxlength="256"
                show-word-limit
                v-model.trim="editNotes.description.value"
                :placeholder="$t('deviceDetail.editDesc.placeholder2')"
            />
        </div>
    </custom-dia-log>
</template>
  
<script lang="ts" setup>
import {msgTip} from 'utils/index.ts';
import {language} from 'utils/publicClass.ts';
import {devicesEditAPI} from 'api/device/request.ts';

interface PropsType {
    diaLog: boolean;
    deviceId?: string;
    description: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// defineEmits API 用来代替emit
const emitValue = defineEmits(['dia-log-close', 'determine-click']);

/**
 * 编辑备注
*/
class EditNotes {
    // 备注内容
    description: Ref<string> = ref<string>('');
    // loading加载态
    isLoading: Ref<boolean> = ref<boolean>(false);

    // 构造器
    constructor () {
        this.description.value = props.description;
        document.onkeyup = (event: {keyCode: number;}) => {
            event.keyCode === 13 && this.determineClick();
        };
    };

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    cancelClick = (): void => {
        // 回传父组件当前页码，可以处理相关操作
        emitValue('dia-log-close', false);
    };

    /**
     * 确定按钮弹窗
    */
    determineClick: () => Promise<void> = async () => {
        this.isLoading.value = true;
        await this.requestEditNotes();
    };

    /**
     * 请求编辑备注接口
    */
    requestEditNotes = () => {
        devicesEditAPI(
            {
                id: props.deviceId,
                description: this.description.value
            }
        )
        .then(({data} : {data: {result: {success: boolean;}}}) => {
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
// 实例化-编辑备注
const editNotes: EditNotes = new EditNotes();

</script>
<style lang="scss">
@import './editNotes.scss';
</style>
