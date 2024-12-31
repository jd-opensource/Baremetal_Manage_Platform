class FilterStyleOperate {
    classNameData = [
        {
            name: 'def-type-status8',
            type: 'name'
        },
        {
            name: 'def-type-status9',
            type: 'deviceSeries'
        },
        {
            name: 'def-type-status10',
            type: 'manageStatus'
        },
        {
            name: 'def-type-status11',
            type: 'collectStatus'
        }
    ];
    filterStatus: any = reactive({
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
        watch(() => 
            props.deviceList.searchTip, (newValue) => {
                if (!newValue.value) {
                    for (let index in this.filterStatus) {
                        this.filterStatus[index] = false;
                    }
                }
                else {
                    this.setFilterStatus(props.filterOperate.reactiveArr.filterParams);
                }
            }, {
                deep: true
            });
    };

    watchFilter = (props: any) => {
        watch(() => props.filterOperate.reactiveArr.filterParams, (newValue) => {
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

    customFilter = (props: any) => {
        let filterName: HTMLElement | null;
        let filterOpt: HTMLElement | null;
        this.classNameData.forEach((item) => {
            filterName = document.querySelector(`.${item.name} > .cell`);
            filterOpt = document.querySelector(`.${item.name} > .cell > .el-table__column-filter-trigger`)
            filterName?.addEventListener('click', () => {
                // this.filterStatus[item.type] = false;
                for (let index in this.filterStatus) {
                    this.filterStatus[index] = false;
                }
            });
            filterOpt?.addEventListener('click', () => {
                if (!props.deviceList.searchTip.value) {
                    const clickStatus: any = {
                        'name': () => {
                            this.filterStatus['deviceSeries'] = false;
                            this.filterStatus['manageStatus'] = false;
                            this.filterStatus['collectStatus'] = false;
                        },
                        'deviceSeries': () => {
                            this.filterStatus['name'] = false;
                            this.filterStatus['manageStatus'] = false;
                            this.filterStatus['collectStatus'] = false;
                        },
                        'manageStatus': () => {
                            this.filterStatus['deviceSeries'] = false;
                            this.filterStatus['name'] = false;
                            this.filterStatus['collectStatus'] = false;
                        },
                        'collectStatus': () => {
                            this.filterStatus['deviceSeries'] = false;
                            this.filterStatus['manageStatus'] = false;
                            this.filterStatus['name'] = false;
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
                            if (key.type === index && props.deviceList.filter?.reactiveArr?.filterParams[key.type]?.length) {
                                this.filterStatus[index] = true;
                            }
                        }
                    }
                })
            })
            let newType = [];
            document.addEventListener('click', () => {
                newType = [];
                newType.push(item.type);
                newType.forEach((ite) => {
                    this.filterDispose(ite, props);
                })
            })
        })
    };

    filterDispose = (type: string, props: any) => {
        const filterVal = props.filterOperate.reactiveArr.filterParams;
        if (props.deviceList.searchTip.value) {
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
