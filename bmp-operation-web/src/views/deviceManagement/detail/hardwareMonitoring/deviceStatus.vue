<template>
    <div ref="searchRef" class="info-content" v-loading="surveillanceOpt.isLoading.value">
        <div
            class="hardware-device-status"
            v-for="(item, index) in surveillanceOpt.reactiveArr?.tableData"
            :key="index"
        >
            <span>{{item.label}}</span>
            <span class="line"></span>
            <span :class="surveillanceOpt.setBackgourdColor(item.status)"></span>
            <span :class="surveillanceOpt.setTextColor(item.status)">{{item.status}}</span>
        </div>
    </div>
</template>
<script lang="ts" setup>

interface PropsType {
    surveillanceOpt: any;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

const searchRef = ref();

// defineEmits API 来替代 emits
const emitValue = defineEmits(['search-ref']);

onMounted(() => emitValue('search-ref', searchRef));
</script>
