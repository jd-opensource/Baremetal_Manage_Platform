<template>
    <!-- 主体头部内容 -->
    <div class="info-header">
        <span class="info-line"></span>
        <!-- 基本信息 -->
        <span class="info-title">
            {{$t('idcDetail.basicInfo.title')}}
        </span>
    </div>
    <!-- 修饰符，阻止浏览器冒泡和浏览器默认行为 -->
    <ui-form @submit.stop.prevent>
        <!-- 等分基本信息数据 -->
        <ui-row
            class="info-content"
            :gutter="20"
        >
            <ui-col
                v-for="(item, index) in BaseInfoOpt.baseInfo"
                :span="8"
                :key="index"
            >
                <ui-form-item :label="$filter.withClon(item.label)">
                    <div :class="item.class">
                       {{$filter.emptyFilter(idcDetail?.reactiveArr?.detailInfo?.[item.info])}}
                    </div>
                </ui-form-item>
            </ui-col>
        </ui-row>
    </ui-form>
</template>

<script lang="ts" setup>
import {CurrencyType} from '@utils/publicType';
import BaseInfoOpt from './base';

interface PropsType {
    idcDetail: {
        reactiveArr: {
            detailInfo: CurrencyType;
        }
    };
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
</script>
