<template>
    <!-- 登录页面 -->
    <div class="operate-content">
        <!-- 登录页标题 -->
        <header class="operate-content-header">
            <!-- 登录页大标题 -->
            <p class="operate-content-title">
                {{$t('login.header.title2')}}
            </p>
        </header>
        <!-- 登录页登录窗口 -->
        <form-submit-data/>
        <div></div>
    </div>
</template>

<script setup lang="ts">
import FormSubmitData from './formSumbit/formSumbit.vue';

</script>
<style lang="scss">
@import './login.scss';
</style>
