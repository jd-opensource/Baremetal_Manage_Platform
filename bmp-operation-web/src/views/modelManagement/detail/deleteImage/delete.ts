interface DataType {
    deleteImageDiaLog: Ref<boolean>;
    imageId: Ref<string>;
    imageInfo: Ref<string>;
}

const deleteImageOperate = (tableDetail: any) => {
    const data: DataType = {
        // 删除镜像组件弹窗 默认false
        deleteImageDiaLog: ref<boolean>(false),
        // 镜像id
        imageId: ref<string>(''),
        // 镜像信息，删除镜像时传递的信息
        imageInfo: ref<string>('')
    };

    /**
     * 删除镜像点击事件
     * @return {boolean} deleteImageDiaLog.value 删除镜像组件弹窗状态
    */
    const deleteImageClick = (item: {imageName: string; imageId: string;}): boolean => {
        data.imageInfo.value = item.imageName;
        data.imageId.value = item.imageId;
        return data.deleteImageDiaLog.value = !data.deleteImageDiaLog.value;
    };

    /**
     * 删除镜像弹窗取消事件
     * @param {boolean} type false 弹窗关闭
     * @return {boolean} deleteImageDiaLog.value 删除镜像弹窗关闭
    */
    const deleteImageCancel = (type: boolean): boolean => {
        return data.deleteImageDiaLog.value = type;
    };
    
    const deleteImageSure = () => {
        tableDetail.init();
    };

    return {
        ...data,
        deleteImageClick,
        deleteImageCancel,
        deleteImageSure
    }
};

export default deleteImageOperate;
