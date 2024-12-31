interface ParamsType {
    filter: {
        ossStore: {
            getOssData(): Function;
        };
    };
    resetOperate: {
        reset(): void;
    }
}

const importImageOperate = (filter: ParamsType['filter'], resetOperate: ParamsType['resetOperate']) => {
    const importImageDiaLog: Ref<boolean> = ref<boolean>(false);

    /**
     * 导入镜像点击事件
     * @return {boolean} importImageDiaLog.value 导入镜像弹窗状态
    */
    const importImageClick = (): boolean => {
        return importImageDiaLog.value = true;
    };
    /**
     * 导入镜像组件取消事件
     * @param {boolean} type 弹窗状态
     * @return {boolean} importImageDiaLog.value 导入镜像弹窗
    */
    const importImageCancel = (type: boolean): boolean => {
        return importImageDiaLog.value = type;
    };

    /**
     * 导入镜像成功的回调操作
    */
    const importImageSure = () => {
        resetOperate.reset();
        // filter?.tableRef?.value?.clearFilter();
        // filter.reactiveArr.filterParams = {architecture: '', osType: '', source: ''};
        // imageList.getImageList();
        filter.ossStore?.getOssData();
    };

    return {
        importImageDiaLog,
        importImageClick,
        importImageCancel,
        importImageSure
    }
};

export default importImageOperate;

// class ImportImageOperate {
//     // 导入镜像弹窗组件
//     importImageDiaLog: Ref<boolean> = ref<boolean>(false);
//     // repeatData: string[] = [''];

//     /**
//      * 导入镜像点击事件
//      * @return {boolean} importImageDiaLog.value 导入镜像弹窗状态
//     */
//     importImageClick = (): boolean => {
//         return this.importImageDiaLog.value = true;
//     };
//     /**
//      * 导入镜像组件取消事件
//      * @param {boolean} type 弹窗状态
//      * @return {boolean} importImageDiaLog.value 导入镜像弹窗
//     */
//     importImageCancel = (type: boolean): boolean => {
//         return this.importImageDiaLog.value = type;
//     };

//     /**
//      * 导入镜像成功的回调操作
//     */
//     importImageSure = () => {
//         filter?.tableRef?.value?.clearFilter();
//         filter.reactiveArr.filterParams = {architecture: '', osType: '', source: ''};
//         imageList.getImageList();
//         filter.ossStore?.getOssData();
//     };
// };