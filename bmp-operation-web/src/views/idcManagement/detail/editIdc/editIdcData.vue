<template>
    <div v-if="editIdcOperate.editIdcDiaLog.value">
        <edit-idc
            :dia-log="editIdcOperate.editIdcDiaLog.value"
            :item-data="idcDetail.reactiveArr.detailInfo"
            @dia-log-close="editIdcOperate.editIdcCancel"
            @determine-click="editIdcOperate.editIdcSure"
        />
    </div>
</template>
<script lang="ts" setup>
import {CurrencyType} from '@utils/publicType';

interface PropsType {
    editIdcOperate: any;
    idcDetail: {
        reactiveArr: {
            detailInfo: CurrencyType;
        }
    };
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>