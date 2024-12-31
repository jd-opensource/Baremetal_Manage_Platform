class FilterStyleOperate {
    classNameData = [
        {
            name: 'def-type-status13',
            type: 'operation'
        },
        {
            name: 'def-type-status14',
            type: 'result'
        }
    ];
    filterStatus: any = reactive({
        operation: false,
        result: false
    })

    constructor (operateLog: any, filter: any) {
        onMounted(() => {
            nextTick(() => {
                this.customFilter(operateLog, filter);
            })
            this.watchSearchTip(operateLog, filter);
            this.watchFilter(filter);
        })
    };

    watchSearchTip = (operateLog: any, filter: any) => {
        watch(() => 
            operateLog.searchTip, (newValue) => {
                if (!newValue.value) {
                    for (let index in this.filterStatus) {
                        this.filterStatus[index] = false;
                    }
                }
                else {
                    this.setFilterStatus(filter.reactiveArr.filterParams);
                }
            }, {
                deep: true
            });
    };

    watchFilter = (filter: any) => {
        watch(() => filter.reactiveArr.filterParams, (newValue) => {
            this.setFilterStatus(newValue);
        }, {deep: true})
    };

    setFilterStatus = (params: any) => {
        this.classNameData.forEach((item) => {
            if (['', undefined].includes(params[item.type])) {
                this.filterStatus[item.type] = false;
                return;
            }
            this.filterStatus[item.type] = true;
        })
    };

    customFilter = (operateLog: any, filter: any) => {
        let filterName: HTMLElement | null;
        let filterOpt: HTMLElement | null;
        this.classNameData.forEach((item) => {
            filterName = document.querySelector(`.${item.name} > .cell`);
            filterOpt = document.querySelector(`.${item.name} > .cell > .el-table__column-filter-trigger`)
            filterName?.addEventListener('click', (event) => {
                event.stopPropagation();
                for (let index in this.filterStatus) {
                    this.filterStatus[index] = false;
                }
            })
            filterOpt?.addEventListener('click', (event) => {
                event.stopPropagation();
                if (!operateLog.searchTip.value) {
                    const clickStatus: any = {
                        'operation': () => {
                            this.filterStatus['result'] = false;
                        },
                        'result': () => {
                            this.filterStatus['operation'] = false;
                        }
                    };
                    clickStatus[item.type]();
                }
                this.filterStatus[item.type] = !this.filterStatus[item.type];
            });

            const filterOtherRegion = document.querySelectorAll('.el-table-filter');
            filterOtherRegion.forEach((_, index) => {
                filterOtherRegion[index]?.addEventListener('click', (event) => {
                    event.stopPropagation();
                    for (let index in this.filterStatus) {
                        for (const key of this.classNameData) {
                            if (key.type === index && filter?.reactiveArr?.filterParams[key.type]?.length) {
                                this.filterStatus[index] = true;
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
                    this.filterDispose(ite, filter, operateLog);
                })
            })
        })
    };

    filterDispose = (type: string, filter: any, operateLog: any) => {
        const filterVal = filter.reactiveArr.filterParams;
        const status = Object.values(filterVal).some((item) => item !== '');
        if (!status) {
            this.filterStatus[type] = false;
            return;
        };
        if (operateLog.searchTip.value) {
            for (let key in filterVal) {
                if (key === type) {
                    this.filterStatus[key] = true;
                }
                else {
                    if (this.filterStatus[type] && Object.keys(filterVal).length > 1) {
                        this.filterStatus[type] = !this.filterStatus[type];
                    }
                    else {
                        this.filterStatus[type] = false;
                    }
                }
            }
        }
        else {
            this.filterStatus[type] = false;
        }
    }
};

export default FilterStyleOperate;
