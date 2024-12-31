interface DataType {
    apikeyId: Ref<string>;
    name: Ref<string>;
    deleteKeyDiaLog: Ref<boolean>;
};

const deleteKey = (apiKey: any) => {
    const data: DataType = {
        apikeyId: ref<string>(''),
        name: ref<string>(''),
        deleteKeyDiaLog: ref<boolean>(false)
    };

    const deleteKeyClick = (item: {apiKeyId: string; name: string;}) => {
        data.name.value = item.name;
        data.apikeyId.value = item.apiKeyId;
        data.deleteKeyDiaLog.value = !data.deleteKeyDiaLog.value;
    };

    const deleteKeyCancel = (type: boolean): boolean => {
        return data.deleteKeyDiaLog.value = type;
    };

    const deleteKeySure = () => {
        apiKey.getApiKey();
    };

    return {
        ...data,
        deleteKeyClick,
        deleteKeyCancel,
        deleteKeySure
    }
};

export default deleteKey;
