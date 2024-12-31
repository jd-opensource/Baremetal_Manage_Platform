import {AxiosError} from 'axios';
import {msgTip} from 'utils/index.ts';
import {CurrentInstance} from 'utils/publicClass.ts';

class IdcDetail {
    instanceProxy = new CurrentInstance().proxy;

    constructor (props: any, formRuleOperate: any) {
        this.getIdcDetail(props, formRuleOperate);
    };

    getIdcDetail = async (props: any, formRuleOperate: any) => {
        try {
            const params = {
                idcid: props.itemData?.value?.idcId??props.itemData?.idcId,
                show: 'iloPassword,switchPassword1,switchPassword2'
            };
            const detailApi = await this.instanceProxy.$idcApi.idcDetailAPI(params);
            if (detailApi.data?.result && Object.keys(detailApi.data?.result).length) {
                const {iloPassword, switchPassword1, switchPassword2} = detailApi.data.result;
                formRuleOperate.ruleForm.iloPassword = iloPassword;
                formRuleOperate.ruleForm.switchPassword1 = switchPassword1;
                formRuleOperate.ruleForm.switchPassword2 = switchPassword2;
            }
        }
        catch (e) {
            const err = e as AxiosError;
            msgTip.error(err.message);
        }
    }
};

export default IdcDetail;
