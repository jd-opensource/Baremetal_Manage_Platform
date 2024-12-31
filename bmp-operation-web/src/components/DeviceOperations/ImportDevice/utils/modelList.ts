import HeightStatus from './utils';
import {CurrencyType} from '@utils/publicType';
import {CurrentInstance} from 'utils/publicClass';
import {modelListAPI} from 'api/model/request.ts';
import {msgTip} from 'utils/index.ts';

class ModelList extends HeightStatus {
    hasRotate: Ref<boolean> = ref<boolean>(false);
    modelNameLoading: Ref<boolean> = ref<boolean>(false);
    reactiveArr: {
        modelData: CurrencyType[];
    } = reactive<{
        modelData: CurrencyType[];
    }>({
        modelData: [] // 机型名称
    });
    idcList: any;
    formRulesOperate: any;
    heightStatus: any;
    instanceProxy = new CurrentInstance().proxy;

    constructor(idcList: any, formRulesOperate: any, heightStatus: any) {
        super(formRulesOperate);
        this.formRulesOperate = formRulesOperate;
        this.idcList = idcList;
        this.heightStatus = heightStatus;
    };

    getModelList = () => {
        modelListAPI(
            {
                isAll: '1',
                idcId: this.idcList.reactiveArr.idcName
            }
        )
        .then(({data} : {data: {result: {deviceTypes: CurrencyType[];}}}) => {
            if (data?.result?.deviceTypes?.length) {
                this.reactiveArr.modelData = data.result.deviceTypes;
                return;
            }
            this.reactiveArr.modelData = [];
        })
        .catch(({message} : {message: string}) => {
            this.idcList.reactiveArr.idcData.length && msgTip.error(message);
            this.reactiveArr.modelData = [];
        })
        .finally(() => {
            this.modelNameLoading.value = false;
            this.hasRotate.value = false;
        });
    };

    jumpPage = () => {
        const path: string = this.instanceProxy!.$defInfo.routerPath('modelList') as unknown as string;
        window.open(path, '_blank')
    };

    resetData = () => {
        if (!this.formRulesOperate.ruleForm.idcName.length) return;
        if (this.hasRotate.value) return;
        this.formRulesOperate.ruleForm.modelName = '';
        this.reactiveArr.modelData = [];
        this.hasRotate.value = true;
        this.formRulesOperate.ruleFormRef.value!.validate((valid: boolean) => {
            if (!valid) this.heightStatus.hasClick.value = true;
        })
        this.getModelList();
    };
};

export default ModelList;
