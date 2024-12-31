import {defineStore} from 'pinia'; // 定义容器名
import {deepCopy} from 'utils/index.ts';
import storeName from 'store/storeName.ts'; // 容器名
import {SystemPartitionType} from '@utils/publicType';
import {language} from 'utils/publicClass.ts';

const sysPartitionStore = defineStore(storeName.sysPartition, {
    state: () => {
        return {}
    },

    actions: {
        /**
         * 自定义分区模块数据展示
         * @param {Array} data 需要处理的数据
         * @return {Array} defaultPartitions 处理完成的数据
        */
        systemPartitionData (data: string, hasLabel: boolean = true) {
            // 声明一个新数组
            let defaultPartitions: {size: number; format: string; point: string;}[] = [];
            // try - catch捕获异常
            try {
                // 正常情况下，深拷贝数据，把数据变为可遍历的
                defaultPartitions = deepCopy(JSON.parse(data));
                // 遍历数据
                defaultPartitions.map((item: {size: number;}) => {
                    // size 处理
                    item.size = this.setSize(item.size);
                });
                // 处理分区模块展示
                return defaultPartitions.reduce((pre: string, cur: {size: number; point: string; format: string;}) => {
                    if (hasLabel) return `${pre}<p>${cur.point}：<span>${cur.size},${cur.format};</span></p>`;
                    return `${pre}${cur.point}：${cur.size},${cur.format}；`;
                }, '');
            }
            catch {
                if (hasLabel) return defaultPartitions = [];
                return '--';
            }
        },

        systemPartitionHtml (val: SystemPartitionType[]) {
            const data: SystemPartitionType[] = deepCopy(val);
            data.map((item: SystemPartitionType) => {
                // size 处理
                item.size = this.setSize(item.size);
            });
            return data.reduce((pre: string, cur: SystemPartitionType) => {
                return `${pre}<span>${cur.point}：${cur.size},${cur.format}；</span>`;
            }, '');
        },

        setSize (size: number) {
            const sizeMap = new Map([
                [999999999, language.t('size.num')]
            ]);
            return sizeMap.get(size)?? `${size / 1024}GiB`;
        },

        setSizeNoUnit (size: number) {
            const sizeMap = new Map([
                [999999999, language.t('size.num')]
            ]);
            return sizeMap.get(size)?? size;
        }
    }
});

export default sysPartitionStore;
