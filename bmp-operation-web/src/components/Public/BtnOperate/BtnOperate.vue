<template>
    <!-- 左侧操作 -->
    <div
        ref="operateBtn"
        class="operate-management-btn"
    >
        <ui-button
            class="btn-operate"
            type="primary"
            @click="btnClick"
        >
            +  {{$t(type)}}
        </ui-button>
    </div>
</template>

<script setup lang="ts">

interface PropsType {
    type?: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {
    type: 'userList.operate.addUser'
});

const operateBtn = ref();

// defineEmits API 来替代 emits
const emitValue = defineEmits(['btn-operate', 'btn-ref']);

onMounted(() => {
    emitValue('btn-ref', operateBtn)
})

const btnClick = () => {
    emitValue('btn-operate')
};
</script>

<style lang="scss" scoped>
@import '@/assets/css/communal.scss';

 .operate-management-btn {
    margin-top: 22px;

    .btn-operate {
        border-radius: 15px;
        font-size: 12px;
        // color: $color3;
        font-weight: 500;
    }
}
</style>