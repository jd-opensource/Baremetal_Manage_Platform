import {methodsTotal} from 'utils/index.ts';
import {CustomOperate} from 'utils/publicClass.ts';
import AllStaticData from 'staticData/staticData.ts'; // 自定义列表信息 自定义列表数据

const {deviceFaultLogInfo} = AllStaticData.allCustomListInfo;
const {deviceFaultLogList} = AllStaticData.allCustomList;

const name: string = methodsTotal.humpConversion('deviceAlertLogList');

const custom = {
    name,
    CustomOperate,
    deviceFaultLogInfo,
    deviceFaultLogList
}

export default custom;
