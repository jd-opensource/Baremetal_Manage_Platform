const createKey = (apiKey: any) => {
    const createKeyDiaLog: Ref<boolean> = ref<boolean>(false);

    const createKeyClick = () => {
        createKeyDiaLog.value = !createKeyDiaLog.value;
    };

    const createKeyCancel = (type: boolean): boolean => {
        return createKeyDiaLog.value = type;
    };

    const createKeySure = () => {
        apiKey.getApiKey();
    };

    return {
        createKeyDiaLog,
        createKeyClick,
        createKeyCancel,
        createKeySure
    };
};

export default createKey;
