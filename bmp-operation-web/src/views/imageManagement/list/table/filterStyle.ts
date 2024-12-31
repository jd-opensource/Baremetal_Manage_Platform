/**
 * @file
 * @author
*/

interface FilterStatus {
    source: boolean;
    architecture: boolean;
    osType: boolean;
};

class FilterStyleOperate {
    classNameData = [
        {
            name: 'def-type-status5',
            type: 'source'
        },
        {
            name: 'def-type-status6',
            type: 'architecture'
        },
        {
            name: 'def-type-status7',
            type: 'osType'
        }
    ];
    filterStatus: FilterStatus = reactive<FilterStatus>({
        source: false,
        architecture: false,
        osType: false
    })

    constructor (props: any) {
        onMounted(() => {
            nextTick(() => {
                this.customFilter(props);
            })
            this.watchSearchTip(props);
            this.watchFilter(props);
        })
    };

    watchSearchTip = (props: any) => {
        watch(() => props.imageList.searchTip, (newValue) => {
            const that = this;
            if (!newValue.value) {
                for (let index in this.filterStatus) {
                    (this.filterStatus[`${index}` as keyof typeof that.filterStatus] as boolean) = false;
                }
            }
            else {
                this.setFilterStatus(props.filter.reactiveArr.filterParams);
            }
        }, {
            deep: true
        });
    };

    watchFilter = (props: any) => {
        watch(() => props.filter.reactiveArr.filterParams, (newValue) => {
            this.setFilterStatus(newValue);
        }, {deep: true})
    };

    setFilterStatus = (params: any) => {
        const that = this;
        this.classNameData.forEach((item) => {
            if (['', void 0].includes(params[item.type])) {
                (this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean) = false;
                return;
            }
            (this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean) = true;
        })
    };

    customFilter = (props: any) => {
        let filterName: HTMLElement | null;
        let filterOpt: HTMLElement | null;
        const that = this;
        this.classNameData.forEach((item) => {
            filterName = document.querySelector(`.${item.name} > .cell`);
            filterOpt = document.querySelector(`.${item.name} > .cell > .el-table__column-filter-trigger`)
            filterName?.addEventListener('click', (event) => {
                event.stopPropagation();
                for (let index in this.filterStatus) {
                    (this.filterStatus[`${index}` as keyof typeof that.filterStatus] as boolean) = false;
                }
            })
            filterOpt?.addEventListener('click', (event) => {
                event.stopPropagation();
                if (!props.imageList.searchTip.value) {
                    const clickStatus: any = {
                        'source': () => {
                            this.filterStatus['architecture'] = false;
                            this.filterStatus['osType'] = false;
                        },
                        'architecture': () => {
                            this.filterStatus['source'] = false;
                            this.filterStatus['osType'] = false;
                        },
                        'osType': () => {
                            this.filterStatus['architecture'] = false;
                            this.filterStatus['source'] = false;
                        }
                    };
                    clickStatus[item.type]();
                }
                this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] = !this.filterStatus[`${item.type}` as keyof typeof that.filterStatus];
            });

            const filterOtherRegion = document.querySelectorAll('.el-table-filter');
            filterOtherRegion.forEach((_, index) => {
                filterOtherRegion[index]?.addEventListener('click', (event) => {
                    event.stopPropagation();
                    for (let index in this.filterStatus) {
                        for (const key of this.classNameData) {
                            if (key.type === index && props.imageList.filter?.reactiveArr?.filterParams[key.type]?.length) {
                                (this.filterStatus[`${index}` as keyof typeof that.filterStatus] as boolean) = true;
                            }
                        }
                    }
                })
            })
            let newType = [];
            document.addEventListener('click', (event) => {
                event.stopPropagation();
                newType = [];
                newType.push(item.type);
                newType.forEach((ite) => {
                    this.filterDispose(ite, props);
                })
            })
        })
    };

    filterDispose = (type: string, props: any) => {
        const that = this;
        const filterVal = props.filter.reactiveArr.filterParams;
        const status = Object.values(filterVal).some((item) => item !== '');
        if (!status) {
            this.filterStatus[`${type}` as keyof typeof that.filterStatus] = false;
            return;
        };
        if (props.imageList.searchTip.value) {
            for (let key in filterVal) {
                if (Object.is(key, type)) {
                    this.filterStatus[`${key}` as keyof typeof that.filterStatus] = true;
                }
                else {
                    if (this.filterStatus[`${type}` as keyof typeof that.filterStatus] && Object.keys(filterVal).length > 1) {
                        this.filterStatus[`${type}` as keyof typeof that.filterStatus] = !this.filterStatus[`${type}` as keyof typeof that.filterStatus];
                    }
                    else {
                        this.filterStatus[`${type}` as keyof typeof that.filterStatus] = false;
                    }
                }
            }
        }
        else {
            (this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) = false;
        }
    }
};

export default FilterStyleOperate;
