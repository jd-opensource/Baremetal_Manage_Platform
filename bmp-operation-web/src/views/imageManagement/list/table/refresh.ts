
type ImageLiatType = {
    getImageList(): Function;
    reactiveArr: {
        tableData: {
            length: number;
        };
    }
};

type ResetOperateType = {
    reset(): Function;
};

interface ParamsType {
    imageList: ImageLiatType;
    resetOperate: ResetOperateType;
};

const refreshOperate = (imageList: ParamsType['imageList'], resetOperate: ParamsType['resetOperate']) => {
    const refresh = () => {
        if (!imageList?.reactiveArr?.tableData?.length) {
            resetOperate.reset();
            return;
        }
        imageList.getImageList();
    };

    return {
        refresh
    };
};

export default refreshOperate;
