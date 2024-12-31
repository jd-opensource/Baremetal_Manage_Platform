import {AxiosError} from 'axios';
import {msgTip} from 'utils/index.ts'; // 详情表格class类
import {language, CurrentInstance} from 'utils/publicClass.ts';
import {NewTableType, CurrencyType, ReactiveArrType, LocationQueryValue} from '../typeManagement';
import RegularContent from 'utils/regular.ts';
import AllStaticData from 'staticData/staticData.ts';

/**
 * 机房详情
*/
class IdcDetail {
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(true);
    // 上一个页面携带的idcid，用来发送详情请求
    idcId: LocationQueryValue | LocationQueryValue[] = '';
    reactiveArr: ReactiveArrType = reactive<ReactiveArrType>({
        detailInfo: {},
        tableData: [] // 表格全局管理数据
    });
    // 是否显示密码
    password: Ref<string> = ref<string>('');
    instanceProxy = new CurrentInstance().proxy;

    constructor (routerOperate: any) {
        this.idcId = routerOperate.router?.currentRoute?.value?.query?.idcId?? '';
        this.getDetailData(routerOperate);
    };

    getDetailData = async (routerOperate: any) => {
        const {value} : {value: string;} = this.password;
        const params: {show?: string;} = value?.length ? {show: this.password.value} : {};
        try {
            const detailApi = await this.instanceProxy.$idcApi.idcDetailAPI({idcid: this.idcId, ...params});
            if (detailApi?.data?.result && Object.keys(detailApi.data?.result).length) {
                this.reactiveArr.detailInfo = detailApi.data.result;
                const newTableData: NewTableType[][] = this.#processingData(detailApi.data, params);
                this.reactiveArr.tableData = newTableData[0];
                return;
            }
            throw new Error('');
        }
        catch (e) {
            const err = e as AxiosError;
            this.reactiveArr.tableData = [];
            this.reactiveArr.detailInfo = {};
            const path: string = this.instanceProxy.$defInfo.routerPath('idcList');
            if (AllStaticData.noFoundData.includes(err.message)) {
                routerOperate.router.push(path);
                return;
            }
            err.message && msgTip.error(err.message);
        }
        finally {
            this.isLoading.value = false;
        }
    };

    #processingData = (data: {result: CurrencyType;}, params: {show?: string}) => {
        return [data.result].map((item: CurrencyType): NewTableType[] => {
            return [
                {
                    name: language.t('idcDetail.allAdministration.loginAccount'),
                    user: item.iloUser,
                    password: this.#setPassword('iloPassword', params, item),
                    type: 'iloPassword',
                    hasPassword: 'iloPassword' === params?.show
                },
                {
                    name: language.t('idcDetail.allAdministration.switchboardOneNum'),
                    user: item.switchUser1,
                    password: this.#setPassword('switchPassword1', params, item),
                    type: 'switchPassword1',
                    hasPassword: 'switchPassword1' === params?.show
                },
                {
                    name: language.t('idcDetail.allAdministration.switchboardTwoNum'),
                    user: item.switchUser2,
                    password:this.#setPassword('switchPassword2', params, item),
                    type: 'switchPassword2',
                    hasPassword: 'switchPassword2' === params?.show
                }
            ]
        });
    };

    #setPassword = (type: string, ...args: any) => {
        const {$filter} = this.instanceProxy;
        return RegularContent.hasEqualReg(type).test(args[0]?.show) ? $filter.emptyFilter((args[1][type])) : '';
    };
};

export default IdcDetail;
