import {paginationOperate} from 'utils/publicClass.ts';

interface DataType {
    imageId: Ref<string>;
    imageInfo: Ref<string>;
    deleteImageDiaLog: Ref<boolean>;
};

interface ParamsType {
    imageList: {
        reactiveArr: {
            tableData: {
                length: number;
            }
        };
        getImageList(): Function;
    };
    filter: {
        tableRef: {
            value: {
                clearFilter(): Function;
            }
        };
        reactiveArr: {
            filterParams: {
                architecture: string;
                osType: string;
                source: string;
            }
        }
    }
}

const deleteImageOperate = (imageList: ParamsType['imageList'], filter: ParamsType['filter']) => {
    const data: DataType = {
        // 镜像id
        imageId: ref<string>(''),
        // 镜像信息
        imageInfo: ref<string>(''),
        // 删除镜像弹窗组件
        deleteImageDiaLog: ref<boolean>(false)
    };

    /**
     * 点击删除镜像
     * @param {string} item 当前点击的镜像信息
     * @return {boolean} deleteImageDiaLog.value 删除镜像弹窗
    */
    const deleteImageClick = (item: {imageName: string; imageId: string;}): boolean => {
        data.imageInfo.value = item.imageName;
        data.imageId.value = item.imageId;
        return data.deleteImageDiaLog.value = !data.deleteImageDiaLog.value;
    };

    /**
     * 删除镜像取消操作
     * @param {boolean} type 删除镜像组件状态
    */
    const deleteImageCancel = (type: boolean): boolean => {
        return data.deleteImageDiaLog.value = type;
    };

    /**
     * 删除镜像操作成功后的处理
    */
    const deleteImageSure = () => {
        if (paginationOperate.pageNumber.value > 1 && !imageList.reactiveArr.tableData.length) {
            paginationOperate.pageNumber.value = paginationOperate.pageNumber.value - 1;
            return;
        }
        if (imageList.reactiveArr.tableData.length === 1) {
            filter?.tableRef?.value?.clearFilter();
            filter.reactiveArr.filterParams = {architecture: '', osType: '', source: ''};
        }
        imageList.getImageList();
    };

    const imageDeleteError = () => {
        imageList.getImageList();
    };

    return {
        ...data,
        deleteImageClick,
        deleteImageCancel,
        deleteImageSure,
        imageDeleteError
    }
};

export default deleteImageOperate;

// /**
//  * 删除镜像操作
// */
// class DeleteImageOperate {
//     // 镜像id
//     imageId: Ref<string> = ref<string>('');
//     // 镜像信息
//     imageInfo: Ref<string> = ref<string>('');
//     // 删除镜像弹窗组件
//     deleteImageDiaLog: Ref<boolean> = ref<boolean>(false);
    
//     /**
//      * 点击删除镜像
//      * @param {string} item 当前点击的镜像信息
//      * @return {boolean} deleteImageDiaLog.value 删除镜像弹窗
//     */
//     deleteImageClick = (item: {imageName: string; imageId: string;}): boolean => {
//         this.imageInfo.value = item.imageName;
//         this.imageId.value = item.imageId;
//         return this.deleteImageDiaLog.value = !this.deleteImageDiaLog.value;
//     };

//     /**
//      * 删除镜像取消操作
//      * @param {boolean} type 删除镜像组件状态
//     */
//     deleteImageCancel = (type: boolean): boolean => {
//         return this.deleteImageDiaLog.value = type;
//     };

//     /**
//      * 删除镜像操作成功后的处理
//     */
//     deleteImageSure = () => {
//         if (paginationOperate.pageNumber.value > 1 && !imageList.reactiveArr.tableData.length) {
//             paginationOperate.pageNumber.value = paginationOperate.pageNumber.value - 1;
//             return;
//         }
//         if (imageList.reactiveArr.tableData.length === 1) {
//             filter?.tableRef?.value?.clearFilter();
//             filter.reactiveArr.filterParams = {architecture: '', osType: '', source: ''};
//         }
//         imageList.getImageList();
//     };
// };