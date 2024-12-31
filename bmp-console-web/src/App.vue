<template>
    <div class="content">
        <div v-if="!$route.meta.showvalues">
            <bmp-header></bmp-header>
            <menu-bar></menu-bar>
        </div> 
        <router-view></router-view>
    </div>
</template>
<script setup lang="ts">
import bmpHeader from 'components/BmpHeader/bmpHeader.vue'
import menuBar from 'components/MenuBar/menuBar.vue'

</script>
<style lang="scss">
@import 'assets/css/clear';
@import 'assets/css/global';


#app {
    width: 100%;
    height: 100%;
}
.content{
    width: 100%;
    height: 100%;
}
body{
    overflow: hidden;
    padding-right: 0 !important;
}
</style>
