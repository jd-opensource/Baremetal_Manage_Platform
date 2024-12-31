<template>
    <!-- 编辑备注 -->
    <div v-if="operate.editNotesDiaLog.value">
        <edit-notes
            :dia-log="operate.editNotesDiaLog.value"
            :device-Id="routerOperate.router?.currentRoute?.value?.query?.deviceId"
            :description="description"
            @dia-log-close="operate.editNotesCacnel"
            @determine-click="operate.editNotesSure"
        />
    </div>
</template>
<script lang="ts" setup>
import {RouterOperate} from 'utils/publicClass.ts';

// defineEmits API 来替代 emits
interface PropsType {
    operate: any;
    description: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    description: ''
});

const routerOperate = new RouterOperate();
</script>