<template>
    <div class="hardware-status" :class="getLocationItem === 'zh_CN' ? 'gap-content' : ''" >
        <div
            v-for="(hardWare) in props.hardWareData" 
            :key="hardWare.name" 
            class="hardware-card"
            :span="8" 
        >
            <div class="hardware-card-body" >
                <span class="device-name">{{hardWare.name}}</span>
                <el-divider direction="vertical" class="div-content"/>
                <!-- 0代表正常 1代表异常 -->
                <p 
                    v-if="hardWare.status === 0" 
                    class="status-success"
                >
                    <span class="status-icon"></span>
                    <span class="status-text">{{ $t('instance.detail.normal') }}</span>	
                </P>
                <p 
                    v-if="hardWare.status === 1" 
                    class="status-error"
                >
                    <span class="status-icon"></span>
                    <span class="status-text">{{ $t('instance.detail.abnormal') }}</span>	
                </p>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import VueCookies from 'vue-cookies'; // cookie
/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    hardWareData: any;
};

const props: PropsType = withDefaults(defineProps<PropsType>(), {
    hardWareData: []
});

/**
 * 从cookie中获取语言类型，没有默认zh
*/
const getLocationItem: any = (VueCookies as any).get('X-Jdcloud-Language') || 'zh_CN';
</script>

<style lang="scss" scoped>
@import './hardwareStatus.scss';
</style>