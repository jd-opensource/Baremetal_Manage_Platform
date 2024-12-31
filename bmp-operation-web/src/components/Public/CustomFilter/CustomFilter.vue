<template>
    <div
        class="custom-table-filter-a"
        v-if="customFilterOpt.customTableFilterDiaLog.value"
        :style="{
            right: customFilterOpt.customRight.value + otherWidth + 'px',
            top: topDefaultHeight + customFilterOpt.searchTipValHeight.value + 'px'
        }"
    >
        <div class="custom-table-filter-b">
            <ui-input
                v-model.trim="customFilterOpt.input.value"
                :placeholder="$t('customFilter.placeholder')"
                :suffix-icon="Search"
                @click.stop=""
                @input="customFilterOpt.inputClick"
            >
            </ui-input>
            <div class="filter-scroll-count">
                <div
                    class="custom-filter-count"
                    v-for="(item, index) in operate.reactiveArr.customFilterList"
                    :key="index"
                    @click.stop="customFilterOpt.selectClick(item, index)"
                >
                    <ui-checkbox
                        size="small"
                        v-model="item.checkboxStatus"
                        @click.stop=""
                        @change="customFilterOpt.checkBoxChange"
                    >
                        <span class="text-tip">
                            {{item.label}}
                        </span>
                    </ui-checkbox>
                </div>
            </div>
            <div
                class="custom-footer-btn"
                @click.stop=""
            >
                <span
                    :class="[customFilterOpt.disabledStatus.value ? 'select-opt' : 'disabled-select']"
                    @click.stop="customFilterOpt.filterClick"
                >
                    {{$t('customFilter.btn.confirm')}}
                </span>
                <span @click.stop="customFilterOpt.reset">
                    {{$t('customFilter.btn.reset')}}
                </span>
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import {Search} from '@element-plus/icons-vue';
import MethodsTotal from 'utils/index.ts';
import {CurrentInstance, locationItem} from 'utils/publicClass.ts';

interface PropsType {
    operate: any,
    typeClass: string;
    otherWidth?: number;
    topDefaultHeight: number;
    otherFilterStyle?: any;
    filterOperate?: any;
}

type ItemType = {label: string; yin: string};

// withDefaults API 用来代替默认值，defineProps API 用来替代 props
const props = withDefaults(defineProps<PropsType>(), {
    otherWidth: 0
});

let span = document.createElement('span');

class CustomFilterOpt {
    customRight: Ref<number> = ref<number>(0);
    disabledStatus: Ref<boolean> = ref<boolean>(false);
    customTableFilterDiaLog: Ref<boolean> = ref<boolean>(false);
    input: Ref<string> = ref<string>('');
    mitt = new CurrentInstance();
    timer: any;
    searchTipValHeight: Ref<number> = ref<number>(0);

    constructor() {
        watch(() => props.operate.reactiveArr.customFilterList, (newValue) => {
            this.disabledStatus.value = newValue.some((item: {checkboxStatus: boolean}) => item.checkboxStatus);
        }, {
            deep: true
        })
        watch(() => this.customTableFilterDiaLog.value, (newValue) => {
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('custom-filter-dia-log', newValue);
            if (newValue) {
                this.pageResize();
            }
        })
        onUnmounted(() => {
            clearTimeout(this.timer);
            this.mitt.instanceMitt?.proxy?.$Bus?.off('custom-filter-input-clear', this.inputClear);
            this.mitt.instanceMitt?.proxy?.$Bus?.off('search-tip-value', this.watchSearchTip)
            span.remove();
        });
        this.init();
        this.watchInputClear();
   
        this.mitt.instanceMitt?.proxy?.$Bus?.on('search-tip-value', this.watchSearchTip)
    }

    watchSearchTip = (value: number) => {
        this.searchTipValHeight.value = value;
    }

    watchInputClear = () => {
        this.mitt.instanceMitt?.proxy?.$Bus.on('custom-filter-input-clear', this.inputClear)
    }

    inputClear = () => {
        this.input.value = '';
    }

    init = () => {
        nextTick(() => {
            const defClsCell = document.querySelector(`.${props.typeClass} .cell`);
            defClsCell?.appendChild(span);
            span.className = 'el-table__column-filter-trigger';
            span.addEventListener('click', (event) => {
                event.stopPropagation();
                this.resetOtherFilterStatus();
                this.initFilterClick();
            })
            document.body.addEventListener('click', () => {
                this.customTableFilterDiaLog.value = false;
            })
        })
    }

    pageResize = async () => {

        await nextTick(() => {
            this.mitt.instanceMitt?.proxy?.$Bus?.on('custom-filter-onresize', () => {
                // @ts-ignore
                const defCls: {offsetWidth: number} = document.querySelector(`.${props.typeClass}`);
                if (!defCls?.offsetWidth) return;
                this.timer = setTimeout(() => {
                    this.customRight.value = Object.is(locationItem.getLocationItem, 'zh_CN') ? defCls.offsetWidth / 2 - 50 : defCls.offsetWidth / 2 - 30;
                }, 0)
            });
        });
    }

    checkBoxChange = () => {
        this.resetOtherFilterStatus();   
    }

    resetOtherFilterStatus = () => {
        for (let index in props.otherFilterStyle.filterStatus) {
            if (!props.filterOperate.reactiveArr?.filterParams[index]?.length) {
                props.otherFilterStyle.filterStatus[index] = false;
            }
        }
    }

    reset = async () => {
        this.input.value = '';
        this.customTableFilterDiaLog.value = false;
        await nextTick(() => {
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('custom-filter-reset');
        });
    }

    filterClick = async () => {
        const data = props.operate.reactiveArr.customFilterList.filter((item: {checkboxStatus: boolean}) => item.checkboxStatus);
        this.customTableFilterDiaLog.value = false;
        await nextTick(() => {
            this.mitt.instanceMitt?.proxy?.$Bus?.emit('custom-filter-click', data);
        });
    }

    selectClick = (item: {checkboxStatus: boolean}, index: number) => {
        item.checkboxStatus = !item.checkboxStatus;
        props.operate.reactiveArr.copyDataVal[index].checkboxStatus = item.checkboxStatus;
    }

    inputClick = MethodsTotal.debounceFun((e: string) => {
        if (e?.length) {
            const copyData = props.operate.reactiveArr.copyDataVal;
            const newVal = copyData.filter((item: ItemType) => (this.handleSearch(item.label) || this.handleSearch(item.yin)));
            props.operate.reactiveArr.customFilterList = newVal;
            return;
        }
        props.operate.reactiveArr.customFilterList = props.operate.reactiveArr.copyDataVal;
    })

    handleSearch = (type: string) => {
        const ipt = this.input.value.toLowerCase();
        return (type.toLowerCase().indexOf(ipt) > -1);
    }

    initFilterClick = () => {
        // @ts-ignore
        const init: {offsetWidth: number} = document.querySelector(`.${props.typeClass}`);
        this.customRight.value = Object.is(locationItem.getLocationItem, 'zh_CN') ? init.offsetWidth / 2 - 50 : init.offsetWidth / 2 - 30;
        this.customTableFilterDiaLog.value = !this.customTableFilterDiaLog.value;
    }
}

const customFilterOpt = new CustomFilterOpt();

</script>
<style lang="scss" scoped>
@import './index.scss';
</style>