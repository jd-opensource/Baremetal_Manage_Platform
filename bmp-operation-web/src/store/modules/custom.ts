import {defineStore} from 'pinia'; // 定义容器名
import {customListAPI, surveillanceCustomListAPI} from 'api/public/request.ts';
import {methodsTotal} from 'utils/index.ts';
import {IndexStateType, CheckListArrType, CurrencyStatusType} from '../typeManagement';
import AllStaticData from 'staticData/staticData.ts';
import storeName from 'store/storeName.ts'; // 容器名

const {allCustomListInfo} = AllStaticData;

const customInfoStore = defineStore(storeName.customList, {
    state: (): IndexStateType => {
        return {
            customDataInfo: [
                {
                    pageKey: 'idcList',
                    data: allCustomListInfo.idcInfo
                },
                {
                    pageKey: 'deviceTypeList',
                    data: allCustomListInfo.modelInfo
                },
                {
                    pageKey: 'imageList',
                    data: allCustomListInfo.imageInfo
                },
                {
                    pageKey: 'deviceList',
                    data: allCustomListInfo.deviceInfo
                },
                {
                    pageKey: methodsTotal.humpConversion('deviceAlertLogList'),
                    data: allCustomListInfo.deviceFaultLogInfo
                },
                {
                    pageKey: methodsTotal.humpConversion('deviceStatusList'),
                    data: allCustomListInfo.hardwareInfo
                },
                {
                    pageKey: methodsTotal.humpConversion('alertLogList'),
                    data: allCustomListInfo.faultLogInfo
                },
                {
                    pageKey: methodsTotal.humpConversion('alertRuleList'),
                    data: allCustomListInfo.faultRulesInfo
                },
                {
                    pageKey: 'roleList',
                    data: allCustomListInfo.roleInfo
                },
                {
                    pageKey: 'userList',
                    data: allCustomListInfo.userInfo
                },
                {
                    pageKey: 'messageList',
                    data: allCustomListInfo.messageInfo
                },
                {
                    pageKey: 'auditLogsList',
                    data: allCustomListInfo.auditLogsInfo
                },
                {
                    pageKey: 'monitorRulesList',
                    data: allCustomListInfo.allAlarmRulesInfo
                }
            ]
        }
    },
    actions: {
        /**
         * 自定义列表
         * @param {string} pageKey 不同列表不同key
         * @param {Object} reactiveArr.hasCustomInfo 对应列表信息，用来接收接口返回的状态
         * @param {Array} checkListArr 根据接口返回的信息做数据处理
        */
        customList(
            pageKey: string,
            reactiveArr: {
                hasCustomInfo: Record<string, CurrencyStatusType>;
                checkListArr: CheckListArrType;
            }
        ) {
            let params: {pageKey?: string;};
            let requestAPI: Function;
            const requestParamsAPI = [
                [
                    (key: string) => key.indexOf('_') > -1,
                    () => {
                        params = methodsTotal.toLine({pageKey});
                        requestAPI = surveillanceCustomListAPI;
                    }
                ],
                [
                    (key: string) => !(key.includes('_')),
                    () => {
                        params = {pageKey};
                        requestAPI = customListAPI;
                    }
                ]
            ];
            for (const key of requestParamsAPI) {
                if (key[0](pageKey)) {
                    key[1](pageKey);
                    break;
                }
            }
            nextTick(() => {
                requestAPI(
                    {
                        ...params
                    }
                ).then(({data} : {data: {result: {customInfo: Record<string, CurrencyStatusType>;}}}) => {
                    const result = methodsTotal.lineConverting(data?.result);
                    if (result?.customInfo) {
                        // 列表展示数据
                        const checkListArr: CheckListArrType = reactiveArr.checkListArr;
                        // 遍历列表展示数据
                        for (const key of checkListArr) {
                            // 遍历自定义信息数据
                            for (const index in reactiveArr.hasCustomInfo) {
                                // 点击的当前项等于对应数据的那一项
                                if (index === key.label) {
                                    // 状态赋值
                                    key.disabled = reactiveArr.hasCustomInfo[index].required;
                                    key.selected = reactiveArr.hasCustomInfo[index].selected;
                                }
                                else {
                                    reactiveArr.hasCustomInfo = result.customInfo;
                                }
                            }
                        }
                        return;
                    }
                    return Promise.reject();
                })
                .catch(() => {
                    reactiveArr.hasCustomInfo = this.customDataInfo?.filter((item: {pageKey: string}) => item.pageKey === pageKey)[0]?.data??[];
                });
            })
        }
    }
});

export default customInfoStore;