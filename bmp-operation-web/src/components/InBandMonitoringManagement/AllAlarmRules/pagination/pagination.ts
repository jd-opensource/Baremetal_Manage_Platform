interface DataType {
    pageSize: Ref<number>;
    pageNumber: Ref<number>;
    total: Ref<number>;
};

const paginationOperate = (fn: any) => {
    const data: DataType = {
        pageSize: ref<number>(20),
        pageNumber: ref<number>(1),
        total: ref<number>(0)
    };

    watch(() => [
        data.pageSize.value,
        data.pageNumber.value
    ], (_: any[]): void => {
        fn();
    });

    /**
     * 分页操作
     * @param {number} count 每页展示条数、页数
     * @param {string} type size、num类型
    */
    const paginationChange = (count: number, type: string) => {
        switch (type) {
            case 'num':
                data.pageNumber.value = count;
                break;
            case 'size':
                data.pageSize.value = count;
                data.pageNumber.value = 1;
                break;
        };
    };

    return {
        ...data,
        paginationChange
    };
};

export default paginationOperate;
