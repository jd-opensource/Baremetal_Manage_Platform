/**
 * @file
 * @author
*/
import { 
    ref,
    Ref
} from 'vue';
import {RuleFormType} from '../type';
import publicConfigOpt from './publicConfig';

class EchartsBigOpt {
    diaLog: Ref<boolean> = ref<boolean>(false);
    sizeBigTitle: Ref<string> = ref<string>('');
    sizeRef: Ref<string> = ref<string>('');
    bigSizeLoading: Ref<boolean> = ref<boolean>(true);
    props;
    
    constructor(props: any) {
        this.props = props;
    }

    closeClick = (val: string) => {
        const that = this;
        this.diaLog.value = !this.diaLog.value;
        this.props.ruleForm[`${val}` as keyof typeof that.props.ruleForm] = this.props.ruleForm[`${this.sizeRef.value}` as keyof typeof that.props.ruleForm];
        const newVals = this.props.reactiveArr.echartsParamsData.filter((item: {type: string}) => item.type === val)[0];
        newVals.model = this.props.ruleForm[`${val}` as keyof typeof that.props.ruleForm];
        const newData = this.props.reactiveArr.echartsData.filter((item: {typeVal: string}) => item.typeVal === val)
        const {diskVal, mountpointVal, nicVal} = this.props;
        const deviceVal = diskVal.value ? {device: diskVal.value} : mountpointVal.value ? {device: mountpointVal.value} : nicVal.value ? {device: nicVal.value} : {};
        publicConfigOpt.clearEmptyObj(deviceVal);
        for (const key of newData) {
            key.aggregationMethod = this.props.ruleForm[`${this.sizeRef.value}` as keyof typeof that.props.ruleForm];
            this.props.getEcharts(key, 'small', deviceVal);
        }
    }

    bigClickOpt = (value: string, ...args: string[]) => {
        this.sizeRef.value = value;
        this.sizeBigTitle.value = args[0];
        this.diaLog.value = !this.diaLog.value;
        publicConfigOpt.smallBigSet(this.props, value, 'bigSize', '')
    }

    bigSizeChange = (val: string, obj: RuleFormType, value: string) => {
        publicConfigOpt.smallBigSet(this.props, val, 'bigSize', obj[`${value}` as keyof typeof obj])
    }
}

export default EchartsBigOpt;