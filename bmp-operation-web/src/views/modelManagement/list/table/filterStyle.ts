/**
 * @file
 * @author
*/

class FilterStyleOperate {
    classNameData = [
        {
            name: 'def-type-status0',
            // name: 'def-model-type-status',
            type: 'deviceSeries'
        }
    ];
    filterStatus: {
        deviceSeries: boolean
    } = reactive({
        deviceSeries: false
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
        watch(() => 
            props.modelList.searchTip, (newValue) => {
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
                (this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean) = !(this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean);
            });
            const filterOtherRegion = document.querySelectorAll('.el-table-filter');
            filterOtherRegion.forEach((_, index) => {
                filterOtherRegion[index]?.addEventListener('click', (event) => {
                    event.stopPropagation();
                    for (let index in this.filterStatus) {
                        for (const key of this.classNameData) {
                            if (index === key.type && props.modelList.filter.reactiveArr.filterParams[key.type]) {
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
        const filterVal = props.filter.reactiveArr.filterParams;
        const that = this;
        if (props.modelList.searchTip.value) {
            for (let key in filterVal) {
                if (key === type) {
                    (this.filterStatus[`${key}` as keyof typeof that.filterStatus] as boolean) = true;
                    // this.filterStatus[key] = true;
                }
                else {
                    // this.filterStatus[type]
                    if ((this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) && Object.keys(filterVal).length > 1) {
                        (this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) = !(this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean);
                    }
                    else {
                        (this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) = false;
                        // this.filterStatus[type] = false;
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