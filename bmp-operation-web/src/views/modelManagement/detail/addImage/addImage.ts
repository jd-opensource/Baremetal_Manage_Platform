const addImageOperate = (resetOperate: any) => {
    // 添加镜像弹窗状态-默认false
    const addImageDialog: Ref<boolean> = ref<boolean>(false);

    /**
     * 添加镜像点击事件
     * @return {boolean} addImageDialog.value 添加镜像组件弹窗状态
    */
    const addImageClick = (): boolean => {
        return addImageDialog.value = !addImageDialog.value;
    };

    /**
     * 添加镜像取消事件，关闭弹窗
     * @param {boolean} type 弹窗状态
     * @return {boolean} addImageDialog.value 弹窗状态
    */
    const addImageCancel = (type: boolean): boolean => {
        return addImageDialog.value = type;
    };

    const addImageSure = () => {
        resetOperate();
    };

    return {
        addImageDialog,
        addImageClick,
        addImageCancel,
        addImageSure
    }
};

export default addImageOperate;
