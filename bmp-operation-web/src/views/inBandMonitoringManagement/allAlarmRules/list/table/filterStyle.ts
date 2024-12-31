class FilterStyleOperate {
    classNameData = [
        {
            name: 'def-type-status15',
            type: 'status'
        }
    ];
    filterStatus: any = reactive({
        status: false,
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
            props.allAlarmRulesList.searchTip, (newValue) => {
                if (!newValue.value) {
                    for (let index in this.filterStatus) {
                        this.filterStatus[index] = false;
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
            filterName?.addEventListener('click', (event) => {
                event.stopPropagation();
                for (let index in this.filterStatus) {
                    this.filterStatus[index] = false;
                }
            })
            filterOpt?.addEventListener('click', (event) => {
                event.stopPropagation();
                // if (!props.allAlarmRulesList.searchTip.value) {
                //     const clickStatus: any = {
                //         'status': () => {
                //             this.filterStatus['architecture'] = false;
                //             this.filterStatus['osType'] = false;
                //         },
                //         'architecture': () => {
                //             this.filterStatus['source'] = false;
                //             this.filterStatus['osType'] = false;
                //         },
                //         'osType': () => {
                //             this.filterStatus['architecture'] = false;
                //             this.filterStatus['source'] = false;
                //         }
                //     };
                //     clickStatus[item.type]();
                // }
                this.filterStatus[item.type] = !this.filterStatus[item.type];
            });

            const filterOtherRegion = document.querySelectorAll('.el-table-filter');
            filterOtherRegion.forEach((_, index) => {
                filterOtherRegion[index]?.addEventListener('click', (event) => {
                    event.stopPropagation();
                    for (let index in this.filterStatus) {
                        for (const key of this.classNameData) {
                            if (key.type === index && props.allAlarmRulesList.filter?.reactiveArr?.filterParams[key.type]?.length) {
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
                    this.filterDispose(ite, props);
                })
            })
        })
    };

    filterDispose = (type: string, props: any) => {
        const filterVal = props.filter.reactiveArr.filterParams;
        const status = Object.values(filterVal).some((item) => item !== '');
        if (!status) {
            this.filterStatus[type] = false;
            return;
        };
        if (props.allAlarmRulesList.searchTip.value) {
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

/**
 * @file
 * @author
*/

// class FilterStyleOperate {
//     classNameData = [
//         {
//             name: 'def-type-status16',
//             type: 'status'
//         }
//     ];
//     filterStatus: {
//         status: boolean
//     } = reactive({
//         status: false
//     })

//     constructor (props: any) {
//         onMounted(() => {
//             nextTick(() => {
//                 this.customFilter(props);
//             })
//             this.watchSearchTip(props);
//             this.watchFilter(props);
//         })
//     };

//     watchSearchTip = (props: any) => {
//         watch(() => 
//             props.allAlarmRulesList.searchTip, (newValue) => {
//                 const that = this;
//                 if (!newValue.value) {
//                     for (let index in this.filterStatus) {
//                         (this.filterStatus[`${index}` as keyof typeof that.filterStatus] as boolean) = false;
//                     }
//                 }
//                 else {
//                     this.setFilterStatus(props.filter.reactiveArr.filterParams);
//                 }
//             }, {
//                 deep: true
//             });
//     };

//     watchFilter = (props: any) => {
//         watch(() => props.filter.reactiveArr.filterParams, (newValue) => {
//             this.setFilterStatus(newValue);
//         }, {deep: true})
//     };

//     setFilterStatus = (params: any) => {
//         const that = this;
//         this.classNameData.forEach((item) => {
//             if (['', void 0].includes(params[item.type])) {
//                 (this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean) = false;
//                 return;
//             }
//             (this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean) = true;
//         })
//     };

//     customFilter = (props: any) => {
//         let filterName: HTMLElement | null;
//         let filterOpt: HTMLElement | null;
//         const that = this;
//         this.classNameData.forEach((item) => {
//             filterName = document.querySelector(`.${item.name} > .cell`);
//             filterOpt = document.querySelector(`.${item.name} > .cell > .el-table__column-filter-trigger`)
//             filterName?.addEventListener('click', (event) => {
//                 event.stopPropagation();
//                 for (let index in this.filterStatus) {
//                     (this.filterStatus[`${index}` as keyof typeof that.filterStatus] as boolean) = false;
//                 }
//             })
//             filterOpt?.addEventListener('click', (event) => {
//                 event.stopPropagation();
//                 (this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean) = !(this.filterStatus[`${item.type}` as keyof typeof that.filterStatus] as boolean);
//             });
//             const filterOtherRegion = document.querySelectorAll('.el-table-filter');
//             filterOtherRegion.forEach((_, index) => {
//                 filterOtherRegion[index]?.addEventListener('click', (event) => {
//                     event.stopPropagation();
//                     for (let index in this.filterStatus) {
//                         for (const key of this.classNameData) {
//                             if (index === key.type && props.allAlarmRulesList.filter.reactiveArr.filterParams[key.type]) {
//                                 (this.filterStatus[`${index}` as keyof typeof that.filterStatus] as boolean) = true;
//                             }
//                         }
//                     }
//                 })
//             })
            
//             let newType = [];
//             document.addEventListener('click', (event) => {
//                 event.stopPropagation();
//                 newType = [];
//                 newType.push(item.type);
//                 newType.forEach((ite) => {
//                     this.filterDispose(ite, props);
//                 })
//             })
//         })
//     };

//     filterDispose = (type: string, props: any) => {
//         const filterVal = props.filter.reactiveArr.filterParams;
//         const that = this;
//         if (props.allAlarmRulesList.searchTip.value) {
//             for (let key in filterVal) {
//                 if (key === type) {
//                     (this.filterStatus[`${key}` as keyof typeof that.filterStatus] as boolean) = true;
//                     // this.filterStatus[key] = true;
//                 }
//                 else {
//                     // this.filterStatus[type]
//                     if ((this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) && Object.keys(filterVal).length > 1) {
//                         (this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) = !(this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean);
//                     }
//                     else {
//                         (this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) = false;
//                         // this.filterStatus[type] = false;
//                     }
//                 }
//             }
//         }
//         else {
//             (this.filterStatus[`${type}` as keyof typeof that.filterStatus] as boolean) = false;
//         }
//     }
// };

// export default FilterStyleOperate;
