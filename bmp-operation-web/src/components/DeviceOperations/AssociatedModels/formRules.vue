<template>
      <ui-form
            label-position="left"
            :label-width="setDiffClass('100px', '80px')"
            :model="upFormOpt.ruleForm"
            :rules="upFormOpt.rules"
            @submit.native.prevent
            @rule-form-ref="upFormOpt.getFormRef"
        >
            <ui-form-item
                prop="model"
                :class="[
                    upFormOpt.ruleForm.model ? '' : associatedModelsStatus ? 'set-empty' : ''
                ]"
                :label="$t('associatedModel.content.select')"
            >
                <ui-select
                    v-model="upFormOpt.ruleForm.model"
                    style="width: 180px"
                    :placeholder="$t('associatedModel.content.placeholder')"
                    @change="upFormOpt.modelChange"
                >
                    <ui-option
                        style="fontSize: 12px"
                        v-for="item in upFormOpt.reactiveArr.modelData"
                        :key="item.value"
                        :label="item.text"
                        :value="item.filterParams"
                    />
                </ui-select>
                <p class="add-new-model">
                    <span @click="upFormOpt.jumpPage">
                        {{$t('associatedModel.content.addNewModel')}}
                    </span>
                </p>
            </ui-form-item>
        </ui-form>
</template>
  
<script lang="ts" setup>
import {methodsTotal, tableClass, setDiffClass} from 'utils/index.ts';
import {uiLocale} from 'utils/publicClass.ts';
import store from 'store/index.ts';
// import UpFormOpt from './utils/upForm';
// import AssociatedModelsOpt from './utils/associatedModelsOpt';

// props类
interface PropsType {
    upFormOpt: any;
    associatedModelsStatus: boolean;
    // associatedModelsOpt: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props: PropsType = withDefaults(defineProps<PropsType>(), {});

// const upFormOpt = new UpFormOpt(props);

// // defineEmits API 用来代替emit
// const emitValue = defineEmits(['dia-log-close', 'determine-click']);

// const associatedModelsOpt = new AssociatedModelsOpt(upFormOpt, emitValue, props.deviceId);
const useListenMsg = methodsTotal.listenMsg((info: {type: string, content: string}) => {
    const {type, content} = info;
    if (['model-add-success', ' model-edit-success', 'model-delete-success', 'model-add-template-success'].indexOf(type) > -1) {
        if (content === 'success') {
            store.ossDataInfo().getModelListRes().then((data: {text: string; filterParams: string; value: number}[]) => {
                props.upFormOpt.reactiveArr.modelData = data;
            })
        }
    }
})

onUnmounted(useListenMsg);

</script>
<style lang="scss">
@import './index.scss';
</style>
