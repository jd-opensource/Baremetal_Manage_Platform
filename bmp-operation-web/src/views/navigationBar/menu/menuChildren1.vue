<template>
    <ui-sub-menu
        v-for="(ite, index) in navData"
        :index="ite.firstIndex"
        :key="index"
        :class="[
            store.navigationBarStatus.hasNavigationBar ? 'custom-children-menu': 'custom-children-menu-short'
        ]"
    >
        <template #title>
            <img
                alt=""
                :class="[
                    'default-img',
                    navOperate.initShort.value ? 'short-img' : '',
                    navOperate.init.value ? 'init-img' : ''
                ]"
                :src="ite.path.includes($route.path) ? ite.changeImg : ite.defaultImg"
            />
            <ui-tooltip
                placement="right"
                :disabled="locationItem.getLocationItem === 'zh_CN'"
            >
                <template #content>
                    <span>
                        {{ite.title}}
                    </span>
                </template>
                <p
                    :class="[
                        'nav-title-default',
                        setDiffClass('nav-sub-long', '')
                    ]"
                    :style="ite.path.includes($route.path) ? {color: '#108ee9'} : {color: '#333'}"
                >
                    {{ite.title}}
                </p>
            </ui-tooltip>
        </template>
        <ui-menu-item
            v-for="(item, index) in ite.children"
            :key="index"
            :class="[
                store.navigationBarStatus.hasNavigationBar ? 'item-long-width' : 'custom-short-nav'
            ]"
            :index="item.path"
        >
            <template #title>
                <p
                    class="chilren-radius"
                    :style="[item.path].concat(item.otherPath).includes($route.path) ? {background: '#108fe9'} : {background: '#000'}"
                >
                </p>
                <ui-tooltip
                    placement="right"
                    :disabled="store.navigationBarStatus.tootlipDisabled(locationItem, item.title)"
                >
                    <template #content>
                        <span>
                            {{item.title}}
                        </span>
                    </template>
                    <p :class="setDiffClass('nav-sub-long', '')">
                        {{item.title}}
                    </p>
                </ui-tooltip>
            </template>
        </ui-menu-item>
    </ui-sub-menu>
</template>
  
<script setup lang="ts">
import store from 'store/index.ts';
import {setDiffClass} from 'utils/index.ts';
import {locationItem} from 'utils/publicClass.ts';

type IteChildrenType = {
    path: string;
    title: string;
    otherPath: string[];
    defaultImg: string;
    changeImg: string;
}

interface PropsType {
    navData: {
        title: string;
        defaultImg: string;
        changeImg: string;
        otherPath: string[];
        path: string[];
        firstIndex: string;
        children: IteChildrenType[];
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
