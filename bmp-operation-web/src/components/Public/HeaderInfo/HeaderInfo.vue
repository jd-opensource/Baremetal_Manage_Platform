<template>
    <header
        ref="headerTitle"
        class="header-count"
    >
        <img
            class="header-count-img"
            alt=""
            :src="obj.img"
        />
        <!-- 标题 -->
        <span class="header-count-title">
            {{$t(`${type}.header.${type}`)}}
        </span>
        <slot></slot>
    </header>
</template>

<script lang="ts" setup>
import {HeightType2} from '@utils/publicType';
import {CurrentInstance} from 'utils/publicClass.ts';

interface PropsType {
    type?: string;
    // imgSrc?: string;
}
/**
 * withDefaults API 用来代替默认值，defineProps API 用来替代 props
*/
const props: PropsType = withDefaults(defineProps<PropsType>(), {
    type: 'userList'
});

const proxy = new CurrentInstance().proxy;

const {imagePath} = proxy.$defInfo;

const obj: any = reactive({
    imgSrc: {
        'idcList': imagePath('idcManageDef'),
        'modelList': imagePath('modelDefault'),
        'imageList': imagePath('imageDefault'),
        'deviceList': imagePath('deviceDefault'),
        'hardwareStatusList': imagePath('surveillanceDef'),
        'faultLogList': imagePath('surveillanceDef'),
        'resourceMonitor': imagePath('defaultInBandMonitoringDef'),
        'historyAlarmInfo': imagePath('defaultInBandMonitoringDef'),
        'allAlarmRulesList': imagePath('defaultInBandMonitoringDef'),
        'faultRulesList': imagePath('surveillanceDef'),
        'roleList': imagePath('roleDefault'),
        'userList': imagePath('userDefault'),
        'myProfile': imagePath('userCenter'),
        'securitySettings': imagePath('userCenter'),
        'apiKey': imagePath('userCenter'),
        'license': imagePath('defaultCard'),
        'myMessage': imagePath('messageDef'),
        'messageSettings': imagePath('messageDef')
    },
    img: ''
})
for (const index in obj.imgSrc) {
    if (index === props.type) {
        obj.img = obj.imgSrc[props.type]
    }
}
// defineEmits API 来替代 emits
const emitValue = defineEmits(['header-ref']);

const headerTitle: Ref<HeightType2> = ref<HeightType2>({offsetHeight: 0});

onMounted(() => {
    emitValue('header-ref', headerTitle);
});

</script>

<style lang="scss" scoped>
@import './headerInfo';
</style>
