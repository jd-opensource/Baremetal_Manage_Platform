import {ReactiveArrType} from '../typeManagement';
import {imageOssAPI} from 'api/image/request.ts';
import {importImageAPI} from 'api/image/request.ts';
import {CurrencyType, CurrencyType5} from '@utils/publicType';
import {msgTip} from 'utils/index.ts';
import MethodsTotal from 'utils/index.ts';
import ImageStaticData from 'staticData/image/index.ts';
import {language} from 'utils/publicClass.ts';

class ImportImageOperate {
    hasClick: Ref<boolean> = ref<boolean>(false);
    imageFormatClick: Ref<boolean> = ref<boolean>(false);
    sysDefault: Ref<string> = ref<string>('');
    hasBootModeFlag: Ref<boolean> = ref<boolean>(false);
    // loading 加载态
    isLoading: Ref<boolean> = ref<boolean>(false);
    init: number = 1;
    // 复杂数据类型
    reactiveArr: ReactiveArrType = reactive<ReactiveArrType>({
        defaultBootMode: [],
        fileParams: {},
        architectureData: [],
        totalData: {},
        architecture: [], // 体系架构
        osType: [], // 操作系统平台
        osVersion: [], // 操作系统版本
        bootModeData: [],
        imageFormatData: []
    });
    formRulesOperate: any;
    rulesVerification: any;
    chooseImageOperate: any;
    emitValue: any;

    constructor (formRulesOperate: any, rulesVerification: any, chooseImageOperate: any, emitValue: any) {
        this.#getImageOss();
        this.formRulesOperate = formRulesOperate;
        this.rulesVerification = rulesVerification;
        this.chooseImageOperate = chooseImageOperate;
        this.emitValue = emitValue;
    };

    #getImageOss = () => {
        imageOssAPI({})
        .then(({data} : {data: {result: {architecture: CurrencyType;}}}) => {
            if (data?.result?.architecture) {
                this.reactiveArr.totalData = data.result.architecture;
                this.reactiveArr.defaultBootMode = (this.reactiveArr.totalData.bootMode as unknown as string[]);
                this.reactiveArr.imageFormatData = (this.reactiveArr.totalData.format as unknown as string[]);
                this.#processData('architecture', data.result.architecture, Object.keys);
                if (this.init === 1) {
                    this.init ++;
                    this.#setInitData();
                }
            }
        });
    };

    #setInitData = () => {
        const osType = (this.reactiveArr.totalData['x86_64'][`${['osType']}` as keyof typeof this.reactiveArr.totalData['x86_64']] as never);
        this.reactiveArr.architectureData = osType;
        this.formRulesOperate.ruleForm.architecture = 'x86_64';
        this.formRulesOperate.ruleForm.osType = Object.keys(osType)[0];
        this.formRulesOperate.ruleForm.version = osType['CentOS'][0];
        this.#processData('osType', osType, Object.keys);
        this.#processData('osVersion', osType['CentOS'], Object.values);
    };

    imageFormatChange = () => {
        this.reactiveArr.bootModeData = this.reactiveArr.defaultBootMode;
        if (this.hasBootModeFlag.value) {
            this.hasBootModeFlag.value = false;
        }
    }

    bootModeChange = (val: string | string[]) => {
        this.rulesVerification.rules.bootMode[0].trigger = 'change';
        this.hasBootModeFlag.value = !(val?.length);
    };

    architectureChange = (architecture: string & number) => {
        this.hasClick.value = true;
        this.formRulesOperate.ruleForm.osType = '';
        this.formRulesOperate.ruleForm.version = '';
        const osType = (this.reactiveArr.totalData[architecture][`${['osType']}` as keyof typeof this.reactiveArr.totalData['architecture']] as never);
        this.reactiveArr.architectureData = osType;
        this.#processData('osType', osType, Object.keys);
    };

    osTypeChange = (osType: number & string) => {
        this.hasClick.value = true;
        this.formRulesOperate.ruleForm.version = '';
        const newData = this.reactiveArr.architectureData[osType];
        this.#processData('osVersion', newData, Object.values);
    };

    #processData = (type: string, newData: string | CurrencyType, obj: Function) => {
        (this.reactiveArr[`${type}` as keyof typeof this.reactiveArr] as string) = obj(newData).map((item: string, index: number): CurrencyType5 => {
            return {
                value: index + 1,
                label: item
            }
        });
    };

    /**
     * 确定按钮弹窗
    */
    determineClick = async (): Promise<void> => {
        this.isLoading.value = true;
        // 判断输入项是否符合条件
        await nextTick(() => {
            this.formRulesOperate.ruleFormRef.value!.validate(async (valid: boolean) => {
                // 输入符合条件后才可以继续后续操作
                if (valid && this.reactiveArr?.fileParams?.result) {
                    this.#requestImportImage(await this.#setParams());
                    return;
                }
                this.isLoading.value = false;
                this.hasClick.value = true;
                this.imageFormatClick.value = true;
                this.hasBootModeFlag.value = !this.formRulesOperate.ruleForm.bootMode?.length
                this.rulesVerification.customOSVersion.value = true;
                this.#fileErrorStatus(valid);
            });
        });
    };

    #setParams = async () => {
        const params: CurrencyType = MethodsTotal.deepCopy(this.formRulesOperate.ruleForm);
        params.systemPartition = JSON.stringify(params.systemPartition);
        if (ImageStaticData.otherSystemData.includes(params.osType)) {
            // params.systemPartition = '';
            if (!['Windows'].includes(params.osType)) {
                params.osType = params.customOperatePlatform
                params.version = params.customVersion;
            }
        }
        else {
            // params.systemPartition = JSON.stringify(params.systemPartition);
        }
        return params;
    };

    #fileErrorStatus = (valid: boolean) => {
        if (!this.formRulesOperate.ruleForm.imageFile?.length) {
            this.chooseImageOperate.hasError.value = true;
            this.chooseImageOperate.chooseImageFlag.value = true;
            return;
        }
        if (!valid) return;
        this.chooseImageOperate.hasError.value = false;
        msgTip.warning(language.t('importImage.errTip.upload'));
    };

    /**
     * 请求导入镜像接口，成功后把事件回传，关闭弹窗
    */
    #requestImportImage = (params: CurrencyType): void => {
        let bootMode = params.imageFormat === 'tar' ? (params.bootMode as any).join(',') : params.bootMode;
        params.bootMode = bootMode;
        importImageAPI({
            ...params,
            ...this.reactiveArr.fileParams.result
        })
        .then(({data} : {data: {result: {imageId: string;}}}) => {
            if (data?.result?.imageId) {
                msgTip.success(language.t('operate.success'));
                this.emitValue('determine-click');
                this.cancelClick();
            }
        })
        .finally(() => {
            this.isLoading.value = false;
        });
    };

    /**
     * 取消按钮点击操作
    */
    cancelClick = (): void => {
        // 关闭蒙层
        this.diaLogClose();
    };

    /**
     * defineEmits API 来替代 emits，子传父，事件回传
    */
    diaLogClose = (): void => {
        this.chooseImageOperate.uploads.value!.abort();
        this.chooseImageOperate.uploads.value!.clearFiles();
        this.emitValue('dia-log-close', false);
    };
};

export default ImportImageOperate;
