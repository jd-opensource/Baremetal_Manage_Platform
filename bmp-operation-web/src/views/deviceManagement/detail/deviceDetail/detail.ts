/**
 * @file
 * @author
*/

import {TimerOperate, CurrentInstance, language} from 'utils/publicClass.ts';
import {DetailType, TipStatusType, ShowPasswordType} from '../typeManagement';
import {CurrencyType} from '@utils/publicType';
import {msgTip} from 'utils/index.ts';
import DeviceStaticData from 'staticData/device/index.ts'; // 中间态1 中间态2
import {devicesDetailAPI, desrcibeAgentStatusAPI} from 'api/device/request.ts';
import router from 'router/index.ts';
import AllStaticData from 'staticData/staticData.ts';
import store from 'store/index.ts';

class DeviceDetail extends TimerOperate {
    isIntervalRequest: boolean = false;
    isLoading: Ref<boolean> = ref<boolean>(true);
    password: Ref<string> = ref<string>('');
    screenWidth: Ref<number> = ref<number>(0);
    statusName: Ref<string> = ref<string>('');
    reactiveArr: DetailType & TipStatusType & ShowPasswordType = reactive<DetailType & TipStatusType & ShowPasswordType>({
        detail: {}, // 详情数据
        internetPort1: false, // 网口1交换机登录密码
        internetPort2: false, // 网口2交换机登录密码
        outLoginPassword: false, // 带外登录密码
        info: '', // 信息，用来显示对应明文显
        tooltipStatus: {
            'svInfo': {
                showTooltip: false
            },
            'dvInfo': {
                showTooltip: false
            },
            'switchPort1': {
                showTooltip: false
            },
            'switchPort2': {
                showTooltip: false
            }
        }
    });
    verification: any;
    proxy = getCurrentInstance()!.proxy;
    instanceProxy = new CurrentInstance().proxy;

    constructor (verification: any) {
        super();
        onUnmounted(() => clearTimeout((this.timer as number)));
        onBeforeUnmount(() => clearTimeout((this.timer as number)));
        this.verification = verification;
        this.screenWidth.value = window.screen.width;
        // this.getDeviceDetail();
    };

    routerJump = () => {
        // if (router.options.history.state.back !== null) {
        //     router.go(-1);
        //     return;
        // }
        const path: string = this.instanceProxy.$defInfo.routerPath('deviceList');
        router.push(path);
    };

    initData = () => {
        this.isLoading.value = true;
        this.getDeviceDetail();
    };

    setCollected = (status: string) => {
        const type = new Map([
            ['1', language.t('deviceDetail.status.collected')],
            ['2', language.t('deviceDetail.status.noCollected')],
            ['3', language.t('deviceDetail.status.collecting')],
            ['4', language.t('deviceDetail.status.collectionFailed')]
        ]);
        return type.get(status);
    }

    getDeviceDetail = () => {
        clearTimeout((this.timer as number));
        const {value}: {value: string;} = this.password;
        const params: {show?: string;} = value?.length ? {show: value} : {};
        devicesDetailAPI(
            {
                deviceId: router?.currentRoute?.value?.query?.deviceId,
                ...params
            }
        )
        .then(({data} : {data: {result: CurrencyType;}}) => {
            if (data?.result && Object.keys(data?.result)?.length) {
                this.reactiveArr.detail = data.result;
                if (params?.show) {
                    this.getPassword(params.show);
                    this.setPasswordStatus(params.show);
                    // new ShowPassword(this, this.verification).setPasswordStatus(params.show);
                }
                const hasLoopReq: boolean = [...DeviceStaticData.intermediate1, ...DeviceStaticData.intermediate2].some((item: string) => item === this.reactiveArr.detail.instanceStatus);
                if (['putawaying'].some((item: string) => item === this.reactiveArr.detail.manageStatus) || hasLoopReq) {
                    this.timer = setTimeout(() => {
                        this.isLoading.value = false;
                        this.getDeviceDetail();
                        this.isIntervalRequest = true;
                    }, 5000)
                    return;
                }
                this.isIntervalRequest = false;
                return;
            }
            return Promise.reject();
        }).catch(({message} : {message: string;}) => {
            const path: string = this.instanceProxy.$defInfo.routerPath('deviceList');
            if (AllStaticData.noFoundData.includes(message)) {
                router.push(path);
                return;
            }
            if (!this.isIntervalRequest) {
                typeof message === 'string' && msgTip.error(message);
                store.deviceDetailErrorTipInfo.errorStatus = Object.is(typeof message, 'string');
                this.reactiveArr.detail = {};
                return;
            }
            this.timer = setTimeout(() => {
                this.isLoading.value = false;
                this.getDeviceDetail();
            }, 5000)
        })
        .finally(() => {
            this.isLoading.value = false;
        });
    };

    desrcibeAgentStatus = () => {
        desrcibeAgentStatusAPI({
            instanceId: router?.currentRoute?.value?.query?.instanceId
        })
        .then(({data}: {data: {result: {agentStatus: {statusName: string; length: string; map(arg0: (item: { statusName: string; }) => string): {join(arg0: string): void}}}}}) => {
            if (data?.result?.agentStatus?.length) {
                // @ts-ignore
                this.statusName.value = data.result.agentStatus.map((item: {statusName: string}) => item.statusName).join(',');
                return;
            }
            return Promise.reject('');
        })
        .catch(() => {
            this.statusName.value = '';
        })
    }

    setPasswordStatus = (show: string) => {
        this.reactiveArr.outLoginPassword = show === 'iloPassword';
        this.reactiveArr.internetPort2 = show === 'switchPassword2';
        this.reactiveArr.internetPort1 = show === 'switchPassword1';
    };

    /**
     * 是否明文显示密码
     * @param {Object} item 当前点击的这一项
     * @param {string} type 当前点击的状态是 open or close
    */
    hasPasswordClick = (bol: never | boolean, type: string, status: string, dataType: string = ''): void => {
        clearTimeout((this.timer as number));
        this.password.value = dataType;
        if (type === 'close') {
            (this.reactiveArr[`${status}` as keyof typeof this.reactiveArr] as boolean) = bol;
            this.password.value = '';
            this.reactiveArr.detail[dataType] = '********';
            return;
        }
        this.verification.securityVerificationDiaLog.value = !this.verification.securityVerificationDiaLog.value;
        this.reactiveArr.info = status;
    };

    getPassword = (show: string) => {
        const {switchPassword1, switchPassword2, iloPassword} = this.reactiveArr.detail;
        const empty = this.proxy!.$filter.emptyFilter;
        const params: Map<string, any> = new Map([
            ['switchPassword1', empty(switchPassword1)],
            ['switchPassword2', empty(switchPassword2)],
            ['iloPassword', empty(iloPassword)]
        ]);
        this.reactiveArr.detail[show] = params.get(show);
    };
};

export default DeviceDetail;
