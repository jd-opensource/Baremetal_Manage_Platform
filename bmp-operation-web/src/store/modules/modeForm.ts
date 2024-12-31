import {defineStore} from 'pinia'; // 定义容器名
// import {filterData} from 'utils/index.ts'; // 过滤重复数据
import {language} from 'utils/publicClass.ts';
import {getFilterListAPI} from 'api/model/request.ts'; // raids接口 filter接口
import {CpuDataType, MemDataType} from '@utils/publicType';
import {StateType, FormSubmitType, SetCpuMemFormSubmitType} from '../typeManagement'; // state类型 表单类型
import storeName from 'store/storeName.ts'; // 容器名

/**
 * 表单验证
*/
class FormCheck {
    /**
     * 输入项/选择项为空时
     * @param {string} str 提示项
     * @param {string} trigger 判断时机-change/blur
    */
    inputSelectEmtpy = (str = '', trigger = 'change', required = true): FormSubmitType[] => {
        return [{required, message: str, trigger}];
    };
};
const formCheck: FormCheck = new FormCheck();

// defineStore定义容器
// 参数1：是对仓库的命名，名称必须具备唯一性；
// 参数2：配置的选项对象，即state、getters、actions
// 其中state的写法必须是函数，为了避免在服务端交叉请求导致的状态数据污染，而且必须是箭头函数，为了更好的TS类型推导
const modeFormStore = defineStore(storeName.modeForm, {
    state: (): StateType => {
        return {
            raidsData: [], // raids数据
            data: {
                radioBtn: [ // 机型类型
                    language.t('modelForm.radioBtn.computed'),
                    language.t('modelForm.radioBtn.storage'),
                    language.t('modelForm.radioBtn.GPU'),
                    language.t('modelForm.radioBtn.other')
                ],
                modelCPUBtn: [ // CPU规格
                    language.t('modelForm.specificationsType.presetSpecifications'),
                    language.t('modelForm.specificationsType.otherSpecifications')
                ],
                storageData: [ // 内存规格
                    language.t('modelForm.specificationsType.presetSpecifications'),
                    language.t('modelForm.specificationsType.otherSpecifications')
                ]
            },
            ruleForm: {
                bootMode: [],
                raidCan: '',
                templateName: '',
                deviceTypeName: '',
                name: '', // 机型名称
                architecture: 'x86_64', // 体系架构
                deviceType: '', // 机型规格
                description: '', // 描述
                modelCPU: language.t('modelForm.specificationsType.presetSpecifications'), // 机型CPU
                cpuInfo: '', // CPU
                cpuManufacturer: '', // 处理器厂商
                cpuModel: '', // 型号
                cpuCores: 1, // 物理核数
                cpuFrequency: '', // 主频
                cpuAmount: 1, // cpu-数量
                modelStorage: language.t('modelForm.specificationsType.presetSpecifications'), // 机型内存
                memInfo: '', // 内存
                memType: '', // 接口
                memFrequency: '', // 主频MHz
                memSize: '', // 容量(GB)
                memAmount: 1, // 数量
                volumeManager: [],
                nicRate: '', // 网卡速度(GE)
                nicAmount: 1, // 网口数量(个)
                interfaceMode: '', // 网络设置
                gpuManufacturer: '', // GPU(品牌)
                gpuModel: '', // GPU(型号)
                gpuAmount: 0, // GPU数量(个)
                height: '', // 高度(U)
                cpuData: [], // cpu数据
                memData: [] // 内存数据
            },
            rules: {
                architecture: formCheck.inputSelectEmtpy(), // 体系架构
                cpuInfo: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.cpu')), // CPU
                cpuManufacturer: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.processorManufacturer')), // 处理器厂商
                memInfo: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.storage')), // 内存
                memType: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.interface')), // 接口
                memFrequency: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.dominantFrequencyMHz')), // 主频(MHz)
                memSize: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.capacity')), // 容量(GB)
                nicRate: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.networkSpeed')), // 网卡速度(GE)
                interfaceMode: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.networkSettings')), // 网络设置
                gpuManufacturer: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.gpuBrand'), 'change', false), // GPU(品牌)
                gpuModel: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.gpuModel'), 'change', false), // GPU(型号)
                height: formCheck.inputSelectEmtpy(language.t('modelForm.emptyTip.heightU')) // 高度(U)
            }
        }
    },
    // 类似computed
    getters: {
        // 获取raids
        getRaids (state) {
            return state.raidsData;
        }
    },
    actions: {
    
        initModelType (deviceSeries: string) {
            const modelType: any = {
                'computer': () => language.t('modelForm.radioBtn.computed'),
                'storage': () => language.t('modelForm.radioBtn.storage'),
                'gpu': () => language.t('modelForm.radioBtn.GPU'),
                'other': () => language.t('modelForm.radioBtn.other')
            };
            if (modelType[deviceSeries]) {
                return modelType[deviceSeries]();
            }
            return modelType['computer']();
        },

        /**
         * 获取预置数据
         * @return {Function} Promise.resolve(data.result); resolve成功的返回
        */
        getCommonData () {
            return getFilterListAPI()
                .then(({data} : {data: {result: {cpuSpec: CpuDataType[]; memSpec: MemDataType[]; bootMode: string[]}}}) => {
                    if (data?.result?.cpuSpec) {
                        this.ruleForm.cpuData = data.result.cpuSpec;
                        this.ruleForm.memData = data.result.memSpec;
                        this.ruleForm.bootMode = data.result.bootMode;
                        return Promise.resolve(data.result);
                    }
                    // return Promise.reject('')
                })
                // .catch(({message}: {message: string;}) => {
                //     message && msgTip.error(message);
                // });
        },

        /**
         * 设置机型类型
        */
        setModelType (type: string): string {
            const obj = new Map([
                [language.t('modelForm.radioBtn.computed'), 'computer'],
                [language.t('modelForm.radioBtn.storage'), 'storage'],
                [language.t('modelForm.radioBtn.GPU'), 'gpu'],
            ]);
            return obj.get(type)?? 'other';
        },

        setCpuMemData (cpuBol: boolean, memBol: boolean, formSubmit: SetCpuMemFormSubmitType) {
            const memData: string[] = ['memType', 'memFrequency', 'memSize', 'memAmount'];
            const cpuData: string[] = ['cpuAmount', 'cpuCores', 'cpuFrequency', 'cpuManufacturer', 'cpuModel'];
            const cpuMemData = [
                [
                    (status1: boolean, _: boolean) => !(status1),
                    () => {
                        const val = formSubmit?.cpuData?.find((item: {label: string}) => item.label === formSubmit.cpuInfo)?.info || [];
                        for(const key of cpuData) {
                            (formSubmit[`${key}` as keyof typeof formSubmit]) = (val[`${key}` as keyof typeof val]);
                        }
                    }
                ],
                [
                    (_: boolean, status2: boolean) => !(status2),
                    () => {
                        const mem = formSubmit?.memData?.find((item: {label: string}) => item.label === formSubmit.memInfo)?.info || [];
                        for(const key of memData) {
                            (formSubmit[`${key}` as keyof typeof formSubmit]) = (mem[`${key}` as keyof typeof mem]);
                        }
                    }
                ]
            ];
            for (const key of cpuMemData) {
                if (key[0](cpuBol, memBol)) {
                    key[1](cpuBol, memBol);
                }
            }
            return Promise.resolve();
        }
    }
});

export default modeFormStore;
