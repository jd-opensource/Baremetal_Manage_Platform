<template>
    <div class="header-language-change">
        <!-- 中文 -->
        <span
            :class="navBarHeader.locale.value === 'zh_CN' ? 'header-optional' : 'header-disabled'"
            @click="navBarHeader.changeLang('zh_CN')"
        >
            中文
        </span>
        <!-- 分割线 -->
        <span style="color: #979797">/</span>
        <!-- 英文 -->
        <span
            style="marginRight: 26px"
            :class="navBarHeader.locale.value === 'en_US' ? 'header-optional' : 'header-disabled'"
            @click="navBarHeader.changeLang('en_US')"
        >
            EN
        </span>
    </div>
</template>
<script setup lang="ts">
// import {navBarHeader} from '../selectOperate/select';
interface PropsType {
    navBarHeader: any;
};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});

</script>