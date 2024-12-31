<template>
    <!-- 导航头 -->
    <header class="operate-header">
        <div
            class="operate-header-info"
            @click="navBarHeader.jumpDefaultList"
        >
            <!-- 导航头大标题 -->
            <p class="operate-header-title">
                {{$t('navigationHeader.title2')}}
            </p>
        </div>
        <!-- 标题信息 -->
        <div class="operate-header-information">
            <!-- 国际化切换 -->
            <i18n-switch :nav-bar-header="navBarHeader"/>
            <img
                alt=""
                :src="navBarHeader.hasNoRead.value ? ($defInfo.imagePath('messageTip') as unknown as string) : ($defInfo.imagePath('messageDef') as unknown as string)"
                @click="navBarHeader.messageListClick"
            />
            <span style="marginRight: 14px">|</span>
            <!-- 操作 -->
            <select-operate :nav-bar-header="navBarHeader"/>
        </div>
    </header>
    <!-- 退出登录 -->
    <div v-if="navBarHeader.loginOutDiaLog.value">
        <login-out
            :dia-log="navBarHeader.loginOutDiaLog.value"
            @dia-log-close="navBarHeader.loginOutCancel"
        />
    </div>
</template>
<script setup lang="ts">
import NavBarHeader from './select';

const navBarHeader = new NavBarHeader();

</script>
<style lang="scss" scoped>
@import './navigationHeader.scss';

</style>
