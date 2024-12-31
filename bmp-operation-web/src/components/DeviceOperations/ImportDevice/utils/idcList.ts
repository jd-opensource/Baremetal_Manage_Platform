import HeightStatus from './utils';
import {AxiosError} from 'axios';
import {CurrencyType} from '@utils/publicType';
import {filterData, msgTip} from 'utils/index.ts';
import {locationItem, CurrentInstance} from 'utils/publicClass.ts';

class IdcList extends HeightStatus {
    idcListLoading: Ref<boolean> = ref<boolean>(true);
    reactiveArr: {
        idcName: string;
        idcData: CurrencyType[];
    } = reactive<{
        idcName: string;
        idcData: CurrencyType[];
    }>({
        idcName: '',
        idcData: [] // 机房名称
    });
    proxy = new CurrentInstance().proxy;

    constructor (formRulesOperate: any, fn: any) {
        super(formRulesOperate);
        this.getIdcList();
        this.watchIdcName(formRulesOperate, fn);
    };

    getIdcList = async () => {
        try {
            const idcApi = await this.proxy.$idcApi.idcListAPI({isAll: '1'});
            if (idcApi?.data?.result?.idcs?.length) {
                const newIdcsData = filterData((idcApi.data.result.idcs), 'name').map((item: CurrencyType) => {
                    return {
                        ...item,
                        newIdcName: this.#setName(locationItem.getLocationItem, item)
                    }
                });
                this.reactiveArr.idcData = newIdcsData;
                return;
            }
            throw new Error('');
        }
        catch (e) {
            const err = e as AxiosError;
            this.reactiveArr.idcData = [];
            err.message && msgTip.error(err.message);
        }
        finally {
            this.idcListLoading.value = false;
        }
    };

    #setName = (name: string, item: CurrencyType) => {
        const keyName: Map<string, string> = new Map([
            ['zh_CN', item.name],
            ['en_US', item.nameEn]
        ]);
        return keyName.get(name);
    };

    watchIdcName = (formRulesOperate: any, fn: any) => {
        watch(() => formRulesOperate.ruleForm.idcName, (newValue: string) => {
            if (newValue) {
                this.reactiveArr.idcName = newValue;
                if (formRulesOperate.ruleForm.modelName) {
                    formRulesOperate.ruleForm.modelName = '';
                    this.hasClick.value = true;
                }
                fn();
                // modelList.modelNameLoading.value = true;
                // modelList.getModelList();
            }
        })
    };
};

export default IdcList;
