import {RefreshType} from '../typeManagement';

const refreshOperate = (modelList: RefreshType['modelList'], reset: RefreshType['reset']) => {

    // 刷新
    const refresh = () => {
        if (!modelList?.reactiveArr?.tableData?.length) {
            reset.reset();
            return;
        }
        modelList.getModelList();
    };

    return {
        refresh
    }
};

export default refreshOperate;
