/**
 * @file
 * @author
*/

import store from 'store/index.ts';
import {ParamsType} from '../typeManagement';
import {CurrencyType} from '@utils/publicType';
import {msgTip, methodsTotal} from 'utils/index.ts';
import {paginationOperate} from 'utils/publicClass.ts';
import ImageStaticData from 'staticData/image/index.ts';

// 镜像类型
type ImagesType = {[x: string]: string | boolean;};

// 镜像返回类型
type ImageResponseType = {images: {[x: string]: string}[], totalCount: number;};

class ImageList {
    systemPartitionInfo: any = store.sysPartitionInfo;
    searchTip: Ref<boolean> = ref<boolean>(false);
    // loading态
    isLoading: Ref<boolean> = ref<boolean>(true);
    type: string = '';
    search: any = {};
    filter: any = {};
    imagesInfo: { 
        getImage(
            arg0: ParamsType | {isAll: string},
            arg1: {tableData: {[x: string]: string | boolean;}[]}
        ): Promise<ImageResponseType>;
        // systemPartitionData(arg0: string, arg1: boolean): never;
    } = store.imagesInfo;
    // status: any = {
    //     value: false
    // };

    // 复杂响应式数据集合
    reactiveArr: {
        tableData: {
            [x: string]: string | boolean;}[];
        } = reactive<{
            tableData: ImagesType[]
        }>({
            tableData: [] // 表格数据-机房列表
    });

    // 构造器
    constructor (search: any, filter: any) {
        this.search = search;
        this.filter = filter;
        this.getImageList();
    };

    /**
     * 获取镜像列表数据
     * @param {Object} params 请求参数，默认{}
    */
    getImageList = () => {
        this.isLoading.value = true;
        const param: ParamsType = {
            pageNumber: paginationOperate.pageNumber.value,
            pageSize: paginationOperate.pageSize.value,
            ...this.filter.reactiveArr.filterParams,
            ...this.search.reactiveArr.searchParams
        };
        store.filterEmpty.deleteEmtpyData(param);
        this.imagesInfo.getImage({...param}, this.reactiveArr).then(({images, totalCount}: ImageResponseType) => {
            if (totalCount > 0) {
                this.reactiveArr.tableData = images.map((item: CurrencyType, index: number) => {
                    ImageStaticData.imageListTipData.forEach((t: string) => {
                        Object.assign(item, {[`${t}${index}`]: {showTooltip: false}})
                    });
                    const val = JSON.parse((item?.systemPartition?.length ? item?.systemPartition : '[]'));
                    return {
                        ...item,
                        newSystemPartition: val
                    }
                });
                paginationOperate.total.value = totalCount;
                return;
            }
            return Promise.reject('');
        })
        .catch((message: string) => {
            this.reactiveArr.tableData = [];
            paginationOperate.total.value = 0;
            methodsTotal.initScrollLeft(this.filter.tableRef.value);
            message && msgTip.error(message);
        })
        .finally(() => {
            this.isLoading.value = false;
            this.search.hasClear.value = false;
            this.searchTip.value = Object.keys(param).length > 2;
        });
    };
};

export default ImageList;
