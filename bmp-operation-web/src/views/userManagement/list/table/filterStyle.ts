class FilterStyleOperate {
    classNameData = [
        {
            name: 'def-type-status12',
            type: 'role'
        }
    ];
    filterStatus: any = reactive({
        role: false
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
            props.userList.searchTip, (newValue) => {
                if (!newValue.value) {
                    for (let index in this.filterStatus) {
                        this.filterStatus[index] = false;
                    }
                }
                else {
                    this.setFilterStatus(props.filterOperates.reactiveArr.filterParams);
                }
            }, {
                deep: true
            });
    };

    watchFilter = (props: any) => {
        watch(() => props.filterOperates.reactiveArr.filterParams, (newValue) => {
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
        const that = this;
        this.classNameData.forEach((item) => {
            filterName = document.querySelector(`.${item.name} > .cell`);
            filterOpt = document.querySelector(`.${item.name} > .cell > .el-table__column-filter-trigger`)
            filterName?.addEventListener('click', (event) => {
                event.stopPropagation();
                // this.filterStatus[item.type] = false;
                for (let index in this.filterStatus) {
                    this.filterStatus[index] = false;
                }
            })
            filterOpt?.addEventListener('click', (event) => {
                event.stopPropagation();
                this.filterStatus[item.type] = !this.filterStatus[item.type];
            });
            const filterOtherRegion = document.querySelectorAll('.el-table-filter');
            filterOtherRegion.forEach((_, index) => {
                filterOtherRegion[index]?.addEventListener('click', (event) => {
                    event.stopPropagation();
                    for (let index in this.filterStatus) {
                        for (const key of this.classNameData) {
                            if (index === key.type && props.filterOperates.reactiveArr.filterParams[key.type]) {
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
        const filterVal = props.filterOperates.reactiveArr.filterParams;
        if (props.userList.searchTip.value) {
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
