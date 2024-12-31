import {defineStore} from 'pinia'; // 定义容器名
import {imagesAPI} from 'api/image/request.ts';
import {CurrencyType, CurrencyType2, PaginationType} from '@utils/publicType';
import storeName from 'store/storeName.ts'; // 容器名

const imagesStore = defineStore(storeName.imagesData, {
    state: () => {
        return {
            loginUserName: '',
            email: ''
        }
    },
    actions: {
        /**
         * 获取镜像数据
         * @param {Object} params 镜像参数
         * @param {Object} reactiveArr.filterParams 镜像搜索参数
         * @return {Function} resolve, promise.resolve()方法
         * @return {Function} reject promise.reject()方法
        */
        getImage(
            params: PaginationType | PaginationType & CurrencyType,
            reactiveArr: {
                filterParams: CurrencyType;
            }
        ) {
            return imagesAPI(
                {
                    ...params,
                    ...reactiveArr.filterParams
                }
            ).then(({data}: {data: {result: {images: CurrencyType2[]}}}) => {
                if (data?.result?.images?.length) {
                    return Promise.resolve(data.result);
                }
                return Promise.resolve([]);
            })
            .catch(({message}: {message: string;}) => {
                return Promise.reject(message);
            });
        },
    }
});

export default imagesStore;