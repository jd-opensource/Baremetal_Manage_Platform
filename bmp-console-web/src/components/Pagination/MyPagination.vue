<template>
    <!-- 默认使用组件的分页 -->
    <el-pagination v-if="hasUseDefault"/>
    <!-- 使用自己封装的分页组件 -->
    <div
        v-else
        :style="styleCss ? paginationComponents.setStyleCss: ''"
        class="footer-count"
    >
        <slot></slot>
        <div class="my-pagination">
            <!-- 总共多少条 -->
            <p class="my-pagination-total">
                {{$t('uiPagination.total', [total])}}
            </p>
            <!-- 每页显示多少条 -->
            <div
                :class="[
                    total > 0 ? 'my-pagination-select' : 'short-select',
                    locationItem.getLocationItem === 'en_US' ? 'set-width' : '',
                    styleCss ? 'set-input-background' : ''
                ]"
            >
                <el-select
                    size="large"
                    v-model="paginationComponents.size.value"
                >
                    <el-option
                        v-for="(item, index) in pageSizes"
                        :key="index"
                        :label="item + paginationComponents.defaultText"
                        :value="item"
                        @click="paginationComponents.pageSizeChange(item)"
                    />
                </el-select>
            </div>
            <!-- 左侧箭头，上一页 -->
            <span
                class="my-pagination-prev"
                v-if="total > 0"
            >
                <el-icon
                    color="#333"
                    :size="12"
                    :class="paginationComponents.currentPage.value === 1 ? 'my-pagination-disabled' : 'my-pagination-pointer'"
                    @click="paginationComponents.changePage(false)"
                >
                    <ArrowLeftBold />
                </el-icon>
            </span>
            <span v-if="(paginationComponents.currentPage.value > 4) && paginationComponents.pageCount.value > 8">
                <span
                    :class="[
                        'my-pagination-number', paginationComponents.currentPage.value === item ? 'active' : ''
                    ]"
                    v-for="item in [1, 2]"
                    :key="item"
                    @click="paginationComponents.changePage(item)"
                >
                    {{item}}
                </span>
            </span>
            <!-- 页码切换 -->
            <span
                v-if="(paginationComponents.currentPage.value > 5) && paginationComponents.pageCount.value > 8"
            >
                <!-- ...省略号 -->
                <span
                    style="cursor: pointer; margin: 0 3px"
                    v-if="!paginationComponents.hasPrevIptNumber.value"
                    @click="paginationComponents.hasPrevIptNumberClick"
                >
                    {{$filter.paginationEllipsis('')}}
                </span>
                <!-- 点击省略号出现input-number输入框，切换页码 -->
                <div v-else>
                    <el-input-number
                        size="small"
                        controls-position="right"
                        v-model="paginationComponents.iptNumberPreValue.value"
                        :min="1"
                        :max="paginationComponents.pageCount.value"
                        @change="paginationComponents.inputNumberChange"
                    />
                
                </div>
            </span>
            <!-- 点击对应页码，进行页码切换 -->
            <span
                :class="[
                    'my-pagination-number', paginationComponents.currentPage.value === item ? 'active' : ''
                ]"
                v-for="item in paginationComponents.changeNumberData.value.slice(0, -2)"
                :key="item"
                @click="paginationComponents.changePage(item)"
            >
                {{item}}
            </span>
            <!-- 右侧箭头 -->
            <span
                v-if="paginationComponents.currentPage.value !== paginationComponents.pageCount.value - 4 && paginationComponents.currentPage.value !== paginationComponents.pageCount.value -3 && paginationComponents.currentPage.value !== paginationComponents.pageCount.value -1 && paginationComponents.currentPage.value !== paginationComponents.pageCount.value - 2 && paginationComponents.currentPage.value !== paginationComponents.pageCount.value && paginationComponents.pageCount.value > 8"
            >
                <!-- 右侧省略号... -->
                <div
                    style="cursor: pointer; margin: 0 3px"
                    v-if="!paginationComponents.hasNextIptNumber.value"
                    @click="paginationComponents.hasNextEllipsisClick"
                >
                    {{$filter.paginationEllipsis('')}}
                </div>
                <!-- 点击省略号出现input-number输入框，切换页码 -->
                <div v-else>
                    <el-input-number
                        size="small"
                        controls-position="right"
                        v-model="paginationComponents.iptNumberNextValue.value"
                        :min="1"
                        :max="paginationComponents.pageCount.value"
                        @change="paginationComponents.inputNumberChange"
                    />
                </div>
            </span>
            <span v-if="paginationComponents.pageCount.value > 1">
                <span
                    :class="[
                        'my-pagination-number', paginationComponents.currentPage.value === item ? 'active' : ''
                    ]"
                    v-for="item in [paginationComponents.pageCount.value - 1, paginationComponents.pageCount.value]"
                    :key="item"
                    @click="paginationComponents.changePage(item)"
                >
                    {{item}}
                </span>
            </span>
            <span v-else>
                <span
                    :class="[
                        'my-pagination-number', paginationComponents.currentPage.value === item ? 'active' : ''
                    ]"
                    v-for="item in [paginationComponents.pageCount.value]"
                    :key="item"
                    @click="paginationComponents.changePage(item)"
                >
                    {{item}}
                </span>
            </span>
            <!-- 右侧箭头，点击切换下一页 -->
            <span
                class="my-pagination-next"
                v-if="total > 0"
            >
                <el-icon
                    color="#333"
                    :size="12"
                    :class="[
                        paginationComponents.currentPage.value === paginationComponents.pageCount.value
                        ? 'my-pagination-disabled'
                        : 'my-pagination-pointer'
                    ]"
                    @click="paginationComponents.changePage(true)"
                >
                    <ArrowRightBold />
                </el-icon>
            </span>
            <!-- 总共多少页 -->
            <p :class="[total > 0 ? 'my-pagination-altogether' : 'short-altogether']">
                {{$t('uiPagination.page', [paginationComponents.pageCount.value])}}
            </p>
        </div>
    </div>
</template>
<script setup lang="ts">
import { 
    ref,
    Ref,
    reactive,
    onMounted, 
    computed,
    onUnmounted,
    ComputedRef,
    onBeforeUnmount, 
    getCurrentInstance
} from 'vue';
import {ElPagination} from 'element-plus'; // paginnation
import VueCookies from 'vue-cookies'; // cookie
import i18n from 'lib/i18n/index.ts'; // 国际化
/**
 * 国际化
*/
class LocationItem {
    // cookie ts规范校验
    cookie: {
        [x: string]: unknown;
        get?: Function;
        set?: Function;
    } = VueCookies;
    getLocationItem = (this.cookie?.get && this.cookie.get('X-Jdcloud-Language'))?? 'zh_CN';
};
const locationItem: LocationItem = new LocationItem();

// 语言-使用国际化
class Language {
    t = i18n.global.t
};
const language = new Language();

/**
 * 父组件传递的props数据类型
*/
interface PropsType {
    styleCss?: boolean;
    pageSizes: number[];
    pageNumber: number;
    total: number;
    pageSize: number;
    hasUseDefault: boolean;
};

// defineEmits API 来替代 emits，子传父，事件回传
const emitValue = defineEmits(['change-page', 'page-sizes-change']);

/**
 * withDefaults API 可以指定默认值，defineProps API 来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    styleCss: false,
    pageSizes: () => [20, 50, 100], // 可选择的每页多少条
    pageNumber: 1, // 当前页
    total: 0, // 总条数
    pageSize: 20, // 当前页当前条数
    hasUseDefault: true, // 是否使用默认分页组件
});

/**
 * 分页组件
*/
class PaginationComponents {
    // 设置style样式
    setStyleCss: {[x: string]: string;} = {
        width: 'calc(100% - 46px) !important',
        // background: 'rgb(250, 250, 251)'
    };
    // 默认文案信息
    defaultText: string = language.t('uiPagination.piecesAndPage');
    // 左侧inputNumber输入框value
    iptNumberPreValue: Ref<number> = ref<number>(0);
    // 右侧inputNumber输入框value
    iptNumberNextValue: Ref<number> = ref<number>(0);
    // 是否显示右侧inputNumber输入框
    hasNextIptNumber: Ref<boolean> = ref<boolean>(false);
    // 是否显示左侧inputNumber输入框
    hasPrevIptNumber: Ref<boolean> = ref<boolean>(false);
    // 当前条数
    size: Ref<number> = ref<number>(props?.pageSize) || 20;
    // 当前页
    currentPage: Ref<number> = ref<number>(props.pageNumber);
    // 计算总页数
    pageCount: ComputedRef<number> = computed(() => Math.ceil(props.total / this.size.value));

    /**
     * 页码显示组合
    */
    changeNumberData = computed(() => {
        const result: number[] = [];
        const numberShowData = [
            [
                (pageCount: ComputedRef<number>) => (pageCount.value <= 8),
                (pageCount: {value: number}) => {
                    for (let i: number = 1; i <= pageCount.value; i++) {
                        result.push(i);
                    }
                }
            ],
            [
                (_: ComputedRef<number>, currentPage: {value: number}) => (currentPage.value <= 4),
                () => {
                    // 总页数大于5页的时候，控制左右两边的省略号显示隐藏，页码的显示个数与选中页码居中
                    for (let i: number = 1; i <= 7; i++) {
                        result.push(i);
                    }
                }
            ],
            [
                (pageCount: ComputedRef<number>, currentPage: {value: number}) => (currentPage.value >= 3 && currentPage.value <= pageCount.value - 5),
                (_: ComputedRef<number>, currentPage: {value: number}) => {
                    for (let i: number = currentPage.value - 2; i <= currentPage.value + 4; i++) {
                        result.push(i);
                    }
                }
            ],
            [
                (_: ComputedRef<number>, currentPage: {value: number}) => (currentPage.value === 5),
                (pageCount: ComputedRef<number>) => {
                    for (let i: number = 3; i < pageCount.value + 1; i++) {
                        result.push(i);
                    }
                }
            ],
            [
                (pageCount: ComputedRef<number>, currentPage: {value: number}) => (currentPage.value > pageCount.value - 5),
                (pageCount: ComputedRef<number>) => {
                    for (let i: number = pageCount.value - 5; i < pageCount.value + 1; i++) {
                        result.push(i);
                    }
                }
            ]
        ];
        for (const key of numberShowData) {
            if (key[0](this.pageCount, this.currentPage)) {
                key[1](this.pageCount, this.currentPage);
                break;
            }
        }
        return result;
    });

    /**
     * 当前每页展示条数变化，页数都要置为1，相当于从第一页开始切换
    */
    pageSizeChange = (size: number): void => {
        this.currentPage.value = 1;
        emitValue('page-sizes-change', size, 'size');
    };

    /**
     * 右侧省略号点击操作-inputNumber输入框显示
     * @return {boolean} hasNextIptNumber.value 省略号的显示隐藏
    */
    hasNextEllipsisClick = (): boolean => {
        return this.hasNextIptNumber.value = !this.hasNextIptNumber.value;
    };

    /**
     * 左侧省略号点击操作-inputNumber输入框显示
     * @return {boolean} hasNextIptNumber.value 省略号的显示隐藏
    */
    hasPrevIptNumberClick = (): boolean => {
        return this.hasPrevIptNumber.value = !this.hasPrevIptNumber.value;
    };

    /**
     * 点击上一页下一页页码改变页码
     * @param {boolean | string} type 【左右】切换是boolean，页码点击是判断类型
    */
    changePage = (type: boolean | number): void => {
        // 如果点击的是页码
        if (typeof type === 'number') {
            // 当前页的值等于页码
            this.currentPage.value = type;
        }
        // 点击上一页按钮
        else if (!type) {
            // 当前页小于等于1，说明不能再点击了
            if (this.currentPage.value <= 1) {
                // return，后续不执行
                return;
            }
            // 页数-1
            this.currentPage.value -= 1;
        }
        else {
            // 点击下一页按钮，当前页数大于等于总页数，说明不能点击了
            if (this.currentPage.value >= this.pageCount.value) {
                // return，后续不执行
                return;
            }
            // 页数+1
            this.currentPage.value += 1;
        }
        // 设置inputNumber输入框的值
        this.setInputNumberValue();
        // 回传父组件当前页码，可以处理相关操作
        emitValue('change-page', this.currentPage.value, 'num');
    };

    /**
     * 输入框数据更新，当前页数也要随之切换
     * @param {number} value inputNumber输入框输入的值
    */
    inputNumberChange = (value: number): void => {
        this.currentPage.value = value;
        emitValue('change-page', value, 'num');
        // 设置inputNumber输入框的值
        this.setInputNumberValue();
    };

    /**
     * 设置输入框value值，页码切换，值要跟着变换
    */
    setInputNumberValue = (): void => {
        this.iptNumberNextValue.value = this.currentPage.value;
        this.iptNumberPreValue.value = this.currentPage.value;
    };
};
// 实例化-分页组件
const paginationComponents: PaginationComponents = new PaginationComponents();

</script>


<style lang="scss">
$center: center;
$color1: #108EF9;
$color2: #333;
$color3: #999;
$color7: #999;
$fontSize: 12px;
$display: flex;
$fontWeight: 400;
$percentageWidth:100%;
$fontSmallWeight:400;
$defaultColor:#333;
$fontLargeWeight: 600;

.footer-count {
    //margin-top: 20px;
    position: relative;
    width: $percentageWidth;
}

.footer-count {
    .my-pagination {
        display: $display;
        justify-content: $center;
        align-items: $center;

        &-total {
            font-size: $fontSize;
            color: $color7;
            display: $display;
            align-items: $center;
            font-weight: $fontSmallWeight;
            margin: 0 25px 2px 0;

            span {
                margin: 0 5px;
            }
        }

        &-prev {
            margin-right: 7px;
        }

        &-disabled {
            cursor: not-allowed;
            opacity: .4;
        }

        &-pointer {
            cursor: pointer;
            opacity: 1;
        }

        &-next {
            margin-left: 7px;
        }

        &-number {
            display: inline-block;
            padding: 5px 7px;
            font-size: $fontSize;
            cursor: pointer;
            color: $defaultColor;

            &:hover {
                color: $color1 !important;
            }
        }

        .active {
            color: $color1 !important;
            pointer-events: none;
            font-weight: $fontLargeWeight;
        }

        .el-input-number--small {
            width: 55px;

            input {
                text-align: left;
                padding-left: 5px !important;
            }
        }

        .el-input__inner {
            font-size: $fontSize;
            padding: 0 !important;
        }

        .el-input__suffix {
            background: url('~@/assets/img/arrowDown.png') $center $center no-repeat;
            background-size: 100% auto;
            right: 0px;

            svg {
                display: none;
            }
        }

        &-select,
        .short-select {
            display: $display;
            align-items: $center;

            input {
                padding: 0 20px;
                width: 70px;
                border: 0;
            }
        }

        .set-width {
            display: $display;
            align-items: $center;

            input {
                padding: 0 20px;
                // width: 100px !important;
                border: 0;
            }
        }

        // .set-input-background {
        //     input {
        //         background: rgb(250, 250, 251);
        //     }
        // }

        &-select {
            margin-right: 25px;
            margin-top: 1px;
        }

        // .short-select,
        // .set-width {
        //     margin-right: 5px;
        // }

        &-altogether,
        .short-altogether {
            font-size: $fontSize;
            color: $color7;
            font-weight: $fontSmallWeight;
        }

        &-altogether {
            margin-left: 24px;
        }

        .short-altogether {
            margin-left: 5px;
        }
    }
}


</style>
