<template>
    <div class="operate-header-login-info">
        <ui-dropdown :popper-class="type" @visible-change="dropHover">
            <!-- 用户信息 -->
            <template #dropdownTitle>
                <!-- 用户头像 -->
                <img
                    class="user-header-img"
                    alt=""
                    :src="($defInfo.imagePath('loginUser') as unknown as string)"
                />
                <!-- 用户名称 -->
                <span class="user-header-name">
                    <span>
                        Hello,
                    </span>
                    <span>
                        {{$filter.emptyFilter(navBarHeader.loginUserName.value)}}
                    </span>
                </span>
                <!-- 下拉箭头 -->
                <ui-icon
                    :class="['user-header-icon', dropStatus ? 'route' : 'default-route']"
                    color="#333"
                    :size="18"
                >
                    <arrow-down />
                </ui-icon>
            </template>
            <ui-dropdown-item>
                <p class="user-count">
                    <img
                        class="user-img"
                        alt=""
                        :src="($defInfo.imagePath('loginUser') as unknown as string)"
                    />

                    <ui-tooltip
                        placement="left"
                        v-if="navBarHeader.loginUserName.value.length > 25"
                    >
                        <template #content>
                            <span>
                                {{navBarHeader.loginUserName.value}}
                            </span>
                        </template>
                        <span class="user-name">
                            {{navBarHeader.loginUserName.value}}
                        </span>
                    </ui-tooltip>
                    <span
                        class="user-name"
                        v-else
                    >
                        {{navBarHeader.loginUserName.value}}
                    </span>
                    <ui-tooltip
                        placement="left"
                        v-if="navBarHeader.email.value.length > 25"
                    >
                        <template #content>
                            <span class="user-email">
                                {{navBarHeader.email.value}}
                            </span>
                        </template>
                        <span class="user-email">
                            {{navBarHeader.email.value}}
                        </span>
                    </ui-tooltip>
                    <span
                        class="user-email"
                        v-else
                    >
                        {{$filter.emptyFilter(navBarHeader.email.value)}}
                    </span>
                    <span
                        class="user-center right line"
                        @click="navBarHeader.userClick('myProfile')"
                    >
                            {{$t('navigationHeader.myProfile')}}
                    </span>
                </p>
            </ui-dropdown-item>
            <ui-dropdown-item>
                <span
                    class="user-center"
                    @click="navBarHeader.userClick('securitySettings')"
                >
                    {{$t('navigationHeader.securitySettings')}}
                </span>
            </ui-dropdown-item>
            <ui-dropdown-item>
                <span
                    class="user-center"
                    @click="navBarHeader.userClick('apiKey')"
                >
                    {{$t('navigationHeader.apiKey')}}
                </span>
            </ui-dropdown-item>
            <!-- 退出登录操作 -->
            <ui-dropdown-item>
                <span
                    class="login-out"
                    @click="navBarHeader.loginOut"
                >
                    {{$t('navigationHeader.loginOut')}}
                </span>
            </ui-dropdown-item>
        </ui-dropdown>
    </div>
</template>
<script setup lang="ts">

interface PropsType {
    navBarHeader: any;
    // userOperate: any;
};
// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
const type: string = 'user';
const dropStatus: Ref<boolean> = ref<boolean>(false);
const dropHover = (type: boolean) => {
    dropStatus.value = type;
}

</script>
<style lang="scss" scoped>
@import './selectOperate';
</style>
