/**
 * @file
 * @author
*/
import ModelStaticData from 'staticData/model/index.ts'

class PageScrollOpt {

    pageScroll = (dataCount: {ruleForm: {modelCPU: string; modelStorage: string;}}) => {
        let defaultType = ['name', 'deviceType'];
        let otherType = ['cpuInfo', 'memInfo'];
        let defaultType2 = ['nicRate', 'nicAmount', 'interfaceMode', 'height', 'volumeManager'];
        let type: string[] = [...defaultType, ...otherType];
        const cpuStatus = ModelStaticData.modelSpecData.includes(dataCount.ruleForm.modelCPU);
        const memStatus = ModelStaticData.modelSpecData.includes(dataCount.ruleForm.modelStorage);
        const scrollArr = [
            [
                (cpuStatus: boolean, memStatus: boolean) => (cpuStatus && memStatus),
                () => {
                    type = [...type, ...defaultType2];
                    this.scrollOpt(type, dataCount)
                }
            ],
            [
                (cpuStatus: boolean, memStatus: boolean) => (!cpuStatus && memStatus),
                () => {
                    otherType = ['cpuManufacturer', 'cpuModel', 'cpuCores', 'cpuFrequency', 'cpuAmount'];
                    type = [...defaultType, ...otherType, 'memInfo', ...defaultType2];
                    this.scrollOpt(type, dataCount);
                }
            ],
            [
                (cpuStatus: boolean, memStatus: boolean) => (cpuStatus && !memStatus),
                () => {
                    otherType = ['memType', 'memFrequency', 'memSize', 'memAmount'];
                    type = [...defaultType, 'cpuInfo', ...otherType, ...defaultType2];
                    this.scrollOpt(type, dataCount);
                }
            ],
            [
                (cpuStatus: boolean, memStatus: boolean) => (!cpuStatus && !memStatus),
                () => {
                    otherType = ['cpuManufacturer', 'cpuModel', 'cpuCores', 'cpuFrequency', 'cpuAmount', 'memType', 'memFrequency', 'memSize', 'memAmount'];
                    type = [...defaultType, ...otherType, ...defaultType2];
                    this.scrollOpt(type, dataCount);
                }
            ]
        ];
        for (const key of scrollArr) {
            if (key[0](cpuStatus, memStatus)) {
                key[1](cpuStatus, memStatus);
                break;
            }
        }
    };

    scrollOpt = (type: string[], dataCount: { ruleForm: {[x: string]: string}; }) => {
        for (const key of type) {
            // @ts-ignore
            if (!dataCount.ruleForm[key]) {
                (document as any).getElementById(key).scrollIntoView({
                    block: 'start',
                    behavior: 'smooth' // 代表平滑滚动
                });
                break;
            }
        }
    };
}

export default PageScrollOpt;
