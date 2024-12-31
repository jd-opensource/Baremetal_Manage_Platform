/**
 * @file
 * @author
*/
import {language} from 'utils/publicClass.ts';
import {disksDetailAPI} from 'api/device/request.ts';
import store from 'store/index.ts';
import DeviceStaticData from 'staticData/device/index.ts';
import {msgTip} from 'utils/index.ts';

class DiskTable {
    isLoading: Ref<boolean> = ref<boolean>(true);
    reactiveArr = reactive({
        disksData: [],
        tableData: [],
        volumesData: []
    });
    query: {
        deviceId: string
    }
    
    constructor(query: {deviceId: string}) {
        this.query = query;
        // this.init()
    }

    init = () => {
        this.isLoading.value = true;
        this.getDiskDetail();
    }

    getDiskDetail = () => {
        disksDetailAPI({
            deviceId: this.query.deviceId
        })
        .then(({data}: any) => {
            if (data?.result && Object.keys(data?.result).length) {
                const {disks, panfu, volumes} = data.result;
                volumes.forEach((item: any, index: number) => {
                    DeviceStaticData.modelInfoTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    })
                });
                this.reactiveArr.volumesData = volumes.map((item: any) => {
                    return {
                        ...item,
                        newDisksVal: this.setDisks(item.disks)
                    }
                });
                this.reactiveArr.disksData = disks;
                this.reactiveArr.tableData = panfu;
            }
        })
        .catch(({message}: {message: string}) => {
            if (!store.deviceDetailErrorTipInfo.errorStatus) {
                message && msgTip.error(message);
            }
        })
        .finally(() => {
            this.isLoading.value = false;
        })
    }

    setDisks = (item: {name: string; size: string; sizeUnit: string}[]) => {
        return item.map((t: {name: string; size: string; sizeUnit: string}) => t.name + ': ' + t.size + t.sizeUnit).join('; ');
    }

    setRaids = (item: {name: string; raidId: string}[]) => {
        return item.map((t: {name: string}) => t.name).join(',');
    }

    setDiskInterfaceType = (item: string) => {
        return item === 'notLimited' ? language.t('modelForm.unrestricted') : item;
    }
}

export default DiskTable;
