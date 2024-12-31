<template>
    <ui-menu-item
        v-for="(item, index) in navData"
        :key="index"
        :class="[
            store.navigationBarStatus.hasNavigationBar ? 'item-long-width' : 'item-short-width'
        ]"
        :index="(item.path as string)"
    >
        <img
            alt=""
            v-if="$route.path === item.path"
            :class="[
                'default-img',
                navOperate.initShort.value ? 'short-img' : '',
                navOperate.init.value ? 'init-img' : ''
            ]"
            :src="$route.path === item.path ? item.changeImg : item.defaultImg"
        />
        <img
            alt=""
            v-else
            :class="[
                'default-img',
                navOperate.initShort.value ? 'short-img' : '',
                navOperate.init.value ? 'init-img' : ''
            ]"
            :src="(item.otherPath.includes($route.path)) ? item.changeImg : item.defaultImg"
        />
        <p
            class="nav-title-default"
            v-if="!store.navigationBarStatus.hasNavigationBar"
        >
            {{item.title}}
        </p>
        <template #title>
            <p class="nav-title-default">
                {{item.title}}
            </p>
        </template>
    </ui-menu-item>
</template>
  
<script setup lang="ts">
import store from 'store/index.ts';

interface PropsType {
    // purviewInfoOpt: {
    //     navLeft1(): {};
    // };
    navData: {
        title: string;
        defaultImg: string;
        changeImg: string;
        otherPath: string[];
        path: string;
        firstIndex: string;
        // children: IteChildrenType[];
    }[];
    navOperate: {
        initShort: {
            value: boolean;
        };
        init: {
            value: boolean;
        }
    }
}

/**
 * props参数
*/
withDefaults(defineProps<PropsType>(), {});

</script>

<style lang="scss" scoped>
@import '../navigationBar.scss';
</style>
