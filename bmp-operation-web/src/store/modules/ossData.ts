import {defineStore} from 'pinia'; // 定义容器名
import {getOssAPI} from 'api/public/request.ts'; // oss接口
import {modelListAPI} from 'api/model/request.ts';
import messageApiCount from 'api/message/request.ts';
import {filterData} from 'utils/index.ts'; // 过滤重复数据
import {CurrencyType} from '@utils/publicType';
import {OsStoreType, OssStateType} from '../typeManagement';
import DeviceStaticData from 'staticData/device/index.ts'; // 机型类型、管理状态
import ModelStaticData from 'staticData/model/index.ts';
import ImageStaticData from 'staticData/image/index.ts';
import MessageStaticData from 'staticData/message/index.ts';
import AllAlarmRulesStatic from 'staticData/inBandMonitoring/allAlarmRules/index.ts'
import storeName from 'store/storeName.ts'; // 容器名
import {auditLogsTypesAPI} from 'api/device/request.ts';
import { pinyin } from 'pinyin-pro';

const ossDataStore = defineStore(storeName.ossData, {
    state: (): OssStateType => {
        return {
            osType: [],
            architecture: [],
            messageType: [], 
            customFilterList: [],
            source: ImageStaticData.source,
            manageStatus: DeviceStaticData.manageMentStatus,
            deviceSeries: ModelStaticData.modelTypeData,
            collectStatus: DeviceStaticData.collectStatusData,
            operation: [],
            deviceTypeId: [],
            result: DeviceStaticData.operationalFeedbackData,
            status: AllAlarmRulesStatic.allAlarmRulesStatusData
        }
    },

    actions: {

        getMessagesType() {
            messageApiCount.messageTypesAPI({})
            .then(({data}: {data: {result: {[x: string]: string}}}) => {
                // @ts-ignore
                const val: string[] = Object.keys(data.result).flat(1);
                this.messageType = val?.length ? val.map((item, index) => {
                    return {
                        text: item,
                        filterParams: item,
                        value: index + 1
                    }
                }) : MessageStaticData.messageType;
                const newData = Object.values(data.result).flat(1);
                this.customFilterList = newData.map(item => {
                    return {
                        label: item,
                        yin: pinyin(item, {toneType: 'none'}).trim().replace(/\s/g, ""),
                        value: item,
                        checkboxStatus: false
                    }
                });
            })
        },
    
        getModelList() {
            modelListAPI({
                isAll: '1'
            })
            .then(({data}: any) => {
                if (data?.result?.deviceTypes?.length) {
                    this.deviceTypeId = data.result.deviceTypes.map((item: {name: string; deviceTypeId: string;}, index: number) => {
                        return {
                            text: item.name,
                            filterParams: item.deviceTypeId,
                            value: index + 1
                        }
                    })
                    return;
                }
                return Promise.reject();
            })
            .catch(() => {
                this.deviceTypeId = [];
            });
        },

        getModelListRes() {
            return new Promise((resolve, reject) => {
                return modelListAPI({isAll: '1'})
                .then(({data}: any) => {
                    if (data?.result?.deviceTypes?.length) {
                        const newData = data.result.deviceTypes.map((item: {name: string; deviceTypeId: string;}, index: number) => {
                            return {
                                text: item.name,
                                filterParams: item.deviceTypeId,
                                value: index + 1
                            }
                        });
                        return resolve(newData);
                    }
                    return Promise.reject();
                })
                .catch(() => {
                    return reject([]);
                })
            })
        },

        getAuditLogsTypes() {
            // return new Promise((resolve, reject) => {
            return auditLogsTypesAPI({})
                .then(({data}: any) => {
                    if (data?.result?.length) {
                        this.operation = data.result.map((item: {name: string; operation: string;}, index: number) => {
                            return {
                                filterParams: item.operation,
                                text: item.name,
                                value: index + 1
                            }
                        })
                        return;
                        // return resolve(data.result);
                    }
                    return Promise.reject();
                })
                .catch(() => {
                    this.operation = []
                    // return reject([]);
                })
            // })
        },

        /**
         * 获取oss数据
         * @return {Function} xxx 返回一个函数
        */
        getOssData(param = {}) {
            const params = Object.keys(param).length ? param : {
                isAll: '1'
            };
            return getOssAPI (
                    {
                        ...params
                    }
                )
                .then(({data} : {data: {result: {oss: OsStoreType}}}) => {
                    if (data?.result?.oss) {
                        this.osType = this.processingData(data, 'osType');
                        this.architecture = this.processingData(data, 'architecture');
                        return;
                    }
                    return Promise.reject();
                })
                .catch(() => {
                    this.osType = [];
                    this.architecture = [];
                });
        },

        /**
         * 处理数据
         * @param {Array} oss oss 数据
         * @param {string} type filter的参数名
        */
        processingData(data: {result: {oss: OsStoreType}}, type: string) {
            return filterData(data.result.oss, type).map((item: CurrencyType, index: number) => {
                return {
                    ...item,
                    text: item[type],
                    value: index + 1,
                    filterParams: item[type]
                };
            });
        }
    }
});

export default ossDataStore;
