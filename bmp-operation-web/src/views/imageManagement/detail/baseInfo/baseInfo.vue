<template>
    <!-- 阻止浏览器默认事件以及阻止冒泡事件 -->
    <section
        class="operate-management-detail-info"
        v-if="Object.is(show, 'basicInfo')"
    >
        <ui-form
            v-loading="loading"
            @submit.stop.prevent
        >
            <!-- Layout布局 -->
            <ui-row
                class="info-content"
                :gutter="20"
            >
                <ui-col v-for="(item, index) in baseInfoData" :span="8" :key="index">
                    <ui-form-item
                        :label="$filter.withClon(item.label)"
                        :class="item.formItemClass ? setDiffClass(item.formItemClass[0], item.formItemClass[1]) : ''"
                    >
                        <div
                            class="set-wrap info-name"
                            v-if="item.hasEdit"
                        >
                            <p class="edit-opt-wrap">
                                <span>
                                    {{$filter.emptyFilter(reactiveArr?.detail?.[item.info])}}
                                    <img
                                        alt=""
                                        class="edit-icon"
                                        :src="($defInfo.imagePath('descEdit') as unknown as string)"
                                        @click="emitValue('desc')"
                                    />
                                </span>
                            </p>
                        </div>
                        <div
                            class="set-wrap info-name"
                            v-else
                        >
                            {{$filter.emptyFilter(reactiveArr?.detail?.[item.info])}}
                        </div>
                    </ui-form-item>
                </ui-col>
            </ui-row>
        </ui-form>
    </section>

</template>
<script lang="ts" setup>
import {setDiffClass} from 'utils/index.ts';
import { CurrencyType } from '@utils/publicType';

interface PropsType {
    reactiveArr: {
        detail: CurrencyType;
    };
    loading: boolean;
    show: string;
    baseInfoData: {
        label: string;
        info: string;
        class: string;
        formItemClass?: string[];
        hasEdit: boolean;
    }[];
    // imgSrc?: string;
}

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
withDefaults(defineProps<PropsType>(), {});
const emitValue = defineEmits(['desc']);
</script>