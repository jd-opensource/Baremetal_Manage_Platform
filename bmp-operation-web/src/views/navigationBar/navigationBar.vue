<template>
    <navigation-header/>
    <!-- 侧边栏 -->
    <div class="operate-navigationbar">
        <div :class="[store.navigationBarStatus.hasNavigationBar ? 'long-set-nav' : 'set-nav']">
            <div class="nav-col">
                <ui-scroll-bar>
                    <!-- 菜单栏 -->
                    <ui-menu
                        :default-active="purviewInfoOpt.defaultActive($route)"
                        :router="false"
                        :collapse="!store.navigationBarStatus.hasNavigationBar"
                        @get-menu-ref="navOperate.getMenuRef"
                        @open="navOperate.openClick"
                        @close="navOperate.closeClick"
                        @select="navOperate.select"
                    >
                        <menu-left
                            :nav-data="purviewInfoOpt.navLeft1()"
                            :nav-operate="navOperate"
                        >
                        </menu-left>
                        <menu-children1
                            :nav-data="purviewInfoOpt.outOfMonitoring()"
                            :nav-operate="navOperate"
                        >
                        </menu-children1>
                        <menu-left
                            :nav-data="purviewInfoOpt.navLeft2()"
                            :nav-operate="navOperate"
                        >
                        </menu-left>
                        <menu-children1
                            :nav-data="purviewInfoOpt.messagePersonal()"
                            :nav-operate="navOperate"
                        >
                        </menu-children1>
                    </ui-menu>
                </ui-scroll-bar>
                <div class="custom-height"></div>
                <!-- 侧边栏分割线 -->
                <div
                    :class="[
                        store.navigationBarStatus.hasNavigationBar ? 'nav-long-line' : 'nav-short-line'
                    ]"
                >
                </div>
                <!-- 侧边栏-展开收起icon -->
                <div
                    :class="[store.navigationBarStatus.hasNavigationBar ? 'long-nav-icon' : 'short-nav-icon']"
                    @click="navOperate.unfoldCollapseClick"
                >
                    <img
                        alt=""
                        :src="($defInfo.imagePath('shrinkExpand') as unknown as string)"
                        :class="[store.navigationBarStatus.hasNavigationBar ? 'img' : 'img-rotate']" 
                    />
                </div>
            </div>
        </div>
        <!-- 子路由-侧边栏导航对应的路由页面 -->
        <div
            :class="[
                store.navigationBarStatus.hasNavigationBar ? 'long-content' : 'short-content'
            ]"
        >
            <!-- 子路由 -->
            <router-view></router-view>
        </div>
    </div>
</template>
  
<script setup lang="ts">
import store from 'store/index.ts';
import MenuLeft from './menu/menuLeft.vue';
import NavOperate from './utils/navOperate';
import PurviewInfoOpt from './utils/purviewInfo';
import MenuChildren1 from './menu/menuChildren1.vue';

const purviewInfoOpt = new PurviewInfoOpt();

const navOperate = new NavOperate(purviewInfoOpt);

</script>

<style lang="scss" scoped>
@import './navigationBar.scss';

</style>
